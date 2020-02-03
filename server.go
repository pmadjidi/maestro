package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"net"
	"sync"
	"time"
)

type appEnvelope struct {
	appName string
	app     chan *App
}

func newAppEnvelope(appName string) *appEnvelope {
	return &appEnvelope{appName, make(chan *App, 1)}
}

type CMD struct {
	terminate     chan notify
	deleteAllApps chan notify
	numberOfApps  chan chan int
	getApp        chan *appEnvelope
	createApp     chan *appEnvelope
}

type Server struct {
	*Flag
	grpcs    *grpc.Server
	apps     map[string]*App
	services map[string]Service
	status   map[string]*Flag
	cfg      *ServerConfig
	sync.RWMutex
	*CMD
}

func newCMD(cfg *ServerConfig) *CMD {
	return &CMD{
		make(chan notify, cfg.SERVER_QEUEU_LENGTH),
		make(chan notify, cfg.SERVER_QEUEU_LENGTH),
		make(chan chan int, cfg.SERVER_QEUEU_LENGTH),
		make(chan *appEnvelope, cfg.SERVER_QEUEU_LENGTH*10),
		make(chan *appEnvelope, cfg.SERVER_QEUEU_LENGTH),
	}
}

func NewServer() *Server {
	cfg := createServerConfig()
	return &Server{NewFlag(), grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor)), make(map[string]*App), make(map[string]Service), make(map[string]*Flag), cfg,
		sync.RWMutex{}, newCMD(cfg)}
}

func (s *Server) CMDcreateApp(appName string) (*App, error) {
	return s.GetOrCreateApp(appName, true)
}

func (s *Server) CMDgetApp(appName string) (*App, error) {
	return s.GetOrCreateApp(appName, false)
}

func (s *Server) GetOrCreateApp(appName string, create bool) (*App, error) {
	var app *App
	aenv := newAppEnvelope(appName)
	if create {
		 s.createApp <- aenv     
	} else {
		 s.getApp <- aenv
	}

	app = <-aenv.app

	if app == nil && !create {
		return nil, fmt.Errorf(Status_NOTFOUND.String())
	}
	if app == nil && create {
		return nil, fmt.Errorf(Status_EXITSTS.String())
	}
	return app, nil
}

func (s *Server) registerServices() {
	ls := newLoginService(s)
	s.services["loginServices"] = ls
	rs := newRegisterService(s)
	s.services["registerServices"] = rs
	ms := newMsgService(s)
	s.services["messageService"] = ms
}

func (s *Server) commandCenter() {
	s.log(fmt.Sprintf("Processing commands for [%s]", s.cfg.SYSTEM_NAME))
	for {
		select {
		case req := <-s.getApp:
			s.processGetApp(req)
		default:
			select {
			case req := <-s.getApp:
				s.processGetApp(req)
			case req := <-s.createApp:
				s.processCreateApp(req)
			case <-time.After(10 * time.Second):
				s.processDirectoryWatchDog()
				s.log("End of directory scan")
			}
		}
	}
}

func (s *Server) Start(reset bool) {

	PrettyPrint(s.cfg)
	if reset {
		s.removeAllApps()
	}
	s.processDirectoryWatchDog()
	s.registerServices()
	for serviceName, srv := range s.services {
		switch serviceName {
		case "loginServices":
			s.log("RPC Registring loginService")
			RegisterLoginServer(s.grpcs, srv.(LoginServer))
		case "registerServices":
			s.log("RPC Registring RegisterService")
			RegisterRegisterServer(s.grpcs, srv.(RegisterServer))
		case "messageService":
			s.log("RPC Registring MessageService")
			RegisterMsgServer(s.grpcs, srv.(MsgServer))
		}
	}

	go s.commandCenter()

	lis, err := net.Listen("tcp", s.cfg.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.log(fmt.Sprintf("Listening to port[%s]", s.cfg.PORT))
	if err := s.grpcs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) log(message string) {
	fmt.Printf("@[%d][%s] %s ...\n", time.Now().Second(), s.cfg.SYSTEM_NAME, message)
}

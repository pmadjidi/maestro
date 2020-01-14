package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"net"
	"sync"
)

type Server struct {
	grpcs    *grpc.Server
	apps     map[string]*App
	services map[string]Service
	cfg      *ServerConfig
	sync.RWMutex
	terminate chan struct{}
}

func NewServer() *Server {
	cfg := createServerConfig()
	return &Server{grpc.NewServer(),make(map[string]*App),make(map[string]Service),cfg,
		sync.RWMutex{},make(chan struct{})}
}


func (s *Server) NewApp(appName string) (*App,error) {
	s.Lock()
	defer s.Unlock()
	_,ok := s.apps[appName]
	if !ok {
		app,err := newApp(appName)
		if err!= nil {
			return nil,err
		}
		s.apps[appName] = app
		app.start()
		return app,nil

	}
	return nil,fmt.Errorf(Status_EXITSTS.String())
}

func (s *Server) GetApp(appName string) (*App,error) {
	s.RLock()
	defer s.RUnlock()
	app,ok := s.apps[appName]
	if !ok {
		return nil,fmt.Errorf(Status_INVALID_APPNAME.String())
	}
	return app,nil
}



func (s *Server) registerServices() {
	s.Lock()
	defer s.Unlock()

	ls := newLoginService(s)
	s.services["loginServices"] = ls
	rs := newRegisterService(s)
	s.services["registerServices"] = rs
	ms := newMsgService(s)
	s.services["messageService"] = ms
}

func (s *Server) StartApps() {
	s.Lock()
	defer s.Unlock()

	for appName, app := range s.apps {
		fmt.Printf("Starting App[%s]", appName)
		app.start()
	}
}

func (s *Server) Start() {



	PrettyPrint(s.cfg)
	s.StartApps()
	s.registerServices()

	for serviceName, srv := range s.services {
		switch serviceName {
		case "loginServices":
			fmt.Printf("RPC Registring loginService...\n")
			RegisterLoginServer(s.grpcs, srv.(LoginServer))
		case "registerServices":
			fmt.Printf("RPC Registring RegisterService...\n")
			RegisterRegisterServer(s.grpcs, srv.(RegisterServer))
		case "messageService":
			fmt.Printf("RPC Registring MessageService...\n")
			RegisterMessageServer(s.grpcs, srv.(MessageServer))
		}
	}

	lis, err := net.Listen("tcp", s.cfg.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening to port[%s]\n", s.cfg.PORT)
	if err := s.grpcs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

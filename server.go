package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	. "maestro/api"
	"net"
	"os"
	"path"
	"sync"
	"time"
)

type Server struct {
	*Flag
	grpcs    *grpc.Server
	apps     map[string]*App
	services map[string]Service
	status   map[string]*Flag
	cfg      *ServerConfig
	sync.RWMutex
	terminate chan struct{}
}

func NewServer() *Server {
	cfg := createServerConfig()
	return &Server{NewFlag(), grpc.NewServer(), make(map[string]*App), make(map[string]Service), make(map[string]*Flag), cfg,
		sync.RWMutex{}, make(chan struct{})}
}

func (s *Server) NewApp(appName string) (*App, error) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.apps[appName]
	if !ok {
		app, err := newApp(appName)
		if err != nil {
			return nil, err
		}
		s.apps[appName] = app
		s.status[appName] = NewFlag()
		s.status[appName].Set(NEW)
		return app, nil

	}
	return nil, fmt.Errorf(Status_EXITSTS.String())
}

func (s *Server) GetApp(appName string) (*App, error) {
	s.RLock()
	defer s.RUnlock()
	app, ok := s.apps[appName]
	if !ok {
		return nil, fmt.Errorf(Status_INVALID_APPNAME.String())
	}
	status, ok := s.status[appName]
	if ok && status.Is(BLOCKED) {
		return nil, fmt.Errorf("Blocked")
	}
	return app, nil
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
	files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		appNAME := file.Name()
		app, err := s.NewApp(appNAME)
		if err != nil {
			fmt.Printf("Could not start app [%s] skipping... \n", appNAME)
		} else {
			exists := fileExists( s.cfg.SYSTEM_PATH+appNAME+"/blocked")
			if exists {
				fmt.Printf("%s blocks request to App %s \n", s.cfg.SYSTEM_NAME, appNAME)
				s.status[appNAME].Set(BLOCKED)
			} else {
				fmt.Printf("Registring app [%s]\n", file.Name())
				fmt.Printf("Starting app [%s]... \n", appNAME)
				app.start()
			}
		}
	}
}

func (s *Server) monitor() {
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Printf("Monitor: Tick..\n")
			files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
			if err != nil {
				fmt.Printf("monitor error, can not read system directory...\n: %s", err.Error())
			} else {
				for _, file := range files {
					appNAME := file.Name()
					exists := fileExists(s.cfg.SYSTEM_PATH+appNAME+"/blocked")
					if exists {
						s.Lock()
						if !s.status[appNAME].Is(BLOCKED) {
							fmt.Printf("%s Blocking app %s \n", s.cfg.SYSTEM_NAME, appNAME)
							s.status[appNAME].Set(BLOCKED)
						}
						s.Unlock()
					} else {
						s.Lock()
						if s.status[appNAME].Is(BLOCKED) {
							fmt.Printf("Unblocking app [%s]\n", file.Name())
							s.status[appNAME].Clear(BLOCKED)
						}
						s.Unlock()
					}
				}
			}
		}
	}
}

func (s *Server) removeAllApps() {
	fmt.Printf("Reset flag recived, removing all apps from system node....")
	files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Printf("Removing app [%s]...\n", file.Name())
		os.RemoveAll(path.Join([]string{s.cfg.SYSTEM_PATH, file.Name()}...))
	}

}

func (s *Server) Start(reset bool) {

	PrettyPrint(s.cfg)

	if reset {
		s.removeAllApps()
	}

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

	go s.monitor()

	lis, err := net.Listen("tcp", s.cfg.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening to port[%s]\n", s.cfg.PORT)
	if err := s.grpcs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

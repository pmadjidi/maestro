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
	"sort"
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

	return &Server{NewFlag(), grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor)), make(map[string]*App), make(map[string]Service), make(map[string]*Flag), cfg,
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

/*
func (s *Server) StartApps() {
	files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(s.cfg.SYSTEM_PATH, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		appNAME := file.Name()
		app, err := s.NewApp(appNAME)
		if err != nil {
			s.log(fmt.Sprintf("Could not start app [%s] skipping", appNAME))
		} else {
			exists := fileExists(s.cfg.SYSTEM_PATH + appNAME + "/blocked")
			if exists {
				s.log(fmt.Sprintf("blocks request to App [%s]", appNAME))
				s.status[appNAME].Set(BLOCKED)
			} else {
				s.log(fmt.Sprintf("Registring app [%s]", file.Name()))
				s.log(fmt.Sprintf("Starting app [%s]", appNAME))
				app.start()
			}
		}
	}
}

 */

func (s *Server) getNumberOfApps() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.apps)
}

func (s *Server) monitor() {
	var blockedAppNames []string
	for {
		blockedAppNames = nil
		select {
		case <-time.After(10 * time.Second):
			s.log(fmt.Sprintf("Monitoring [%d] Apps tick ... Zzzz", len(s.apps)))
			files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
			if err != nil {
				s.log(fmt.Sprintf("monitor error, can not read system directory [%s]", err.Error()))
			} else {
				for _, file := range files {
					appNAME := file.Name()
					if fileExists(s.cfg.SYSTEM_PATH + appNAME + "/blocked") {
						blockedAppNames = append(blockedAppNames, appNAME)
					}
				}
				sort.Strings(blockedAppNames)
				s.Lock()
				for _, app := range s.apps {
					if contains(blockedAppNames, app.name) {
						if !s.status[app.name].Is(BLOCKED) {
							s.log(fmt.Sprintf("Hum... Blocking app [%s]", app.name))
							s.status[app.name].Set(BLOCKED)
							s.apps[app.name].Stop()
						}
					} else {
						if s.status[app.name].Is(BLOCKED) {
							s.log(fmt.Sprintf("Hum... Unblocking app [%s]", app.name))
							s.status[app.name].Clear(BLOCKED)
							s.apps[app.name].startSubSystems()
						}
					}
				}
				s.Unlock()
				for _, possibleNewAppName := range files {
					s.RLock()
					_,ok := s.apps[possibleNewAppName.Name()]
					s.RUnlock()
					if !ok {
						app, err := s.NewApp(possibleNewAppName.Name())
						if err != nil {
							s.log(fmt.Sprintf("Could not start app [%s] skipping", possibleNewAppName.Name()))
						} else {
							app.start()
						}
					}
				}
			}
		}
	}
}

func (s *Server) removeAllApps() {
	fmt.Printf("Reset flag recived, removing all apps from system node....")
	files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(s.cfg.SYSTEM_PATH, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
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

	go s.monitor()
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
			RegisterMessageServer(s.grpcs, srv.(MessageServer))
		}
	}



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

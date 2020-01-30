package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
)

func (s *Server) processGetApp(req *appEnvelope) {
	s.log(fmt.Sprintf("processGetApp called [%s]", req.appName))
	if app, ok := s.apps[req.appName]; ok {
		req.app <- app
	} else {
		s.log(fmt.Sprintf("App [%s] does not exist", req.appName))
		req.app <- nil
	}
}

func (s *Server) processCreateApp(req *appEnvelope) {
	s.log(fmt.Sprintf("processCreateApp called [%s]", req.appName))
	appName := req.appName
	_, ok := s.apps[appName]
	if !ok {
		app, err := newApp(s.cfg, appName)
		if err != nil {
			req.app <- nil
			return
		}
		s.apps[appName] = app
		s.status[appName] = NewFlag()
		s.status[appName].Set(NEW)
		go app.start()
		req.app <- app
	} else {
		s.log(fmt.Sprintf("Failed to create App [%s]", req.appName))
		req.app <- nil
	}
}

func (s *Server) processDirectoryWatchDog() {
	var blockedAppNames []string
	s.log(fmt.Sprintf("Monitoring [%d] Apps tick ... Zzzz", len(s.apps)))
	files, err := ioutil.ReadDir(s.cfg.SYSTEM_PATH)
	if err != nil {
		s.log(fmt.Sprintf("Can not read system directory [%s] creating system directory", err.Error()))
		err := os.Mkdir(s.cfg.SYSTEM_NAME, os.ModePerm)
		handleError(err)
	} else {
		for _, file := range files {
			appNAME := file.Name()
			if fileExists(s.cfg.SYSTEM_PATH + appNAME + "/blocked") {
				blockedAppNames = append(blockedAppNames, appNAME)
			}
		}
		sort.Strings(blockedAppNames)
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
					s.apps[app.name].reStartSubSystems()
				}
			}
		}
		for _, possibleNewAppName := range files {
			_, ok := s.apps[possibleNewAppName.Name()]
			if !ok {
				env := newAppEnvelope(possibleNewAppName.Name())
				s.createApp <- env
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

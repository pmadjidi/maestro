package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"net"
)

type Service interface {
	Name() string
}

type App struct {
	quit      chan bool
	services  map[string]Service
	cfg       *ServerConfig
	server    *grpc.Server
	users     *usersdb
	messages  *messagesdb
	loginQ    chan *loginEnvelope
	registerQ chan *registerEnvelope
	msgQ      chan *msgEnvelope
	UserDbQ   chan []*User
	MsgDBQ    chan[]*Message
	DATABASE  *sql.DB
}

func newApp() *App {
	cfg := createLoginServerConfig()
	app := App{make(chan bool), make(map[string]Service), cfg, grpc.NewServer(),
		newUserdb(cfg.ARRAY_PRE_ALLOCATION_LIMIT),
		newMessageDb(),
		make(chan *loginEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *registerEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *msgEnvelope,cfg.SERVER_QEUEU_LENGTH),
		make(chan []*User, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*Message, cfg.SERVER_QEUEU_LENGTH),
		newDatabase(cfg),
	}
	app.readUsersFromDatabase()
	app.registerServices()
	return &app
}

func (a *App) registerServices() {
	ls := newLoginService(a.loginQ, a.cfg)
	a.services["loginServices"] = ls
	rs := newRegisterService(a.registerQ, a.cfg)
	a.services["registerServices"] = rs
	ms := newMsgService(a.msgQ,a.cfg)
	a.services["messageService"] = ms
}

func (a *App) start() {
	PrettyPrint(a.cfg)
	for serviceName, s := range a.services {
		switch serviceName {
		case "loginServices":
			fmt.Printf("RPC Registring loginService...\n")
			RegisterLoginServer(a.server, s.(LoginServer))
		case "registerServices":
			fmt.Printf("RPC Registring RegisterService...\n")
			RegisterRegisterServer(a.server, s.(RegisterServer))
		case "messageService":
			fmt.Printf("RPC Registring MessageService...\n")
			RegisterMessageServer(a.server,s.(MessageServer))
		}
	}

	a.Run()

	lis, err := net.Listen("tcp", a.cfg.PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening to port[%s]\n", a.cfg.PORT)
	if err := a.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Run() {
	a.readUsersFromDatabase()
	//a.readMessagesFromDatabase()
	//a.readSubscriptionsFromDatabase()
	go a.userManager()
	go a.messageManager()
	go a.databaseManager()

}



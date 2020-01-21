package main

import (
	"database/sql"
	"fmt"
	. "maestro/api"
	"time"
)

type Service interface {
	Name() string
}

type App struct {
	name     string
	quit      chan bool
	cfg       *ServerConfig
	users     *usersdb
	messages  *messagesdb
	loginQ    chan *loginEnvelope
	registerQ chan *registerEnvelope
	msgSendQ      chan *msgEnvelope
	msgRecQ      chan *msgEnvelope
	UserDbQ   chan []*User
	MsgDBQ    chan []*Message
	DATABASE  *sql.DB
}

func newApp(name string) (*App, error) {
	if name == ""  {
		return nil, fmt.Errorf(Status_INVALID_APPNAME.String())
	}

	cfg := createServerConfig(name)
	app := App{name,make(chan bool), cfg,
		newUserdb(cfg.ARRAY_PRE_ALLOCATION_LIMIT),
		newMessageDb(),
		make(chan *loginEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *registerEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *msgEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *msgEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*User, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*Message, cfg.SERVER_QEUEU_LENGTH),
		newDatabase(cfg),
	}
	app.readUsersFromDatabase()
	return &app,nil
}


func (a *App) start() *App {
	PrettyPrint(a.cfg)
	a.log("Reading from database")
	a.readUsersFromDatabase()
	a.readMessagesFromDatabase()
	//a.readSubscriptionsFromDatabase()
	a.log("Starting subsystems")
	a.startSubSystems()
	return a
}

func (a *App) startSubSystems() {
	go a.userManager()
	go a.messageManager()
	go a.databaseManager()
}

func (a *App) log(message string) {
	fmt.Printf("@[%d]-[%s] App: [%s] %s ...\n",time.Now().Second(),a.cfg.SYSTEM_NAME,a.cfg.APP_NAME,message)
}


func (a *App) StopLoginService() {
	a.log("Stoping user manger")
	close(a.loginQ)
	close(a.registerQ)
}



func (a *App) StopMessageService() {
	a.log("Stoping messaging service")
	close(a.msgRecQ)
}

func (a *App) StopDatabaseService() {
	a.log("Stoping Database service")
	close(a.MsgDBQ)
	close(a.UserDbQ)
}



func (a *App) Stop() {
	a.log("Stoping")
	a.StopLoginService()
	a.StopMessageService()
	a.StopDatabaseService()
}








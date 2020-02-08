package main

import (
	"database/sql"
	"fmt"
	. "maestro/api"
	"time"
)

type Service interface {
	getname() string
}


type appCommands struct {
	loginQ    chan *loginEnvelope
	registerQ chan *registerEnvelope
	userQ  chan *userEnvelope
	timeLineQ      chan *msgEnvelope
	msgQ      chan *msgEnvelope
	topicSubQ     chan *topicEnvelope
	topicUnSubQ    chan *topicEnvelope
	topicListQ chan *topicEnvelope
	UserDbQ   chan []*User
	MsgDBQ    chan []*Message
}

func newAppCommands(cfg *ServerConfig) *appCommands {
	return &appCommands{
		make(chan *loginEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *registerEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *userEnvelope),
		make(chan *msgEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *msgEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *topicEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *topicEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *topicEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*User, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*Message, cfg.SERVER_QEUEU_LENGTH),
	}
}

type App struct {
	name     string
	quit      chan bool
	cfg       *AppConfig
	*appCommands
	*usersdb
	*messagesdb
	*sql.DB
}



func newApp(cfg *ServerConfig,name string) (*App, error) {
	if name == ""  {
		return nil, fmt.Errorf(Status_INVALID_APPNAME.String())
	}

	appCfg := createAppConfig(cfg,name)
	app := App{name,make(chan bool),
		appCfg,
		newAppCommands(cfg),
		newUserdb(cfg.ARRAY_PRE_ALLOCATION_LIMIT),
		newMessageDb(),
		newDatabase(appCfg),
	}
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

func (a *App) reStartSubSystems() { // dump closed channels on restart...
	a.appCommands = newAppCommands(a.cfg.ServerConfig)
	go a.userManager()
	go a.messageManager()
	go a.databaseManager()
}



func (a *App) log(message string) {
	fmt.Printf("@[%d]---[%s]App: [%s] %s ...\n",time.Now().Second(),a.cfg.SYSTEM_NAME,a.cfg.APP_NAME,message)
}


func (a *App) StopLoginService() {
	close(a.loginQ)
	close(a.registerQ)
	close(a.timeLineQ)
	close(a.userQ)
	a.log("Stoping user manger")
}



func (a *App) StopMessageService() {
	close(a.msgQ)
	close(a.topicUnSubQ)
	close(a.topicSubQ)
	a.log("Stoping messaging service")
}

func (a *App) StopDatabaseService() {
	close(a.MsgDBQ)
	close(a.UserDbQ)
	a.log("Stoping Database service")
}



func (a *App) Stop() {
	a.StopLoginService()
	a.StopMessageService()
	a.StopDatabaseService()
	a.log("Stoping")
}








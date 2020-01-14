package main

import (
	"database/sql"
	"fmt"
	. "maestro/api"
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
	msgQ      chan *msgEnvelope
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
		make(chan []*User, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*Message, cfg.SERVER_QEUEU_LENGTH),
		newDatabase(cfg),
	}
	app.readUsersFromDatabase()
	return &app,nil
}


func (a *App) start() {
	PrettyPrint(a.cfg)
	a.readUsersFromDatabase()
	a.readMessagesFromDatabase()
	//a.readSubscriptionsFromDatabase()

	go a.userManager()
	go a.messageManager()
	go a.databaseManager()



}



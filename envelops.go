package main

import . "maestro/api"

type appEnvelope struct {
	appName string
	app     chan *App
	Status
}

func newAppEnvelope(appName string) *appEnvelope {
	return &appEnvelope{appName, make(chan *App, 1),Status_NEW}
}

type topicEnvelope struct {
	*TopicReq
	Username *string
	resp chan notify
	Status
}

func newTopicEnvelope() *topicEnvelope {
	return &topicEnvelope{nil, nil,make(chan notify),Status_NEW}
}


type registerEnvelope struct {
	req  chan *RegisterReq
	resp chan notify
	token  *string
	Status
}


func newRegisterEnvelope() *registerEnvelope {
	return &registerEnvelope{make(chan *RegisterReq, 1), make(chan notify, 1),nil,Status_NEW}
}

type loginEnvelope struct {
	resp     chan notify
	token    *string
	username *string
	password *string
	device   *string
	Status
}

func newLoginEnvelope() *loginEnvelope {
	return &loginEnvelope{make(chan notify, 1), nil, nil, nil, nil, Status_NEW}
}

type userEnvelope struct {
	Username *string
	User *User
	resp chan notify
	Status
}

func newUserEnvelope () *userEnvelope {
	return &userEnvelope{nil,nil,make(chan notify,1),Status_NEW}
}








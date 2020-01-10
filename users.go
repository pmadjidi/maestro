package main

import (
	"github.com/google/uuid"
	. "maestro/api"
	"sync"
	"time"
)

type User struct {
	*RegisterReq
	*sync.RWMutex
	uid string
	status *Flag
	modified time.Time
	loginAttempts int
	messages chan *Message
	topics []string
}


func newUser(req *RegisterReq) *User {
	id := uuid.New()
	//fmt.Printf("%s Creating user %s with id %s \n",time.Now(),req.UserName,id)
	req.PassWord = hashAndSalt(req.PassWord)
	newUser := User{req,&sync.RWMutex{},id.String(),NewFlag(),
		time.Now(),0,make(chan *Message),make([]string,0)}
	return &newUser
}


package main

import (
	"fmt"
	"github.com/google/uuid"
	"maestro/api"
	"sync"
	"time"
)

type User struct {
	*api.RegisterReq
	*sync.RWMutex
	uid string
	status *Flag
	modified time.Time
	loginAttempts int
}


func newUser(req *api.RegisterReq) *User {
	id := uuid.New()
	fmt.Printf("%s Creating user %s with id %s \n",time.Now(),req.UserName,id)
	req.PassWord = hashAndSalt(req.PassWord)
	newUser := User{req,&sync.RWMutex{},id.String(),NewFlag(),
		time.Now(),0}
	newUser.status.Set(DIRTY)
	return &newUser
}


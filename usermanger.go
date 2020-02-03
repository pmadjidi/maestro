package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	. "maestro/api"
	"time"
)

func (a *App) issueToken(userName, device string) (string, error) {
	//fmt.Printf("Got %s, %s,%s %s %s \n",a.cfg.SYSTEM_NAME,userName,device,a.cfg.SYSTEM_NAME,a.cfg.APP_NAME)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer":    a.cfg.SYSTEM_NAME,
		"username":  userName,
		"superuser": "false",
		"device":    device,
		"stamp":     time.Now().String(), // vaid for a week
		"appname":   a.cfg.APP_NAME,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(a.cfg.SYSTEM_SECRET))
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func (a *App) userManager() {
	a.log("LoginServer, Entering processing loop")
	signalLoginQ := false
	signalRegisterQ := false
	signalMegSendQ := false

loop:
	for {
		select {
		case env, ok := <-a.loginQ:
			if ok {
				status := a.tryLogin(*env.username, []byte(*env.password))
				if status == Status_SUCCESS {
					token, err := a.issueToken(*env.username, *env.device)
					if err != nil {
						env.Status = Status_ERROR
					} else {
						env.token = &token
						env.Status = Status_SUCCESS
					}
				} else {
					env.Status = status
					fmt.Printf("In here %s\n", env.Status.String())
				}
				env.resp <- notify{}
			} else {
				signalLoginQ = true
			}

		case <-time.After(a.cfg.WRITE_LATENCY * time.Millisecond):
			//fmt.Printf("loginServer: Looking for changes in user database...\n")
			for _, aUser := range a.udb {
				aUser.Lock()
				if aUser.status.Is(DIRTY) == true {
					a.udirty = append(a.udirty, aUser)
					a.udirtyCounter += 1
					aUser.status.Clear(DIRTY)
				}
				aUser.Unlock()
			}
			if a.udirtyCounter > 0 {
				fmt.Printf("LoginServer, dirty users found [%d]...\n", len(a.udirty))
				select {
				case a.UserDbQ <- a.udirty:
					a.udirty = make([]*User, 0)
					a.udirtyCounter = 0
				case <-time.After(2 * time.Second):
					fmt.Printf("loginServer: database server blocked ...\n")
				}
			}

		case env, ok := <-a.registerQ:
			if ok {
				if len(a.udb) > a.cfg.MAX_NUMBER_OF_USERS {
					env.status = Status_MAXIMUN_NUMBER_OF_USERS_REACHED
				} else {
					req := <-env.req
					_, exists := a.udb[req.UserName]
					if exists {
						fmt.Printf("User %+v exists\n", req)
						env.status = Status_EXITSTS
					} else {
						newUser := newUser(req)
						newUser.status.Set(DIRTY)
						a.udb[newUser.UserName] = newUser
						token, err := a.issueToken(req.GetUserName(), req.GetDevice())
						if err != nil {
							fmt.Printf("Error %+v", err)
							env.status = Status_NOAUTH
						} else {
							env.token = &token
							env.status = Status_SUCCESS
						}
					}
					env.resp <- notify{}
				}
			} else {
				signalRegisterQ = true
			}

		case env, ok := <-a.msgSendQ:
			if ok {
				aUser, ok := a.udb[env.userName]
				if !ok {
					env.Status = Status_INVALID_USERNAME
				} else {
					env.Status = Status_SUCCESS
					env.messages = aUser.timeLine
				}
				env.resp <- notify{}
			} else {
				signalMegSendQ = true
				if signalLoginQ && signalRegisterQ && signalMegSendQ {
					break loop
				}
			}

		case <-a.quit:
			break loop

		}
	}

	a.log("Exiting user Manager")

}

func (a *App) tryLogin(userName string, pass []byte) Status {

	aUser, userExists := a.udb[userName]
	if !userExists {
		return Status_INVALID_USERNAME
	}

	aUser.Lock()
	defer aUser.Unlock()

	if aUser.status.Is(BLOCKED) {
		return Status_BLOCKED
	}

	if bcrypt.CompareHashAndPassword(aUser.PassWord, pass) != nil {
		aUser.loginAttempts += 1
		if aUser.loginAttempts >= a.cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT-1 {
			a.log(fmt.Sprintf("Blocking user [%s]", userName))
			aUser.status.Set(BLOCKED)
			aUser.status.Set(DIRTY)
			return Status_BLOCKED
		}
		a.log(fmt.Sprintf("Login failed, wrong password [%s]", userName))
		return Status_NOAUTH
	}
	return Status_SUCCESS
}

func (a *App) presistUser(users []*User) {
	tx, err := a.Begin()
	handleError(err)
	for i := 0; i < len(users); i++ {
		u := users[i]
		u.RLock()
		_, err := tx.Exec("INSERT OR REPLACE INTO users (uid, status,UserName,Password,FirstName,LastName,Email, Phone,Device) VALUES (?, ?,?,?,?,?,?,?,?)",
			u.uid, u.status.Get(), u.UserName, u.PassWord, u.FirstName, u.LastName,
			u.Email, u.Phone, u.Device)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		_, err = tx.Exec("INSERT OR REPLACE INTO address (uid,Zip,Street,City,State) VALUES (?,?,?,?,?)",
			u.uid, u.Address.Zip, u.Address.Street, u.Address.City, u.Address.State)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		u.RUnlock()
	}
	handleError(tx.Commit())
	/*
		for _, u := range users {
			fmt.Printf("Presisted user[%s]\n", u.UserName)
		}
	*/
	a.log(fmt.Sprintf("[%s] Pressised %d Users in batch", a.name, len(users)))
}

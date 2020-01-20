package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	. "maestro/api"
	"time"
)




func (a *App) issueToken(userName,device string) (string,error) {
	//fmt.Printf("Got %s, %s,%s %s %s \n",a.cfg.SYSTEM_NAME,userName,device,a.cfg.SYSTEM_NAME,a.cfg.APP_NAME)


	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"issuer": a.cfg.SYSTEM_NAME,
			"username": userName,
			"superuser": "false",
			"device": device,
			"stamp": time.Now().String(), // vaid for a week
			"appname": a.cfg.APP_NAME,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(a.cfg.SYSTEM_SECRET))
	fmt.Printf("?? %s",tokenString)
	if err != nil {
		return "", err
	} else {
		return tokenString,nil
	}
}


func (a *App) userManager() {

	fmt.Println("LoginServer, Entering processing loop...")
	for {
		select {
		case env := <-a.loginQ:
			status :=  a.tryLogin(*env.username,[]byte(*env.password))
			if status == Status_SUCCESS {
				token,err := a.issueToken(*env.username,*env.device)
				if err != nil {
					env.Status = Status_ERROR
				} else {
					env.token = &token
					env.Status = Status_SUCCESS
				}
			} else {
				env.Status = status
				fmt.Printf("In here %s\n",env.Status.String())
			}
			env.resp <- notify{}
		case <-time.After(a.cfg.WRITE_LATENCY * time.Millisecond):
			//fmt.Printf("loginServer: Looking for changes in user database...\n")
			for _, aUser := range a.users.db {
				aUser.Lock()
				if aUser.status.Is(DIRTY) == true {
					a.users.dirty = append(a.users.dirty,aUser)
					a.users.dirtyCounter += 1
					aUser.status.Clear(DIRTY)
				}
				aUser.Unlock()
			}
			if a.users.dirtyCounter > 0 {
				fmt.Printf("LoginServer, dirty users found [%d]...\n", len(a.users.dirty))
				select {
				case a.UserDbQ <- a.users.dirty:
					a.users.dirty = make([]*User, 0)
					a.users.dirtyCounter = 0
				case <-time.After(2 * time.Second):
					fmt.Printf("loginServer: database server blocked ...\n")
				}
			}

		case env := <-a.registerQ:
			if len(a.users.db) > a.cfg.MAX_NUMBER_OF_USERS {
				env.status = Status_MAXIMUN_NUMBER_OF_USERS_REACHED
			} else {
				req := <-env.req
				_, exists := a.users.db[req.UserName]
				if exists {
					fmt.Printf("User %+v exists\n", req)
					env.status = Status_EXITSTS
				} else {
					newUser := newUser(req)
					newUser.status.Set(DIRTY)
					a.users.db[newUser.UserName] = newUser
					token,err := a.issueToken(req.GetUserName(),req.GetDevice())
					if err != nil {
						fmt.Printf("Error %+v",err)
						env.status = Status_NOAUTH
					} else {
						env.token = &token
						env.status = Status_SUCCESS
					}
				}
				env.resp <- struct{}{}
			}

			case env := <- a.msgSendQ:
				aUser,ok := a.users.db[env.userName]
				if !ok {
					env.Status = Status_INVALID_USERNAME
				} else {
					env.Status = Status_SUCCESS
					env.messages <- aUser.timeLine
				}
				env.resp <- notify{}


		case <-a.quit:
			break
		}

	}

	fmt.Println("LoginServer, Exit processing loop...")
}

func (a *App) tryLogin(userName string, pass []byte) Status {

	aUser, userExists := a.users.db[userName]
	if !userExists {
		return Status_INVALID_USERNAME
	}

	aUser.Lock()
	defer aUser.Unlock()

	if aUser.status.Is(BLOCKED) {
		return Status_BLOCKED
	}

	if bcrypt.CompareHashAndPassword(aUser.PassWord,pass) != nil  {
		aUser.loginAttempts += 1
		if aUser.loginAttempts >= a.cfg.MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT - 1 {
			fmt.Printf("Blocking user %s\n",userName)
			aUser.status.Set(BLOCKED)
			aUser.status.Set(DIRTY)
			return Status_BLOCKED
		}
		fmt.Printf("[%s] Login failed, wrong password %s,%s\n",a.name,userName,pass)
		return Status_NOAUTH
	}
	return Status_SUCCESS
}

func (a *App) presistUser(users []*User) {
	tx, err := a.DATABASE.Begin()
	handleError(err)
	for i := 0 ; i <  len(users) ; i++ {
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
	fmt.Printf("[%s] Pressised %d Users in batch....\n",a.name,len(users))
}



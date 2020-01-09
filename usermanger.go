package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	. "maestro/api"
)

func (a *App) issueToken(durationInSeconds int,userName,device string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer": a.cfg.APP_NAME,
		"Username": userName,
		"device": device,
		"VALID": time.Now().Second() +  durationInSeconds, // vaid for a week
	})
	tokenString, err := token.SignedString([]byte(a.cfg.APP_SECRET))
	if err != nil {
		fmt.Printf("Error: %v", err)
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
				token,err := a.issueToken(7 * 24 * 60 * 60,*env.username,*env.device)
				if err != nil {
					env.resp <- &LoginResp{Status: Status_ERROR}
				} else {
					env.token <- &token
					env.resp <- &LoginResp{Status: Status_SUCCESS}
				}
			} else {
				env.resp <- &LoginResp{Status: status}
			}
		case <-time.After(20 * time.Second):
			fmt.Printf("loginServer: Looking for changes in user database...\n")
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
			} else {
				fmt.Printf("loginServer: No change to user database...\n")
			}
		case env := <-a.registerQ:
			if len(a.users.db) > a.cfg.MAX_NUMBER_OF_USERS {
				env.resp <- &RegisterResp{Status: Status_MAXIMUN_NUMBER_OF_USERS_REACHED}
			} else {
				req := <-env.req
				_, exists := a.users.db[req.UserName]
				if exists {
					fmt.Printf("User %+v exists\n", req)
					env.resp <- &RegisterResp{Status: Status_EXITSTS}
				} else {
					newUser := newUser(req)
					a.users.db[newUser.UserName] = newUser
					token,err := a.issueToken(7 * 24 * 60 * 60,req.GetUserName(),req.GetDevice())
					if err != nil {
						fmt.Printf("Error %+v",err)
						env.resp <-  &RegisterResp{Id: newUser.uid, Status: Status_NOAUTH}
					} else {
						env.token <- &token
						env.resp <- &RegisterResp{Id: newUser.uid, Status: Status_SUCCESS}
					}
				}
			}
		case <-a.quit:
			break
		}

	}

	fmt.Println("LoginServer, Exit processing loop...")
}

func (a *App) tryLogin(userName string, pass []byte) Status {

	aUser, userExists := a.users.db[userName]
	if !userExists {
		return Status_FAIL
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
		fmt.Printf("Login failed, wrong password %s,%s\n",userName,pass)
		return Status_FAIL
	}
	return Status_SUCCESS
}

func (a *App) presistUser(users []*User) {
	tx, err := a.DATABASE.Begin()
	handleError(err)
	fmt.Printf("presistUser: Presisting %d users ",len(users))
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
	for _, u := range users {
		fmt.Printf("Presisted user[%s]\n", u.UserName)
	}
	fmt.Printf("Pressised %d Users in batch....",len(users))
}



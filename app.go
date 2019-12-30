package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"log"
	. "maestro/api"
	"net"
	"sync"
	"time"
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
	loginQ    chan *loginEnvelope
	registerQ chan *registerEnvelope
	UserDbQ   chan []*User
	DATABASE  *sql.DB
}

func newApp() *App {
	cfg := createLoginServerConfig()
	app := App{make(chan bool), make(map[string]Service), cfg, grpc.NewServer(), newUserdb(cfg.ARRAY_PRE_ALLOCATION_LIMIT),
		make(chan *loginEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan *registerEnvelope, cfg.SERVER_QEUEU_LENGTH),
		make(chan []*User, cfg.SERVER_QEUEU_LENGTH),
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
}

func (a *App) start() {
	PrettyPrint(a.cfg)
	for serviceName, s := range a.services {
		switch serviceName {
		case "loginServices":
			fmt.Printf("RPC Registring loginService...\n")
			RegisterLoginServer(a.server, s.(LoginServer))
		case "registerServices":
			fmt.Printf("RPC Registring RegisterServer...\n")
			RegisterRegisterServer(a.server, s.(RegisterServer))
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
	go a.userManager()
	go a.databaseServer()
}

func (a *App) databaseServer() {
	fmt.Println("Database Server, Entering processing loop...")
	for {
		select {
		case users := <-a.UserDbQ:
			a.presistUser(users)
		default:
		}
	}
}

func (a *App) userManager() {
	fmt.Println("LoginServer, Entering processing loop...")
	for {
		select {
		case env := <-a.loginQ:
			req := <-env.req
			env.resp <- &LoginResp{Status: a.tryLogin(req.GetUserName(),
				req.GetPassWord())}
		case <-time.After(3 * time.Second):
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
					a.users.dirty = make([]*User, 5)
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
					env.resp <- &RegisterResp{Id: newUser.uid, Status: Status_SUCCESS}
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
}

/*
type User struct {
	*api.RegisterReq
	sync.RWMutex
	uid string
	status *Flag
	modified time.Time
	loginAttempts int
}

*/

func (a *App) readUsersFromDatabase() {
	fmt.Printf("Cashing user database....\n")
	rows, err := a.DATABASE.Query("SELECT users.uid, users.status,users.UserName,users.Password,users.FirstName," +
		"users.LastName,users.Email,users.Phone,users.Device, address.Zip,address.Street,address.City,address.State FROM users left  join address using(uid)   ")
	handleError(err)
	for rows.Next() {
		u := User{&RegisterReq{}, &sync.RWMutex{}, "", NewFlag(), time.Now(), 0}
		ad := RegisterReq_Address{}
		u.Address = &ad
		var  status int
		err = rows.Scan(&u.uid, &status, &u.UserName, &u.PassWord, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.Device, &ad.Zip, &ad.Street, &ad.City, &ad.State)
		handleError(err)
		u.status.Set(uint(status))
		if !u.status.Is(DELETED) {
			fmt.Printf("Reading user %+v\n",u)
			a.users.db[u.UserName] = &u
		} else {
			fmt.Printf("User %+v is marked deleted skipping...\n", u)
		}
	}
}

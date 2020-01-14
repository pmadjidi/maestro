package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"sync"
	"time"
	. "maestro/api"
)




type usersdb struct {
	db map[string]*User
	dirty []*User
	locked []*User
	dirtyCounter int64
	lockedCounter int64
}

func newUserdb(sliceLimit int) *usersdb {
	return &usersdb{
		make(map[string]*User),
		make([]*User,sliceLimit),
		make([]*User,sliceLimit),
		0,
		0,
	}
}

func newDatabase(cfg *ServerConfig) * sql.DB{

	createUserDb := "CREATE TABLE IF NOT EXISTS users (" +
		"id INTEGER PRIMARY KEY," +
		"uid TEXT," +
		"status INTEGER," +
		"UserName TEXT," +
		"Password BLOB," +
		"FirstName TEXT," +
		"LastName TEXT," +
		"Email TEXT," +
		"Phone TEXT," +
		"stamp NUMERIC," +
		"Device TEXT )"

	createAddressDb := "CREATE TABLE IF NOT EXISTS address (" +
		"id INTEGER PRIMARY KEY," +
		"uid TEXT," +
		"Street TEXT," +
		"City TEXT," +
		"State TEXT," +
		"STATUS  INTEGER," +
		"stamp NUMERIC," +
		"Zip TEXT)"


	createTopicDb := "CREATE TABLE IF NOT EXISTS topics (" +
		"id INTEGER PRIMARY KEY," +
		"userid TEXT," +
		"status INTEGER," +
		"stamp NUMERIC," +
		"topic TEXT)"


	createMessagesDb := "CREATE TABLE IF NOT EXISTS messages (" +
		"id INTEGER PRIMARY KEY," +
		"mid TEXT," +
		"topic TEXT," +
		"Pic BLOB," +
		"parentid TEXT," +
		"status INTEGER," +
		"stamp NUMERIC)"


	fmt.Printf("Storage path is set to [%s]",cfg.STORAGEPATH)
	dbPath := cfg.STORAGEPATH + "db/"
	err := os.MkdirAll(dbPath, os.ModePerm)
	handleError(err)

	dbName := dbPath + cfg.APP_NAME + ".db"

	if cfg.RESETDATABASE_ON_START {
		fmt.Printf("RESETDATABASE_ON_START is set to true, removing old database...\n")
		os.Remove(dbName)
	}

	db, err := sql.Open("sqlite3", dbName)
	handleError(err)
	_, err = db.Exec(createUserDb)
	handleError(err)
	_,err = db.Exec(createAddressDb)
	handleError(err)
	_,err = db.Exec(createTopicDb)
	handleError(err)
	_,err = db.Exec(createMessagesDb)
	handleError(err)

	return db
}


func (a *App) readUsersFromDatabase() {
	fmt.Printf("Cashing user database....\n")
	rows, err := a.DATABASE.Query("SELECT users.uid, users.status,users.UserName,users.Password,users.FirstName," +
		"users.LastName,users.Email,users.Phone,users.Device, address.Zip,address.Street,address.City,address.State FROM users left  join address using(uid)   ")
	handleError(err)
	for rows.Next() {
		u := User{&RegisterReq{}, &sync.RWMutex{}, "", NewFlag(), time.Now(), 0,make([]*Message,0),
		make(map[string]int64)}
		ad := RegisterReq_Address{}
		u.Address = &ad
		var  status int
		err = rows.Scan(&u.uid, &status, &u.UserName, &u.PassWord, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.Device, &ad.Zip, &ad.Street, &ad.City, &ad.State)
		handleError(err)
		u.status.Set(uint(status))
		if !u.status.Is(DELETED) {
			//fmt.Printf("Reading user %+v\n",u)
			a.users.db[u.UserName] = &u
		} else {
			fmt.Printf("User %+v is marked deleted skipping...\n", u)
		}
	}
	fmt.Printf("%d users for %s\n",len(a.users.db),a.cfg.APP_NAME)
}


func (a *App) readMessagesFromDatabase() {
	var messageCounter int
	fmt.Printf("Cashing messaging database....\n")
	rows, err := a.DATABASE.Query("SELECT mid, topic ,Pic,parentid,status,stamp FROM messages")
	handleError(err)
	for rows.Next() {
		m := newMessage()
		var  status  int
		var stamp int64
		err = rows.Scan(&m.Id,&m.Topic,&m.Pic,&m.ParentId,&status,&stamp)
		handleError(err)
		m.Set(uint(status))
		if !m.Is(DELETED) && m.Topic != ""{
			fmt.Printf("Reading message %s",m.Id)
			_,ok := a.messages.msg[m.Topic]
			if !ok {
				a.messages.msg[m.Topic] = make([]*Message,10)
			}
			a.messages.msg[m.Topic]= append(a.messages.msg[m.Topic],m)
			messageCounter++
		} else {
			fmt.Printf("User %+v is marked deleted skipping...\n", m.Id)
		}
	}
	fmt.Printf("%d messages for %s\n",messageCounter,a.cfg.APP_NAME)
}




func (a *App) databaseManager() {
	fmt.Println("Database Server, Entering processing loop...")
	for {
		select {
		case users := <-a.UserDbQ:
			a.presistUser(users)
			case messages := <- a.MsgDBQ:
			a.presistMessage(messages)
		default:
		}
	}
}




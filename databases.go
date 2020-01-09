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

	/*
	type MsgReq struct {
		Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
		Text                 []string             `protobuf:"bytes,2,rep,name=text,proto3" json:"text,omitempty"`
		Pic                  [][]byte             `protobuf:"bytes,3,rep,name=pic,proto3" json:"pic,omitempty"`
		ParentId             string               `protobuf:"bytes,4,opt,name=parentId,proto3" json:"parentId,omitempty"`
		Topic                string               `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
		TimeName             *timestamp.Timestamp `protobuf:"bytes,6,opt,name=time_name,json=timeName,proto3" json:"time_name,omitempty"`
		XXX_NoUnkeyedLiteral struct{}             `json:"-"`
		XXX_unrecognized     []byte               `json:"-"`
		XXX_sizecache        int32                `json:"-"`
	}
	*/

	createMessagesDb := "CREATE TABLE IF NOT EXISTS messages (" +
		"id INTEGER PRIMARY KEY," +
		"mid TEXT," +
		"topic TEXT," +
		"Pic BLOB," +
		"parentid TEXT," +
		"status INTEGER," +
		"stamp NUMERIC)"


	err := os.MkdirAll("./db/", os.ModePerm)
	handleError(err)

	dbName := "./db/" + cfg.APP_NAME + ".db"

	if cfg.RESETDATABASE_ON_START {
		fmt.Printf("RESETDATABASE_ON_START is set to true, removing old database...\n")
		os.Remove(dbName)
	}

	db, _ := sql.Open("sqlite3", dbName)
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
		u := User{&RegisterReq{}, &sync.RWMutex{}, "", NewFlag(), time.Now(), 0,make(chan *Message),
		make([]string,0)}
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
	fmt.Printf("%d users for %s\n",len(a.users.db),a.cfg.APP_NAME)
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




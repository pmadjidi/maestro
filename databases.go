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
		"Device TEXT )"

	createAddressDb := "CREATE TABLE IF NOT EXISTS address (" +
		"id INTEGER PRIMARY KEY," +
		"uid TEXT," +
		"Street TEXT," +
		"City TEXT," +
		"State TEXT," +
		"Zip TEXT)"


	err := os.MkdirAll("./db/", os.ModePerm)
	handleError(err)

	db, _ := sql.Open("sqlite3", "./db/" + cfg.APP_NAME + ".db")
	_, err = db.Exec(createUserDb)
	handleError(err)
	_,err = db.Exec(createAddressDb)
	handleError(err)
	return db
}


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
	fmt.Printf("%d users for %s\n",len(a.users.db),a.cfg.APP_NAME)
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




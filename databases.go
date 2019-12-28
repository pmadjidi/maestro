package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
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

package main

import (
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	SYSTEM_NAME string
	APP_NAME string
	PORT string
	MAX_NUMBER_OF_USERS                int
	MAX_NUMBER_OF_APPS int
	MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT int
	SERVER_QEUEU_LENGTH          int
	NAME_LENGTH_LIMIT int
	MINIMUM_PASSWORD_LENGTH int
	ARRAY_PRE_ALLOCATION_LIMIT int
	MESSAGE_RETENTION_PERIOD int
	SYSTEM_SECRET string
	RESETDATABASE_ON_START bool
	MAX_NUMBER_OF_MESSAGES_PER_TOPIC int
	WRITE_LATENCY time.Duration
	MAX_NUMBER_OF_TOPICS int
	STORAGEPATH string
}



func createServerConfig(name ...string) *ServerConfig {

	var APP_NAME string

	SYSTEM_NAME := os.Getenv("SYSTEM_NAME")
	if SYSTEM_NAME == "" {
		SYSTEM_NAME = "SYSTEM-MAESTRO"
	}




	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":50051"
	}


	SYSTEM_SECRET := os.Getenv("APP_SECRET")
	if SYSTEM_SECRET == "" {
		SYSTEM_SECRET = "ABRAKADABRA"
	}



	MAX_NUMBER_OF_USERS, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_USERS"))
	if err != nil {
		MAX_NUMBER_OF_USERS = 100
	}

	MAX_NUMBER_OF_APPS, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_APPS"))
	if err != nil {
		MAX_NUMBER_OF_APPS = 100
	}

	MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT"))
	if err != nil {
		MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT = 3
	}

	SERVER_QEUEU_LENGTH, err := strconv.Atoi(os.Getenv("SERVER_QEUEU_LENGTH"))
	if err != nil {
		SERVER_QEUEU_LENGTH = 100000
	}

	NAME_LENGTH_LIMIT , err := strconv.Atoi(os.Getenv("NAME_LENGTH_LIMIT"))
	if err != nil {
		NAME_LENGTH_LIMIT = 32
	}

	MINIMUM_PASSWORD_LENGTH , err := strconv.Atoi(os.Getenv("MINIMUM_PASSWORD_LENGTH"))
	if err != nil {
		MINIMUM_PASSWORD_LENGTH = 8
	}

	ARRAY_PRE_ALLOCATION_LIMIT , err := strconv.Atoi(os.Getenv("ARRAY_PRE_ALLOCATION_LIMIT"))
	if err != nil {
		ARRAY_PRE_ALLOCATION_LIMIT	 = 0
	}

	MESSAGE_RETENTION_PERIOD, err := strconv.Atoi(os.Getenv("MESSAGE_RETENTION_PERIOD"))
	if err != nil {
		MESSAGE_RETENTION_PERIOD	 =   60 * 60 * 24 * 30 // 30 days
	}

	RESETDATABASE_ON_START , err := strconv.ParseBool(os.Getenv("RESETDATABASE_ON_START"))
	if err  != nil  {
		RESETDATABASE_ON_START = true
	}

	MAX_NUMBER_OF_MESSAGES_PER_TOPIC, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_MESSAGES_PER_TOPIC"))
	if err != nil {
		MAX_NUMBER_OF_MESSAGES_PER_TOPIC = 10000
	}

	WRITE_LATENCY, err := time.ParseDuration(os.Getenv("WRITE_LATENCY"))
	if err != nil {
		WRITE_LATENCY = 300
	}

	MAX_NUMBER_OF_TOPICS , err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_TOPICS"))
	if err != nil {
		MAX_NUMBER_OF_TOPICS = 10000
	}


	SYSTEM_PATH :=  os.Getenv("SYSTEM_PATH")
	if SYSTEM_PATH == "" {
		SYSTEM_PATH = "./" + SYSTEM_NAME + "/"
	}

	if len(name) != 0  && name[0] != "" {
		APP_NAME = name[0]
	}  else {
		APP_NAME = "TEST"
	}

	STORAGEPATH  :=  os.Getenv("STORAGEPATH")
	if STORAGEPATH == "" {
		STORAGEPATH = SYSTEM_PATH + APP_NAME + "/"
	}









	return &ServerConfig{
		SYSTEM_NAME,
		APP_NAME,
		PORT,
		MAX_NUMBER_OF_USERS,
		MAX_NUMBER_OF_APPS,
		MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT,
		SERVER_QEUEU_LENGTH,
		NAME_LENGTH_LIMIT,
		MINIMUM_PASSWORD_LENGTH,
		ARRAY_PRE_ALLOCATION_LIMIT,
		MESSAGE_RETENTION_PERIOD,
		SYSTEM_SECRET,
		RESETDATABASE_ON_START,
		MAX_NUMBER_OF_MESSAGES_PER_TOPIC,
		WRITE_LATENCY,
		MAX_NUMBER_OF_TOPICS,
		STORAGEPATH,
	}
}


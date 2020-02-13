package main

import (
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	SYSTEM_NAME string
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
	SYSTEM_PATH string
	SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT time.Duration
	SYSTEM_MONITOR_FREQUENCY time.Duration
}

type  AppConfig  struct {
	*ServerConfig
	APP_NAME string
	STORAGEPATH string
}


func createAppConfig(config *ServerConfig,appName ...string) *AppConfig {

	var APP_NAME string

	if len(appName) != 0  && appName[0] != "" {
		APP_NAME = appName[0]
	}  else {
		APP_NAME = "TEST"
	}

	STORAGEPATH  :=  os.Getenv("STORAGEPATH")
	if STORAGEPATH == "" {
		STORAGEPATH = config.SYSTEM_PATH + APP_NAME + "/"
	}


	return &AppConfig{
		config,
		APP_NAME,
		STORAGEPATH,
	}
}




func createServerConfig(systemName ...string) *ServerConfig {

	var SYSTEM_NAME string


	if len(systemName) != 0  && systemName[0] != "" {
		SYSTEM_NAME = systemName[0]
	}  else if  systemNameFromEenv := os.Getenv("SYSTEM_NAME")  ; SYSTEM_NAME != "" {
		SYSTEM_NAME = systemNameFromEenv
	} else {
		SYSTEM_NAME = "SYSTEM-MAESTRO"
		os.Setenv("SYSTEM_NAME",SYSTEM_NAME)
	}


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":50051"
		os.Setenv("PORT",":50051")
	}


	SYSTEM_SECRET := os.Getenv("APP_SECRET")
	if SYSTEM_SECRET == "" {
		SYSTEM_SECRET = "ABRAKADABRA"
		os.Setenv("APP_SECRET","ABRAKADABRA")
	}



	MAX_NUMBER_OF_USERS, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_USERS"))
	if err != nil {
		MAX_NUMBER_OF_USERS = 2000
		os.Setenv("MAX_NUMBER_OF_USERS","100")
	}

	MAX_NUMBER_OF_APPS, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_APPS"))
	if err != nil {
		MAX_NUMBER_OF_APPS = 500
		os.Setenv("MAX_NUMBER_OF_APPS","100")
	}

	MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT"))
	if err != nil {
		MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT = 3
		os.Setenv("MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT","3")
	}

	SERVER_QEUEU_LENGTH, err := strconv.Atoi(os.Getenv("SERVER_QEUEU_LENGTH"))
	if err != nil {
		SERVER_QEUEU_LENGTH = 100000
		os.Setenv("SERVER_QEUEU_LENGTH","100000")
	}

	NAME_LENGTH_LIMIT , err := strconv.Atoi(os.Getenv("NAME_LENGTH_LIMIT"))
	if err != nil {
		NAME_LENGTH_LIMIT = 32
		os.Setenv("NAME_LENGTH_LIMIT","32")
	}

	MINIMUM_PASSWORD_LENGTH , err := strconv.Atoi(os.Getenv("MINIMUM_PASSWORD_LENGTH"))
	if err != nil {
		MINIMUM_PASSWORD_LENGTH = 8
		os.Setenv("MINIMUM_PASSWORD_LENGTH","8")
	}

	ARRAY_PRE_ALLOCATION_LIMIT , err := strconv.Atoi(os.Getenv("ARRAY_PRE_ALLOCATION_LIMIT"))
	if err != nil {
		ARRAY_PRE_ALLOCATION_LIMIT	 = 0
		os.Setenv("ARRAY_PRE_ALLOCATION_LIMIT","0")
	}

	MESSAGE_RETENTION_PERIOD, err := strconv.Atoi(os.Getenv("MESSAGE_RETENTION_PERIOD"))
	if err != nil {
		MESSAGE_RETENTION_PERIOD	 =   60 * 60 * 24 * 30 // 30 days = 2592000 seconds
		os.Setenv("MESSAGE_RETENTION_PERIOD","2592000")
	}

	RESETDATABASE_ON_START , err := strconv.ParseBool(os.Getenv("RESETDATABASE_ON_START"))
	if err  != nil  {
		RESETDATABASE_ON_START = false
		os.Setenv("RESETDATABASE_ON_START","false")
	}

	MAX_NUMBER_OF_MESSAGES_PER_TOPIC, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_MESSAGES_PER_TOPIC"))
	if err != nil {
		MAX_NUMBER_OF_MESSAGES_PER_TOPIC = 10000
		os.Setenv("MAX_NUMBER_OF_MESSAGES_PER_TOPIC","10000")
	}

	WRITE_LATENCY, err := time.ParseDuration(os.Getenv("WRITE_LATENCY"))
	if err != nil {
		WRITE_LATENCY = 300
		os.Setenv("WRITE_LATENCY","300")
	}

	MAX_NUMBER_OF_TOPICS , err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_TOPICS"))
	if err != nil {
		MAX_NUMBER_OF_TOPICS = 10000
		os.Setenv("MAX_NUMBER_OF_TOPICS","10000")
	}


	SYSTEM_PATH :=  os.Getenv("SYSTEM_PATH")
	if SYSTEM_PATH == "" {
		SYSTEM_PATH = "./" + SYSTEM_NAME + "/"
		os.Setenv("SYSTEM_PATH",SYSTEM_PATH)
	}


	SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT, err := time.ParseDuration(os.Getenv("SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT"))
	if err != nil {
		SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT = 3
		os.Setenv("SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT","3")
	}


	SYSTEM_MONITOR_FREQUENCY, err := time.ParseDuration(os.Getenv("SYSTEM_MONITOR_FREQUENCY"))
	if err != nil {
		SYSTEM_MONITOR_FREQUENCY = 30
		os.Setenv("SYSTEM_MONITOR_FREQUENCY","30")
	}



	return &ServerConfig{
		SYSTEM_NAME,
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
		SYSTEM_PATH,
		SYSTEM_QUEUE_WAIT_BEFORE_TIME_OUT,
		SYSTEM_MONITOR_FREQUENCY,
	}
}


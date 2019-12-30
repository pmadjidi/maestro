package main

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	PORT string
	APP_NAME string
	MAX_NUMBER_OF_USERS                int
	MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT int
	SERVER_QEUEU_LENGTH          int
	NAME_LENGTH_LIMIT int
	MINIMUM_PASSWORD_LENGTH int
	ARRAY_PRE_ALLOCATION_LIMIT int
	MESSAGE_RETENTION_PERIOD int
}



func createLoginServerConfig() *ServerConfig {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":50051"
	}

	APP_NAME := os.Getenv("APP_NAME")
	if APP_NAME == "" {
		APP_NAME = "TEST"
	}


	MAX_NUMBER_OF_USERS, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_USERS"))
	if err != nil {
		MAX_NUMBER_OF_USERS = 100
	}
	MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT, err := strconv.Atoi(os.Getenv("MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT"))
	if err != nil {
		MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT = 3
	}

	SERVER_QEUEU_LENGTH, err := strconv.Atoi(os.Getenv("SERVER_QEUEU_LENGTH"))
	if err != nil {
		SERVER_QEUEU_LENGTH = 0
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






	return &ServerConfig{
		PORT,
		APP_NAME,
		MAX_NUMBER_OF_USERS,
		MAX_NUMBER_OF_FAILED_LOGIN_ATTEMPT,
		SERVER_QEUEU_LENGTH,
		NAME_LENGTH_LIMIT,
		MINIMUM_PASSWORD_LENGTH,
		ARRAY_PRE_ALLOCATION_LIMIT,
		MESSAGE_RETENTION_PERIOD,

	}
}


package utils

import (
	"log"
	"os"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func InitLogger() {
	var err error
	LogFile, err = os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	Logger = log.New(LogFile, "", log.LstdFlags)
}

func CloseLogger() {
	LogFile.Close()
}

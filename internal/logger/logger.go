package logger

import (
	"fmt"
	"log"
)

var (
	loggerErr *log.Logger
)

func Init() {
	loggerErr = log.Default()
	loggerErr.SetPrefix("[Error]-")
}

func Error(at string, err error, data interface{}) {
	if loggerErr == nil {
		fmt.Println("Logger was not created")
		panic(err)
	}
	loggerErr.Println(fmt.Sprintf("at: %s, message: %s, data: %v", at, err.Error(), data))
}

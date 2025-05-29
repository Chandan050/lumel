package utils

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	f, _ := os.OpenFile("logs/refresh.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return log.New(f, "REFRESH: ", log.LstdFlags)
}

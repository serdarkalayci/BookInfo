package logger

import (
	"log"
	"os"
	"time"
)

func Log(message string, level Level, err ...error) {
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Printf("[%s] [%s] %s\nError:[%s]", level, time.Now(), message, err)
	} else {
		log.SetOutput(os.Stdout)
		log.Printf("[%s] [%s] %s", level, time.Now(), message)
	}
}

package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once sync.Once
	logger *log.Logger
)

func Logger() *log.Logger {
	once.Do(func() {
		logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	})
	return logger
}

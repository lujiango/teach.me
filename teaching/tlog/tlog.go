package tlog

import (
	"fmt"
	"log"
	"os"
)

var logName string = "tlog/teaching.log"
var isInit bool = false
var logger *log.Logger

func initLog(prefix string) {
	if !isInit {
		logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
		if err != nil {
			log.Fatalln("open log file error !", err)
		}
		logger = log.New(logFile, prefix, log.LstdFlags|log.Lshortfile)
		if logger != nil {
			isInit = true
		}
	} else {
		logger.SetPrefix(prefix)
	}
}
func Info(v ...interface{}) {
	initLog("[INFO]")
	logger.Output(2, fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	initLog("[DEBUG]")
	logger.Output(2, fmt.Sprintln(v...))
}
func Error(v ...interface{}) {
	initLog("[ERROR]")
	logger.Output(2, fmt.Sprintln(v...))
}
func Fatal(v ...interface{}) {
	initLog("[FATAL]")
	logger.Output(2, fmt.Sprintln(v...))
}

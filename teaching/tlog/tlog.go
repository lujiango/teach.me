package tlog

import (
	"log"
	"os"
)

var logName string = "teaching.log"
var isInit bool = false
var logger *log.Logger

func initLog(prefix string) {
	if !isInit {
		logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatalln("open log file error !")
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
	logger.Println(v)
}

func Debug(v ...interface{}) {
	initLog("[DEBUG]")
	logger.Println(v)

}
func Error(v ...interface{}) {
	initLog("[ERROR]")
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	initLog("[FATAL]")
	logger.Fatalln(v)
}

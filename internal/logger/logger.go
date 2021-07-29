package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

func InitLog() *logrus.Logger {

	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	Logger.SetReportCaller(true)
	setFileForLog()

	return Logger
}

func setFileForLog() {
	timeNow := time.Now()

	year, month, day := timeNow.Date()
	nameLogFile := fmt.Sprintf("logFrom_%d_%d_%d_%d.Logger", day, month, year, timeNow.Nanosecond())
	err := os.MkdirAll("./log/", 0666)

	if err != nil {
		Logger.Warn("Failed to Logger to file. I cannot make dir ")
		return
	}


	newFile, err := os.OpenFile("./log/"+nameLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Warn("Failed to Logger to file, using default stderr ")
		return
	}

	Logger.Out = newFile
}


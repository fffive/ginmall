package utils

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := SetOutPutFile()
	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func SetOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""

	// 创建文件
	if wd, err := os.Getwd(); err == nil {
		logFilePath = wd + "/logs/"
	}

	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath + logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fileName, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	// 写文件
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return file, nil
}

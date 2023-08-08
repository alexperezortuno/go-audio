package commons

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func GetLogger() *logrus.Logger {
	logFile := "log.txt"
	var f *os.File
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	var standardLogger = logrus.New()

	standardLogger.SetLevel(logrus.DebugLevel)
	standardLogger.SetFormatter(&logrus.JSONFormatter{})
	//standardLogger.SetFormatter(&easy.Formatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	LogFormat:       "%time% [%lvl%] line:%line% - %msg%",
	//})
	standardLogger.SetOutput(io.MultiWriter(f, os.Stdout))

	return standardLogger
}

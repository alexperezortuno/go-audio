package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func GetLogger(c *cli.Context) *logrus.Logger {
	logFile := "log.txt"
	var f *os.File
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	var standardLogger = logrus.New()

	if c.Bool("debug") {
		standardLogger.SetLevel(logrus.DebugLevel)
	} else {
		standardLogger.SetLevel(logrus.InfoLevel)
	}
	if c.String("log_format") == "json" {
		standardLogger.SetFormatter(&logrus.JSONFormatter{})
	} else if c.String("log_format") == "txt" {
		standardLogger.SetFormatter(&logrus.TextFormatter{})
	}

	standardLogger.SetOutput(io.MultiWriter(f, os.Stdout))

	return standardLogger
}

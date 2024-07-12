package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var logger *logrus.Logger
var loggerEntry *logrus.Entry

func NewLogger() {
	logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: true},
		Level:     logrus.InfoLevel,
	}
	buildTag := "develop"
	loggerEntry = logrus.NewEntry(logger)
	loggerEntry = loggerEntry.WithField("build", buildTag)
}

func Error(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Error(args...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

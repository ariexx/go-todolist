package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerInterface[T any] interface {
	SaveLog(level logrus.Level, msg T)
}

func NewLoggerService() LoggerInterface[any] {
	return &Logger{}
}

type Logger struct {
}

func (log *Logger) SaveLog(level logrus.Level, msg any) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Log(level, msg)
}

func (log *Logger) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (log *Logger) Fire(entry *logrus.Entry) error {
	fmt.Println(entry.Level, entry.Message)
	return nil
}

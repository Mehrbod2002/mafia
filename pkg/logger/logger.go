package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Init(level string) {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	l, _ := logrus.ParseLevel(level)
	logrus.SetLevel(l)
}

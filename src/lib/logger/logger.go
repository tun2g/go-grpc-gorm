package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(context string) *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
		PadLevelText:  true,
	})
	logger.SetLevel(logrus.DebugLevel)
	return logger.WithField("context", context)
}
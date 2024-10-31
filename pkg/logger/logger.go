package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New(logLevel string) *logrus.Logger {
    log := logrus.New()
    log.Out = os.Stdout

    level, err := logrus.ParseLevel(logLevel)
    if err != nil {
        log.Warnf("Invalid log level '%s', defaulting to 'info'", logLevel)
        level = logrus.InfoLevel
    }
    log.SetLevel(level)

    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

    return log
}

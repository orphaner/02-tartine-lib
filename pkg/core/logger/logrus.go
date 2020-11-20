package logger

import (
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	defaultLogLevel = "info"
	flagLoglevel    string
	Log             *logrus.Entry
)

func init() {
	flag.StringVar(&flagLoglevel, "log-level", defaultLogLevel, "log level value: info, debug")
}

func InitLogger() {
	Log = createLogger()
}

func createLogger() *logrus.Entry {
	logger := logrus.New()
	logger.Level = getLogLevel()
	return logrus.NewEntry(logger)
}

func getLogLevel() logrus.Level {
	logLevel, err := logrus.ParseLevel(flagLoglevel)
	if err != nil {
		logrus.
			WithError(err).
			Errorf("log level parse failed. %s is set as default.", defaultLogLevel)
		return logrus.InfoLevel
	}
	return logLevel
}

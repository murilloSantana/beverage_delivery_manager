package log

import "github.com/sirupsen/logrus"

//go:generate mockery --name Logger --case=underscore --output ../../mocks

type LoggerFields map[string]interface{}

// Encapsulates the way logs are made, decoupling from external libs
type Logger interface {
	Info(values map[string]interface{}, msg interface{})
	Error(values map[string]interface{}, msg interface{})
}

type loggerWrap struct {
	log *logrus.Logger
}

func NewLogger() Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return loggerWrap{
		log: l,
	}
}

func (l loggerWrap) Info(values map[string]interface{}, msg interface{}) {
	l.log.WithFields(values).Info(msg)
}

func (l loggerWrap) Error(values map[string]interface{}, msg interface{}) {
	l.log.WithFields(values).Error(msg)
}

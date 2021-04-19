package log

import "github.com/sirupsen/logrus"

//go:generate mockery --name Logger --case=underscore --output ../../mocks

type LoggerFields map[string]interface{}

// Encapsulates the way logs are made, decoupling from external libs
type Logger interface {
	Info(values map[string]interface{}, msg interface{})
	Error(values map[string]interface{}, msg interface{})
}

type loggerWrap struct{}

func NewLogger() Logger {
	return loggerWrap{}
}

func (l loggerWrap) Info(values map[string]interface{}, msg interface{}) {
	logrus.WithFields(values).Info(msg)
}

func (l loggerWrap) Error(values map[string]interface{}, msg interface{}) {
	logrus.WithFields(values).Error(msg)
}

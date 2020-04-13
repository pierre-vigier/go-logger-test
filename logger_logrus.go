package main

import (
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger logrus.FieldLogger
}

// NewFromLogrus return a zap powered logger satisfying Logger interface
func NewFromLogrus(logger logrus.FieldLogger) Logger {
	return logrusLogger{logger: logger}
}

func (l logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l logrusLogger) Debugf(str string, args ...interface{}) {
	l.logger.Debugf(str, args...)
}

func (l logrusLogger) With(fields ...interface{}) Logger {
	if len(fields) == 0 {
		return logrusLogger{logger: l.logger}
	}

	f := logrus.Fields{}
	for i := 0; i < len(fields); {
		if i == len(fields)-1 {
			// last element, odd number...
			f["loggerError"] = "Value missing, Odd number of elements"
			break
		}
		key, val := fields[i], fields[i+1]
		if keyStr, ok := key.(string); !ok {
			f["loggerError"] = "Issue in list of fields"
			break
		} else {
			f[keyStr] = val
		}
		i += 2
	}
	return logrusLogger{logger: l.logger.WithFields(f)}
}

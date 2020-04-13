package main

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

// NewFromZap return a zap powered logger satisfying Logger interface
func NewFromZap(logger *zap.Logger) Logger {
	return zapLogger{logger: logger.Sugar()}
}

func (z zapLogger) Debug(args ...interface{}) {
	z.logger.Debug(args...)
}

func (z zapLogger) Debugf(str string, args ...interface{}) {
	z.logger.Debugf(str, args...)
}

func (z zapLogger) With(fields ...interface{}) Logger {
	return zapLogger{logger: z.logger.With(fields...)}
}

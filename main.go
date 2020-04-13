package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	lz := zap.NewExample()
	logzap := NewFromZap(lz)
	logzap.Debug("With zap, Debug")
	logzap.Debugf("With %s, Debugf", "zap")
	logzap.With("logger", "zap", "Answer", 42).Debug("Zap and fields")

	lr := logrus.New()
	lr.Level = logrus.TraceLevel
	loglogrus := NewFromLogrus(lr)
	loglogrus.Debug("With Logrus, Debug")
	loglogrus.Debugf("With %s, Debugf", "logrus")
	loglogrus.With("logger", "logrus", "Answer", 42).Debug("Logrus and fields")
}

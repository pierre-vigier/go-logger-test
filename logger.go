package main

// Logger is the interface used to manage log internally
type Logger interface {
	Debug(args ...interface{})
	Debugf(str string, args ...interface{})
	With(fields ...interface{}) Logger
}

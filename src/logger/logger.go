package logger

import (
	"fmt"
	"log"

	"github.com/getsentry/raven-go"
)

type Logger struct {
	verbose, initialized bool
	tags                 map[string]string
}

var defaultLogger = New(true, nil)

func New(verbose bool, tags map[string]string) *Logger {
	return &Logger{
		tags:    tags,
		verbose: verbose,
	}
}

func (l *Logger) Info(message string) {
	Info(message)
}

func (l *Logger) Infof(message string) {
	Infof(message)
}

func (l *Logger) Error(err error) {
	Error(err)
}

func (l *Logger) Fatal(err error) {
	Fatal(err)
}

func Info(message string) {
	if defaultLogger.verbose {
		log.Println(message)
	}

	//go raven.CaptureMessage(message, l.tags)
}

func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}

func Error(err error) {
	if defaultLogger.verbose {
		log.Println(err.Error())
	}

	go raven.CaptureError(err, defaultLogger.tags)
}

func Fatal(err error) {
	raven.CaptureErrorAndWait(err, nil)

	if defaultLogger.verbose {
		log.Fatal(err)
	}
}

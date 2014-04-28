package log

import (
	"time"
	"fmt"
	"os"
	"errors"
)

type LogLevel int

func (this LogLevel) String() string {
	switch this {
		case FATAL: return "FATAL"
		case CRITICAL: return "CRITICAL"
		case ERROR: return "ERROR"
		case WARNING: return "WARNING"
		case NOTICE: return "NOTICE"
		case INFO: return "INFO"
		case DEBUG: return "DEBUG"
	}
	return "unknown"
}

const (
	FATAL LogLevel = iota
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

const timeFormat = "2006-01-02 15:04:05"

var globalLogLevel LogLevel = DEBUG

func SetLevel(logLevel LogLevel) {
	globalLogLevel = logLevel
}

func GetLevel() LogLevel {
	return globalLogLevel
}

func logFormattedEntry(logLevel LogLevel, message string, args ...interface{}) string {
	if logLevel > globalLogLevel {
		return ""
	} 
	entryString := fmt.Sprintf("%s %s %s", time.Now().Format(timeFormat), logLevel, fmt.Sprintf(message, args...))
	fmt.Fprintln(os.Stderr, entryString)
	return entryString
}

func logEntry(logLevel LogLevel, message string, args ...interface{}) string {
	entryString := message
	for _, s := range args {
		entryString += fmt.Sprintf(" %s", s)
	}
	return logFormattedEntry(logLevel, entryString)
}

func logErrorEntry(logLevel LogLevel, err error) error {
	if err == nil {
		// No error
		return nil;
	}
	entryString := fmt.Sprintf("%+v", err)
	logEntry(logLevel, entryString)
	return err	
}

func Debug(message string, args ...interface{}) string {
	return logEntry(DEBUG, message, args...)
}

func Debugf(message string, args ...interface{}) string {
	return logFormattedEntry(DEBUG, message, args...)
}

func Info(message string, args ...interface{}) string {
	return logEntry(INFO, message, args...)
}

func Infof(message string, args ...interface{}) string {
	return logFormattedEntry(INFO, message, args...)
}

func Notice(message string, args ...interface{}) string {
	return logEntry(NOTICE, message, args...)
}

func Noticef(message string, args ...interface{}) string {
	return logFormattedEntry(NOTICE, message, args...)
}

func Warning(message string, args ...interface{}) error {
	return errors.New(logEntry(WARNING, message, args...))
}

func Warningf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(WARNING, message, args...))
}

func Error(message string, args ...interface{}) error {
	return errors.New(logEntry(ERROR, message, args...))
}

func Errorf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(ERROR, message, args...))
}

func Errore(err error) error {
	return logErrorEntry(ERROR, err)
}

func Critical(message string, args ...interface{}) error {
	return errors.New(logEntry(CRITICAL, message, args...))
}

func Criticalf(message string, args ...interface{}) error {
	return errors.New(logFormattedEntry(CRITICAL, message, args...))
}

func Criticale(err error) error {
	return logErrorEntry(CRITICAL, err)
}

func Fatal(message string, args ...interface{}) error {
	logEntry(FATAL, message, args...)
	os.Exit(1)
	return errors.New(logEntry(CRITICAL, message, args...))
}

func Fatalf(message string, args ...interface{}) error {
	logEntry(FATAL, message, args...)
	os.Exit(1)
	return errors.New(logFormattedEntry(CRITICAL, message, args...))
}

func Fatale(err error) error {
	logErrorEntry(FATAL, err)
	os.Exit(1)
	return err
}


// This is log wrapper to enable log level:
// use SetLevel(loglevel) to control display log, all level below this setting will be ignored.
package slog

import (
	"fmt"
	"io"
	"log"
	"os"
)

// define log level
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

var loggers [OFF]*log.Logger // Store each log level logger

// Default logger output
var logouts = [...]io.Writer{
	DEBUG: os.Stdout,
	INFO:  os.Stdout,
	WARN:  os.Stdout,
	ERROR: os.Stdout,
	FATAL: os.Stderr,
}

// Default logger prefix
var logpref = [...]string{
	DEBUG: "[DEBUG]:",
	INFO:  "[INFO]:",
	WARN:  "[WARN]:",
	ERROR: "[ERROR]:",
	FATAL: "[FATAL]:",
}

// Default logger flag
var logflags = [...]int{
	DEBUG: log.LstdFlags | log.Lshortfile,
	INFO:  log.LstdFlags,
	WARN:  log.LstdFlags | log.Lshortfile,
	ERROR: log.LstdFlags | log.Lshortfile,
	FATAL: log.LstdFlags | log.Lshortfile,
}

// Default logger lever to display all
var logLevel = DEBUG

func init() {
	// init all logger based on default settting
	for idx := DEBUG; idx < OFF; idx++ {
		loggers[idx] = log.New(logouts[idx], logpref[idx], logflags[idx])
	}
}

// set output log level. If set to WARN, DEBUG and INFO logs will be ignored
func SetLevel(l int) error {
	var err error = nil

	if l < DEBUG || l > OFF {
		err = fmt.Errorf("Error Log Level:%d", l)
	} else {
		logLevel = l
	}
	return err
}

func Debug(v ...interface{}) {
	if DEBUG >= logLevel {
		loggers[DEBUG].Output(2, fmt.Sprint(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if DEBUG >= logLevel {
		loggers[DEBUG].Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if INFO >= logLevel {
		loggers[INFO].Output(2, fmt.Sprint(v...))
	}
}

func Infof(format string, v ...interface{}) {
	if INFO >= logLevel {
		loggers[INFO].Output(2, fmt.Sprintf(format, v...))
	}
}

func Warn(v ...interface{}) {
	if WARN >= logLevel {
		loggers[WARN].Output(2, fmt.Sprint(v...))
	}
}

func Warnf(format string, v ...interface{}) {
	if WARN >= logLevel {
		loggers[WARN].Output(2, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	if ERROR >= logLevel {
		loggers[ERROR].Output(2, fmt.Sprint(v...))
	}
}

func Errorf(format string, v ...interface{}) {
	if ERROR >= logLevel {
		loggers[ERROR].Output(2, fmt.Sprintf(format, v...))
	}
}

func Fatal(v ...interface{}) {
	if FATAL >= logLevel {
		loggers[FATAL].Output(2, fmt.Sprint(v...))
		panic(fmt.Sprint(v...))
	}
}

func Fatalf(format string, v ...interface{}) {
	if FATAL >= logLevel {
		loggers[FATAL].Output(2, fmt.Sprintf(format, v...))
		panic(fmt.Sprintf(format, v...))
	}
}

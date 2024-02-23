package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents the logging level
type LogLevel int

const (
	// DEBUG level logs everything
	DEBUG LogLevel = iota
	// INFO level logs informational messages
	INFO
	// WARN level logs warnings
	WARN
	// ERROR level logs errors
	ERROR
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string, err error)
}

// Logger is a custom logger
type Log struct {
	level       LogLevel
	serviceName string
}

// New creates a new Logger with the specified level
func New(level LogLevel, serviceName string) *Log {
	return &Log{level: level, serviceName: serviceName}
}

// Debug logs a debug message
func (l *Log) Debug(msg string) {
	if l.level <= DEBUG {
		l.log(DEBUG, msg, l.serviceName)
	}
}

// Info logs an informational message
func (l *Log) Info(msg string) {
	if l.level <= INFO {
		l.log(INFO, msg, l.serviceName)
	}
}

// Warn logs a warning message
func (l *Log) Warn(msg string) {
	if l.level <= WARN {
		l.log(WARN, msg, l.serviceName)
	}
}

// Error logs an error message
func (l *Log) Error(msg string, err error) {
	l.logFatal(ERROR, fmt.Sprintf("%s: %v", msg, err), l.serviceName)
}

// log formats and prints a log message
func (l *Log) log(level LogLevel, msg string, serviceName string) {
	// Get filename and line number
	_, file, line, _ := runtime.Caller(2)
	file = strings.TrimPrefix(file, "/go/src/") // Optional: Remove "/go/src/" prefix

	// Format log message
	var formattedMessage string = ""

	if level == INFO {
		formattedMessage = fmt.Sprintf("[%s][%s][%s][%s][%s]",
			levelToString(level), time.Now().Format("2006-01-02 15:04:05"), serviceName, file, msg)
	} else {
		formattedMessage = fmt.Sprintf("[%s][%s][%s][%s][%d][%s]",
			levelToString(level), time.Now().Format("2006-01-02 15:04:05"), serviceName, file, line, msg)
	}
	// Print log message
	fmt.Println(formattedMessage)
}

// log formats and prints a log message
func (l *Log) logFatal(level LogLevel, msg string, serviceName string) {
	// Get filename and line number
	_, file, line, _ := runtime.Caller(2)
	file = strings.TrimPrefix(file, "/go/src/") // Optional: Remove "/go/src/" prefix

	// Format log message
	formattedMessage := fmt.Sprintf("[%s][%s][%s][%s][%d][%s]",
		levelToString(level), time.Now().Format("2006-01-02 15:04:05"), serviceName, file, line, msg)

	// Print log message
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.Fatal(formattedMessage)
	//fmt.Println(formattedMessage)
}

// levelToString converts LogLevel to string
func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

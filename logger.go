package logger

import (
	"io"
	"log"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var logNames = map[int]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type Logger struct {
	level  int
	logger *log.Logger
}

func NewLogger(level int, output io.Writer) *Logger {

	return &Logger{
		level:  level,
		logger: log.New(output, "", 0),
	}
}

func (l *Logger) log(level int, msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	l.logger.Printf("%s [%s] %s\n", timestamp, logNames[level], msg)
}

func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg)
}
func (l *Logger) Info(msg string) {
	l.log(INFO, msg)
}
func (l *Logger) Warn(msg string) {
	l.log(WARN, msg)
}
func (l *Logger) Error(msg string) {
	l.log(ERROR, msg)
}

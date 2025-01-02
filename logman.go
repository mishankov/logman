package logman

import (
	"io"
	"os"
)

type LogLevel string

const (
	Debug = LogLevel("Debug")
	Info  = LogLevel("Info")
	Warn  = LogLevel("Warn")
	Error = LogLevel("Error")
	Fatal = LogLevel("Fatal")
)

type TimeProvider interface {
	Time() string
}

type Formatter interface {
	Format(logLevel LogLevel, dateTime string, message string) string
}

type Logger struct {
	Writer    io.Writer
	Timer     TimeProvider
	Formatter Formatter
}

func NewLogger(output io.Writer, timer TimeProvider, formatter Formatter) *Logger {
	return &Logger{Writer: output, Timer: timer, Formatter: formatter}
}

func NewDefaultLogger() *Logger {
	return &Logger{
		Writer:    os.Stdout,
		Timer:     NewDefaultTimeProvider(DefaultTimeFormat),
		Formatter: NewDefaultFormatter(DefaultFormat),
	}
}

func (l *Logger) Debug(message string) {
	l.Writer.Write([]byte(l.Formatter.Format(Debug, l.Timer.Time(), message)))
}

func (l *Logger) Info(message string) {
	l.Writer.Write([]byte(l.Formatter.Format(Info, l.Timer.Time(), message)))
}

func (l *Logger) Warn(message string) {
	l.Writer.Write([]byte(l.Formatter.Format(Warn, l.Timer.Time(), message)))
}

func (l *Logger) Error(message string) {
	l.Writer.Write([]byte(l.Formatter.Format(Error, l.Timer.Time(), message)))
}

func (l *Logger) Fatal(message string) {
	l.Writer.Write([]byte(l.Formatter.Format(Fatal, l.Timer.Time(), message)))
}

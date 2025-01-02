package logman

import (
	"fmt"
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

func (l *Logger) Log(logLevel LogLevel, message ...any) {
	l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.Timer.Time(), string(fmt.Appendln([]byte{}, message...)))))
}

func (l *Logger) Logf(logLevel LogLevel, message string, formats ...any) {
	l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.Timer.Time(), fmt.Sprintf(message, formats...)+"\n")))
}

func (l *Logger) Debug(message ...any) {
	l.Log(Debug, message...)
}

func (l *Logger) Debugf(message string, formats ...any) {
	l.Logf(Debug, message, formats...)
}

func (l *Logger) Info(message ...any) {
	l.Log(Info, message...)
}

func (l *Logger) Infof(message string, formats ...any) {
	l.Logf(Info, message, formats...)
}

func (l *Logger) Warn(message ...any) {
	l.Log(Warn, message...)
}

func (l *Logger) Warnf(message string, formats ...any) {
	l.Logf(Warn, message, formats...)
}

func (l *Logger) Error(message ...any) {
	l.Log(Error, message...)
}

func (l *Logger) Errorf(message string, formats ...any) {
	l.Logf(Error, message, formats...)
}

func (l *Logger) Fatal(message ...any) {
	l.Log(Fatal, message...)
}

func (l *Logger) Fatalf(message string, formats ...any) {
	l.Logf(Fatal, message, formats...)
}

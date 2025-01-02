package logman

import (
	"io"
	"os"
)

type TimeProvider interface {
	Time() string
}

type Formatter interface {
	Format(logLevel string, dateTime string, message string) string
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
	l.Writer.Write([]byte(l.Formatter.Format("DEBUG", l.Timer.Time(), message)))
}

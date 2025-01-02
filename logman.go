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
	writer    io.Writer
	timer     TimeProvider
	formatter Formatter
}

func NewLogger(output io.Writer, timer TimeProvider, formatter Formatter) *Logger {
	return &Logger{writer: output, timer: timer, formatter: formatter}
}

func NewDefaultLogger() *Logger {
	return &Logger{
		writer:    os.Stdout,
		timer:     &DefaultTimeProvider{},
		formatter: NewDefaultFormatter(defaultFormat),
	}
}

func (l *Logger) Debug(message string) {
	l.writer.Write([]byte(l.formatter.Format("DEBUG", l.timer.Time(), message)))
}

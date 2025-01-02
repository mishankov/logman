package logman

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	output io.Writer
	timer  TimeProvider
}

type TimeProvider interface {
	Time() string
}

type DefaultTimeProvider struct{}

func (dt *DefaultTimeProvider) Time() string {
	return time.Now().Format("2006-01-02 15:04:05 GMT-0700")
}

func NewLogger(output io.Writer, timer TimeProvider) *Logger {
	return &Logger{output: output, timer: timer}
}

func NewDefaultLogger() *Logger {
	return &Logger{output: os.Stdout, timer: &DefaultTimeProvider{}}
}

func (l *Logger) Debug(message string) {
	l.output.Write([]byte(fmt.Sprintf("[%v] [DEBUG] - %v", l.timer.Time(), message)))
}

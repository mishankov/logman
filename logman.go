package logman

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"time"
)

type LogLevel int

const (
	Debug = LogLevel(0)
	Info  = LogLevel(1)
	Warn  = LogLevel(2)
	Error = LogLevel(3)
	Fatal = LogLevel(4)
)

func (ll LogLevel) String() string {
	switch ll {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return strconv.Itoa(int(ll))
	}
}

type Formatter interface {
	Format(logLevel LogLevel, dateTime time.Time, callLocation string, message string) string
}

type Filter interface {
	Filter(logLevel LogLevel, callLocation string, message string) bool
}

func callLocation() string {
	const skip = 5
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	_, _, line, _ := runtime.Caller(skip - 1)

	return fmt.Sprintf("%v:%v", frame.Function, line)
}

type Logger struct {
	Writer    io.Writer
	Formatter Formatter
	Filter    Filter
}

func NewLogger(output io.Writer, formatter Formatter, filter Filter) *Logger {
	return &Logger{Writer: output, Formatter: formatter, Filter: filter}
}

func (l *Logger) log(logLevel LogLevel, message string) {
	cl := callLocation()

	if l.Filter == nil || l.Filter.Filter(logLevel, cl, message) {
		//TODO-docs: Here errors are not meant to be handled. It should be the concern of Logger.Writer
		l.Writer.Write([]byte(l.Formatter.Format(logLevel, time.Now(), cl, message) + "\n"))
	}
}

func (l *Logger) Log(logLevel LogLevel, message ...any) {
	m := string(fmt.Appendln([]byte{}, message...))
	// Remove new line at the end of a message. Adds it later at the end of the formatted line
	m = m[:len(m)-1]

	l.log(logLevel, m)
}

func (l *Logger) Logf(logLevel LogLevel, message string, formats ...any) {
	m := fmt.Sprintf(message, formats...)
	l.log(logLevel, m)
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

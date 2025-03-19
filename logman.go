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
	Format(logLevel LogLevel, dateTime time.Time, callLocation string, message string, params ...any) string
}

type Filter interface {
	Filter(logLevel LogLevel, callLocation string, message string) bool
}

func callLocation(skipCorrection int) string {
	skip := 4 + skipCorrection
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	_, _, line, _ := runtime.Caller(skip - 1)

	return fmt.Sprintf("%v:%v", frame.Function, line)
}

func messageFromParts(message ...any) string {
	m := string(fmt.Appendln([]byte{}, message...))
	// Remove new line at the end of a message. Adds it later at the end of the formatted line
	return m[:len(m)-1]
}

type Logger struct {
	Writer    io.Writer
	Formatter Formatter
	Filter    Filter
}

func NewLogger(output io.Writer, formatter Formatter, filter Filter) *Logger {
	return &Logger{Writer: output, Formatter: formatter, Filter: filter}
}

func (l *Logger) log(skipCorrection int, logLevel LogLevel, message string, params ...any) {
	cl := callLocation(skipCorrection)

	if l.Filter == nil || l.Filter.Filter(logLevel, cl, message) {
		//TODO-docs: Here errors are not meant to be handled. It should be the concern of Logger.Writer
		_, _ = l.Writer.Write([]byte(l.Formatter.Format(logLevel, time.Now(), cl, message, params...) + "\n"))
	}
}

func (l *Logger) Log(logLevel LogLevel, message ...any) {
	l.log(0, logLevel, messageFromParts(message...))
}

func (l *Logger) Logf(logLevel LogLevel, message string, formats ...any) {
	l.log(0, logLevel, fmt.Sprintf(message, formats...))
}

func (l *Logger) Logs(logLevel LogLevel, message string, params ...any) {
	l.log(0, logLevel, message, params...)
}

func (l *Logger) Debug(message ...any) {
	l.log(0, Debug, messageFromParts(message...))
}

func (l *Logger) Debugf(message string, formats ...any) {
	l.log(0, Debug, fmt.Sprintf(message, formats...))
}

func (l *Logger) Debugs(message string, params ...any) {
	l.log(0, Debug, message, params...)
}

func (l *Logger) Info(message ...any) {
	l.log(0, Info, messageFromParts(message...))
}

func (l *Logger) Infof(message string, formats ...any) {
	l.log(0, Info, fmt.Sprintf(message, formats...))
}

func (l *Logger) Infos(message string, params ...any) {
	l.log(0, Info, message, params...)
}

func (l *Logger) Warn(message ...any) {
	l.log(0, Warn, messageFromParts(message...))
}

func (l *Logger) Warnf(message string, formats ...any) {
	l.log(0, Warn, fmt.Sprintf(message, formats...))
}

func (l *Logger) Warns(message string, params ...any) {
	l.log(0, Warn, message, params...)
}

func (l *Logger) Error(message ...any) {
	l.log(0, Error, messageFromParts(message...))
}

func (l *Logger) Errorf(message string, formats ...any) {
	l.log(0, Error, fmt.Sprintf(message, formats...))
}

func (l *Logger) Errors(message string, params ...any) {
	l.log(0, Error, message, params...)
}

func (l *Logger) Fatal(message ...any) {
	l.log(0, Fatal, messageFromParts(message...))
}

func (l *Logger) Fatalf(message string, formats ...any) {
	l.log(0, Fatal, fmt.Sprintf(message, formats...))
}

func (l *Logger) Fatals(message string, params ...any) {
	l.log(0, Fatal, message, params...)
}

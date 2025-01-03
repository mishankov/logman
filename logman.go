package logman

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
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
		return ""
	}
}

type TimeFormatter interface {
	Format(time time.Time) string
}

type Formatter interface {
	Format(logLevel LogLevel, dateTime string, callLocation string, message string) string
}

type Filter interface {
	Filter(logLevel LogLevel, callLocation string, message string) bool
}

func callLocation() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	_, _, line, _ := runtime.Caller(3)

	bi, ok := debug.ReadBuildInfo()
	if !ok || bi.Main.Path == "" {
		return fmt.Sprintf("%v:%v", frame.Function, line)
	}

	loc := fmt.Sprintf("%v:%v", strings.ReplaceAll(frame.Function, bi.Main.Path+"/", ""), line)

	return loc
}

type Logger struct {
	Writer        io.Writer
	TimeFormatter TimeFormatter
	Formatter     Formatter
	Filter        Filter
}

func NewLogger(output io.Writer, timer TimeFormatter, formatter Formatter, filter Filter) *Logger {
	return &Logger{Writer: output, TimeFormatter: timer, Formatter: formatter, Filter: filter}
}

func NewDefaultLogger() *Logger {
	return &Logger{
		Writer:        os.Stdout,
		TimeFormatter: NewDefaultTimeFormatter(DefaultTimeFormat),
		Formatter:     NewDefaultFormatter(DefaultFormat),
	}
}

func (l *Logger) Log(logLevel LogLevel, message ...any) {
	cl := callLocation()
	m := string(fmt.Appendln([]byte{}, message...))
	if l.Filter == nil || l.Filter.Filter(logLevel, cl, m) {
		//TODO-docs: Here and in Logf errors are not ment to be handled. It should be concern of Logger.Writer
		l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.TimeFormatter.Format(time.Now()), cl, m)))
	}
}

func (l *Logger) Logf(logLevel LogLevel, message string, formats ...any) {
	cl := callLocation()
	m := fmt.Sprintf(message, formats...) + "\n"
	if l.Filter == nil || l.Filter.Filter(logLevel, cl, m) {
		l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.TimeFormatter.Format(time.Now()), cl, m)))
	}
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

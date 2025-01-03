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

type LogLevel string

const (
	Debug = LogLevel("Debug")
	Info  = LogLevel("Info")
	Warn  = LogLevel("Warn")
	Error = LogLevel("Error")
	Fatal = LogLevel("Fatal")
)

type TimeFormatter interface {
	Format(time time.Time) string
}

type Formatter interface {
	Format(logLevel LogLevel, dateTime string, callLocation string, message string) string
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
}

func NewLogger(output io.Writer, timer TimeFormatter, formatter Formatter) *Logger {
	return &Logger{Writer: output, TimeFormatter: timer, Formatter: formatter}
}

func NewDefaultLogger() *Logger {
	return &Logger{
		Writer:        os.Stdout,
		TimeFormatter: NewDefaultTimeFormatter(DefaultTimeFormat),
		Formatter:     NewDefaultFormatter(DefaultFormat),
	}
}

func (l *Logger) Log(logLevel LogLevel, message ...any) {
	//TODO-docs: Here and in Logf errors are not ment to be handled. It should be concern of Logger.Writer
	l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.TimeFormatter.Format(time.Now()), callLocation(), string(fmt.Appendln([]byte{}, message...)))))
}

func (l *Logger) Logf(logLevel LogLevel, message string, formats ...any) {
	l.Writer.Write([]byte(l.Formatter.Format(logLevel, l.TimeFormatter.Format(time.Now()), callLocation(), fmt.Sprintf(message, formats...)+"\n")))
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

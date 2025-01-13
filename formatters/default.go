package formatters

import (
	"bytes"
	"text/template"
	"time"

	"github.com/mishankov/logman"
)

const DefaultFormat = "[{{.DateTime}}] [{{.CallLocation}}] [{{.LogLevel}}] - {{.Message}}"
const DefaultTimeLayout = "2006-01-02 15:04:05 GMT-0700"

// DefaultFormatter implements Formatter interface.
type DefaultFormatter struct {
	template       *template.Template
	dateTimeFormat string
}

// NewDefaultFormatter creates a new DefaultFormatter with the given format string.
// The format string may contain special tags: _logLevel_, _dateTime_, _callLocation_ and _message_.
// These tags will be replaced with the corresponding values when formatting log messages.
func NewDefaultFormatter(format string, dateTimeFormat string) DefaultFormatter {
	templ, _ := template.New("formatter").Parse(format)

	return DefaultFormatter{template: templ, dateTimeFormat: dateTimeFormat}
}

// Format formats a log message according to the format string of the DefaultFormatter.
func (df DefaultFormatter) Format(logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string) string {
	var res = &bytes.Buffer{}
	_ = df.template.Execute(res, struct {
		DateTime     string
		LogLevel     string
		CallLocation string
		Message      string
	}{
		DateTime:     dateTime.Format(df.dateTimeFormat),
		LogLevel:     logLevel.String(),
		CallLocation: callLocation,
		Message:      message,
	},
	)

	return res.String()
}

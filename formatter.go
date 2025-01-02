package logman

import "strings"

const defaultFormat = "[_dateTime_] [_logLevel_] - _message_"

// DefaultFormatter implements Formatter interface
type DefaultFormatter struct {
	format string
}

// NewDefaultFormatter creates a new DefaultFormatter with the given format string.
// The format string may contain special tags: _logLevel_, _dateTime_, and _message_.
// These tags will be replaced with the corresponding values when formatting log messages.
// If the format string is empty, the default format is used.
func NewDefaultFormatter(format string) DefaultFormatter {
	if format == "" {
		format = defaultFormat
	}

	return DefaultFormatter{format: format}
}

// Format formats a log message according to the format string of the DefaultFormatter.
func (df DefaultFormatter) Format(logLevel string, dateTime string, message string) string {
	res := strings.ReplaceAll(df.format, "_logLevel_", logLevel)
	res = strings.ReplaceAll(res, "_dateTime_", dateTime)
	res = strings.ReplaceAll(res, "_message_", message)

	return res
}

package logman

import "strings"

const DefaultFormat = "[_dateTime_] [_callLocation_] [_logLevel_] - _message_"

// DefaultFormatter implements Formatter interface
type DefaultFormatter struct {
	format string
}

// NewDefaultFormatter creates a new DefaultFormatter with the given format string.
// The format string may contain special tags: _logLevel_, _dateTime_, _callLocation_ and _message_.
// These tags will be replaced with the corresponding values when formatting log messages.
func NewDefaultFormatter(format string) DefaultFormatter {
	return DefaultFormatter{format: format}
}

// Format formats a log message according to the format string of the DefaultFormatter.
func (df DefaultFormatter) Format(logLevel LogLevel, dateTime string, callLocation string, message string) string {
	res := strings.ReplaceAll(df.format, "_logLevel_", logLevel.String())
	res = strings.ReplaceAll(res, "_dateTime_", dateTime)
	res = strings.ReplaceAll(res, "_callLocation_", callLocation)
	res = strings.ReplaceAll(res, "_message_", message)

	return res
}

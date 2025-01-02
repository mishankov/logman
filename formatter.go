package logman

import "strings"

type DefaultFormatter struct {
	format string
}

func NewDefaultFormatter(format string) DefaultFormatter {
	return DefaultFormatter{format: format}
}

func (df DefaultFormatter) Format(logLevel string, dateTime string, message string) string {
	res := strings.ReplaceAll(df.format, "_logLevel_", logLevel)
	res = strings.ReplaceAll(res, "_dateTime_", dateTime)
	res = strings.ReplaceAll(res, "_message_", message)

	return res
}

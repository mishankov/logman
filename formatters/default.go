package formatters

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mishankov/logman"
)

const DefaultFormat = "[_dateTime_] [_callLocation_] [_logLevel_] - _message_ _params_"
const DefaultTimeLayout = "2006-01-02 15:04:05 GMT-0700"

// DefaultFormatter implements Formatter interface.
type DefaultFormatter struct {
	format         string
	dateTimeFormat string
}

// NewDefaultFormatter creates a new DefaultFormatter with the given format string.
// The format string may contain special tags: _logLevel_, _dateTime_, _callLocation_, _message_ and _params_.
// These tags will be replaced with the corresponding values when formatting log messages.
func NewDefaultFormatter(format string, dateTimeFormat string) DefaultFormatter {
	return DefaultFormatter{format: format, dateTimeFormat: dateTimeFormat}
}

// Format formats a log message according to the format string of the DefaultFormatter.
// This PR shows why we are using ReplaceAll here instead of templates - https://github.com/mishankov/logman/pull/8
func (df DefaultFormatter) Format(_ context.Context, logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string, params ...any) string {
	var key any
	var paramsStr string
	for i, param := range params {
		if i%2 == 0 {
			key = param
			continue
		}

		paramsStr += fmt.Sprintf("%v=%v ", key, param)
	}

	paramsStr = strings.TrimSpace(paramsStr)

	res := strings.ReplaceAll(df.format, "_logLevel_", logLevel.String())
	res = strings.ReplaceAll(res, "_dateTime_", dateTime.Format(df.dateTimeFormat))
	res = strings.ReplaceAll(res, "_callLocation_", callLocation)
	res = strings.ReplaceAll(res, "_message_", message)
	res = strings.ReplaceAll(res, "_params_", paramsStr)
	res = strings.TrimSpace(res)

	return res
}

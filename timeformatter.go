package logman

import "time"

// const DefaultTimeFormat = "2006-01-02 15:04:05 GMT-0700"

// DefaultTimeFormatter implements TimeProvider interface
type DefaultTimeFormatter struct {
	timeFormat string
}

// NewDefaultTimeProvider creates a new DefaultTimeProvider with the specified time format.
func NewDefaultTimeFormatter(timeFormat string) DefaultTimeFormatter {
	return DefaultTimeFormatter{timeFormat: timeFormat}
}

// Time returns current time formatted according to the time format of the DefaultTimeProvider.
func (dt DefaultTimeFormatter) Format(time time.Time) string {
	return time.Format(dt.timeFormat)
}

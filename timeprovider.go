package logman

import "time"

const defaultTimeFormat = "2006-01-02 15:04:05 GMT-0700"

// DefaultTimeProvider implements TimeProvider interface
type DefaultTimeProvider struct {
	timeFormat string
}

// NewDefaultTimeProvider creates a new DefaultTimeProvider with the specified time format.
func NewDefaultTimeProvider(timeFormat string) DefaultTimeProvider {
	return DefaultTimeProvider{timeFormat: timeFormat}
}

// Time returns current time formatted according to the time format of the DefaultTimeProvider.
func (dt DefaultTimeProvider) Time() string {
	return time.Now().Format(dt.timeFormat)
}

package logman

import "time"

type DefaultTimeProvider struct{}

func (dt *DefaultTimeProvider) Time() string {
	return time.Now().Format("2006-01-02 15:04:05 GMT-0700")
}

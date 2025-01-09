package formatters

import (
	"encoding/json"
	"time"

	"github.com/mishankov/logman"
)

// JSONFormatter implements Formatter interface
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSONFormatter
func NewJSONFormatter() JSONFormatter {
	return JSONFormatter{}
}

type jsonLog struct {
	LogLevel     string `json:"log_level"`
	DateTime     string `json:"date_time"`
	CallLocation string `json:"call_location"`
	Message      string `json:"message"`
}

// Format formats log message as JSON with keys: log_level, date_time, call_location and message
func (jf JSONFormatter) Format(logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string) string {
	res, _ := json.Marshal(jsonLog{logLevel.String(), dateTime.Format(DefaultTimeLayout), callLocation, message})
	return string(res)
}

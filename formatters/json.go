package formatters

import (
	"encoding/json"
	"time"

	"github.com/mishankov/logman"
)

// JSONFormatter implements Formatter interface.
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSONFormatter.
func NewJSONFormatter() JSONFormatter {
	return JSONFormatter{}
}

type jsonLog struct {
	LogLevel     string `json:"logLevel"`
	DateTime     string `json:"dateTime"`
	CallLocation string `json:"callLocation"`
	Message      string `json:"message"`
}

// Format formats log message as JSON with keys: logLevel, dateTime, callLocation and message.
func (jf JSONFormatter) Format(logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string) string {
	res, _ := json.Marshal(jsonLog{logLevel.String(), dateTime.Format(DefaultTimeLayout), callLocation, message})

	return string(res)
}

package formatters

import (
	"context"
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

// Format formats log message as JSON with keys: log_level, date_time, call_location and message.
func (jf JSONFormatter) Format(_ context.Context, logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string, params ...any) string {
	resMap := map[string]any{
		"logLevel":     logLevel.String(),
		"time":         dateTime.Format(DefaultTimeLayout),
		"callLocation": callLocation,
		"message":      message,
	}

	var key string
	for i, param := range params {
		if i%2 == 0 {
			key = param.(string)
			continue
		}

		resMap[key] = param
	}

	res, _ := json.Marshal(resMap)

	return string(res)
}

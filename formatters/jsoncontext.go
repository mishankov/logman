package formatters

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mishankov/logman"
)

// JSONFormatter implements Formatter interface.
type JSONContextFormatter struct {
	dateTimeFormat string
	ctxFields      []fmt.Stringer
}

// NewJSONFormatter creates a new JSONFormatter.
func NewJSONContextFormatter(dateTimeFormat string, ctxFields []fmt.Stringer) JSONContextFormatter {
	return JSONContextFormatter{dateTimeFormat: dateTimeFormat, ctxFields: ctxFields}
}

// Format formats log message as JSON with keys: log_level, date_time, call_location and message.
func (jcf JSONContextFormatter) Format(ctx context.Context, logLevel logman.LogLevel, dateTime time.Time, callLocation string, message string, params ...any) string {
	resMap := map[string]any{
		"level":    logLevel.String(),
		"time":     dateTime.Format(jcf.dateTimeFormat),
		"location": callLocation,
		"msg":      message,
	}

	for _, field := range jcf.ctxFields {
		resMap[field.String()] = ctx.Value(field)
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

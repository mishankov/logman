package formatters

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mishankov/logman"
)

type DefaultContextFormatter struct {
	dateTimeFormat string
	ctxFields      []fmt.Stringer
}

func NewDefaultContextFormatter(dateTimeFormat string, ctxFields []fmt.Stringer) DefaultContextFormatter {
	return DefaultContextFormatter{dateTimeFormat: dateTimeFormat, ctxFields: ctxFields}

}

func writeKeyValue(b *strings.Builder, key, value string) {
	b.WriteString(key)
	b.WriteString("=")

	if strings.Contains(value, " ") || strings.Contains(value, "=") {
		value = fmt.Sprintf("\"%s\"", value)
	}
	b.WriteString(value)
	b.WriteString(" ")
}

func (dcf DefaultContextFormatter) Format(ctx context.Context, level logman.LogLevel, dateTime time.Time, callLocation string, message string, params ...any) string {
	result := strings.Builder{}

	writeKeyValue(&result, "time", dateTime.Format(dcf.dateTimeFormat))
	writeKeyValue(&result, "level", level.String())
	writeKeyValue(&result, "msg", message)

	for _, field := range dcf.ctxFields {
		ctxValueStr := fmt.Sprintf("%v", ctx.Value(field))
		writeKeyValue(&result, field.String(), ctxValueStr)
	}

	var key string
	for i, param := range params {
		if i%2 == 0 {
			key = param.(string)
			continue
		}

		writeKeyValue(&result, key, fmt.Sprintf("%v", param))
	}

	return strings.TrimSpace(result.String())
}

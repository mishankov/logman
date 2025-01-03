package loggers

import (
	"os"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/formatters"
)

func NewDefaultLogger() *logman.Logger {
	return logman.NewLogger(os.Stdout, formatters.NewDefaultFormatter(formatters.DefaultFormat, formatters.DefaultTimeFormat), nil)
}

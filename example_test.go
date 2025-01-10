package logman_test

import (
	"github.com/mishankov/logman"
	"github.com/mishankov/logman/filters"
	"github.com/mishankov/logman/formatters"
	"github.com/mishankov/logman/writers"
)

func ExampleNewLogger() {
	fw, _ := writers.NewFileWriter("error.log")
	formatter := formatters.NewJSONFormatter()
	filter := filters.NewLevelFilter(logman.Error)
	logger := logman.NewLogger(fw, formatter, filter)

	logger.Error("Hello,", "world!")
	logger.Debug("I am not logged")
}

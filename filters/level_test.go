package filters_test

import (
	"fmt"
	"testing"

	"github.com/mishankov/logman"
	"github.com/mishankov/logman/filters"
	"github.com/mishankov/logman/internal/testutils"
)

func TestLevelFilter(t *testing.T) {

	for _, logLevelFilter := range logLevels() {
		levelFilter := filters.NewLevelFilter(logLevelFilter)

		for _, logLevelMessage := range logLevels() {
			t.Run(fmt.Sprintf("filter %v message %v", logLevelFilter, logLevelMessage), func(t *testing.T) {
				if logLevelFilter > logLevelMessage {
					testutils.AssertEqual(t, levelFilter.Filter(logLevelMessage, "", ""), false)
				} else {
					testutils.AssertEqual(t, levelFilter.Filter(logLevelMessage, "", ""), true)
				}
			})
		}
	}
}

func logLevels() [6]logman.LogLevel {
	return [6]logman.LogLevel{
		logman.Debug, logman.Info, logman.Warn, logman.Error, logman.Fatal,
	}
}

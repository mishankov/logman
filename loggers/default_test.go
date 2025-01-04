package loggers_test

import (
	"bytes"
	"testing"

	"github.com/mishankov/logman/internal/testutils"
	"github.com/mishankov/logman/loggers"
)

func TestDefaultLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := loggers.NewDefaultLogger()
	logger.Writer = buffer

	logger.Debug("some message")
	testutils.AssertRegex(t, buffer.String(), `\[\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2} GMT[\+\-]\d{4}\] \[github\.com\/mishankov\/logman\/loggers_test\.TestDefaultLogger:\d+\] \[Debug\] \- some message\n`)
}

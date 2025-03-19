package loggers_test

import (
	"bytes"
	"testing"

	"github.com/mishankov/logman/loggers"
	"github.com/mishankov/testman/assert"
)

func TestDefaultLogger(t *testing.T) {
	buffer := &bytes.Buffer{}
	logger := loggers.NewDefaultLogger()
	logger.Writer = buffer

	logger.Debug("some message")
	assert.Regex(t, buffer.String(), `\[\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2} GMT[\+\-]\d{4}\] \[github\.com\/mishankov\/logman\/loggers_test\.TestDefaultLogger:\d+\] \[Debug\] \- some message\n`)
}

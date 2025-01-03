package logman_test

import "github.com/mishankov/logman"

func ExampleLogger_Debug() {
	logger := logman.NewDefaultLogger()

	logger.Debug("message")
}

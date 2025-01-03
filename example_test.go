package logman_test

import "github.com/mishankov/logman"

func ExampleLogger_Debug() {
	logger := logman.NewDefaultLogger()
	// Using fake time provider for test to pass. Remove it in your code
	logger.TimeFormatter = &FakeTimeFormatter{}

	logger.Debug("message")

	// Output: [2006-01-02 15:04:05 GMT-0700] [github.com/mishankov/logman_test.ExampleLogger_Debug:10] [Debug] - message
}

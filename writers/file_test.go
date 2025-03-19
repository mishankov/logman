package writers_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/logman/writers"
)

func TestFileWriter(t *testing.T) {
	path := t.TempDir() + "/test.log"

	fileWriter, err := writers.NewFileWriter(path)
	if !assert.NoError(t, err) {
		return
	}

	_, err = fileWriter.Write([]byte("some data\n"))
	if !assert.NoError(t, err) {
		return
	}

	_, err = fileWriter.Write([]byte("some more data\n"))
	if !assert.NoError(t, err) {
		return
	}

	data, err := os.ReadFile(path)
	if !assert.NoError(t, err) {
		return
	}

	assert.DeepEqual(t, data, []byte("some data\nsome more data\n"))
}

func TestInvalidPath(t *testing.T) {
	brokenPath := ""
	if runtime.GOOS == "windows" {
		brokenPath = ":://brokenName"
	}

	writer, err := writers.NewFileWriter(brokenPath)
	// TODO: find a way to make a broken path for unix
	if runtime.GOOS == "windows" && err == nil {
		assert.NoError(t, err)
	}

	n, err := writer.Write([]byte("message"))

	assert.Equal(t, n, 0)
	assert.Error(t, err)
}

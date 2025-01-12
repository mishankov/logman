package writers_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/mishankov/logman/internal/testutils"
	"github.com/mishankov/logman/writers"
)

func TestFileWriter(t *testing.T) {
	path := t.TempDir() + "/test.log"

	fileWriter, err := writers.NewFileWriter(path)
	if err != nil {
		t.Error("Error creating FileWriter:", err)

		return
	}

	_, err = fileWriter.Write([]byte("some data\n"))
	if err != nil {
		t.Error("Error writing to test file:", err)

		return
	}

	_, err = fileWriter.Write([]byte("some more data\n"))
	if err != nil {
		t.Error("Error writing to test file second time:", err)

		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Error("Error reading test file:", err)

		return
	}

	testutils.AssertDeepEqual(t, data, []byte("some data\nsome more data\n"))
}

func TestInvalidPath(t *testing.T) {
	brokenPath := ""
	if runtime.GOOS == "windows" {
		brokenPath = ":://brokenName"
	}

	writer, err := writers.NewFileWriter(brokenPath)
	// TODO: find a way to make a broken path for unix
	if runtime.GOOS == "windows" && err == nil {
		t.Error("Error expected to be not nil")
	}

	n, err := writer.Write([]byte("message"))

	testutils.AssertEqual(t, n, 0)

	if err == nil {
		t.Error("Error expected to be not nil")
	}

}

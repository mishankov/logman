package writers_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mishankov/logman/internal/testutils"
	"github.com/mishankov/logman/writers"
)

func TestFileWriter(t *testing.T) {
	path, err := getTempFilePath(t, "test.log")
	if err != nil {
		t.Error("Error contructing file path:", err)
		return
	}

	fw, err := writers.NewFileWriter(path)
	if err != nil {
		t.Error("Error creating FileWriter:", err)
		return
	}

	_, err = fw.Write([]byte("some data\n"))
	if err != nil {
		t.Error("Error writing to test file:", err)
		return
	}

	_, err = fw.Write([]byte("some more data\n"))
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

	w, err := writers.NewFileWriter(brokenPath)
	// TODO: find a way to make a broken path for unix
	if runtime.GOOS == "windows" && err == nil {
		t.Error("Error expected to be not nil")
	}

	n, err := w.Write([]byte("message"))

	testutils.AssertEqual(t, n, 0)
	if err == nil {
		t.Error("Error expected to be not nil")
	}

}

// getTempFilePath returns path to temp file and cleans it up after test finishes
func getTempFilePath(t *testing.T, name string) (string, error) {
	if name != "" {
		name = "file.txt"
	}
	path, err := filepath.Abs("../tmp/" + name)
	if err != nil {
		return "", err
	}

	t.Cleanup(func() { os.RemoveAll(filepath.Dir(path)) })

	return path, err
}

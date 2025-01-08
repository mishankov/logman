package writers_test

import (
	"os"
	"path/filepath"
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

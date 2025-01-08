package writers

import (
	"bufio"
	"os"
	"path/filepath"
)

// FileWriter implements io.Writer interface
type FileWriter struct {
	path string
}

// NewFileWriter creates new FileWriter and creates necessary folders
func NewFileWriter(path string) (FileWriter, error) {
	err := os.MkdirAll(filepath.Dir(path), 0600)
	if err != nil {
		return FileWriter{}, err
	}

	return FileWriter{path: path}, nil
}

// Write writes message to file with path at FileWriter.path
func (fr FileWriter) Write(message []byte) (int, error) {
	f, err := os.OpenFile(fr.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	return w.Write(message)
}

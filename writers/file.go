package writers

import (
	"bufio"
	"os"
	"path/filepath"
)

// FileWriter implements io.Writer interface.
type FileWriter struct {
	path string
}

const permissions = 0777

// NewFileWriter creates new FileWriter and creates necessary folders.
func NewFileWriter(path string) (FileWriter, error) {
	err := os.MkdirAll(filepath.Dir(path), permissions)
	if err != nil {
		return FileWriter{}, err
	}

	return FileWriter{path: path}, nil
}

// Write writes message to file with path at FileWriter.path.
func (fr FileWriter) Write(message []byte) (int, error) {
	file, err := os.OpenFile(fr.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, permissions)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	return w.Write(message)
}

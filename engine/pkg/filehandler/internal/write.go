package internal

import (
	"fmt"
	"os"
)

func (fh *FileHandler) Write(bytes []byte, path string) error {
	if err := os.WriteFile(
		path,
		bytes,
		os.FileMode(0644),
	); err != nil {
		return fmt.Errorf("os.WriteFile: %+v", err)
	}

	return nil
}

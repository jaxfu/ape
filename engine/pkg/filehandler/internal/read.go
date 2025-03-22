package internal

import (
	"os"
)

func (fh *FileHandler) ReadFile(path string) (RawFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return RawFile{}, err
	}

	return RawFile{
		Filepath: path,
		RawBytes: data,
	}, nil
}

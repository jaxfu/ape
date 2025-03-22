package filehandler

import "github.com/jaxfu/ape/engine/pkg/filehandler/internal"

type FileHandler interface {
	ReadFile(string) (RawFile, error)
	GetDirMap(string) (DirFilesMap, error)
	Write([]byte, string) error
}

func NewFileHandler() FileHandler {
	return internal.DefaultFileHandler()
}

type (
	RawFile     = internal.RawFile
	DirFilesMap = internal.DirFilesMap
)

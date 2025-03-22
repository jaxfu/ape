package internal

type FileHandler struct{}

type DirFilesMap = map[string][]string

func DefaultFileHandler() *FileHandler {
	return &FileHandler{}
}

type RawFile struct {
	Filepath string
	RawBytes []byte
}

func (rf RawFile) Path() string {
	return rf.Filepath
}

func (rf RawFile) Bytes() []byte {
	return rf.RawBytes
}

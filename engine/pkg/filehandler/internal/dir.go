package internal

import (
	"os"
	"path/filepath"
)

func (fh *FileHandler) GetDirMap(rootFp string) (DirFilesMap, error) {
	dirMap := map[string][]string{}

	err := filepath.WalkDir(rootFp, func(
		path string,
		d os.DirEntry,
		err error,
	) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(rootFp, path)
		if err != nil {
			return err
		}

		if !d.IsDir() {
			dir := filepath.Dir(relPath)
			dirMap[dir] = append(dirMap[dir], path)
		}

		return nil
	})
	if err != nil {
		return dirMap, err
	}

	// fmt.Printf("%+v\n", dirMap)
	return dirMap, nil
}

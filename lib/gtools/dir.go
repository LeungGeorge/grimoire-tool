package gtools

import (
	"os"
	"path/filepath"
)

// DirSize ...
func DirSize(path string) (size uint64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += uint64(info.Size())
		}
		return err
	})

	return
}

package code

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("stat %s: %w", path, err)
	}

	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	var totalSize int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("read dir %s: %w", path, err)
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		size, err := GetPathSize(childPath)
		if err != nil {
			return 0, err
		}
		totalSize += size
	}

	return totalSize, nil
}

package code

import (
	"os"
)

func GetPathSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var totalSize int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			continue
		}
		if !entryInfo.IsDir() {
			totalSize += entryInfo.Size()
		}
	}
	return totalSize, nil

}

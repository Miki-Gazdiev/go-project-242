package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}
		var totalSize int64
		for _, entrie := range entries {
			if entrie.IsDir() {
				continue
			}
			info, err := entrie.Info()
			if err != nil {
				continue
			}
			totalSize += info.Size()
		}
		return fmt.Sprintf("%dB", totalSize), nil
	}
	return fmt.Sprintf("%dB", fileInfo.Size()), nil
}

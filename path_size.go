package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		var totalSize int64
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			info, err := entry.Info()
			if err != nil {
				continue
			}
			totalSize += info.Size()
		}
		return totalSize, nil
	}
	return fileInfo.Size(), nil
}

func FormatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	value := float64(size)
	if human {
		for value >= 1024 && unitIndex < len(units)-1 {
			value /= 1024
			unitIndex++
		}
		return fmt.Sprintf("%.1f%s", value, units[unitIndex])
	}
	return fmt.Sprintf("%dB", size)
}

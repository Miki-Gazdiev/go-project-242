package code

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetPathSize_File(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := []byte("hello")
	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	size, err := GetPathSize(tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := int64(len(content))
	if size != expected {
		t.Errorf("expected %d bytes, got %d", expected, size)
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "testdir*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	files := map[string]string{
		"file1.txt": "hello",
		"file2.txt": "world",
		"file3.txt": "go",
	}

	var expectedSize int64
	for name, content := range files {
		path := filepath.Join(tmpDir, name)
		os.WriteFile(path, []byte(content), 0644)
		expectedSize += int64(len(content))
	}
	size, err := GetPathSize((tmpDir))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if size != expectedSize {
		t.Errorf("expected %d bytes, got %d", expectedSize, size)
	}
}

package copy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src, dest string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err 
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(srcPath)
		if err != nil {
			return err 
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := createDir(destPath, 0755); err != nil {
				return err 
			}
			if err := CopyDir(srcPath, destPath); err != nil {
				return err
			}
		default:
			if err := copy(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func createDir(dir string, perm os.FileMode) error {
	if exists(dir) {
		return nil 
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s' error: '%s'", dir, err.Error())
	}

	return nil
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false 
	}

	return true
}

func copy(src, dest string) error {
	stats, err := os.Stat(src)
	if err != nil {
		return err 
	}

	if !stats.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)		
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err 
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
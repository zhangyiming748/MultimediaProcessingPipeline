package stepbystep

import (
	"os"
	"path/filepath"
)

// FindFiles finds all files (not directories) in a given path, recursively.
// It's similar to the `find . -type f` command.
// It returns a slice of strings with the absolute paths of the files.
func FindSubtitleFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it is a regular file.
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if ext := filepath.Ext(absPath); ext == ".srt" {
				files = append(files, absPath)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

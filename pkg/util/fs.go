package util

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// CreateFolderIfNotExist creates a new folder at the specified path if it
// does not already exist with the parent folder's permissions. It will only
// create a folder if the parent folder exists.
func CreateFolderIfNotExist(path string) error {
	// Checking if the folder already exists
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return nil // Folder already exists, no need to create one
	}

	// Getting the permissions of the parent folder
	parentFolder := filepath.Dir(path)
	parentInfo, err := os.Stat(parentFolder)
	if err != nil {
		return fmt.Errorf(
			"unable to get permissions of parent folder: %w",
			err,
		)
	}
	permissions := parentInfo.Mode().Perm()

	// Creating the folder with the same permissions as the parent folder
	err = os.MkdirAll(path, permissions)
	if err != nil {
		return fmt.Errorf("unable to create the folder: %w", err)
	}

	// Folder created successfully
	return nil
}

// CopyDirToDir copies an embedded filesystem to a real OS directory
func CopyDirToDir(src fs.FS, dest string) error {
	err := fs.WalkDir(
		src, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// Generate full paths
			embeddedPath := filepath.Join(".", path)
			targetPath := filepath.Join(dest, path)

			// Handle directories
			if d.IsDir() {
				return os.MkdirAll(targetPath, 0755)
			}

			// Copy files
			data, err := fs.ReadFile(src, embeddedPath)
			if err != nil {
				return err
			}

			return os.WriteFile(targetPath, data, 0755)
		},
	)

	return err
}

package util

import (
	"fmt"
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

// Name: init.go
// Desc:
//   Initializing the logging portion of this application.
package logging

import (
	"errors"
	"github.com/crockeo/dotfile-manager/files"
	"os"
)

var (
	fileHandle *os.File = nil
)

// Initializing the logger such that it has a file handle open to use.
func Open(path string) error {
	if files.Exists(path) {
		file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)

		if err != nil {
			return errors.New("Failed to open file at " + path)
		}

		fileHandle = file
		return nil
	} else {
		file, err := os.Create(path)

		if err != nil {
			return errors.New("Failed to create logging file at " + path)
		}

		fileHandle = file
		return nil
	}
}

// Cleaning up the logger.
func Close() error {
	if fileHandle == nil {
		return errors.New("fileHandle is already closed.")
	} else if _, err := fileHandle.Stat(); err != nil {
		fileHandle = nil
		return errors.New("There as an error accessing the file handle.")
	}

	err := fileHandle.Close()
	if err != nil {
		fileHandle = nil
		return errors.New("Failed to close the file.")
	}

	fileHandle = nil
	return nil
}

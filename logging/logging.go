// Name: logging.go
// Desc:
//   This portion of the package implements the main logging functionality -
//   printing out both to the user and to the log file.
package logging

import (
	"errors"
	"fmt"
)

// Logging some piece of data.
func Write(str string) error {
	if fileHandle == nil {
		return errors.New("The logging file handle does not exist. (Did you forget to run logging.Init()?)")
	}

	_, err := fileHandle.WriteString(str + "\n")
	if err != nil {
		return errors.New("Could not write to the fileHandle.")
	}

	fmt.Println(str)

	return nil
}

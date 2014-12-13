// Name: move.go
// Desc:
//   This file provides the function to move a file to another location.
package files

import (
	"errors"
	"os/exec"
)

// This function moves a file from one place to another.
func Move(src, dst string) error {
	cmd := exec.Command("mv", src, dst)

	err := cmd.Run()
	if err != nil {
		return errors.New("Failed to move the file.")
	}

	return nil
}

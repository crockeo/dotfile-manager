// Name: copy.go
// Desc:
//   This portion of a package copies a file from one place to another.
package files

import (
	"errors"
	"os/exec"
)

// The copy function itself.
func Copy(src, dst string) error {
	cmd := exec.Command("cp", "-r", src, dst)

	err := cmd.Run()
	if err != nil {
		return errors.New("Could not copy " + src + " to " + dst)
	}

	return nil
}

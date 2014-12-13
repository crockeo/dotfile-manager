// Name: exists.go
// Desc:
//   This provides the functionality to check whether or not a file or folder
//   exists on the file system.
package files

import "os"

// Checking if a file exists and can be opened.
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

// Name: loadpackage.go
// Desc:
//   A package to load metadata about a dotfile package.
package pkgfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crockeo/dotfile-manager/files"
	"os"
)

// Loading a package from a given location.
func LoadPackage(name string) (PkgFile, error) {
	path := name + "/" + RelativeLocation
	fmt.Println(path)

	if !files.Exists(path) {
		return PkgFile{}, errors.New("Package file does not exist!")
	}

	file, err := os.Open(path)

	if err != nil {
		return PkgFile{}, errors.New("Failed to open the package file.")
	}

	var pkgFile PkgFile
	fmt.Println("Loading from " + path + "!")
	decoder := json.NewDecoder(file)
	decoder.Decode(&pkgFile)

	return pkgFile, nil
}

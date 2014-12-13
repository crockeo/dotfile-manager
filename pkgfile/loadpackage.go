// Name: loadpackage.go
// Desc:
//   A package to load metadata about a dotfile package.
package pkgfile

import "fmt"

// Loading a package from a given location.
func LoadPackage(name string) (PkgFile, error) {
	path := name + "/" + RelativeLocation

	fmt.Println("Loading from " + path + "!")

	return PkgFile{
		name,
		make([]PkgOperation, 0),
	}, nil
}

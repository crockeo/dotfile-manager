// Name: loadpackage.go
// Desc:
//   A package to load metadata about a dotfile package.
package pkgfile

// Loading a package from a given location.
func LoadPackage(name string) (PkgFile, error) {
	return PkgFile{
		name,
		make([]PkgOperation, 0),
	}, nil
}

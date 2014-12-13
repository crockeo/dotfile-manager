// Name: installpackage.go
// Desc:
//   The portion of this package to deal with the InstallPackage type.
package pkgfile

const (
	InstallPackageType string = "InstallPackage"
)

// A type to represent downloading another package from GitHub.
type InstallPackage struct {
	Url string
}

// Converting an InstallPackage to a nice string that can be printed out later.
func (ip InstallPackage) String() string {
	return "InstallPackage " + ip.Url
}

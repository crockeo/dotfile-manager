// Name: pkgfile.go
// Desc:
//   The types used in this module to represent package files.
package pkgfile

const (
	RelativeLocation string = "dotfile.json"
)

// A type that represents the form a package installation can take.
type PkgFile struct {
	Name            string           `json:"name"`
	MoveFiles       []CopyFile       `json:"moveFiles"`
	RunScripts      []RunScript      `json:"runScripts"`
	InstallPackages []InstallPackage `json:"installPackages"`
}

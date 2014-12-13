// Name: pkgfile.go
// Desc:
//   The types used in this module to represent package files.
package pkgfile

const (
	MoveFileType       string = "MoveFile"
	RunScriptType      string = "RunScript"
	InstallPackageType string = "InstallPackage"

	RelativeLocation string = "dotfile.json"
)

type PkgOperation interface {
	GetType() string
}

// A type to represent moving a file from one place to another.
type MoveFile struct {
	Src string
	Dst string
}

func (mf MoveFile) GetType() string {
	return MoveFileType
}

// A type to represent running a script at a given location (if it has run
// permissions.)
type RunScript struct {
	Src string
}

func (rs RunScript) GetType() string {
	return RunScriptType
}

// A type to represent downloading another package from GitHub.
type InstallPackage struct {
	Url string
}

func (dp InstallPackage) GetType() string {
	return InstallPackageType
}

// The type that represents a whole package file.
type PkgFile struct {
	Name string         // The name of the package.
	Ops  []PkgOperation // Each operation that the file wants to perform.
}

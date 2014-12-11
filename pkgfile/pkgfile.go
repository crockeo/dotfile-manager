// Name: pkgfile.go
// Desc:
//   The types used in this module to represent package files.
package pkgfile

const (
	MoveFileType        string = "MoveFile"
	RunScriptType       string = "RunScript"
	DownloadPackageType string = "DownloadPackage"
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
type DownloadPackage struct {
	Url string
}

func (dp DownloadPackage) GetType() string {
	return DownloadPackageType
}

// The type that represents a whole package file.
type PkgFile []PkgOperation

// Name: movefile.go
// Desc:
//   This type describes moving a file to another location.
package pkgfile

const (
	CopyFileType string = "MoveFile"
)

// The type to describe a file move.
type CopyFile struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

// Converting a move file to a string.
func (mf CopyFile) String() string {
	return "CopyFile '" + mf.Src + "' -> '" + mf.Dst + "'"
}

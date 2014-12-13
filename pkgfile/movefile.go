// Name: movefile.go
// Desc:
//   This type describes moving a file to another location.
package pkgfile

const (
	MoveFileType string = "MoveFile"
)

// The type to describe a file move.
type MoveFile struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

// Converting a move file to a string.
func (mf MoveFile) String() string {
	return "MoveFile " + mf.Src + " -> " + mf.Dst
}

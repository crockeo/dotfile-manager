// Name: runscript.go
// Desc:
//   This portion of the package describes the type to represent running a
//   script.
package pkgfile

const (
	RunScriptType string = "RunScript"
)

// The type itself.
type RunScript struct {
	Src string `json:"src"`
}

// Converting a RunScript into a nice string such that it can be printed later.
func (rs RunScript) String() string {
	return "RunScript '" + rs.Src + "'"
}

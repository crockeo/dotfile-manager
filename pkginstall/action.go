// Name: action.go
// Desc:
//   This portion of the package manager is meant to go through each
//   PkgOperation and perform them.
package pkginstall

import (
	"errors"
	"fmt"
	"github.com/crockeo/dotfile-manager/files"
	"github.com/crockeo/dotfile-manager/pkgfile"
	"os/exec"
)

var (
	MalformedPkgOperationError error = errors.New("Malformed package operation.")
)

// Performing a MoveFile operation.
func performMoveFile(moveFile pkgfile.MoveFile) error {
	return files.Move(moveFile.Src, moveFile.Dst)
}

// Performing a RunScript operation.
func performRunScript(runScript pkgfile.RunScript) error {
	cmd := exec.Command("sh", runScript.Src)
	err := cmd.Run()

	if err != nil {
		return errors.New("The script '" + runScript.Src + "' returned with an error.")
	}

	return nil
}

// Performing an InstallPackage operation.
func performInstallPackage(installPackage pkgfile.InstallPackage) error {
	return InstallPackage(installPackage.Url)
}

// Performing a single package operation.
func PerformPackageOperation(pkgOperation pkgfile.PkgOperation) error {
	t := pkgOperation.GetType()

	switch t {
	case pkgfile.MoveFileType:
		v, succ := pkgOperation.(pkgfile.MoveFile)
		fmt.Println("Running: " + v.String())
		if !succ {
			return MalformedPkgOperationError
		}

		return performMoveFile(v)
	case pkgfile.RunScriptType:
		v, succ := pkgOperation.(pkgfile.RunScript)
		fmt.Println("Running: " + v.String())
		if !succ {
			return MalformedPkgOperationError
		}

		return performRunScript(v)
	case pkgfile.InstallPackageType:
		v, succ := pkgOperation.(pkgfile.InstallPackage)
		fmt.Println("Running: " + v.String())
		if !succ {
			return MalformedPkgOperationError
		}

		return performInstallPackage(v)
	}

	return MalformedPkgOperationError
}

// Performing a whole set of package operations (as described in a PkgFile).
func PerformPackageOperations(pkgFile pkgfile.PkgFile) error {
	var err error
	for _, v := range pkgFile.Ops {
		err = PerformPackageOperation(v)
		if err != nil {
			return err
		}
	}

	return nil
}

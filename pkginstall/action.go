// Name: action.go
// Desc:
//   This portion of the package manager is meant to go through each
//   PkgOperation and perform them.
package pkginstall

import (
	"fmt"
	"github.com/crockeo/dotfile-manager/pkgfile"
)

//MoveFile
//RunScript
//InstallPackage

// Performing a MoveFile operation.
func performMoveFile(moveFile pkgfile.MoveFile) error {
	return nil
}

// Performing a RunScript operation.
func performRunScript(runScript pkgfile.RunScript) error {
	return nil
}

// Performing an InstallPackage operation.
func performInstallPackage(installPackage pkgfile.InstallPackage) error {
	return nil
}

// Performing a single package operation.
func PerformPackageOperation(pkgOperation pkgfile.PkgOperation) error {
	t := pkgOperation.GetType()

	switch t {
	case pkgfile.MoveFileType:
		v, succ := pkgOperation.(pkgfile.MoveFile)
		if !succ {
			fmt.Println(":(")
		}

		return performMoveFile(v)
	case pkgfile.RunScriptType:
		v, succ := pkgOperation.(pkgfile.RunScript)
		if !succ {
			fmt.Println(":(")
		}

		return performRunScript(v)
	case pkgfile.InstallPackageType:
		v, succ := pkgOperation.(pkgfile.InstallPackage)
		if !succ {
			fmt.Println(":(")
		}

		return performInstallPackage(v)
	}

	return nil
}

// Performing a whole set of package operations (as described in a PkgFile).
func PerformPackageOperations(pkgFile pkgfile.PkgFile) error {
	var err error
	for _, v := range pkgFile {
		err = PerformPackageOperation(v)
		if err != nil {
			return err
		}
	}

	return nil
}

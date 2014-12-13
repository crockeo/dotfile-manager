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
func performMoveFile(name string, moveFile pkgfile.MoveFile) error {
	return files.Move(name+"/"+moveFile.Src, moveFile.Dst)
}

// Performing a RunScript operation.
func performRunScript(name string, runScript pkgfile.RunScript) error {
	cmd := exec.Command("sh", name+"/"+runScript.Src)
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

// Performing a whole set of package operations (as described in a PkgFile).
func PerformPackageOperations(pkgFile pkgfile.PkgFile) error {
	var err error

	for _, v := range pkgFile.MoveFiles {
		fmt.Println("Performing: " + v.String())
		err = performMoveFile(pkgFile.Name, v)
		if err != nil {
			return err
		}
	}

	for _, v := range pkgFile.RunScripts {
		fmt.Println("Performing: " + v.String())
		err = performRunScript(pkgFile.Name, v)
		if err != nil {
			return err
		}
	}

	for _, v := range pkgFile.InstallPackages {
		fmt.Println("Performing: " + v.String())
		err = performInstallPackage(v)
		if err != nil {
			return err
		}
	}

	return nil
}

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
	"os"
	"os/exec"
	"strings"
)

var (
	MalformedPkgOperationError error = errors.New("Malformed package operation.")
)

// Performing a MoveFile operation.
func performCopyFile(name string, copyFile pkgfile.CopyFile) error {
	return files.Copy(name+"/"+copyFile.Src, strings.Replace(copyFile.Dst, "~", os.Getenv("HOME"), -1))
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

	for _, v := range pkgFile.CopyFiles {
		fmt.Println("Performing: " + v.String())
		err = performCopyFile(pkgFile.Name, v)
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

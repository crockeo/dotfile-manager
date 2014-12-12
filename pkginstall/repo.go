// Name: repo.go
// Desc:
//   This module deals with installing a whole package hosted on a GitHub repo.
package pkginstall

import (
	"github.com/crockeo/dotfile-manager/pkgfile"
	"os/exec"
)

// Formatting the URL for a git clone.
func formatURL(name string) string {
	return "https://github.com/" + name
}

// Performing a Git clone on a given repo name.
func cloneRepo(name string) error {
	cmd := exec.Command("git", "clone", formatURL(name))
	return cmd.Run()
}

// Installing a package from a Git repository at a given location.
func InstallPackage(name string) error {
	err := cloneRepo(name)

	if err != nil {
		return err
	}

	pkg, err := pkgfile.LoadPackage(name)

	if err != nil {
		return err
	}

	PerformPackageOperations(pkg)

	return nil
}

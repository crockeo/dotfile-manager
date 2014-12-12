// Name: repo.go
// Desc:
//   This module deals with installing a whole package hosted on a GitHub repo.
package pkginstall

import (
	"errors"
	"github.com/crockeo/dotfile-manager/pkgfile"
	"os/exec"
	"strings"
)

// Formatting the URL for a git clone.
func formatURL(name string) string {
	return "https://github.com/" + name
}

// Performing a Git clone on a given repo name.
func cloneRepo(name string) error {
	cmd := exec.Command("git", "clone", formatURL(name))
	err := cmd.Run()

	if err != nil {
		return errors.New("Failed to clone repository '" + name + "'!")
	}

	return nil
}

// Getting the directory of the repo by its name.
func getRepoName(name string) string {
	ss := strings.Split(name, "/")
	return strings.Trim(ss[len(ss)-1], ".git")
}

// Installing a package from a Git repository at a given location.
func InstallPackage(name string) error {
	err := cloneRepo(name)

	if err != nil {
		return err
	}

	pkg, err := pkgfile.LoadPackage(getRepoName(name))

	if err != nil {
		return err
	}

	PerformPackageOperations(pkg)

	return nil
}

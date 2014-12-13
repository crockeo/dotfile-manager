// Name: repo.go
// Desc:
//   This module deals with installing a whole package hosted on a GitHub repo.
package pkginstall

import (
	"errors"
	"fmt"
	"github.com/crockeo/dotfile-manager/files"
	"github.com/crockeo/dotfile-manager/pkgfile"
	"os/exec"
	"strings"
)

// Formatting the URL for a git clone.
func formatURL(name string) string {
	return "http://github.com/" + name
}

// Performing a Git clone on a given repo name.
func cloneRepo(name string) error {
	fmt.Println("Attempting to clone repository '" + name + "'!")

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
	return strings.TrimSuffix(ss[len(ss)-1], ".git")
}

// Installing a package from a Git repository at a given location.
func InstallPackage(name string) error {
	if !files.Exists(getRepoName(name)) {
		err := cloneRepo(name)

		if err != nil {
			return err
		}
	} else {
		fmt.Println("Repository '" + name + "' already exists - using cached files!")
	}

	pkg, err := pkgfile.LoadPackage(getRepoName(name))

	if err != nil {
		return err
	}

	return PerformPackageOperations(pkg)
}

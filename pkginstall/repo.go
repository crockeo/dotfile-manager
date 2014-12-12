// Name: repo.go
// Desc:
//   This module deals with installing a whole package hosted on a GitHub repo.
package pkginstall

import (
	"log"
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
func InstallPackage(name string) {
	err := cloneRepo(name)

	if err != nil {
		log.Fatal("Could not clone GitHub repo at '" + name + "'!")
	}
}

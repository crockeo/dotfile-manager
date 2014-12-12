// Name: main.go
// Desc:
//   The entry point to the application.
package main

import (
	"github.com/crockeo/dotfile-manager/pkginstall"
	"log"
	"os"
)

// The REAL entry point to the application.
func main() {
	if len(os.Args) != 3 {
		log.Println("Proper command usage: dotfile-manager install <pkg location>")
	} else {
		pkginstall.InstallPackage(os.Args[2])
	}
}

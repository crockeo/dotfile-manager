// Name: main.go
// Desc:
//   The entry point to the application.
package main

import (
	"fmt"
	"github.com/crockeo/dotfile-manager/logging"
	"github.com/crockeo/dotfile-manager/pkginstall"
	"os"
)

// The REAL entry point to the application.
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Proper command usage: dotfile-manager install <pkg location>")
	} else {
		err := logging.Open(".dotfile-manager.log")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer logging.Close()

		err = pkginstall.InstallPackage(os.Args[2])
		if err != nil {
			logging.Write(err.Error())
		} else {
			logging.Write("Package successfully installed!")
		}
	}
}

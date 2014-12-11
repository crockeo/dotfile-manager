// Name: main.go
// Desc:
//   The entry point to the application.
package main

import (
	"fmt"
	"os"
	"os/exec"
)

// The REAL entry point to the application.
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Proper command usage: dotfile-manager install <pkg location>")
	} else {
		url := "http://github.com/" + os.Args[2] + ".git"
		cmd := exec.Command("git", "clone", url)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("ERROR!")
			fmt.Println(" " + err.Error())
		} else {
			fmt.Println("SUCCESS!")
			fmt.Println(" " + string(out))
		}

		fmt.Println(url)
		// TODO: Git cloning of a repo.
		fmt.Println("You used it right!")
	}
}

// Name: repo_test.go
// Desc:
//   Testing the repo.go file.
package pkginstall

import (
	"testing"
)

// Function TestFormatURL tests the function formatURL.
func TestFormatURL(t *testing.T) {
	ins := []string{
		"crockeo/vimconfig.git",
		"crockeo/prelude.git",
		"what/testing-things",
	}

	outs := []string{
		"https://github.com/crockeo/vimconfig.git",
		"https://github.com/crockeo/prelude.git",
		"https://github.com/what/testing-things",
	}

	for n := 0; n < len(ins); n++ {
		if formatURL(ins[n]) != outs[n] {
			t.Error(formatURL(ins[n]) + " did not equal " + outs[n])
		}
	}
}

// Function TestGetRepoName tests the function getRepoName.
func TestGetRepoName(t *testing.T) {
	ins := []string{
		"crockeo/vimconfig.git",
		"crockeo/prelude.git",
		"what/testing-things",
	}

	outs := []string{
		"vimconfig",
		"prelude",
		"testing-things",
	}

	for n := 0; n < len(ins); n++ {
		if getRepoName(ins[n]) != outs[n] {
			t.Error(getRepoName(ins[n]) + " did not equal " + outs[n])
		}
	}
}

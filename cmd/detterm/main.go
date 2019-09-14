// detterm starts with its own process id and walks the process tree
// until it finds one of the passed terminal emulators.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

var defaultEmulators = []string{"alacritty", "konsole", "yakuake"} //nolint: gochecknoglobals

func makeTermLookup(names []string) map[string]struct{} {
	expectedTerms := make(map[string]struct{}, len(names))
	for _, s := range names {
		expectedTerms[s] = struct{}{}
	}
	return expectedTerms
}

func emulatorNamesFromEnv() []string {
	emulators := os.Getenv("DETTERM_EMULATORS")
	if emulators == "" {
		return nil
	}
	return strings.Split(emulators, " ")
}

func main() {
	emulatorNames := emulatorNamesFromEnv()
	if len(emulatorNames) == 0 {
		emulatorNames = defaultEmulators
	}
	expectedTerms := makeTermLookup(emulatorNames)
	pid := os.Getpid()
	for pid != 1 {
		proc, err := ps.FindProcess(pid)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		if _, ok := expectedTerms[proc.Executable()]; ok {
			fmt.Println(proc.Executable())
			os.Exit(0)
		}
		pid = proc.PPid()
	}
}

// determine-term starts with its own process id and walks the process tree
// until it finds one of the passed terminal emulators.
package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-ps"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage:")
		fmt.Printf("\n\t%s <name> [names ...]", os.Args[0])
		os.Exit(1)
	}
	expectedTerms := make(map[string]struct{}, len(os.Args)-1)
	for _, s := range os.Args[1:] {
		expectedTerms[s] = struct{}{}
	}
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

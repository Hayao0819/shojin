package main

import (
	"os"

	"github.com/Hayao0819/shojin/cmd"
)

func main() {
	root := cmd.Root()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

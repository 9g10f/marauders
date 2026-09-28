package main

import (
	"os"

	"codeberg.org/9g10f/marauders/cmd"
)

func main() {
	err := cmd.MarauderS()
	if err != nil {
		os.Exit(1)
	}
}
package main

import (
	"os"

	"github.com/meibel-ai/meibel-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

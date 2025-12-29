package main

import (
	"os"

	"github.com/fruity-loozrz/go-scratchpad/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
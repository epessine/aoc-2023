package main

import (
	"os"

	"github.com/epessine/aoc-2023/cmd"
)

func main() {
	err := cmd.NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

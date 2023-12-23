package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/epessine/aoc-2023/day1"
	"github.com/epessine/aoc-2023/day2"
)

func NewRootCommand() *cobra.Command {
	var start time.Time

	result := &cobra.Command{
		Use:   "aoc-2023",
		Short: "Advent of Code 2023 Solutions",
		Long: `Advent of Code 2023 Solutions
		Run with params {day} {challenge} to print out solution.`,
		Args: cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			start = time.Now()
			fmt.Println("Running...")
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			fmt.Println("Running time:", time.Since(start))
		},
	}

	addCommands(result)

	return result
}

func addCommands(root *cobra.Command) {
	root.AddCommand(day1.Command())
	root.AddCommand(day2.Command())
}

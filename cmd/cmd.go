package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/epessine/aoc-2023/day1"
	"github.com/epessine/aoc-2023/day2"
	"github.com/epessine/aoc-2023/day3"
	"github.com/epessine/aoc-2023/day4"
	"github.com/epessine/aoc-2023/day5"
	"github.com/epessine/aoc-2023/day6"
	"github.com/epessine/aoc-2023/day7"
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
	root.AddCommand(
		day1.Command(),
		day2.Command(),
		day3.Command(),
		day4.Command(),
		day5.Command(),
		day6.Command(),
		day7.Command(),
	)
}

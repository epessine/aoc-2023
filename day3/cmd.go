package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/epessine/aoc-2023/challenge"
)

func Command() *cobra.Command {
	day3 := &cobra.Command{
		Use:   "3",
		Short: "Day 3 Solutions",
	}

	day3.AddCommand(
		&cobra.Command{
			Use:   "a",
			Short: "A Solution",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Solution is:", solveA(challenge.GetInput(cmd)))
			},
		},
		&cobra.Command{
			Use:   "b",
			Short: "B Solution",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Solution is:", solveB(challenge.GetInput(cmd)))
			},
		},
	)

	return day3
}

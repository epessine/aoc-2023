package day6

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/epessine/aoc-2023/challenge"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "6",
		Short: "Day 6 Solutions",
	}

	cmd.AddCommand(
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

	return cmd
}

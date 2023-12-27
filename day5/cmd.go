package day5

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/epessine/aoc-2023/challenge"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "5",
		Short: "Day 5 Solutions",
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

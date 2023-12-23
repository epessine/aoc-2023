package day2

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	day2 := &cobra.Command{
		Use:   "2",
		Short: "Day 2 Solutions",
	}

	day2.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "A Solution",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Solution is:", solveA())
		},
	})

	day2.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "B Solution",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Solution is:", solveB())
		},
	})

	return day2
}

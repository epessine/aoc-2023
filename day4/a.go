package day4

import (
	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	points := 0
	for line := range input.Lines() {
		card := createFromLine(line)
		points += card.calculatePoints()
	}
	return points
}

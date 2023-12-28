package day7

import (
	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	CardValues = map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}
	JokerRule = false
	return parseInput(input).totalWinnings()
}

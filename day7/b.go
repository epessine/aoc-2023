package day7

import (
	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	CardValues = map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	JokerRule = true
	return parseInput(input).totalWinnings()
}

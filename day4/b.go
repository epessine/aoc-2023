package day4

import (
	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	cards := make([]Card, 196)
	for line := range input.Lines() {
		card := createFromLine(line)
		cards = append(cards, *card)
	}
	count := 0
	for ci, c := range cards {
		for i := 0; i < c.copies; i++ {
			n := c.calculateMatchingNumbers()
			if n == 0 {
				break
			}
			for i := ci + 1; i <= ci+n; i++ {
				cards[i].copies++
			}
		}
		count += c.copies
	}
	return count
}

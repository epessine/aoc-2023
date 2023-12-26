package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	numbers        []int
}

func createFromLine(line string) *Card {
	card := Card{}
	numbers := strings.Split(strings.Split(line, ":")[1], "|")
	parseNumberString(&card.winningNumbers, numbers[0])
	parseNumberString(&card.numbers, numbers[1])
	return &card
}

func parseNumberString(rc *[]int, s string) {
	for _, wn := range strings.Fields(strings.TrimSpace(s)) {
		pwn, err := strconv.Atoi(wn)
		if err != nil {
			panic(err)
		}
		*rc = append(*rc, pwn)
	}
}

func (c *Card) calculatePoints() int {
	points := 0
	for _, n := range c.numbers {
		for _, wn := range c.winningNumbers {
			if n == wn {
				if points == 0 {
					points++
				} else {
					points = points * 2
				}
			}
		}
	}
	return points
}

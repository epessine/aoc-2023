package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	copies         int
	winningNumbers []int
	numbers        []int
}

func createFromLine(line string) *Card {
	s := strings.Split(line, ":")
	numbers := strings.Split(s[1], "|")
	return &Card{
		copies:         1,
		winningNumbers: parseNumberString(numbers[0]),
		numbers:        parseNumberString(numbers[1]),
	}
}

func parseNumberString(s string) []int {
	ns := []int{}
	for _, wn := range strings.Fields(strings.TrimSpace(s)) {
		pwn, err := strconv.Atoi(wn)
		if err != nil {
			panic(err)
		}
		ns = append(ns, pwn)
	}
	return ns
}

func (c *Card) calculateMatchingNumbers() int {
	tn := 0
	for _, n := range c.numbers {
		for _, wn := range c.winningNumbers {
			if n == wn {
				tn++
			}
		}
	}
	return tn
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

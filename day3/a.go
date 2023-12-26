package day3

import (
	"slices"
	"unicode"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	y := 0
	numbers := []Number{}
	symbols := []Symbol{}
	for line := range input.Lines() {
		number := Number{}
		l := []rune(line)
		for x, r := range l {
			if unicode.IsDigit(r) {
				number.build(r, x, y)
				if x == len(l)-1 {
					numbers = append(numbers, number.complete(x))
				}
				continue
			}
			if r != rune('.') {
				symbols = append(symbols, Symbol{X: x, Y: y})
			}
			if number.Builder.Len() > 0 {
				numbers = append(numbers, number.complete(x))
			}
			number = Number{}
		}
		y++
	}
	sum := 0
	for _, number := range numbers {
		if slices.ContainsFunc(symbols, number.isAdjacentTo) {
			sum += number.Value
		}
	}
	return sum
}

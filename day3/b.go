package day3

import (
	"unicode"

	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	y := 0
	numbers := []Number{}
	asterisks := []Symbol{}
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
			if r == rune('*') {
				asterisks = append(asterisks, Symbol{X: x, Y: y})
			}
			if number.Builder.Len() > 0 {
				numbers = append(numbers, number.complete(x))
			}
			number = Number{}
		}
		y++
	}
	sum := 0
	for _, asterisk := range asterisks {
		n := []Number{}
		for _, number := range numbers {
			if number.isAdjacentTo(asterisk) {
				n = append(n, number)
			}
		}
		if len(n) == 2 {
			sum += n[0].Value * n[1].Value
		}
	}
	return sum
}

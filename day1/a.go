package day1

import (
	"regexp"
	"strconv"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	var total int = 0
	for line := range input.Lines() {
		numbers := regexp.MustCompile("[1-9]").FindAllString(line, -1)
		parsedDigits, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			panic(err)
		}
		total += parsedDigits
	}
	return total
}

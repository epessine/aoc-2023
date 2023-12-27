package day6

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	race := Race{}
	for line := range input.Lines() {
		if strings.HasPrefix(line, "Time:") {
			duration, err := strconv.Atoi(strings.Map(func(r rune) rune {
				if unicode.IsDigit(r) {
					return r
				}
				return -1
			}, line))
			if err != nil {
				panic(err)
			}
			race.Duration = duration
		}
		if strings.HasPrefix(line, "Distance:") {
			record, err := strconv.Atoi(strings.Map(func(r rune) rune {
				if unicode.IsDigit(r) {
					return r
				}
				return -1
			}, line))
			if err != nil {
				panic(err)
			}
			race.Record = record
		}
	}
	return race.WaysToWin()
}

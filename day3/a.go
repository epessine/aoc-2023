package day3

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	lines := input.AllLines()
	sum := 0
	for i, line := range lines {
		numbers := strings.Fields(strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return rune(' ')
		}, line))
		for _, number := range numbers {
			position := getNumberPosition(&line, number)
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			switch {
			case hasAdjacentSymbolOnLine(lines, i, position):
				sum += n
			case i > 0 && hasAdjacentSymbolOnLine(lines, i-1, position):
				sum += n
			case i < len(lines)-1 && hasAdjacentSymbolOnLine(lines, i+1, position):
				sum += n
			}
		}
	}

	return sum
}

type position struct {
	Min int
	Max int
}

func (position *position) containsIndex(index int) bool {
	return index >= position.Min && index <= position.Max
}

func getNumberPosition(line *string, number string) *position {
	min := utf8.RuneCountInString((*line)[:strings.Index(*line, number)])
	max := min + utf8.RuneCountInString(number) - 1
	if min > 0 {
		min--
	}
	if max < utf8.RuneCountInString(*line) {
		max++
	}
	*line = strings.Replace(*line, number, strings.Repeat(".", utf8.RuneCountInString(number)), 1)
	return &position{Min: min, Max: max}
}

func hasAdjacentSymbolOnLine(lines []string, lineIndex int, position *position) bool {
	symbolIndexes := []int{}
	for i, r := range lines[lineIndex] {
		if !unicode.IsDigit(r) && r != rune('.') {
			symbolIndexes = append(symbolIndexes, i)
		}
	}
	for _, s := range symbolIndexes {
		if position.containsIndex(s) {
			return true
		}
		continue
	}
	return false
}

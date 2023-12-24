package day2

import (
	"strconv"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	possibleGameSum := 0

	for game := range input.Lines() {
		parts := strings.Split(game, ": ")
		gameId, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
		if err != nil {
			panic(err)
		}
		sets := strings.Split(parts[1], ";")
		possibleSets := 0
		for _, set := range sets {
			colorStrings := strings.Split(set, ", ")
			possibleColors := 0
			for _, colorString := range colorStrings {
				color := strings.Fields(strings.TrimSpace(colorString))
				number, err := strconv.Atoi(color[0])
				if err != nil {
					panic(err)
				}
				if number <= limits[color[1]] {
					possibleColors++
				}
			}
			if possibleColors == len(colorStrings) {
				possibleSets++
			}
		}
		if possibleSets == len(sets) {
			possibleGameSum += gameId
		}
	}

	return possibleGameSum
}

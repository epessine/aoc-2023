package day2

import (
	"strconv"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	gamePowerSum := 0

	for line := range input.Lines() {
		minimumSet := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		sets := strings.Split(strings.Split(line, ": ")[1], ";")
		for _, set := range sets {
			colorStrings := strings.Split(set, ", ")
			for _, colorString := range colorStrings {
				color := strings.Fields(strings.TrimSpace(colorString))
				colorName := color[1]
				number, err := strconv.Atoi(color[0])
				if err != nil {
					panic(err)
				}
				if number > minimumSet[colorName] {
					minimumSet[colorName] = number
				}
			}
		}
		gamePower := minimumSet["red"] * minimumSet["blue"] * minimumSet["green"]
		gamePowerSum += gamePower
	}

	return gamePowerSum
}

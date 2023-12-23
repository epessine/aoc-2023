package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solveA() int64 {
	file, err := os.Open("day2/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	limits := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var possibleGameSum int64 = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ": ")
		gameId, err := strconv.ParseInt(strings.Replace(splitted[0], "Game ", "", 1), 10, 8)
		if err != nil {
			panic(err)
		}
		sets := strings.Split(splitted[1], ";")
		possibleSets := 0
		for _, set := range sets {
			colorStrings := strings.Split(set, ", ")
			possibleColors := 0
			for _, colorString := range colorStrings {
				color := strings.Split(strings.TrimSpace(colorString), " ")
				colorName := color[1]
				number, err := strconv.ParseInt(color[0], 10, 8)
				if err != nil {
					panic(err)
				}
				if number <= limits[colorName] {
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

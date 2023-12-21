package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	file, err := os.Open("day-two/input.txt")

	if err != nil {
		log.Fatal(err)
		return
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
			continue
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
					log.Panic(err)
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

	log.Printf("part 1 result is %d", possibleGameSum)
}

func two() {
	file, err := os.Open("day-two/input.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	var gamePowerSum int64 = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			continue
		}
		minimumSet := map[string]int64{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		sets := strings.Split(strings.Split(line, ": ")[1], ";")
		for _, set := range sets {
			colorStrings := strings.Split(set, ", ")
			for _, colorString := range colorStrings {
				color := strings.Split(strings.TrimSpace(colorString), " ")
				colorName := color[1]
				number, err := strconv.ParseInt(color[0], 10, 8)
				if err != nil {
					log.Panic(err)
				}
				if number > minimumSet[colorName] {
					minimumSet[colorName] = number
				}
			}
		}
		gamePower := minimumSet["red"] * minimumSet["blue"] * minimumSet["green"]
		gamePowerSum += gamePower
	}

	log.Printf("part 2 result is %d", gamePowerSum)
}

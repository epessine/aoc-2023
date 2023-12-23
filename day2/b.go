package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solveB() int64 {
	file, err := os.Open("day2/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var gamePowerSum int64 = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
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

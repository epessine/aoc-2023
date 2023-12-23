package day1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solveB() int64 {
	file, err := os.Open("day1/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64 = 0
	regex := "([1-9]|(?:(one|two|three|four|five|six|seven|eight|nine)))"

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		numbers := regexp.MustCompile(regex).FindAllString(parsedLine, -1)
		parsedNumbers := replaceSpelledOutWithDigit(numbers)
		parsedDigits, _ := strconv.ParseInt(parsedNumbers[0]+parsedNumbers[len(parsedNumbers)-1], 10, 32)
		total += parsedDigits
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func parseLine(line string) string {
	var parsedLine string
	parsedLine = strings.ReplaceAll(line, "one", "onee")
	parsedLine = strings.ReplaceAll(parsedLine, "two", "twoo")
	parsedLine = strings.ReplaceAll(parsedLine, "three", "threee")
	parsedLine = strings.ReplaceAll(parsedLine, "four", "fourr")
	parsedLine = strings.ReplaceAll(parsedLine, "five", "fivee")
	parsedLine = strings.ReplaceAll(parsedLine, "six", "sixx")
	parsedLine = strings.ReplaceAll(parsedLine, "seven", "sevenn")
	parsedLine = strings.ReplaceAll(parsedLine, "eight", "eightt")
	parsedLine = strings.ReplaceAll(parsedLine, "nine", "ninee")
	return parsedLine
}

func replaceSpelledOutWithDigit(items []string) []string {
	newItems := make([]string, len(items))

	for index, item := range items {
		switch item {
		case "one":
			newItems[index] = "1"
		case "two":
			newItems[index] = "2"
		case "three":
			newItems[index] = "3"
		case "four":
			newItems[index] = "4"
		case "five":
			newItems[index] = "5"
		case "six":
			newItems[index] = "6"
		case "seven":
			newItems[index] = "7"
		case "eight":
			newItems[index] = "8"
		case "nine":
			newItems[index] = "9"
		default:
			newItems[index] = item
		}
	}

	return newItems
}

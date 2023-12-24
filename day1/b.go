package day1

import (
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	regex := "([1-9]|(?:(one|two|three|four|five|six|seven|eight|nine)))"
	result := make(chan int, 1)
	defer close(result)
	result <- 0
	wg := sync.WaitGroup{}
	for line := range input.Lines() {
		wg.Add(1)
		go func(line string) {
			parsedLine := parseLine(line)
			numbers := regexp.MustCompile(regex).FindAllString(parsedLine, -1)
			parsedNumbers := replaceSpelledOutWithDigit(numbers)
			parsedDigits, _ := strconv.Atoi(parsedNumbers[0] + parsedNumbers[len(parsedNumbers)-1])
			total := <-result
			result <- total + parsedDigits
			wg.Done()
		}(line)
	}
	wg.Wait()
	return <-result
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

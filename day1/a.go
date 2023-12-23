package day1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func solveA() int64 {
	file, err := os.Open("day1/input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := regexp.MustCompile("[1-9]").FindAllString(line, -1)
		parsedDigits, _ := strconv.ParseInt(numbers[0]+numbers[len(numbers)-1], 10, 32)
		total += parsedDigits
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

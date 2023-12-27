package day6

import (
	"strconv"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	races := []Race{}
	for line := range input.Lines() {
		if strings.HasPrefix(line, "Time:") {
			durations := strings.Fields(strings.TrimSpace(strings.TrimPrefix(line, "Time:")))
			for _, duration := range durations {
				d, err := strconv.Atoi(duration)
				if err != nil {
					panic(err)
				}
				races = append(races, Race{Duration: d})
			}
		}
		if strings.HasPrefix(line, "Distance:") {
			records := strings.Fields(strings.TrimSpace(strings.TrimPrefix(line, "Distance:")))
			for i, record := range records {
				r, err := strconv.Atoi(record)
				if err != nil {
					panic(err)
				}
				races[i].Record = r
			}
		}
	}
	ways := 0
	for _, race := range races {
		if ways == 0 {
			ways += race.WaysToWin()
			continue
		}
		ways = ways * race.WaysToWin()
	}
	return ways
}

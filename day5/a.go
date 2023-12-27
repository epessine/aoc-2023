package day5

import (
	"slices"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	var seeds []Seed
	maps := []Map{}
	pm := Map{}
	i := 0
	for line := range input.Lines() {
		switch {
		case i == 0:
			seeds = getSeedsFromLine(line)
		case line == "":
		case strings.HasSuffix(line, "map:"):
			if len(pm.Ranges) != 0 {
				maps = append(maps, pm)
			}
			pm = Map{}
		default:
			pm.addRange(line)
		}
		i++
	}
	if len(pm.Ranges) != 0 {
		maps = append(maps, pm)
	}
	for _, seed := range seeds {
		seed.populate(maps)
	}
	min := slices.MinFunc(seeds, func(a, b Seed) int {
		if a[7] == b[7] {
			return 0
		}
		if a[7] > b[7] {
			return 1
		}
		return -1
	})
	return min[7]
}

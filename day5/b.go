package day5

import (
	"strings"
	"sync"

	"github.com/epessine/aoc-2023/challenge"
)

func solveB(input *challenge.Input) int {
	var seeds <-chan Seed
	maps := []Map{}
	pm := Map{}
	i := 0
	for line := range input.Lines() {
		switch {
		case i == 0:
			seeds = getSeedsFromRangeLine(line)
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
	workers := 100
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}
	min := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for seed := range seeds {
				seed.populate(maps)
				mutex.RLock()
				if min == 0 || seed[7] < min {
					min = seed[7]
				}
				mutex.RUnlock()
			}
		}()
	}
	wg.Wait()
	return min
}

package day5

import (
	"strconv"
	"strings"
	"sync"
)

type Seed []int

func (s Seed) populate(ms []Map) {
	for i, m := range ms {
		s[i+1] = m.getDestination(s[i])
	}
}

func getSeedsFromLine(line string) []Seed {
	seeds := []Seed{}
	for _, f := range strings.Fields(strings.TrimPrefix(line, "seeds: ")) {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		s := make(Seed, 8)
		s[0] = n
		seeds = append(seeds, s)
	}
	return seeds
}

func getSeedsFromRangeLine(line string) <-chan Seed {
	ns := make([]int, 20)
	for _, f := range strings.Fields(strings.TrimPrefix(line, "seeds: ")) {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		ns = append(ns, n)
	}
	seeds := make(chan Seed)
	go func() {
		wg := sync.WaitGroup{}
		for i, n := range ns {
			wg.Add(1)
			go func(i int, n int) {
				defer wg.Done()
				if i%2 == 0 {
					for sn := n; sn < n+ns[i+1]; sn++ {
						s := make(Seed, 8)
						s[0] = sn
						seeds <- s
					}
				}
			}(i, n)
		}
		wg.Wait()
		close(seeds)
	}()
	return seeds
}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Map struct {
	Ranges []Range
}

func (m *Map) getDestination(source int) int {
	for _, r := range m.Ranges {
		srs := r.SourceRangeStart
		sre := srs + r.RangeLength - 1
		if source >= srs && source <= sre {
			return r.DestinationRangeStart + (source - srs)
		}
	}
	return source
}

func (m *Map) addRange(line string) {
	ranges := strings.Fields(line)
	drs, err := strconv.Atoi(ranges[0])
	if err != nil {
		panic(err)
	}
	srs, err := strconv.Atoi(ranges[1])
	if err != nil {
		panic(err)
	}
	rl, err := strconv.Atoi(ranges[2])
	if err != nil {
		panic(err)
	}
	m.Ranges = append(m.Ranges, Range{
		DestinationRangeStart: drs,
		SourceRangeStart:      srs,
		RangeLength:           rl,
	})
}

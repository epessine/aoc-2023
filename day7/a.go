package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

func solveA(input *challenge.Input) int {
	hands := Hands{}
	for line := range input.Lines() {
		fields := strings.Fields(line)
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{Bid: bid, Cards: []rune(fields[0])}
		hand.setType()
		hands = append(hands, hand)
	}
	sort.Sort(hands)
	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += hand.Bid * rank
	}
	return total
}

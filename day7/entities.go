package day7

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/epessine/aoc-2023/challenge"
)

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var (
	CardValues map[rune]int
	JokerRule  bool
)

type Hand struct {
	Bid   int
	Cards []rune
	Type  HandType
}

func parseInput(i *challenge.Input) Hands {
	hands := Hands{}
	for line := range i.Lines() {
		fields := strings.Fields(line)
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{Bid: bid, Cards: []rune(fields[0])}
		hand.setType(JokerRule)
		hands = append(hands, hand)
	}
	sort.Sort(hands)
	return hands
}

func (h *Hand) setType(jokerRule bool) {
	h.matchType()
	if jokerRule {
		h.applyJokerRule()
	}
}

func (h *Hand) unique() int {
	c2 := make([]rune, len(h.Cards))
	copy(c2, h.Cards)
	slices.Sort(c2)
	return len(slices.Compact(c2))
}

func (h *Hand) matchType() {
	unique := h.unique()
	c := make([]rune, len(h.Cards))
	copy(c, h.Cards)
	slices.Sort(c)
	switch {
	case unique == 1:
		h.Type = FiveOfKind
	case unique == 2 && c[1] == c[3]:
		h.Type = FourOfKind
	case unique == 2 && c[1] != c[3]:
		h.Type = FullHouse
	case unique == 3 && ((c[0] == c[2]) || (c[1] == c[3]) || (c[2] == c[4])):
		h.Type = ThreeOfAKind
	case unique == 3 && !((c[0] == c[2]) || (c[1] == c[3]) || (c[2] == c[4])):
		h.Type = TwoPair
	case unique == 4:
		h.Type = OnePair
	case unique == 5:
		h.Type = HighCard
	}
}

func (h *Hand) applyJokerRule() {
	jokers := len([]rune(strings.Map(func(r rune) rune {
		if r == 'J' {
			return r
		}
		return -1
	}, string(h.Cards))))
	if jokers == 0 {
		return
	}
	switch h.Type {
	case HighCard:
		h.Type = OnePair
	case OnePair:
		h.Type = ThreeOfAKind
	case TwoPair:
		switch jokers {
		case 2:
			h.Type = FourOfKind
		case 1:
			h.Type = FullHouse
		}
	case ThreeOfAKind:
		h.Type = FourOfKind
	case FullHouse, FourOfKind:
		h.Type = FiveOfKind
	}
}

type Hands []Hand

func (hs Hands) Len() int {
	return len(hs)
}
func (hs Hands) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func (hs Hands) Less(i, j int) bool {
	if hs[i].Type == hs[j].Type {
		for c := 0; c < len(hs[i].Cards); c++ {
			if hs[i].Cards[c] == hs[j].Cards[c] {
				continue
			}
			return CardValues[hs[i].Cards[c]] < CardValues[hs[j].Cards[c]]
		}
	}
	return hs[i].Type < hs[j].Type
}

func (hs Hands) totalWinnings() int {
	total := 0
	for i, hand := range hs {
		rank := i + 1
		total += hand.Bid * rank
	}
	return total
}

package day7

import (
	"slices"
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

var CardValues = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0}

type Hand struct {
	Bid   int
	Cards []rune
	Type  HandType
}

func (h *Hand) setType() {
	c := make([]rune, len(h.Cards))
	c2 := make([]rune, len(h.Cards))
	copy(c, h.Cards)
	copy(c2, h.Cards)
	slices.Sort(c)
	slices.Sort(c2)
	uc := len(slices.Compact(c2))
	switch {
	case uc == 1:
		h.Type = FiveOfKind
	case uc == 2 && c[1] == c[3]:
		h.Type = FourOfKind
	case uc == 2 && c[1] != c[3]:
		h.Type = FullHouse
	case uc == 3 && ((c[0] == c[2]) || (c[1] == c[3]) || (c[2] == c[4])):
		h.Type = ThreeOfAKind
	case uc == 3 && !((c[0] == c[2]) || (c[1] == c[3]) || (c[2] == c[4])):
		h.Type = TwoPair
	case uc == 4:
		h.Type = OnePair
	case uc == 5:
		h.Type = HighCard
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

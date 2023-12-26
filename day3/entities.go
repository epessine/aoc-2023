package day3

import (
	"strconv"
	"strings"
)

type Number struct {
	Value   int
	X       [2]int
	Y       [2]int
	Builder strings.Builder
}

func (n *Number) build(r rune, x int, y int) {
	if n.Builder.Len() == 0 {
		n.X[0] = x - 1
		n.Y[0] = y - 1
		n.Y[1] = y + 1
	}
	_, err := n.Builder.WriteRune(r)
	if err != nil {
		panic(err)
	}
}

func (n *Number) complete(x int) Number {
	var err error
	n.Value, err = strconv.Atoi(n.Builder.String())
	if err != nil {
		panic(err)
	}
	n.X[1] = x
	return *n
}

func (n *Number) isAdjacentTo(s Symbol) bool {
	xc := s.X >= n.X[0] && s.X <= n.X[1]
	yc := s.Y >= n.Y[0] && s.Y <= n.Y[1]
	return xc && yc
}

type Symbol struct {
	X int
	Y int
}

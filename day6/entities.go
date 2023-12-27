package day6

type Race struct {
	Duration int
	Record   int
}

func (r *Race) WaysToWin() int {
	ways := 0
	for speed := 1; speed < r.Duration; speed++ {
		distance := speed * (r.Duration - speed)
		if distance > r.Record {
			ways++
		}
	}
	return ways
}

package util

type Point struct {
	X, Y int
}

func (p Point) Add(p2 Point) Point {
	return Point{X: p2.X + p.X, Y: p2.Y + p.Y}
}

func Manhattan(p1, p2 Point) int {
	return IntAbs(p2.X - p1.X) + IntAbs(p2.Y - p1.Y)
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

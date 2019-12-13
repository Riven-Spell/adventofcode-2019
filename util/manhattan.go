package util

import (
	"math"
)

type Point struct {
	X, Y int64
}

func (p Point) Add(p2 Point) Point {
	return Point{X: p2.X + p.X, Y: p2.Y + p.Y}
}

func (p Point) Sub(p2 Point) Point {
	return Point{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func Manhattan(p1, p2 Point) int64 {
	return IntAbs(p2.X - p1.X) + IntAbs(p2.Y - p1.Y)
}

func Slope(p1, p2 Point) Fraction {
	f := Fraction{Numerator: p2.Y - p1.Y, Denominator: p2.X - p1.X}.Simplify()

	if f.Numerator == 0 {
		f.Denominator = TernaryInt64(f.Denominator > 0, math.MaxInt64, math.MinInt64)
	}

	if f.Denominator == 0 {
		f.Numerator = TernaryInt64(f.Numerator > 0, math.MaxInt64, math.MinInt64)
	}

	return f
}

type DistSorter struct {
	Center Point
	List []Point
}

func (s *DistSorter) Len() int {
	return len(s.List)
}

func (s *DistSorter) Less(i, j int) bool {
	return Manhattan(s.Center, s.List[i]) < Manhattan(s.Center, s.List[j])
}

func (s *DistSorter) Swap(i, j int) {
	s.List[i], s.List[j] = s.List[j], s.List[i]
}

func IntAbs(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

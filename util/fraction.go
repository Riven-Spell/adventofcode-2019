package util

import (
	"math"
)

type Fraction struct {
	Numerator, Denominator int64
}

func (f Fraction) Simplify() Fraction {
	gcd := GCD(f.Numerator, f.Denominator)

	num := f.Numerator / gcd
	den := f.Denominator / gcd

	return Fraction{Numerator:num, Denominator:den}
}

func (f Fraction) GetAngle() float64 {
	angle := math.Atan2(float64(f.Numerator), float64(f.Denominator)) * 180 / math.Pi

	return TernaryFloat64(angle > 0, angle, angle + 360)
}

func GCD(x1, x2 int64) int64 {
	gcd := int64(1)
	x1, x2 = IntAbs(x1), IntAbs(x2)
	for i := int64(1); i <= x1 && i <= x2; i++ {
		if x1 % i == 0 && x2 % i == 0 {
			gcd = i
		}
	}

	return gcd
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

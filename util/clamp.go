package util

func ClampInt64(n, min, max int64) int64 {
	if n < min {
		return min
	} else if n > max {
		return max
	} else {
		return n
	}
}

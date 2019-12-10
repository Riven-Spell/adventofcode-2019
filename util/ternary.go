package util

func TernaryString(condition bool, tval, fval string) string {
	if condition {
		return tval
	}
	return fval
}

func TernaryInt64(condition bool, tval, fval int64) int64 {
	if condition {
		return tval
	}
	return fval
}

func TernaryFloat64(condition bool, tval, fval float64) float64 {
	if condition {
		return tval
	}
	return fval
}
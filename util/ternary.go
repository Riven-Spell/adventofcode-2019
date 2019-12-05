package util

func TernaryString(condition bool, tval, fval string) string {
	if condition {
		return tval
	}
	return fval
}

func TernaryInteger(condition bool, tval, fval int) int {
	if condition {
		return tval
	}
	return fval
}
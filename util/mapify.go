package util

func MapifyIntList(in []int) map[int]int {
	out := make(map[int]int)

	for k,v := range in {
		out[k] = v
	}

	return out
}

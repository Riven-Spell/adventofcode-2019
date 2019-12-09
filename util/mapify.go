package util

func MapifyIntList(in []int64) map[int64]int64 {
	out := make(map[int64]int64)

	for k,v := range in {
		out[int64(k)] = v
	}

	return out
}

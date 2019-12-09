package util

import (
	"math"
	"strconv"
	"strings"
)

func ByDigit(x int64) (out []int64) {
	out = make([]int64, 0)

	for _,v := range strings.Split(strconv.Itoa(int(x)), "") {
		x, _ := strconv.ParseInt(v, 10, 64)
		out = append(out, x)
	}

	return
}

func DigitsToInt(x []int64) (out int64) {
	for k,v := range x {
		digit := int64(math.Pow10(len(x) - (k+1)))

		out += int64(v) * digit
	}

	return
}
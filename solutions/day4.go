package solutions

import (
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
)

type Day4Input struct{
	Max, Min int64
}

func (s *Day4Input) Prepare(input string) {
	i := strings.Split(input, "-")

	var err error
	s.Min, err = strconv.ParseInt(i[0], 10, 64)
	util.PanicIfErr(err)
	s.Max, err = strconv.ParseInt(i[1], 10, 64)
	util.PanicIfErr(err)
}

func (s *Day4Input) IncrementCode(code []int64) {
	code[len(code)-1]++

	for i := len(code)-1; i >= 0; i-- {
		if code[i] > 9 {
			code[i] = 0
			code[i-1]++
		} else {
			return
		}
	}
}

func (s *Day4Input) MakePerfect(code []int64) {
	highestDigit := int64(0)
	location := 0
	for k, v := range code {
		if v > highestDigit {
			highestDigit = v
			location = k

			if k+1 == len(code) || code[k+1] < highestDigit {
				break
			}
		}
	}

	for i := location + 1; i < len(code); i++ {
		code[i] = highestDigit
	}
}

func (s *Day4Input) Verify(code []int64, onlyTwo bool) bool {
	highest := int64(0)
	last := int64(-1)
	double := false
	runLength := 0
	for _,v := range code {
		if v == last {
			runLength++

			if !onlyTwo && runLength == 2 {
				double = true
			}
		} else {
			if runLength == 2 {
				double = true
			}

			runLength = 1
		}

		if v > highest {
			highest = v
		} else if v < highest {
			return false
		}

		last = v
	}

	if runLength == 2 {
		return true
	}

	return double
}

func (s *Day4Input) Part1() string {
	// First, let's segment our minimum input.
	minDig := util.ByDigit(s.Min)
	maxDig := util.ByDigit(s.Max)

	s.MakePerfect(minDig)

	valid := 0
	for util.DigitsToInt(minDig) < util.DigitsToInt(maxDig) {

		if s.Verify(minDig, false) {
			valid++
		}

		s.IncrementCode(minDig)
		s.MakePerfect(minDig)
	}

	return strconv.Itoa(valid)
}

func (s *Day4Input) Part2() string {
	// First, let's segment our minimum input.
	minDig := util.ByDigit(s.Min)
	maxDig := util.ByDigit(s.Max)

	s.MakePerfect(minDig)

	valid := 0
	for util.DigitsToInt(minDig) < util.DigitsToInt(maxDig) {

		if s.Verify(minDig, true) {
			valid++
		}

		s.IncrementCode(minDig)
		s.MakePerfect(minDig)
	}

	return strconv.Itoa(valid)
}
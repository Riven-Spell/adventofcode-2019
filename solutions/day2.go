package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
)

type Day2Input struct{
	Memory       map[int]int // This is actually just acting as an unbound array.
	baseStr      string
	CustNounVerb bool // Set to true during preparation, considered when running parts.
	Noun int // Default set to 12
	Verb int // Default set to 2
}

func (s *Day2Input) Prepare(input string) {
	s.Memory = make(map[int]int)
	s.baseStr = input

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.Memory[k] = int(pi)
	}

	s.CustNounVerb = true
	s.Noun = 12
	s.Verb = 2
}

func (s *Day2Input) Add(from1, from2, to int) {
	s.Memory[to] = s.Memory[from1] + s.Memory[from2]
}

func (s *Day2Input) Mul(from1, from2, to int) {
	s.Memory[to] = s.Memory[from1] * s.Memory[from2]
}

func (s *Day2Input) Part1() string {
	if s.CustNounVerb {
		s.Memory[1] = s.Noun
		s.Memory[2] = s.Verb
	}

	for i := 0; s.Memory[i] != 99; i += 4 {
		switch s.Memory[i] {
		case 1:
			s.Add(s.Memory[i+1], s.Memory[i+2], s.Memory[i+3])
		case 2:
			s.Mul(s.Memory[i+1],s.Memory[i+2],s.Memory[i+3])
		case 99:
			fmt.Println("Intcode didn't expect to run 99")
		default:
			panic(strconv.Itoa(s.Memory[i]) + " is not a valid opcode")
		}
	}

	return strconv.Itoa(s.Memory[0])
}

func (s *Day2Input) Part2() string {
	noun, verb := 0, 0

	increment := func() {
		noun++

		if noun > 99 {
			verb++
			noun = 0
		}

		if verb > 99 {
			panic("solution should have been found already!")
		}
	}

	reset := func() {
		s.Prepare(s.baseStr)
		increment()
		s.Noun = noun
		s.Verb = verb
	}

	for ; ; reset() {
		s.Part1()
		if s.Memory[0] == 19690720 {
			break
		}
	}

	return strconv.Itoa((100 * noun) + verb)
}
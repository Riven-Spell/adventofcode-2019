package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
)

type Day5Input struct{
	intcode.VM
}

func (s *Day5Input) Prepare(input string) {
	s.VM = intcode.VM{}
	s.Memory = make(map[int]int)

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.Memory[k] = int(pi)
	}
}

func (s *Day5Input) Part1() string {
	// input the ID needed, reset stdout
	s.Ioutil.Stdin = []int{1}
	s.Ioutil.Stdout = []int{}
	// Blacklist all part 2 functions
	s.BlacklistedOps = map[int]bool{
		5:true,
		6:true,
		7:true,
		8:true,
	}

	s.Autorun()

	return fmt.Sprint(s.Ioutil.Stdout[len(s.Ioutil.Stdout)-1])
}

func (s *Day5Input) Part2() string {
	// input the ID needed, reset stdout
	s.Ioutil.Stdin = []int{5}
	s.Ioutil.Stdout = []int{}

	s.Autorun()

	return fmt.Sprint(s.Ioutil.Stdout[len(s.Ioutil.Stdout)-1])
}
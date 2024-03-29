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
	s.Memory = make(map[int64]int64)

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.Memory[int64(k)] = pi
	}

	s.IoMgr = &intcode.PreparedIO{}
}

func (s *Day5Input) Part1() string {
	io := s.IoMgr.(*intcode.PreparedIO)
	// input the ID needed, reset stdout
	io.Stdin = []int64{1}
	io.Stdout = []int64{}
	// Blacklist all part 2 functions
	s.BlacklistedOps = map[int64]bool{
		5:true,
		6:true,
		7:true,
		8:true,
	}

	s.Autorun()

	return fmt.Sprint(io.Stdout[len(io.Stdout)-1])
}

func (s *Day5Input) Part2() string {
	io := s.IoMgr.(*intcode.PreparedIO)
	// input the ID needed, reset stdout
	io.Stdin = []int64{5}
	io.Stdout = []int64{}

	s.Autorun()

	return fmt.Sprint(io.Stdout[len(io.Stdout)-1])
}
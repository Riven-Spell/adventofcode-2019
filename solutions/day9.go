package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
)

type Day9Input struct{
	baseMem map[int64]int64
}

func (s *Day9Input) Prepare(input string) {
	s.baseMem = make(map[int64]int64)

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.baseMem[int64(k)] = pi
	}
}

func (s *Day9Input) Part1() string {
	vm := intcode.VM{}
	vm.Memory = s.baseMem

	vm.IoMgr = &intcode.PreparedIO{
		Stdout: make([]int64, 0),
		Stdin: []int64{1},
	}

	vm.Autorun()

	return fmt.Sprint(vm.IoMgr.(*intcode.PreparedIO).Stdout[0])
}


func (s *Day9Input) Part2() string {
	vm := intcode.VM{}
	vm.Memory = s.baseMem

	vm.IoMgr = &intcode.PreparedIO{
		Stdout: make([]int64, 0),
		Stdin: []int64{2},
	}

	vm.Autorun()

	return fmt.Sprint(vm.IoMgr.(*intcode.PreparedIO).Stdout[0])
}

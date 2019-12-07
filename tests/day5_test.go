package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	chk "gopkg.in/check.v1"
)

type Day5TestSuite struct{}
var _ = chk.Suite(&Day5TestSuite{})

func (s *Day5TestSuite) TestDay5VMExtensions(c *chk.C) {
	inputs := []struct{
		input int
		memory []int
		result int
	}{
		{ // Position mode EQ
			input: 3,
			memory: []int{3,9,8,9,10,9,4,9,99,-1,8},
			result: 0,
		},
		{
			input: 8,
			memory: []int{3,9,8,9,10,9,4,9,99,-1,8},
			result: 1,
		},

		{ // Position mode LT
			input: 3,
			memory: []int{3,9,7,9,10,9,4,9,99,-1,8},
			result: 1,
		},
		{
			input: 8,
			memory: []int{3,9,7,9,10,9,4,9,99,-1,8},
			result: 0,
		},
		{
			input: 9,
			memory: []int{3,9,7,9,10,9,4,9,99,-1,8},
			result: 0,
		},

		{ // Immediate mode EQ
			input: 3,
			memory: []int{3,3,1108,-1,8,3,4,3,99},
			result: 0,
		},
		{
			input: 8,
			memory: []int{3,3,1108,-1,8,3,4,3,99},
			result: 1,
		},

		{ // Immediate mode LT
			input: 3,
			memory: []int{3,3,1107,-1,8,3,4,3,99},
			result: 1,
		},
		{
			input: 8,
			memory: []int{3,3,1107,-1,8,3,4,3,99},
			result: 0,
		},
		{
			input: 9,
			memory: []int{3,3,1107,-1,8,3,4,3,99},
			result: 0,
		},

		{ // Position mode JEZ
			input: 0,
			memory: []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9},
			result: 0,
		},
		{
			input: 2,
			memory: []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9},
			result: 1,
		},

		{ // Position mode JNZ
			input: 0,
			memory: []int{3,12,5,12,15,1,13,14,13,4,13,99,-1,0,1,9},
			result: 1,
		},
		{
			input: 2,
			memory: []int{3,12,5,12,15,1,13,14,13,4,13,99,-1,0,1,9},
			result: 0,
		},
	}

	for _,v := range inputs {
		vm := intcode.VM{Memory:util.MapifyIntList(v.memory), IoMgr: &intcode.PreparedIO{Stdin: []int{v.input}}}

		vm.Autorun()

		io := vm.IoMgr.(*intcode.PreparedIO)

		c.Log("Got output chain: ", io.Stdout, " Expected only output: ", v.result)
		c.Assert(io.Stdout[0], chk.Equals, v.result)
	}
}
package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	chk "gopkg.in/check.v1"
)

type Day9TestSuite struct{}
var _ = chk.Suite(&Day9TestSuite{})

// This test doesn't fail yet.
func (s * Day9TestSuite) TestDay9Part1Additions(c *chk.C) {
	inputs := []struct{
		memory []int64
		expectedStdout []int64
	}{
		{
			memory: []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
			expectedStdout: []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
		},
		{
			memory: []int64{1102,34915192,34915192,7,4,7,99,0},
			expectedStdout: []int64{1219070632396864},
		},
		{
			memory: []int64{104,1125899906842624,99},
			expectedStdout: []int64{1125899906842624},
		},
	}

	for _,v := range inputs {
		vm := intcode.VM{
			Memory:         util.MapifyIntList(v.memory),
			IoMgr:          &intcode.PreparedIO{},
		}

		vm.Autorun()

		c.Assert(vm.IoMgr.(*intcode.PreparedIO).Stdout, chk.DeepEquals, v.expectedStdout)
	}
}

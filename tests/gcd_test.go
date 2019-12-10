package tests

import (
	"github.com/Virepri/adventofcode-2019/util"
	chk "gopkg.in/check.v1"
)

type UtilTestSuite struct{}
var _ = chk.Suite(&UtilTestSuite{})

func (s *UtilTestSuite) TestGCD(c *chk.C) {
	inputs := []struct{
		x1, x2, result int64
	}{
		{
			x1:8,
			x2:12,
			result:4,
		},
		{
			x1: 20,
			x2: 15,
			result:5,
		},
		{
			x1: 6,
			x2: 12,
			result: 6,
		},
		{
			x1: 2,
			x2: 10,
			result: 2,
		},
		{
			x1: -2,
			x2: -4,
			result: 2,
		},
	}

	for _,v := range inputs {
		c.Assert(util.GCD(v.x1, v.x2), chk.Equals, v.result)
	}
}

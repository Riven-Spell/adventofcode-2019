package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day2TestSuite struct{}
var _ = chk.Suite(&Day2TestSuite{})

func (s *Day2TestSuite) TestPart1(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input: "1,9,10,3,2,3,11,0,99,30,40,50",
			result: "3500",
		},
		{
			input: "1,0,0,0,99",
			result: "2",
		},
		{
			input: "2,3,0,3,99",
			result: "2",
		},
		{
			input: "2,4,4,5,99,0",
			result: "2",
		},
		{
			input: "1,1,1,4,99,5,6,0,99",
			result: "30",
		},
	}

	for _,v := range inputs {
		input := solutions.Day2Input{}
		input.Prepare(v.input)
		input.CustNounVerb = false

		out := input.Part1()
		c.Log("Got: ", out, " Expected: ", v.result)
		c.Assert(out, chk.Equals, v.result)
	}
}
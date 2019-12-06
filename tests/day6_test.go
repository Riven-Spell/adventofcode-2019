package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day6TestSuite struct {}
var _ = chk.Suite(&Day6TestSuite{})

func (s *Day6TestSuite) TestPart1(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input:`COM)B
					B)C
					C)D
					D)E
					E)F
					B)G
					G)H
					D)I
					E)J
					J)K
					K)L`,
			result:"42",
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day6Input{}

		dummyInput.Prepare(v.input)

		c.Assert(dummyInput.Part1(), chk.Equals, v.result)
	}
}

func (s *Day6TestSuite) TestPart2(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input:`COM)B
					B)C
					C)D
					D)E
					E)F
					B)G
					G)H
					D)I
					E)J
					J)K
					K)L
					K)YOU
					I)SAN`,
			result:"4",
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day6Input{}

		dummyInput.Prepare(v.input)

		c.Assert(dummyInput.Part2(), chk.Equals, v.result)
	}
}
package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day12TestSuite struct{}
var _ = chk.Suite(&Day12TestSuite{})

func (s *Day12TestSuite) TestDay12Part1(c *chk.C) {
	inputs := []struct{
		input string
		stepCount int
		output string
	}{
		{
			input: `<x=-1, y=0, z=2>
					<x=2, y=-10, z=-7>
					<x=4, y=-8, z=8>
					<x=3, y=5, z=-1>`,
			stepCount: 10,
			output: `179`,
		},
		{
			input: `<x=-8, y=-10, z=0>
					<x=5, y=5, z=10>
					<x=2, y=-7, z=3>
					<x=9, y=-8, z=-3>`,
			stepCount: 100,
			output: `1940`,
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day12Input{}
		dummyInput.Prepare(v.input)
		dummyInput.StepCount = v.stepCount

		c.Assert(dummyInput.Part1(), chk.Equals, v.output)
	}
}

func (s *Day12TestSuite) TestDay12Part2(c *chk.C) {
	inputs := []struct{
		input string
		output string
	}{
		{
			input: `<x=-1, y=0, z=2>
					<x=2, y=-10, z=-7>
					<x=4, y=-8, z=8>
					<x=3, y=5, z=-1>`,
			output: `2772`,
		},
		{
			input: `<x=-8, y=-10, z=0>
					<x=5, y=5, z=10>
					<x=2, y=-7, z=3>
					<x=9, y=-8, z=-3>`,
			output: `4686774924`,
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day12Input{}
		dummyInput.Prepare(v.input)

		c.Assert(dummyInput.Part2(), chk.Equals, v.output)
	}
}
package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day3TestSuite struct{}
var _ = chk.Suite(&Day3TestSuite{})

func (s *Day3TestSuite) TestPart1(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input: `R8,U5,L5,D3
					U7,R6,D4,L4`,
			result: "6",
		},
		{
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
					U62,R66,U55,R34,D71,R55,D58,R83`,
			result: "159",
		},
		{
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
					U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			result: "135",
		},
	}

	for _,v := range inputs {
		d3i := solutions.Day3Input{}
		d3i.Prepare(v.input)

		p1o := d3i.Part1()
		c.Log("Got: ", p1o, " Expected: ", v.result)
		c.Assert(p1o, chk.Equals, v.result)
	}
}

func (s *Day3TestSuite) TestPart2(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input: `R8,U5,L5,D3
					U7,R6,D4,L4`,
			result: "30",
		},
		{
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
					U62,R66,U55,R34,D71,R55,D58,R83`,
			result: "610",
		},
		{
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
					U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			result: "410",
		},
	}

	for _,v := range inputs {
		d3i := solutions.Day3Input{}
		d3i.Prepare(v.input)

		p2o := d3i.Part2()
		c.Log("Got: ", p2o, " Expected: ", v.result)
		c.Assert(p2o, chk.Equals, v.result)
	}
}
package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day1TestSuite struct{}
var _ = chk.Suite(&Day1TestSuite{})

// TODO: Make standard runs for each day's input.

func (s *Day1TestSuite) TestDay1Part1(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input:"12",
			result:"2",
		},
		{
			input:"14",
			result:"2",
		},
		{
			input:`12
14`,
			result:"4",
		},
		{
			input:"1969",
			result:"654",
		},
		{
			input: "100756",
			result: "33583",
		},
	}

	for _,v := range inputs {
		in := solutions.Day1Input{}
		in.Prepare(v.input)

		result := in.Part1()
		c.Log("Input:\n", v.input, "\nOutput: ", result, "\nExpected: ", v.result)
		c.Assert(result, chk.Equals, v.result)
	}
}

func (s *Day1TestSuite) TestDay1Part2(c *chk.C) {
	inputs := []struct{
		input string
		result string
	}{
		{
			input: "14",
			result: "2",
		},
		{
			input: "1969",
			result: "966",
		},
		{
			input: "100756",
			result: "50346",
		},
	}

	for _,v := range inputs {
		in := solutions.Day1Input{}
		in.Prepare(v.input)

		result := in.Part2()
		c.Log("Input:\n", v.input, "\nOutput: ", result, "\nExpected: ", v.result)
		c.Assert(result, chk.Equals, v.result)
	}
}
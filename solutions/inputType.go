package solutions

import "github.com/Virepri/adventofcode-2019/inputs"

// PerDayInput is an interface that exposes a simple structure:
// The runner should call Prepare() on it to prepare the input
// Then call the part 1 and part 2 functions if wanted.
type PerDayInput interface {
	Prepare(input string)
	Part1() string
	Part2() string
}

var RegisteredDays = []struct{
	DummyInput PerDayInput
	StringInput *string
	ExpectedOutputs []string
}{
	{ // Day 1
		DummyInput: &Day1Input{},
		StringInput: &inputs.Day1Data,
		ExpectedOutputs: []string{"3363033", "5041680"},
	},
	{ // Day 2
		DummyInput: &Day2Input{},
		StringInput: &inputs.Day2Data,
		ExpectedOutputs: []string{"3716250", "6472"},
	},
	{ // Day 3
		DummyInput: &Day3Input{},
		StringInput: &inputs.Day3Data,
		ExpectedOutputs: []string{"1195", "91518"},
	},
	{ // Day 4
		DummyInput: &Day4Input{},
		StringInput: &inputs.Day4Data,
		ExpectedOutputs: []string{"1675", "1142"},
	},
	{ // Day 5
		DummyInput: &Day5Input{},
		StringInput: &inputs.Day5Data,
		ExpectedOutputs: []string {"7157989", "7873292"},
	},
	{ // Day 6
		DummyInput: &Day6Input{},
		StringInput: &inputs.Day6Data,
		ExpectedOutputs: []string{"145250","274"},
	},
	{ // Day 7
		DummyInput: &Day7Input{},
		StringInput: &inputs.Day7Data,
		ExpectedOutputs: []string{"440880", "3745599"},
	},
}
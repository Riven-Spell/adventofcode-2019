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
}{
	{ // Day 1
		DummyInput: &Day1Input{},
		StringInput: &inputs.Day1Data,
	},
	{ // Day 2
		DummyInput: &Day2Input{},
		StringInput: &inputs.Day2Data,
	},
	{ // Day 3
		DummyInput: &Day3Input{},
		StringInput: &inputs.Day3Data,
	},
}
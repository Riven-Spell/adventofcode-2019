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
	{ // Day 8
		DummyInput: &Day8Input{},
		StringInput: &inputs.Day8Data,
		ExpectedOutputs: []string{"2080","disabled"},
	},
	{ // Day 9
		DummyInput: &Day9Input{},
		StringInput: &inputs.Day9Data,
		ExpectedOutputs: []string{"2453265701","80805"},
	},
	{ // Day 10
		DummyInput: &Day10Input{},
		StringInput: &inputs.Day10Data,
		ExpectedOutputs: []string{"284", "404"},
	},
	{ // Day 11
		DummyInput: &Day11Input{},
		StringInput: &inputs.Day11Data,
		ExpectedOutputs: []string{"2016", "disabled"},
	},
	{ // Day 12
		DummyInput: &Day12Input{},
		StringInput: &inputs.Day12Data,
		ExpectedOutputs: []string{"14907", "467081194429464"},
	},
	{ // Day 13
		DummyInput: &Day13Input{},
		StringInput: &inputs.Day13Data,
		ExpectedOutputs: []string{"420", "21651"},
	},
	{ // Day 14
		DummyInput: &Day14Input{},
		StringInput: &inputs.Day14Data,
		ExpectedOutputs: []string{"628586", "3209254"},
	},
}
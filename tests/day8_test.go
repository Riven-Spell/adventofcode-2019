package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	"github.com/Virepri/adventofcode-2019/util"
	chk "gopkg.in/check.v1"
)

type Day8TestSuite struct{}
var _ = chk.Suite(&Day8TestSuite{})

func (s *Day8TestSuite) TestDay8Part1(c *chk.C) {
	inputs := []struct{
		imgSize util.Point
		imgStr string
		result string
	}{
		{
			imgSize: util.Point{X: 5, Y: 1},
			imgStr: `10000100021102210022`,
			result: `4`,
		},
		{
			imgSize: util.Point{X: 10, Y: 1},
			imgStr: `1000010002112221002210000100021101210022`,
			result: `15`,
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day8Input{}
		dummyInput.ImgSize = v.imgSize

		dummyInput.Prepare(v.imgStr)

		c.Assert(dummyInput.Part1(), chk.Equals, v.result)
	}
}

func (s *Day8TestSuite) TestDay8Part2(c *chk.C) {
	inputs := []struct{
		imgSize util.Point
		imgStr string
		result string
	}{
		{
			imgSize: util.Point{X: 5, Y: 1},
			imgStr: `10000100021102210022`,
			result: "█    \n",
		},
		{
			imgSize: util.Point{X: 10, Y: 1},
			imgStr: `1000010002112221002210000100021101210020`,
			result: "█    █    \n",
		},
	}

	for _,v := range inputs {
		dummyInput := solutions.Day8Input{}
		dummyInput.ImgSize = v.imgSize

		dummyInput.Prepare(v.imgStr)

		c.Assert(dummyInput.Part2(), chk.Equals, v.result)
	}
}
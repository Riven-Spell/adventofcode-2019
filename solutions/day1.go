package solutions

import (
	"github.com/Virepri/adventofcode-2019/util"
	"math"
	"strconv"
	"strings"
)

type Day1Input struct {
	Input []int64
}

func (d *Day1Input) Prepare(input string) {
	d.Input = make([]int64, 0)

	for _,v := range strings.Split(input, "\n") {
		// Space is trimmed so that multi-line tests can be written for readability
		pi, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		util.PanicIfErr(err)

		d.Input = append(d.Input, pi)
	}
}

func (d *Day1Input) Part1() string {
	totalFuel := int64(0)

	for _,v := range d.Input {
		totalFuel += d.CalculateFuelCost(v)
	}

	return strconv.FormatInt(totalFuel, 10)
}

func (d *Day1Input) Part2() string {
	totalFuel := int64(0)

	for _,v := range d.Input {
		moduleReq := d.CalculateFuelCost(v)

		addedFC := d.CalculateFuelCost(moduleReq)
		for addedFC > 0 {
			moduleReq += addedFC
			addedFC = d.CalculateFuelCost(addedFC)
		}

		totalFuel += moduleReq
	}

	return strconv.FormatInt(totalFuel, 10)
}

func (d Day1Input) CalculateFuelCost(input int64) int64 {
	return util.ClampInt64((input / 3) - 2, 0, math.MaxInt64)
}
package tests

import (
	"github.com/Virepri/adventofcode-2019/solutions"
	chk "gopkg.in/check.v1"
)

type Day4TestSuite struct{}
var _ = chk.Suite(&Day4TestSuite{})

func (s *Day4TestSuite) TestDay4Validator(c *chk.C) {
	inputs := []struct{
		code int64
		valid bool
		exact bool
	}{
		{ // Part 1 validation tests
			code: 112233,
			valid: true,
		},
		{
			code: 111123,
			valid: true,
		},
		{
			code: 122406,
			valid: false,
		},
		{
			code: 123456,
			valid: false,
		},
		{ // Part 2 validation tests
			code: 112222,
			valid: true,
			exact: true,
		},
		{
			code: 122456,
			valid: true,
			exact: true,
		},
		{
			code: 222213,
			valid: false,
			exact: true,
		},
		{
			code: 222234,
			valid: false,
			exact: true,
		},
	}

	dummyInput := solutions.Day4Input{}
	for _,v := range inputs {
		c.Log("(exact: ", v.exact, ") Code ", v.code, " is valid? ", v.valid)

		isValid := dummyInput.Verify(dummyInput.ByDigit(v.code), v.exact)
		c.Log("Validator says: ", isValid)
		c.Assert(isValid, chk.Equals, v.valid)
	}
}

func (s *Day4TestSuite) TestDay4Incrementor(c *chk.C) {
	inputs := []int64{
		69,
		420,
		117,
	}

	dummyInput := solutions.Day4Input{}
	for _,v := range inputs {
		digs := dummyInput.ByDigit(v)
		dummyInput.IncrementCode(digs)

		c.Assert(dummyInput.DigitsToInt(digs), chk.Equals, v + 1)
	}
}

func (s *Day4TestSuite) TestDay4Perfector(c *chk.C) {
	inputs := []struct{
		original int64
		perfected int64
	}{
		{
			original: 111115,
			perfected: 111115,
		},
		{
			original: 175555,
			perfected: 177777,
		},
	}

	dummyInput := solutions.Day4Input{}
	for _,v := range inputs {
		digs := dummyInput.ByDigit(v.original)
		dummyInput.MakePerfect(digs)

		c.Log("Machine perfected: ", dummyInput.DigitsToInt(digs), " Expected: ", v.perfected)
		c.Assert(dummyInput.DigitsToInt(digs), chk.Equals, v.perfected)
	}
}
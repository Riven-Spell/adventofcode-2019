package solutions

import (
	"context"
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"runtime"
	"strconv"
	"strings"
)

type Day2Input struct{
	intcode.VM
	baseStr      string
	CustNounVerb bool // Set to true during preparation, considered when running parts.
	Noun int64 // Default set to 12
	Verb int64 // Default set to 2
}

func (s *Day2Input) Prepare(input string) {
	s.VM = intcode.VM{}
	s.Memory = make(map[int64]int64)
	s.baseStr = input

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.Memory[int64(k)] = pi
	}

	s.CustNounVerb = true
	s.Noun = 12
	s.Verb = 2
}

func (s *Day2Input) Part1() string {
	if s.CustNounVerb {
		s.Memory[1] = s.Noun
		s.Memory[2] = s.Verb
	}

	s.BlacklistedOps = map[int64]bool{
		3:true,
		4:true,
		5:true,
		6:true,
		7:true,
		8:true,
	}

	s.Autorun()

	return fmt.Sprint(s.Memory[0])
}

func (s *Day2Input) Part2() string {
	// Start by firing up a channel and yeeting in every possible noun and verb.
	// X = noun Y = verb
	instructionChan := make(chan util.Point, runtime.NumCPU() * 2)
	searchCtx, Canceller := context.WithCancel(context.Background())

	go func() {
		defer func(){
			// This magic lives here to handle the channel closing.
			// This isn't actually a bug, despite what Go would like to think.
			// The channel is closed to terminate all of day 2.
			recover()
		}()

		var noun, verb int64 = 0, 0

		instructionChan <- util.Point{X: 0, Y: 0}

		increment := func() {
			noun++

			if noun > 99 {
				verb++
				noun = 0
			}
		}

		for verb < 99 {
			increment()
			instructionChan <- util.Point{X: noun, Y: verb}
		}
	}()

	var result util.Point

	for i := runtime.NumCPU(); i > 0; i-- {
		go func() {
			for {
				var instruction util.Point
				select {
				case instruction = <- instructionChan:
				case <-searchCtx.Done():
					return
				}

				in := Day2Input{}

				in.Prepare(s.baseStr)
				in.Noun = instruction.X
				in.Verb = instruction.Y

				in.Part1()
				if in.Memory[0] == 19690720 {
					result = instruction
					Canceller()
					return
				}
			}
		}()
	}

	<-searchCtx.Done()
	close(instructionChan)

	return fmt.Sprint((100 * result.X) + result.Y)
}
package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
)

type Day11Input struct{
	baseMem map[int64]int64
	seenPoints map[util.Point]bool
	upperLeft util.Point
	lowerRight util.Point
}

func (s *Day11Input) Prepare(input string) {
	s.baseMem = make(map[int64]int64)
	s.seenPoints = make(map[util.Point]bool)
	s.upperLeft = util.Point{}
	s.lowerRight = util.Point{}

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.baseMem[int64(k)] = pi
	}
}

func (s *Day11Input) Fuck(vm *intcode.VM, chio *intcode.ChanIO) {
	fnSuicideChan := make(chan bool, 1)
	currentPosition := util.Point{}
	currentRotation := Up

	// This will only ever be 1 or 0
	rotate := func(dir int64) {
		if dir == 0 {
			if currentRotation == Up {
				currentRotation = Left
			} else {
				currentRotation--
			}
		} else {
			if currentRotation == Left {
				currentRotation = Up
			} else {
				currentRotation++
			}
		}

		currentPosition = currentPosition.Add(dPMapping[currentRotation])

		// Left = -1, so, if we're less than the current upper left X, increment.
		s.upperLeft.X = util.TernaryInt64(currentPosition.X < s.upperLeft.X, currentPosition.X, s.upperLeft.X)
		// Up = 1 so if we're higher than the current uper left Y, increment
		s.upperLeft.Y = util.TernaryInt64(currentPosition.Y > s.upperLeft.Y, currentPosition.Y, s.upperLeft.Y)
		// Reverse for lower right
		s.lowerRight.X = util.TernaryInt64(currentPosition.X > s.lowerRight.X, currentPosition.X, s.lowerRight.X)
		s.lowerRight.Y = util.TernaryInt64(currentPosition.Y < s.lowerRight.Y, currentPosition.Y, s.lowerRight.Y)
	}

	holdoverChan := make(chan bool, 1)
	go func() {
		for {
			chio.Stdin <- util.TernaryInt64(s.seenPoints[currentPosition], 1, 0)

			select {
			case <-fnSuicideChan:
				holdoverChan <- true
				return
			case output := <-chio.Stdout:
				s.seenPoints[currentPosition] = output == 1
				rotate(<-chio.Stdout)
			}
		}
	}()

	vm.Autorun()

	fnSuicideChan <- true
	<- holdoverChan
}

func (s *Day11Input) Part1() string {
	vm := intcode.VM{}
	vm.Memory = s.baseMem
	chio := intcode.GenChanIO(3)
	vm.IoMgr = chio

	s.seenPoints[util.Point{}] = false
	s.Fuck(&vm, chio)

	return fmt.Sprint(len(s.seenPoints))
}

func (s *Day11Input) Part2() string {
	vm := intcode.VM{}
	vm.Memory = s.baseMem
	chio := intcode.GenChanIO(3)
	vm.IoMgr = chio

	s.seenPoints[util.Point{}] = true
	s.Fuck(&vm, chio)

	getPanel := func(point util.Point) bool {
		if pan, ok := s.seenPoints[point]; ok {
			return pan
		} else {
			return false
		}
	}

	cPos := s.upperLeft
	result := ``

	for cPos.Y >= s.lowerRight.Y {
		for cPos.X <= s.lowerRight.X {
			if getPanel(cPos) {
				result += `█`
			} else {
				result += ` `
			}

			//fmt.Println(result)

			cPos.X++
		}

		cPos.Y--
		cPos.X = s.upperLeft.X
		result += "\n"
	}

	return result
}

package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"math"
	"strconv"
	"strings"
)

type wireMovement struct {
	Direction direction
	Distance int
}

type direction uint8

const (
	Up direction = iota
	Right
	Down
	Left
)

var dMapping = map[string]direction {
	"U":Up,
	"R":Right,
	"D":Down,
	"L":Left,
}

var dPMapping = map[direction]util.Point {
	Up:{Y:1},
	Right:{X:1},
	Down:{Y:-1},
	Left:{X:-1},
}

type Day3Input struct{
	Wires [][]wireMovement
}

func (s *Day3Input) Prepare(input string) {
	s.Wires = make([][]wireMovement, 0)

	for _,wire := range strings.Split(input, "\n") {
		movList := make([]wireMovement, 0)

		for _, v := range strings.Split(wire, ",") {
			v = strings.ToUpper(strings.TrimSpace(v))

			dist, err := strconv.ParseInt(v[1:], 10, 64)
			util.PanicIfErr(err)

			wm := wireMovement{
				Direction: dMapping[v[0:1]],
				Distance:  int(dist),
			}

			movList = append(movList, wm)
		}

		s.Wires = append(s.Wires, movList)
	}
}

func (s *Day3Input) Part1() string {
	wireframe := map[util.Point]int{}
	intersections := make([]util.Point, 0)

	for currentWire,wire := range s.Wires {
		cPoint := util.Point{X:0,Y:0}

		for _,movement := range wire {
			pMov := dPMapping[movement.Direction]

			for i := movement.Distance; i > 0; i-- {
				cPoint = cPoint.Add(pMov)

				if wireAtPoint, ok := wireframe[cPoint]; ok && cPoint != (util.Point{}) && wireAtPoint != currentWire {
					intersections = append(intersections, cPoint)
				}
				wireframe[cPoint] = currentWire
			}
		}
	}

	var closestManhattan int64 = math.MaxInt64
	for _,v := range intersections {
		m := util.Manhattan(util.Point{}, v)

		if m < closestManhattan {
			closestManhattan = m
		}
	}

	return fmt.Sprint(closestManhattan)
}

func (s *Day3Input) Part2() string {
	wireframe := map[util.Point]struct{wire, steps int}{}
	intersections := make([]struct{
		location util.Point
		sumsteps int
	}, 0)

	for currentWire,wire := range s.Wires {
		cPoint := util.Point{X:0,Y:0}
		stepsTaken := 0

		for _,movement := range wire {
			pMov := dPMapping[movement.Direction]

			for i := movement.Distance; i > 0; i-- {
				stepsTaken++
				cPoint = cPoint.Add(pMov)

				if wireAtPoint, ok := wireframe[cPoint]; ok && cPoint != (util.Point{}) && wireAtPoint.wire != currentWire {
					intersections = append(intersections, struct{
						location util.Point
						sumsteps int
					}{
						location:cPoint,
						sumsteps:stepsTaken + wireAtPoint.steps,
					})
				} else if !ok {
					wireframe[cPoint] = struct{wire, steps int}{wire: currentWire, steps: stepsTaken}
				}
			}
		}
	}

	lowestSteps := math.MaxInt64
	for _,v := range intersections {
		if v.sumsteps < lowestSteps {
			lowestSteps = v.sumsteps
		}
	}

	return strconv.Itoa(lowestSteps)
}
package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"strings"
	"sync"
)

type Day12Input struct{
	moons []physicsObject
	StepCount int
}

type physicsObject struct {
	Position util.Point3D
	Velocity util.Point3D
}

func (p *physicsObject) applyGravity(p2 physicsObject, axes string) {
	if strings.Contains(axes, "x") {
		if p2.Position.X > p.Position.X {
			p.Velocity.X++
		} else if p2.Position.X < p.Position.X {
			p.Velocity.X--
		}
	}

	if strings.Contains(axes, "y") {
		if p2.Position.Y > p.Position.Y {
			p.Velocity.Y++
		} else if p2.Position.Y < p.Position.Y {
			p.Velocity.Y--
		}
	}

	if strings.Contains(axes, "z") {
		if p2.Position.Z > p.Position.Z {
			p.Velocity.Z++
		} else if p2.Position.Z < p.Position.Z {
			p.Velocity.Z--
		}
	}
}

func (s *Day12Input) Prepare(input string) {
	s.moons = make([]physicsObject, 0)
	s.StepCount = 1000

	for _,v := range strings.Split(input, "\n") {
		moon := physicsObject{}
		_, err := fmt.Sscanf(strings.TrimSpace(v), "<x=%d, y=%d, z=%d>", &moon.Position.X, &moon.Position.Y, &moon.Position.Z)
		if err != nil {
			continue // Ignore the line, it's rubbish.
		}

		s.moons = append(s.moons, moon)
	}
}

func (s *Day12Input) SimulateMoons(axes string) {
	for k,m1 := range s.moons {
		for k2,m2 := range s.moons {
			if k != k2 {
				m1.applyGravity(m2, axes)
				s.moons[k] = m1
			}
		}
	}

	for k, m := range s.moons {
		m.Position = m.Position.Add(m.Velocity)
		s.moons[k] = m
	}
}

func (s *Day12Input) Part1() string {
	for i := s.StepCount; i > 0; i-- {
		s.SimulateMoons("xyz")
	}

	result := int64(0)
	for _,m := range s.moons {
		pot := util.IntAbs(m.Position.X) + util.IntAbs(m.Position.Y) + util.IntAbs(m.Position.Z)
		kin := util.IntAbs(m.Velocity.X) + util.IntAbs(m.Velocity.Y) + util.IntAbs(m.Velocity.Z)
		result += pot*kin
	}

	return fmt.Sprint(result)
}

func (s *Day12Input) Part2() string {
	var wg sync.WaitGroup

	matches := make([]int64, 3)

	wg.Add(3)

	for k,v := range strings.Split("xyz", "") {
		go func(axis string, match int) {
			defer wg.Done()
			dummyInput := Day12Input{
				moons: make([]physicsObject, len(s.moons)),
			}

			copy(dummyInput.moons, s.moons)

			history := make(map[string]bool)
			cstate := fmt.Sprint(dummyInput.moons)
			inc := int64(0)

			notInHistory := func() bool {
				_, ok := history[cstate]
				if !ok {
					history[cstate] = true
				}
				return !ok
			}

			for notInHistory() {
				dummyInput.SimulateMoons(axis)
				cstate = fmt.Sprint(dummyInput.moons)
				inc++
			}

			matches[match] = inc
		}(v, k)
	}

	wg.Wait()

	// Now we find the lowest common multiple.
	return fmt.Sprint(util.LCM(matches[0], matches[1], matches[2]))
}

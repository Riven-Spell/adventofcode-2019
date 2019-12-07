package solutions

import (
	"fmt"
	"strings"
)

type Day6Input struct{
	Orbits map[string]string
}

func (s *Day6Input) Prepare(input string) {
	s.Orbits = make(map[string]string)

	for _,orbit := range strings.Split(input, "\n") {
		orbit = strings.TrimSpace(orbit)
		orbitMarker := strings.Index(orbit, ")")

		base := orbit[:orbitMarker]
		object := orbit[orbitMarker+1:]

		s.Orbits[object] = base
	}
}

func (s *Day6Input) Part1() string {
	var orbits int64

	for v := range s.Orbits {
		obj, ok := v, true
		for ok {
			obj, ok = s.Orbits[obj]
			if ok {
				orbits++
			}
		}
	}

	return fmt.Sprint(orbits)
}

func (s *Day6Input) Part2() string {
	dmap := make(map[string]int)
	distance := 0

	obj, ok := "YOU", true
	for ok {
		dmap[obj] = distance
		obj, ok = s.Orbits[obj]
		if ok {
			distance++
		}
	}

	obj, ok = "SAN", true
	distance = -1 // An orbital transfer to SAN doesn't make sense.
	for ok {
		obj, ok = s.Orbits[obj]

		// Why check our distance before we add? Because this distance we'd miss here is already accounted for in the overlap.
		if dist, ok := dmap[obj]; ok {
			return fmt.Sprint(distance + dist)
		}

		if ok {
			distance++
		}
	}

	panic("orbits didn't link up")
}
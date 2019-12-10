package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"sort"
	"strings"
	"sync"
)

type Day10Input struct{
	AstList       []util.Point
	MSView        map[util.Fraction][]util.Point
	DebugAstCount int
}

func (s *Day10Input) Prepare(input string) {
	s.AstList = make([]util.Point, 0)
	s.MSView = nil // We reset MSview for an accurate calculation of how long part 2 will take to run, since it uses part 1 calculations.

	testProcess := util.Point{}
	s.DebugAstCount = 200
	for y,line := range strings.Split(input, "\n") {
		for x,char := range strings.TrimSpace(line) {
			if char != '.' {
				s.AstList = append(s.AstList, util.Point{X: int64(x), Y: int64(y)})
			}

			if char == 'X' {
				testProcess = util.Point{X: int64(x), Y: int64(y)}
			}
		}
	}

	if testProcess != (util.Point{}) {
		s.MSView = s.CalculateAstroidViews(testProcess)

		for k,v := range s.MSView {
			sorter := &util.DistSorter{
				Center: testProcess,
				List:   v,
			}

			sort.Sort(sorter)

			s.MSView[k] = sorter.List
		}
	}
}

// slopes to lines of sight
func (s *Day10Input) CalculateAstroidViews(asteroid util.Point) map[util.Fraction][]util.Point {
	out := make(map[util.Fraction][]util.Point)

	put := func(fraction util.Fraction, point util.Point) {
		if v, ok := out[fraction]; ok {
			out[fraction] = append(v, point)
		} else {
			out[fraction] = []util.Point{point}
		}
	}

	for _, other := range s.AstList {
		if other != asteroid {
			slope := util.Slope(asteroid, other).Simplify()

			put(slope, other)
		}
	}

	return out
}

func (s *Day10Input) Part1() string {
	var wg sync.WaitGroup

	var msLock sync.Mutex
	maxSeen := int64(0)

	for _,asteroid := range s.AstList {
		wg.Add(1)
		go func(ast util.Point) {
			defer wg.Done()

			view := s.CalculateAstroidViews(ast)

			msLock.Lock()
			if maxSeen < int64(len(view)) {
				maxSeen = int64(len(view))
				s.MSView = view

				for k,v := range s.MSView {
					sorter := &util.DistSorter{
						Center: ast,
						List:   v,
					}

					sort.Sort(sorter)

					s.MSView[k] = sorter.List
				}
			}
			msLock.Unlock()
		}(asteroid)
	}

	wg.Wait()

	return fmt.Sprint(maxSeen)
}


func (s *Day10Input) Part2() string {
	if s.MSView == nil {
		s.Part1() // We run part 1 since it gives us important view information about
	}

	// Translate to angles, then sort.
	var Angles []float64
	var newViews = make(map[float64][]util.Point)
	closest := float64(360)
	for k, v := range s.MSView {
		angle := k.GetAngle()

		Angles = append(Angles, angle)
		newViews[angle] = v
		if angle <= closest && angle >= 270 {
			closest = angle
		}
	}

	sort.Float64s(Angles)

	// Locate the golden angle
	i := 0
	for k,v := range Angles {
		if v == closest {
			i = k
			break
		}
	}

	// Start looping
	iterate := func() {
		i++
		if i >= len(Angles) {
			i = 0
		}
	}
	vaporized := 0
	var lastVaporized util.Point

	for vaporized < s.DebugAstCount {
		angle := Angles[i]

		lastVaporized = newViews[angle][0]
		newViews[angle] = newViews[angle][1:]
		vaporized++

		if len(newViews[angle]) == 0 {

			// Delete this from the map and the list
			delete(newViews, angle)

			Angles = append(Angles[:i], Angles[i+1:]...)
		} else {
			iterate()
		}

		if len(Angles) == 0 {
			panic("no solution found!")
		}
	}

	return fmt.Sprint((lastVaporized.X * 100) + lastVaporized.Y)
}

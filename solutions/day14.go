package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"math"
	"strings"
)

type Day14Input struct {
	// element name to its requirement
	Reactions map[string]elementReaction
}

type elementAmount struct {
	count   int64
	element string
}
type elementReaction struct {
	reactants []elementAmount
	result    elementAmount
}

func (s *Day14Input) Prepare(input string) {
	s.Reactions = make(map[string]elementReaction)

	for _, line := range strings.Split(input, "\n") {
		buffer := ""
		reactants := make([]elementAmount, 0)
		awaitResult := false
		result := elementAmount{}
		for _, element := range strings.Split(strings.TrimSpace(line), " ") {
			if element != "=>" {
				buffer += strings.TrimSuffix(element, ",")

				count, ename := int64(0), ""
				if c, err := fmt.Sscanf(buffer, "%d %s", &count, &ename); err == nil && c == 2 {
					if !awaitResult {
						reactants = append(reactants, elementAmount{count: count, element: ename})
					} else {
						result = elementAmount{count: count, element: ename}
					}

					buffer = ""
				} else {
					buffer += " "
				}
			} else {
				awaitResult = true
			}
		}

		s.Reactions[result.element] = elementReaction{
			reactants: reactants,
			result:    result,
		}
	}
}

func (s *Day14Input) Part1() string {
	eQueue := make(chan elementAmount, 500)
	oreCount := int64(0)
	spares := make(map[string]int64)
	reactions := s.Reactions

	putElementInQueue := func(amount elementAmount) {
		eQueue <- amount
	}

	eQueue <- elementAmount{
		count:   1,
		element: "FUEL",
	}

	for {
		var x elementAmount
		var toBreak bool
		select {
		case x = <-eQueue:
		default:
			toBreak = true
		}

		if toBreak {
			break
		}

		recipe, ok := reactions[x.element]
		if !ok {
			panic(x.element + " is not an element in the reactions list.")
		}

		// check for any spares
		var needed = x.count
		if spareCount, ok := spares[x.element]; ok && spareCount > 0 {
			needed = util.ClampInt64(needed-spareCount, 0, math.MaxInt64)

			// Use only what we need
			if needed == 0 {
				spareCount -= x.count
				spares[x.element] = spareCount
			} else {
				spares[x.element] = 0
			}
		}

		// Find multiple needed
		var ct, mul int64 = 0, 0
		for needed > ct {
			mul++
			ct += recipe.result.count
		}

		// Process needed elements
		for _, v := range recipe.reactants {
			if v.element == "ORE" {
				oreCount += v.count * mul
				continue
			}

			if needed > 0 {
				putElementInQueue(elementAmount{
					count:   v.count * mul,
					element: v.element,
				})
			}
		}

		// Mark leftovers
		spares[x.element] += ct - needed
	}

	return fmt.Sprint(oreCount)
}

func (s *Day14Input) GenerateFuel(count int64) int64 {
	eQueue := make(chan elementAmount, 500)
	oreCount := int64(0)
	spares := make(map[string]int64)
	reactions := s.Reactions

	putElementInQueue := func(amount elementAmount) {
		eQueue <- amount
	}

	eQueue <- elementAmount{
		count:   count,
		element: "FUEL",
	}

	for {
		var x elementAmount
		var toBreak bool
		select {
		case x = <-eQueue:
		default:
			toBreak = true
		}

		if toBreak {
			break
		}

		recipe, ok := reactions[x.element]
		if !ok {
			panic(x.element + " is not an element in the reactions list.")
		}

		// check for any spares
		var needed = x.count
		if spareCount, ok := spares[x.element]; ok && spareCount > 0 {
			needed = util.ClampInt64(needed-spareCount, 0, math.MaxInt64)

			// Use only what we need
			if needed == 0 {
				spareCount -= x.count
				spares[x.element] = spareCount
			} else {
				spares[x.element] = 0
			}
		}

		// Find multiple needed
		var ct, mul int64 = 0, 0
		for needed > ct {
			mul++
			ct += recipe.result.count
		}

		// Process needed elements
		for _, v := range recipe.reactants {
			if v.element == "ORE" {
				oreCount += v.count * mul
				continue
			}

			if needed > 0 {
				putElementInQueue(elementAmount{
					count:   v.count * mul,
					element: v.element,
				})
			}
		}

		// Mark leftovers
		spares[x.element] += ct - needed
	}

	return oreCount
}

func (s *Day14Input) Part2() string {
	fmt.Println("Hold tight. Day 14 part 2 takes a long time. (Up to several minutes)")

	lower := int64(0)
	max := int64(1)
	threshold := int64(1000000000000)

	for s.GenerateFuel(max) < threshold {
		max *= 2
		fmt.Println("bumping max search to", max)
	}
	lower = max / 2

	for {
		if lower >= max {
			return fmt.Sprint(lower - 1)
		}

		mid := (lower + max) / 2

		fmt.Println("testing ", mid, "min:", lower, "max:", max)
		result := s.GenerateFuel(mid)
		if result > threshold {
			max = mid
		} else if result < threshold {
			lower = mid + 1
		}
	}
}

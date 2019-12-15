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

func (s *Day14Input) Part2() string {
	target := int64(1000000000000)

	eQueue := make(chan elementAmount, 500)
	oreCount := int64(0)
	lastOreCount := int64(0)
	fuelCount := int64(0)
	fuelDiff := int64(1)
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
		select {
		case x = <-eQueue:
		default:
			// We'll accelerate pace depending on what we think of the current fueling rate.
			// The idea being that as we approach our target closer and closer, we'll have a better and better idea of just how much fuel can be produced.
			// The final run should be just one unit.

			fuelCount += fuelDiff
			oreDiff := oreCount - lastOreCount
			diffPerFuel := oreDiff / fuelDiff
			lastOreCount = oreCount

			// We now have an idea of our diff per fuel.
			// You might be wondering why I left in these debug statements.
			// The code actually runs twice as fast with them than without them... fucking somehow!?!!?
			fmt.Println("Approximately", diffPerFuel, "ORE per FUEL")
			fmt.Println(target - oreCount, "ore remaining", fuelCount, "fuel produced")

			// Let's decide at what pace we can move at.
			pace := (target - oreCount) / diffPerFuel

			// There seems to be an upper and lower bound in terms of success with our back-off rate.
			backOffRate := int64(2)

			if pace < backOffRate {
				fmt.Println("Expecting less fuel than our backoff rate.")
				putElementInQueue(elementAmount{
					count:   1,
					element: "FUEL",
				})

				fuelDiff = 1
				continue
			}

			fmt.Printf("(inferred %d FUEL prior to backoff), ", pace)
			pace /= backOffRate
			fmt.Printf("inferring %d FUEL remaining (((%d - %d) / %d) / %d)\n", pace, target, oreCount, diffPerFuel, backOffRate)

			putElementInQueue(elementAmount{
				count:   pace,
				element: "FUEL",
			})
			fuelDiff = pace
			continue
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

				if oreCount > target {
					return fmt.Sprint(fuelCount)
				}
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
}

package main

import (
	"bytes"
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions"
	"github.com/Virepri/adventofcode-2019/util"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dxpxresult = regexp.MustCompile(`(?i)d\d+p\d`)

func printHelp() {
	fmt.Println("Usage: aoc <dXpX/all> <inputFile>")
	fmt.Println("dXpX: Day X part 1/2/0 (0 = both parts) (OR, specify all to run all days and parts.)")
	fmt.Println("inputFile: Specifies the input file (Optional: Supplies the function with my input)")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	} else {
		if os.Args[1] == "all" {
			fmt.Println("Running all days, recording time and using default input")

			packetStart := time.Now()
			for k,v := range solutions.RegisteredDays {
				if v.ExpectedOutputs[0] != "norun" {
					v.DummyInput.Prepare(*v.StringInput)

					t := time.Now()           // current time
					o := v.DummyInput.Part1() // output
					tx := time.Now().Sub(t)   // runtime

					fmt.Println("Day " + strconv.Itoa(k+1) + " Part 1 (runtime: " + tx.String() + ")")
					if v.ExpectedOutputs[0] != "disabled" {
						fmt.Println(o + " " + util.TernaryString(o == v.ExpectedOutputs[0], "(PASSED)", "(FAILED: "+v.ExpectedOutputs[0]+")"))
					} else {
						fmt.Println(o)
					}
				} else {
					fmt.Println("Day", k+1, "Part 2 must be run alone. (Perhaps it requires input?)")
				}

				if v.ExpectedOutputs[1] != "norun" {
					// Reset the input
					v.DummyInput.Prepare(*v.StringInput)
					t := time.Now()
					o := v.DummyInput.Part2()
					tx := time.Now().Sub(t)

					fmt.Println("Day " + strconv.Itoa(k+1) + " Part 2 (runtime: " + tx.String() + ")")
					if v.ExpectedOutputs[1] != "disabled" {
						fmt.Println(o + " " + util.TernaryString(o == v.ExpectedOutputs[1], "(PASSED)", "(FAILED: "+v.ExpectedOutputs[1]+")"))
					} else {
						fmt.Println(o)
					}
				} else {
					fmt.Println("Day", k+1, "Part 2 must be run alone. (Perhaps it requires input?)")
				}
			}
			tx := time.Now().Sub(packetStart)

			fmt.Println("\nTotal packet runtime: " + tx.String())
		} else if dxpxresult.MatchString(os.Args[1]) {
			pIdx := strings.Index(os.Args[1], "p")
			day, err := strconv.ParseInt(os.Args[1][1:pIdx], 10, 64)
			util.PanicIfErr(err)
			part, err := strconv.ParseInt(os.Args[1][pIdx+1:], 10, 64)
			util.PanicIfErr(err)

			daysolution := solutions.RegisteredDays[day-1]
			dataset := false

			if len(os.Args) >= 3 {
				f, err := os.Open(os.Args[2])
				util.PanicIfErr(err)

				buf := bytes.Buffer{}

				_, err = buf.ReadFrom(f)
				util.PanicIfErr(err)

				*daysolution.StringInput = buf.String()
				dataset = true
			}

			daysolution.DummyInput.Prepare(*daysolution.StringInput)
			fmt.Println("Running day", day)

			switch part {
			case 0:
				if daysolution.ExpectedOutputs[0] != "norun" {
					t := time.Now()
					o := daysolution.DummyInput.Part1()
					tx := time.Now().Sub(t)

					fmt.Println("Part 1 (runtime: " + tx.String() + ")")
					if !dataset && daysolution.ExpectedOutputs[0] != "disabled" {
						fmt.Println(o + " " + util.TernaryString(o == daysolution.ExpectedOutputs[0], "(PASSED)", "(FAILED: "+daysolution.ExpectedOutputs[0]+")"))
					} else {
						fmt.Println(o)
					}
				} else {
					fmt.Println("Part 1 can only be ran on it's own. Perhaps it requires input?")
				}

				if daysolution.ExpectedOutputs[1] != "norun" {
					daysolution.DummyInput.Prepare(*daysolution.StringInput) // Some days modify their input.

					t := time.Now()
					o := daysolution.DummyInput.Part2()
					tx := time.Now().Sub(t)

					fmt.Println("Part 2 (runtime: " + tx.String() + ")")
					if !dataset && daysolution.ExpectedOutputs[1] != "disabled" {
						fmt.Println(o + " " + util.TernaryString(o == daysolution.ExpectedOutputs[1], "(PASSED)", "(FAILED: "+daysolution.ExpectedOutputs[1]+")"))
					} else {
						fmt.Println(o)
					}
				} else {
					fmt.Println("Part 1 can only be ran on it's own. Perhaps it requires input?")
				}
			case 1:
				t := time.Now()
				o := daysolution.DummyInput.Part1()
				tx := time.Now().Sub(t)

				fmt.Println("Part 1 " + util.TernaryString(daysolution.ExpectedOutputs[0] != "norun", "(runtime: " +  tx.String() + ")", ""))
				if !dataset && (daysolution.ExpectedOutputs[0] != "disabled" && daysolution.ExpectedOutputs[0] != "norun") {
					fmt.Println(o + " " + util.TernaryString(o == daysolution.ExpectedOutputs[0], "(PASSED)", "(FAILED: "+daysolution.ExpectedOutputs[0]+")"))
				} else {
					fmt.Println(o)
				}
			case 2:
				t := time.Now()
				o := daysolution.DummyInput.Part2()
				tx := time.Now().Sub(t)

				fmt.Println("Part 2 " + util.TernaryString(daysolution.ExpectedOutputs[1] != "norun", "(runtime: " +  tx.String() + ")", ""))
				if !dataset && (daysolution.ExpectedOutputs[1] != "disabled" && daysolution.ExpectedOutputs[1] != "norun") {
					fmt.Println(o + " " + util.TernaryString(o == daysolution.ExpectedOutputs[1], "(PASSED)", "(FAILED: "+daysolution.ExpectedOutputs[1]+")"))
				} else {
					fmt.Println(o)
				}
			default:
				printHelp()
			}
		} else {
			printHelp()
		}
	}
}

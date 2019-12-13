package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"strconv"
	"strings"
	"sync"
)

type Day13Input struct{
	baseMem map[int64]int64
	screen map[util.Point]gameTile
}

type gameTile uint8

const (
	TileEmpty gameTile = iota
	TileWall
	TileBlock
	TilePaddle
	TileBall
)

func (s *Day13Input) Prepare(input string) {
	s.baseMem = make(map[int64]int64)
	s.screen = make(map[util.Point]gameTile)

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.baseMem[int64(k)] = pi
	}
}

func (s *Day13Input) Part1() string {
	chio := intcode.GenChanIO(3)
	vm := intcode.VM{
		Memory:         s.baseMem,
		IoMgr:          chio,
	}

	// Render screen
	killCh := make(chan bool, 1)
	waitCh := make(chan bool, 1)
	go func(){
		for {
			var pt util.Point
			var tile gameTile
			select {
			case <- killCh:
				waitCh <- true
				return
			case pt.X = <-chio.Stdout:
				pt.Y = <-chio.Stdout
				tile = gameTile(<-chio.Stdout)
			}

			s.screen[pt] = tile
		}
	}()

	vm.Autorun()
	killCh <- true
	<- waitCh

	count := 0
	for _,v := range s.screen {
		if v == TileBlock {
			count++
		}
	}

	return fmt.Sprint(count)
}

func (s *Day13Input) OutputScreen() {
	var mx, my int64 = 0, 0
	for k := range s.screen {
		if k.X > mx {
			mx = k.X
		}

		if k.Y > my {
			my = k.Y
		}
	}

	fmt.Print("\033[H\033[2J")

	for y := int64(0); y <= my; y++ {
		for x := int64(0); x <= mx; x++ {
			pt := util.Point{X:x, Y:y}

			t, ok := s.screen[pt]

			if !ok {
				fmt.Print(" ")
			} else {
				switch t {
				case TileEmpty:
					fmt.Print(" ")
				case TileBall:
					fmt.Print("o")
				case TileWall:
					fmt.Print("W")
				case TileBlock:
					fmt.Print("B")
				case TilePaddle:
					fmt.Print("P")
				}
			}
		}

		fmt.Println()
	}
}

func (s *Day13Input) Part2() string {
	chio := intcode.GenChanIO(3)
	vm := intcode.VM{
		Memory:         s.baseMem,
		IoMgr:          chio,
	}

	vm.Memory[0] = 2

	var plock sync.Mutex
	var ballPos, paddlePos util.Point
	var score int

	// Render screen
	killCh := make(chan bool, 1)
	waitCh := make(chan bool, 1)
	go func(){
		for {
			var pt util.Point
			var tile int
			select {
			case <- killCh:
				killCh <- true
				waitCh <- true
				return
			case pt.X = <-chio.Stdout:
				pt.Y = <-chio.Stdout
				tile = int(<-chio.Stdout)
			}

			if pt != (util.Point{X: -1}) {
				plock.Lock()
				if gameTile(tile) == TilePaddle {
					paddlePos = pt
				}

				if gameTile(tile) == TileBall {
					ballPos = pt

					switch {
					case paddlePos.X > ballPos.X:
						chio.Stdin <- -1
					case paddlePos.X < ballPos.X:
						chio.Stdin <- 1
					default:
						chio.Stdin <- 0
					}
				}
				plock.Unlock()

				s.screen[pt] = gameTile(tile)
			} else {
				score = tile
			}
		}
	}()

	vm.Autorun()
	killCh <- true
	<- waitCh

	return fmt.Sprint(score)
}

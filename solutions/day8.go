package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
	"math"
	"strings"
	"sync"
)

type Day8Input struct{
	ImgSize util.Point // Default 25x6, respected by prepare if not 0,0
	Layers  []string
}

func (s *Day8Input) Prepare(input string) {
	if s.ImgSize == (util.Point{}) {
		s.ImgSize = util.Point{X: 25, Y: 6}
	}

	layerSize := s.ImgSize.X * s.ImgSize.Y

	s.Layers = make([]string, len(input)/layerSize)

	for i := 0; i < len(s.Layers); i++ {
		s.Layers[i] = input[i*layerSize:(i+1)*layerSize]
	}
}

func (s *Day8Input) Part1() string {
	var wg sync.WaitGroup
	var zdlock sync.Mutex
	var zeroDigs = math.MaxInt64
	var oneDigs int
	var twoDigs int

	for _,v := range s.Layers {
		wg.Add(1)
		go func(row string) {
			defer wg.Done()

			zeroes := strings.Count(row, "0")
			ones := strings.Count(row, "1")
			twos := strings.Count(row, "2")

			zdlock.Lock()
			if zeroes < zeroDigs {
				zeroDigs = zeroes
				oneDigs = ones
				twoDigs = twos
			}
			zdlock.Unlock()
		}(v)
	}

	wg.Wait()

	return fmt.Sprint(oneDigs * twoDigs)
}

func (s *Day8Input) CalculatePixel(pIdx int) rune {
	// If you really sound the name of this variable out, it'll tell you exactly what this solution is
	var cAnswer = '2'

	for i := len(s.Layers)-1; i >= 0; i-- {
		v := s.Layers[i]

		if v[pIdx] != '2' {
			if v[pIdx] == '1' {
				cAnswer = '█'
			} else {
				cAnswer = ' '
			}
		}
	}

	return cAnswer
}

func (s *Day8Input) Part2() string {
	output := make([][]rune, s.ImgSize.Y)

	for y := 0; y < s.ImgSize.Y; y++ {
		output[y] = make([]rune, s.ImgSize.X)

		for x := 0; x < s.ImgSize.X; x++ {
			output[y][x] = s.CalculatePixel((y*s.ImgSize.X) + x)
		}
	}

	strOut := ""

	for _,v := range output {
		strOut += string(v) + "\n"
	}

	return strOut
}

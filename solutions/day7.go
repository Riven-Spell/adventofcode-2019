package solutions

import (
	"fmt"
	qp "github.com/Ramshackle-Jamathon/go-quickPerm"
	"github.com/Virepri/adventofcode-2019/solutions/intcode"
	"github.com/Virepri/adventofcode-2019/util"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Day7Input struct{
	baseMem map[int]int
}

func (s *Day7Input) Prepare(input string) {
	s.baseMem = make(map[int]int)

	for k,v := range strings.Split(input, ",") {
		pi, err := strconv.ParseInt(v, 10, 64)
		util.PanicIfErr(err)

		s.baseMem[k] = int(pi)
	}
}

func (s *Day7Input) Part1() string {
	permutations := qp.GeneratePermutationsInt([]int{0,1,2,3,4})
	var wg sync.WaitGroup

	var maxsig int
	var mslock sync.Mutex

	for i := runtime.NumCPU(); i > 0; i-- {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				perm, ok := <- permutations
				if !ok {
					return
				}

				vm := intcode.VM{}
				sig := 0
				for _, phase := range perm {
					vm.Reset(s.baseMem)
					io := vm.IoMgr.(*intcode.PreparedIO) // Reset makes it
					io.Stdin = []int{phase, sig}

					vm.Autorun()

					sig = io.Stdout[0]
				}

				mslock.Lock()
				if sig > maxsig {
					maxsig = sig
				}
				mslock.Unlock()
			}
		}()
	}

	wg.Wait()

	return fmt.Sprint(maxsig)
}

func (s *Day7Input) CreateVMPool() (handles []*intcode.VM) {
	// We need 5 VMs. Spin up the first, and then link 4 more to its output.
	firstVM := &intcode.VM{}
	firstVM.Reset(s.baseMem)
	firstVM.IoMgr = intcode.GenChanIO(2)

	lastVM := firstVM

	handles = []*intcode.VM{firstVM}

	for i := 4; i > 0; i-- {
		newVM := &intcode.VM{}
		newVM.Reset(s.baseMem)
		chio := intcode.GenChanIO(2)
		chio.Stdin = lastVM.IoMgr.(*intcode.ChanIO).Stdout
		newVM.IoMgr = chio

		handles = append(handles, newVM)
		lastVM = newVM
	}

	// Plenty of type wrangling to be had.
	firstVM.IoMgr.(*intcode.ChanIO).Stdin = lastVM.IoMgr.(*intcode.ChanIO).Stdout

	return
}

func (s *Day7Input) Part2() string {
	permutations := qp.GeneratePermutationsInt([]int{5,6,7,8,9})
	var wg sync.WaitGroup

	var maxsig int
	var mslock sync.Mutex

	for {
		perm, ok := <- permutations
		if !ok {
			break
		}

		// CreateVMPool spins up VMs with two buffer slots on the channels.
		// This is great, because we can go right ahead and write the necessary signals ahead of time.
		pool := s.CreateVMPool()

		// Indicate their phase
		for k,v := range perm {
			pool[k].IoMgr.Write(v)
		}

		// Write 0 for the initial signal
		pool[0].IoMgr.Write(0)

		// Spin up all the VMs, sans the first because she's speshul
		for k,v := range pool[1:] {
			wg.Add(1)
			go func(vmi int, x *intcode.VM){
				defer wg.Done()
				x.Autorun()
			}(k,v)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			pool[0].Autorun()

			// We'll deadlock if there's nothing there.
			sig := <- pool[0].IoMgr.(*intcode.ChanIO).Stdout

			mslock.Lock()
			if sig > maxsig {
				maxsig = sig
			}
			mslock.Unlock()
		}()
	}

	wg.Wait()

	return fmt.Sprint(maxsig)
}

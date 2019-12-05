package intcode

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
)

type VM struct {
	PC int
	BlacklistedOps map[int]bool
	Memory map[int]int // This is actually just acting as an unbound array.
	Ioutil Ioutil
}

func (v *VM) GetMemory(i int) int {
	if out, ok := v.Memory[i]; ok {
		return out
	} else {
		return 0
	}
}

func (v *VM) GetMemoryRange(loc, length int) []int {
	out := make([]int, 0)

	for length > 0 {
		out = append(out, v.GetMemory(loc))

		length--
		loc++
	}

	return out
}

func (v *VM) GetArgs(argModes, args []int) []int {
	if len(argModes) != len(args) {
		panic("argmodes must equal args")
	}

	out := make([]int, len(args))

	for k,arg := range args {
		switch argModes[k] {
		case 0: // Position mode
			out[k] = v.GetMemory(arg) // Despite output args NEVER being in position mode, we won't question it to save time, and just trust the input.
		case 1: // Immediate mode
			out[k] = arg
		}
	}

	return out
}

func (v *VM) Autorun() {
	for v.Step() {}
}

// False: Running still.
// True: The VM has stopped.
func (v *VM) Step() bool {
	var argModes []int
	var opCode int
	var o IntcodeOperation

	if v.Memory[v.PC] > 99 {
		// TODO: Cache these.
		digs := util.ByDigit(int64(v.Memory[v.PC]))
		codeDiv := len(digs)-2
		opCode = int(util.DigitsToInt(digs[codeDiv:]))

		o = OperationMap[opCode]

		// Copy the default arg mode so we don't overwrite it
		overrideModes := digs[:codeDiv]
		argModes = make([]int, len(o.DefaultArgMode))

		copy(argModes, o.DefaultArgMode)
		for x := len(overrideModes) - 1; x >= 0; x-- {
			argModes[(len(overrideModes)-1) - x] = overrideModes[x]
		}
	} else {
		opCode = v.Memory[v.PC]

		o = OperationMap[opCode]

		argModes = o.DefaultArgMode
	}

	if _, ok := v.BlacklistedOps[opCode]; ok {
		panic(fmt.Sprintf("blacklisted opcode %d was ran", opCode))
	}

	if opCode == 99 {
		return false
	}

	args := v.GetArgs(argModes, v.GetMemoryRange(v.PC+1, o.ArgCount))

	if o.F(args, v) {
		v.PC += 1 + o.ArgCount
	}

	if opCode == 99 {
		return false
	}

	return true
}
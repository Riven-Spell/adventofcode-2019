package intcode

import (
	"fmt"
	"github.com/Virepri/adventofcode-2019/util"
)

type VM struct {
	PC             int
	BlacklistedOps map[int]bool
	Memory         map[int]int // This is actually just acting as an unbound array.
	IoMgr          IoMgr
}

// Reset fully resets the ioutil state-- be sure to replace it.
func (v *VM) Reset(baseMem map[int]int) {
	v.Memory = make(map[int]int)
	v.PC = 0
	v.IoMgr = &PreparedIO{}

	for loc, value := range baseMem {
		v.Memory[loc] = value
	}
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

func (v *VM) AutorunWIndex(idx int) {
	for v.Step() {
		fmt.Println("VM ", idx, "is running", v.PC, v.Memory[v.PC])
	}
}

// False: Running still.
// True: The VM has stopped.
func (v *VM) Step() bool {
	var argModes []int
	var opCode int
	var o IntcodeOperation
	var ok bool

	if v.Memory[v.PC] > 99 {
		// TODO: Cache these.
		digs := util.ByDigit(int64(v.Memory[v.PC]))
		codeDiv := len(digs)-2
		opCode = int(util.DigitsToInt(digs[codeDiv:]))

		o, ok = OperationMap[opCode]

		// Copy the default arg mode so we don't overwrite it
		overrideModes := digs[:codeDiv]
		argModes = make([]int, len(o.DefaultArgMode))

		copy(argModes, o.DefaultArgMode)
		for x := len(overrideModes) - 1; x >= 0; x-- {
			argModes[(len(overrideModes)-1) - x] = overrideModes[x]
		}
	} else {
		opCode = v.Memory[v.PC]

		o, ok = OperationMap[opCode]

		argModes = o.DefaultArgMode
	}

	if !ok {
		panic(fmt.Sprint(opCode, "is not a valid opcode."))
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
package intcode

import "github.com/Virepri/adventofcode-2019/util"

// Mem is only meant to be used as "put it here".
// Also takes in a VM.
// Returns true if you should increment the PC, false if you do not.
type IntcodeFunc func(args []int64, vm *VM) bool

type IntcodeOperation struct {
	ArgCount       int64
	DefaultArgMode []int64
	Mode2ArgWrite  []bool
	F              IntcodeFunc
}

var OperationMap = map[int64]IntcodeOperation {
	1: { // Intcode add
		ArgCount:       3,
		DefaultArgMode: []int64{0,0,1},
		F: func(args []int64, vm *VM) bool {
			if len(args) != 3 {
				panic("incorrect ArgCount on add")
			}

			vm.Memory[args[2]] = args[0] + args[1]
			return true
		},
	},
	2: { // Intcode mul
		ArgCount:       3,
		DefaultArgMode: []int64{0,0,1},
		F: func(args []int64, vm *VM) bool {
			if len(args) != 3 {
				panic("incorrect ArgCount on mul")
			}

			vm.Memory[args[2]] = args[0] * args[1]
			return true
		},
	},
	3: { // Intcode read
		ArgCount:       1,
		DefaultArgMode: []int64{1},
		F: func(args []int64, vm *VM) bool {
			vm.Memory[args[0]] = vm.IoMgr.Read()
			return true
		},
	},
	4: { // Intcode write
		ArgCount:       1,
		DefaultArgMode: []int64{0},
		F: func(args []int64, vm *VM) bool {
			vm.IoMgr.Write(args[0])
			return true
		},
	},
	5: { // Intcode JNZ (Jump nonzero)
		ArgCount: 2,
		DefaultArgMode: []int64{0, 0},
		F: func(args []int64, vm *VM) bool {
			if args[0] != 0 {
				vm.PC = args[1]
				return false
			}
			return true
		},
	},
	6: { // Intcode JEZ (Jump equals zero)
		ArgCount: 2,
		DefaultArgMode: []int64{0, 0},
		F: func(args []int64, vm *VM) bool {
			if args[0] == 0 {
				vm.PC = args[1]
				return false
			}
			return true
		},
	},
	7: { // Intcode LT (Less than)
		ArgCount: 3,
		DefaultArgMode: []int64{0, 0, 1},
		F: func(args []int64, vm *VM) bool {
			vm.Memory[args[2]] = util.TernaryInt64(args[0] < args[1], 1, 0)

			return true
		},
	},
	8: { // Intcode EQ (Equals)
		ArgCount: 3,
		DefaultArgMode: []int64{0, 0, 1},
		F: func(args []int64, vm *VM) bool {
			vm.Memory[args[2]] = util.TernaryInt64(args[0] == args[1], 1, 0)

			return true
		},
	},
	9: { // Intcode RB (Adjust relative base)
		ArgCount: 1,
		DefaultArgMode: []int64{0},
		F: func(args []int64, vm *VM) bool {
			vm.RelativeBase += args[0]
			return true
		},
	},
	99: { // STOP
		ArgCount: 0,
		F: func(args []int64, vm *VM) bool {
			panic("opcode 99 (STOP) was ran")
			return true
		},
	},
}
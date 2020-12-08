package computer

import (
	"strconv"
)

// operations
const (
	OpAcc = "acc"
	OpJmp = "jmp"
	OpNop = "nop"
)

// Instruction model
type Instruction struct {
	Op  string
	Arg int
}

// Computer model
type Computer struct {
	Acc               int
	History           map[int]bool
	Instructions      []Instruction
	SavedInstructions []Instruction
}

// New com from list of instructions
func New(rows [][]string) Computer {
	instructions := make([]Instruction, len(rows))
	for i, row := range rows {
		instr := Instruction{}
		instr.Op = row[0]
		instr.Arg, _ = strconv.Atoi(row[1])
		instructions[i] = instr
	}
	return Computer{Instructions: instructions}
}

// SetInstr at index
func (com *Computer) SetInstr(index int, op string, arg int) {
	// save instruction set so we can swap back if needed
	com.SavedInstructions = make([]Instruction, len(com.Instructions))
	copy(com.SavedInstructions, com.Instructions)
	com.Instructions[index] = Instruction{Op: op, Arg: arg}
}

// Revert back to saved instructions if exists
func (com *Computer) Revert() {
	if com.SavedInstructions != nil {
		com.Instructions = com.SavedInstructions
	}
}

// Run instructions, returns (acc, isInfinite)
func (com *Computer) Run() (int, bool) {
	// reset history and acc for each run
	com.History = map[int]bool{}
	com.Acc = 0

	index := 0

	// log.Printf("com.Instructions: %v\n", com.Instructions)

	for true {
		if com.History[index] {
			return com.Acc, true
		}
		if index >= len(com.Instructions) {
			return com.Acc, false
		}

		com.History[index] = true
		instr := com.Instructions[index]
		switch instr.Op {
		case OpAcc:
			{
				index++
				com.Acc += instr.Arg
			}
		case OpJmp:
			{
				index += instr.Arg
			}
		case OpNop:
			{
				index++
			}
		}
	}
	return 0, false
}

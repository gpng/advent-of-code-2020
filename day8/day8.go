package day8

import (
	"log"

	"github.com/gpng/advent-of-code-2020/computer"
	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 8")
	defer utils.Timer("Day 8 total")()

	rows := utils.ScanFileLinesToStrings("day8/input.txt", " ")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	com := computer.New(rows)
	acc, _ := com.Run()
	return acc
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()
	com := computer.New(rows)
	for i, instr := range com.Instructions {
		toRun := true
		switch instr.Op {
		case computer.OpNop:
			{
				com.SetInstr(i, computer.OpJmp, instr.Arg)
				break
			}
		case computer.OpJmp:
			{
				com.SetInstr(i, computer.OpNop, instr.Arg)
				break
			}
		default:
			toRun = false
		}
		if toRun {
			acc, isInfinite := com.Run()
			if !isInfinite {
				return acc
			}
			com.Revert()
		}
	}
	return com.Acc
}

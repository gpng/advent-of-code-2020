package day15

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 15")
	defer utils.Timer("Day 15 total")()

	ints := utils.ScanFileLinesToInts("day15/input.txt", ",")

	log.Printf("Part 1 answer: %d", part1(ints))
	log.Printf("Part 2 answer: %d", part2(ints))
}

func part1(ints [][]int) int {
	defer utils.Timer("Part 1")()

	return play(ints, 2020)
}

func part2(ints [][]int) int {
	defer utils.Timer("Part 2")()

	return play(ints, 30000000)
}

func play(ints [][]int, turns int) int {
	row := ints[0]

	nums := map[int]int{}
	for i, num := range row[:len(row)-1] {
		nums[num] = i + 1
	}
	last := row[len(row)-1]

	for i := len(row) + 1; i <= turns; i++ {
		turn, ok := nums[last]
		nums[last] = i - 1
		if !ok {
			last = 0
			continue
		}
		next := i - 1 - turn
		last = next
	}
	return last
}

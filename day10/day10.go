package day10

import (
	"log"
	"sort"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 10")
	defer utils.Timer("Day 10 total")()

	ints := utils.ScanFileLinesToInt("day10/input.txt", " ")
	ints = append(ints, 0)
	sort.Ints(ints)

	log.Printf("Part 1 answer: %d", part1(ints))
	log.Printf("Part 2 answer: %d", part2(ints))
}

func part1(ints []int) int {
	defer utils.Timer("Part 1")()

	num1 := 0
	num3 := 1

	for i := 1; i < len(ints); i++ {
		diff := (ints[i] - ints[i-1])
		switch diff {
		case 1:
			num1++
		case 3:
			num3++
		}
	}

	return num1 * num3
}

func part2(ints []int) int {
	defer utils.Timer("Part 2")()

	adaptors := map[int]int{}
	for _, adaptor := range ints {
		adaptors[adaptor] = 0
	}
	adaptors[ints[0]] = 1

	for _, num := range ints {
		for i := num + 1; i < num+4; i++ {
			if _, ok := adaptors[i]; ok {
				adaptors[i] += adaptors[num]
			}
		}
	}
	return adaptors[ints[len(ints)-1]]
}

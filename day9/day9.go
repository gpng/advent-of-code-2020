package day9

import (
	"log"
	"math"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 9")
	defer utils.Timer("Day 9 total")()

	ints := utils.ScanFileLinesToInt("day9/input.txt", " ")

	invalid := part1(ints, 25)

	log.Printf("Part 1 answer: %d", invalid)
	log.Printf("Part 2 answer: %d", part2(ints, invalid))

}

func part1(ints []int, preambleLength int) int {
	defer utils.Timer("Part 1")()

	invalid := 0
	for i := preambleLength; i < len(ints); i++ {
		num := ints[i]

		preamble := ints[i-preambleLength : i]

		found := false
		for j := 0; j < preambleLength-1; j++ {
			if found {
				break
			}
			for k := j; k < preambleLength; k++ {
				if preamble[j]+preamble[k] == num {
					found = true
					break
				}
			}
		}
		if !found {
			invalid = num
		}
	}
	return invalid
}

func part2(ints []int, invalid int) int {
	defer utils.Timer("Part 2")()

	start := 0
	end := 2
	sum, sumMinMax := sums(ints[start:end])
	for sum != invalid {
		if sum < invalid {
			end++
		}
		if sum > invalid {
			start++
		}
		sum, sumMinMax = sums(ints[start:end])
	}
	return sumMinMax
}

func sums(ints []int) (int, int) {
	sum := 0
	min := math.MaxInt64
	max := 0
	for _, i := range ints {
		sum += i
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return sum, min + max
}

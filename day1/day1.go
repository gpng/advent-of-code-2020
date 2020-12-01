package day1

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 1")
	defer utils.Timer("Day 1 total")()

	nums := utils.ScanFileLinesToInt("day1/input.txt", ",")

	log.Printf("Part 1 answer: %d", part1(nums))
	log.Printf("Part 2 answer: %d", part2(nums))
}

func part1(nums []int) int {
	defer utils.Timer("Part 1")()
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j]
			}
		}
	}
	return 0
}

func part2(nums []int) int {
	defer utils.Timer("Part 1")()
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {

				if nums[i]+nums[j]+nums[k] == 2020 {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}
	return 0
}

package day6

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 6")
	defer utils.Timer("Day 6 total")()

	rows := utils.ScanFileLinesToStrings("day6/input.txt", "")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	total := 0

	charMap := map[string]int{}
	for _, row := range rows {
		if len(row) == 0 || row[0] == "" {
			total += len(charMap)
			charMap = map[string]int{}
			continue
		}

		for _, char := range row {
			if _, ok := charMap[char]; !ok {
				charMap[char] = 0
			}
			charMap[char]++
		}
	}

	return total + len(charMap)
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()

	total := 0
	charMap := map[string]int{}
	numRows := 0
	for _, row := range rows {
		if len(row) == 0 || row[0] == "" {
			for _, v := range charMap {
				if v == numRows {
					total++
				}
			}
			charMap = map[string]int{}
			numRows = 0
			continue
		}

		numRows++
		for _, char := range row {
			if _, ok := charMap[char]; !ok {
				charMap[char] = 0
			}
			charMap[char]++
		}
	}
	for _, v := range charMap {
		if v == numRows {
			total++
		}
	}

	return total
}

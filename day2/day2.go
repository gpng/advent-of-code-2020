package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 2")
	defer utils.Timer("Day 2 total")()

	text := utils.ScanFileLinesToStrings("day2/input.txt", " ")

	log.Printf("Part 1 answer: %d", part1(text))
	log.Printf("Part 2 answer: %d", part2(text))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()
	result := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		minMax := strings.Split(row[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		char := row[1][0:1]
		password := row[2]

		count := 0
		for j := 0; j < len(password); j++ {
			c := password[j : j+1]
			if c == char {
				count++
			}
		}
		if count >= min && count <= max {
			result++
		}
	}
	return result
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()
	result := 0
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		indexes := strings.Split(row[0], "-")
		firstIndex, _ := strconv.Atoi(indexes[0])
		secondIndex, _ := strconv.Atoi(indexes[1])
		char := row[1][0:1]
		password := row[2]

		firstChar := password[firstIndex-1 : firstIndex]
		secondChar := password[secondIndex-1 : secondIndex]

		count := 0
		if firstChar == char {
			count++
		}
		if secondChar == char {
			count++
		}
		if count == 1 {
			result++
		}
	}
	return result
}

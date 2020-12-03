package day3

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 3")
	defer utils.Timer("Day 3 total")()

	rows := utils.ScanFileLinesToStrings("day3/input.txt", "")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func traverse(rows [][]string, dx int, dy int) int {
	x := 0
	y := 0

	trees := 0

	for y < len(rows)-dy {
		y += dy
		x += dx

		row := rows[y]
		rowX := x % len(row)
		if row[rowX] == "#" {
			trees++
		}
	}
	return trees
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()
	return traverse(rows, 3, 1)
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()
	return traverse(rows, 1, 1) * traverse(rows, 3, 1) * traverse(rows, 5, 1) * traverse(rows, 7, 1) * traverse(rows, 1, 2)
}

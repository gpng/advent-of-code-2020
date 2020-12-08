package day5

import (
	"log"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 5")
	defer utils.Timer("Day 5 total")()

	rows := utils.ScanFileLinesToStrings("day5/input.txt", ";")

	max, ids := part1(rows)
	log.Printf("Part 1 answer: %d", max)
	log.Printf("Part 2 answer: %d", part2(ids))
}

func part1(rows [][]string) (int64, map[int64]bool) {
	defer utils.Timer("Part 1")()

	var max int64
	seats := map[int64][]bool{}
	ids := map[int64]bool{}

	for _, r := range rows {
		rowBin := strings.ReplaceAll(strings.ReplaceAll(r[0][:7], "F", "0"), "B", "1")
		row, _ := strconv.ParseInt(rowBin, 2, 64)
		colBin := strings.ReplaceAll(strings.ReplaceAll(r[0][7:], "L", "0"), "R", "1")
		col, _ := strconv.ParseInt(colBin, 2, 64)

		if _, ok := seats[row]; !ok {
			seats[row] = make([]bool, 8)
		}
		seats[row][col] = true

		id := row*8 + col
		if id > max {
			max = id
		}
		ids[id] = true
	}
	return max, ids
}

func part2(IDs map[int64]bool) int64 {
	defer utils.Timer("Part 2")()

	for i := int64(1); i < 127*8+7-1; i++ {
		if !IDs[i] && IDs[i-1] && IDs[i+1] {
			return i
		}
	}
	return 0
}

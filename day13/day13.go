package day13

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 13")
	defer utils.Timer("Day 13 total")()

	rows := utils.ScanFileLinesToStrings("day13/input.txt", ",")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %s", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	earliest, _ := strconv.Atoi(rows[0][0])

	minDiff := math.MaxInt64
	minBusNo := 0

	for _, bus := range rows[1] {
		if bus == "x" {
			continue
		}
		busNo, _ := strconv.Atoi(bus)
		diff := busNo - earliest%busNo
		if diff < minDiff {
			minDiff = diff
			minBusNo = busNo
		}
	}
	return minDiff * minBusNo
}

func part2(rows [][]string) string {
	defer utils.Timer("Part 2")()

	eqs := make([]string, 0)
	for i, bus := range rows[1] {
		if bus == "x" {
			continue
		}
		busNo, _ := strconv.Atoi(bus)
		eq := fmt.Sprintf("(t + %d)%%%d=0", i, busNo)
		eqs = append(eqs, eq)
	}
	return "Copy paste this into www.wolframalpha.com: " + strings.Join(eqs, ", ")
}

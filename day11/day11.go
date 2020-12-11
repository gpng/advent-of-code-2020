package day11

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Seats
const (
	Floor    = "."
	Empty    = "L"
	Occupied = "#"
)

// Run code
func Run() {
	log.Println("Running day 11")
	defer utils.Timer("Day 11 total")()

	rows := utils.ScanFileLinesToStrings("day11/input.txt", "")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	old := deepCopy(rows)
	new := round1(old)
	for changed(old, new) {
		old = new
		new = round1(old)
	}

	return countOccupied(new)
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()

	old := deepCopy(rows)
	new := round2(old)
	for changed(old, new) {
		old = new
		new = round2(old)
	}
	return countOccupied(new)
}

func round1(rows [][]string) [][]string {
	new := deepCopy(rows)

	for y, row := range rows {
		for x, seat := range row {
			if seat == Floor {
				continue
			}
			occupied := 0
			for yy := y - 1; yy <= y+1; yy++ {
				if yy < 0 || yy > len(rows)-1 {
					continue
				}
				for xx := x - 1; xx <= x+1; xx++ {
					if xx < 0 || xx > len(row)-1 || (yy == y && xx == x) {
						continue
					}
					s := rows[yy][xx]
					if s == Occupied {
						occupied++
					}
				}
			}
			switch seat {
			case Empty:
				{
					if occupied == 0 {
						new[y][x] = Occupied
					}
				}
			case Occupied:
				{
					if occupied >= 4 {
						new[y][x] = Empty
					}
				}
			}
		}
	}
	return new
}

func round2(rows [][]string) [][]string {
	new := deepCopy(rows)

	for y, row := range rows {
		for x, seat := range row {
			if seat == Floor {
				continue
			}
			occupied := 0

			// check upwards
			for yy := y - 1; yy >= 0; yy-- {
				val := rows[yy][x]
				if val == Occupied {
					occupied++
				}
				if val != Floor {
					break
				}
			}

			// up right
			xx := x + 1
			for yy := y - 1; yy >= 0; yy-- {
				if xx > len(row)-1 {
					break
				}
				val := rows[yy][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
				xx++
			}

			// right
			for xx := x + 1; xx <= len(row)-1; xx++ {
				val := rows[y][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
			}

			// down right
			xx = x + 1
			for yy := y + 1; yy <= len(rows)-1; yy++ {
				if xx > len(row)-1 {
					break
				}
				val := rows[yy][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
				xx++
			}

			// down
			for yy := y + 1; yy <= len(rows)-1; yy++ {
				val := rows[yy][x]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
			}

			// down left
			xx = x - 1
			for yy := y + 1; yy <= len(rows)-1; yy++ {
				if xx < 0 {
					break
				}
				val := rows[yy][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
				xx--
			}

			// left
			for xx := x - 1; xx >= 0; xx-- {
				val := rows[y][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
			}

			// up left
			xx = x - 1
			for yy := y - 1; yy >= 0; yy-- {
				if xx < 0 {
					break
				}
				val := rows[yy][xx]
				if val == Occupied {
					occupied++
					break
				}
				if val != Floor {
					break
				}
				xx--
			}

			switch seat {
			case Empty:
				{
					if occupied == 0 {
						new[y][x] = Occupied
					}
				}
			case Occupied:
				{
					if occupied >= 5 {
						new[y][x] = Empty
					}
				}
			}
		}
	}
	return new
}

func deepCopy(rows [][]string) [][]string {
	new := make([][]string, len(rows))

	for i, row := range rows {
		newRow := make([]string, 0)
		newRow = append(newRow, row...)
		new[i] = newRow
	}
	return new
}

func changed(a [][]string, b [][]string) bool {
	for x, row := range a {
		for y, val := range row {
			if val != b[x][y] {
				return true
			}
		}
	}
	return false
}

func countOccupied(rows [][]string) int {
	total := 0

	for _, row := range rows {
		for _, val := range row {
			if val == Occupied {
				total++
			}
		}
	}
	return total
}

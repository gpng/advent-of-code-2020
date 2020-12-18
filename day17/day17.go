package day17

import (
	"log"

	"github.com/gpng/advent-of-code-2020/utils"
)

const (
	active   = "#"
	inactive = "."
)

// Run code
func Run() {
	log.Println("Running day 17")
	defer utils.Timer("Day 17 total")()

	rows := utils.ScanFileLinesToStrings("day17/input.txt", "")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	grid := map[int]map[int]map[int]bool{}

	for y, row := range rows {
		for x, status := range row {
			assign(grid, 0, y, x, getStatus(status))
		}
	}
	startLength := len(rows)
	for i := 1; i <= 6; i++ {
		grid = cycle(grid, startLength, i)
	}

	return countActive(grid)
}

func getStatus(str string) bool {
	return str == active
}

func cycle(grid map[int]map[int]map[int]bool, startLength, num int) map[int]map[int]map[int]bool {
	newGrid := copy(grid)
	for z := -num; z <= num; z++ {
		for y := -num; y <= startLength+num; y++ {
			for x := -num; x <= startLength+num; x++ {
				activeCount := activeNeighbours(grid, z, y, x)
				status, ok := grid[z][y][x]
				if !ok {
					assign(newGrid, z, y, x, false)
				}
				if status && activeCount != 2 && activeCount != 3 {
					assign(newGrid, z, y, x, false)
				}
				if !status && activeCount == 3 {
					assign(newGrid, z, y, x, true)
				}
			}
		}
	}
	return newGrid
}

func copy(grid map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	newGrid := map[int]map[int]map[int]bool{}
	for z, zz := range grid {
		for y, yy := range zz {
			for x, status := range yy {
				assign(newGrid, z, y, x, status)
			}
		}
	}
	return newGrid
}

func assign(grid map[int]map[int]map[int]bool, z, y, x int, status bool) {
	if _, ok := grid[z]; !ok {
		grid[z] = map[int]map[int]bool{}
	}
	if _, ok := grid[z][y]; !ok {
		grid[z][y] = map[int]bool{}
	}
	grid[z][y][x] = status
}

func activeNeighbours(grid map[int]map[int]map[int]bool, z, y, x int) int {
	count := 0
	for zz := z - 1; zz <= z+1; zz++ {
		for yy := y - 1; yy <= y+1; yy++ {
			for xx := x - 1; xx <= x+1; xx++ {
				if zz == z && yy == y && xx == x {
					continue
				}
				if grid[zz][yy][xx] {
					count++
				}
			}
		}
	}
	return count
}

func countActive(grid map[int]map[int]map[int]bool) int {
	count := 0
	for _, zz := range grid {
		for _, yy := range zz {
			for _, status := range yy {
				if status {
					count++
				}
			}
		}
	}
	return count
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()

	grid := map[int]map[int]map[int]map[int]bool{}

	for y, row := range rows {
		for x, status := range row {
			assign4d(grid, 0, 0, y, x, getStatus(status))
		}
	}
	startLength := len(rows)

	for i := 1; i <= 6; i++ {
		grid = cycle4d(grid, startLength, i)
	}

	return countActive4d(grid)
}

func cycle4d(grid map[int]map[int]map[int]map[int]bool, startLength, num int) map[int]map[int]map[int]map[int]bool {
	newGrid := copy4d(grid)
	for w := -num; w <= num; w++ {
		for z := -num; z <= num; z++ {
			for y := -num; y <= startLength+num; y++ {
				for x := -num; x <= startLength+num; x++ {
					activeCount := activeNeighbours4d(grid, w, z, y, x)
					status, ok := grid[w][z][y][x]
					if !ok {
						assign4d(newGrid, w, z, y, x, false)
					}
					if status && activeCount != 2 && activeCount != 3 {
						assign4d(newGrid, w, z, y, x, false)
					}
					if !status && activeCount == 3 {
						assign4d(newGrid, w, z, y, x, true)
					}
				}
			}
		}
	}
	return newGrid
}

func copy4d(grid map[int]map[int]map[int]map[int]bool) map[int]map[int]map[int]map[int]bool {
	newGrid := map[int]map[int]map[int]map[int]bool{}
	for w, ww := range grid {
		for z, zz := range ww {
			for y, yy := range zz {
				for x, status := range yy {
					assign4d(newGrid, w, z, y, x, status)
				}
			}
		}
	}
	return newGrid
}

func assign4d(grid map[int]map[int]map[int]map[int]bool, w, z, y, x int, status bool) {
	if _, ok := grid[w]; !ok {
		grid[w] = map[int]map[int]map[int]bool{}
	}
	if _, ok := grid[w][z]; !ok {
		grid[w][z] = map[int]map[int]bool{}
	}
	if _, ok := grid[w][z][y]; !ok {
		grid[w][z][y] = map[int]bool{}
	}
	grid[w][z][y][x] = status
}

func activeNeighbours4d(grid map[int]map[int]map[int]map[int]bool, w, z, y, x int) int {
	count := 0
	for ww := w - 1; ww <= w+1; ww++ {
		for zz := z - 1; zz <= z+1; zz++ {
			for yy := y - 1; yy <= y+1; yy++ {
				for xx := x - 1; xx <= x+1; xx++ {
					if ww == w && zz == z && yy == y && xx == x {
						continue
					}
					if grid[ww][zz][yy][xx] {
						count++
					}
				}
			}
		}
	}
	return count
}

func countActive4d(grid map[int]map[int]map[int]map[int]bool) int {
	count := 0
	for _, ww := range grid {
		for _, zz := range ww {
			for _, yy := range zz {
				for _, status := range yy {
					if status {
						count++
					}
				}
			}
		}
	}
	return count
}

package day12

import (
	"log"
	"math"
	"strconv"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 12")
	defer utils.Timer("Day 12 total")()

	rows := utils.ScanFileLinesToStrings("day12/input.txt", " ")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()

	coords := []int{0, 0}
	angle := 90

	for _, row := range rows {
		instr := row[0][:1]
		val, _ := strconv.Atoi(row[0][1:])

		switch instr {
		case "N":
			{
				coords[1] += val
				break
			}
		case "S":
			{
				coords[1] -= val
				break
			}
		case "E":
			{
				coords[0] += val
				break
			}
		case "W":
			{
				coords[0] -= val
				break
			}
		case "L":
			{
				angle = updateAngle(angle, -val)
				break
			}
		case "R":
			{
				angle = updateAngle(angle, val)
				break
			}
		case "F":
			{
				switch angle {
				case 0:
					{
						coords[1] += val
						break
					}
				case 90:
					{
						coords[0] += val
						break
					}
				case 180:
					{
						coords[1] -= val
						break
					}
				case 270:
					{
						coords[0] -= val
						break
					}
				}
			}
		}
	}

	return distance(coords)
}

func updateAngle(angle, diff int) int {
	new := angle + diff
	if new < 0 {
		new += 360
	} else if new >= 360 {
		new -= 360
	}
	return new
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()

	waypoint := []int{10, 1}
	coords := []int{0, 0}

	for _, row := range rows {
		instr := row[0][:1]
		val, _ := strconv.Atoi(row[0][1:])

		switch instr {
		case "N":
			{
				waypoint[1] += val
				break
			}
		case "S":
			{
				waypoint[1] -= val
				break
			}
		case "E":
			{
				waypoint[0] += val
				break
			}
		case "W":
			{
				waypoint[0] -= val
				break
			}
		case "L":
			{
				waypoint = rotate(waypoint, -val+360)
				break
			}
		case "R":
			{
				waypoint = rotate(waypoint, val)
				break
			}
		case "F":
			{
				coords[0] += (waypoint[0] * val)
				coords[1] += (waypoint[1] * val)
				break
			}
		}
	}

	return distance(coords)
}

func rotate(coords []int, angle int) []int {
	switch angle {
	case 0:
		return coords
	case 90:
		{
			return []int{coords[1], -coords[0]}
		}
	case 180:
		{
			coords = []int{-coords[0], -coords[1]}
		}
	case 270:
		{
			coords = []int{-coords[1], coords[0]}
		}
	}
	return coords
}

func distance(coords []int) int {
	return int(math.Abs(float64(coords[0])) + math.Abs(float64(coords[1])))
}

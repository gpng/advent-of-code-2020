package day16

import (
	"log"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 16")
	defer utils.Timer("Day 16 total")()

	nearby := utils.ScanFileLinesToInts("day16/nearby.txt", ",")
	mine := utils.ScanFileLinesToInts("day16/ticket.txt", ",")[0]
	fieldRows := utils.ScanFileLinesToStrings("day16/fields.txt", ": ")

	fields := parseFields(fieldRows)

	invalidSum, valid := part1(fields, nearby)
	log.Printf("Part 1 answer: %d", invalidSum)
	log.Printf("Part 2 answer: %d", part2(fields, valid, mine))
}

func part1(fields map[string]map[int]bool, nearby [][]int) (int, [][]int) {
	defer utils.Timer("Part 1")()

	invalidSum := 0
	validTickets := make([][]int, 0)
	for _, ticket := range nearby {
		invalid := false
		for _, val := range ticket {
			valValid := false
			for _, field := range fields {
				if field[val] {
					valValid = true
					break
				}
			}
			if !valValid {
				invalid = true
				invalidSum += val
			}
		}
		if invalid {
			continue
		}
		validTickets = append(validTickets, ticket)
	}
	return invalidSum, validTickets
}

func parseFields(rows [][]string) map[string]map[int]bool {
	fields := map[string]map[int]bool{}
	for _, row := range rows {
		key := row[0]
		possibilities := strings.Split(row[1], " or ")
		fields[key] = map[int]bool{}
		for _, poss := range possibilities {
			minMax := strings.Split(poss, "-")
			min, _ := strconv.Atoi(minMax[0])
			max, _ := strconv.Atoi(minMax[1])
			for i := min; i <= max; i++ {
				fields[key][i] = true
			}
		}
	}
	return fields
}

func part2(fields map[string]map[int]bool, nearby [][]int, mine []int) int {
	defer utils.Timer("Part 1")()

	tickets := append(nearby, mine)
	possiblePositions := map[string][]int{}
	numFields := len(fields)

	// each field has multiple valid positions, so we collect all valid positions first
	for i := 0; i < numFields; i++ {
		for field, vals := range fields {
			valid := true
			for _, ticket := range tickets {
				if !vals[ticket[i]] {
					valid = false
					break
				}
			}
			if valid {
				if _, ok := possiblePositions[field]; !ok {
					possiblePositions[field] = make([]int, 0)
				}
				possiblePositions[field] = append(possiblePositions[field], i)
			}
		}
	}

	// eliminate possibilities by looking for fields that only have a single valid position
	// and removing that position from other fields
	final := map[string]int{}
	for len(final) < numFields {
		fixed := ""
		deleteIndex := 0
		for field, poss := range possiblePositions {
			if len(poss) == 1 {
				fixed = field
				deleteIndex = poss[0]
			}
			break
		}
		delete(possiblePositions, fixed)
		final[fixed] = deleteIndex

		for field, poss := range possiblePositions {
			if field == fixed {
				continue
			}
			newPoss := make([]int, 0)
			copy(newPoss, poss)
			for _, val := range poss {
				if val != deleteIndex {
					newPoss = append(newPoss, val)
				}
			}
			possiblePositions[field] = newPoss
		}
	}

	mult := 1
	for key, i := range final {
		if strings.Split(key, " ")[0] != "departure" {
			continue
		}
		mult *= mine[i]
	}

	return mult
}

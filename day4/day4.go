package day4

import (
	"log"
	"strings"

	"github.com/gpng/advent-of-code-2020/passport"
	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 4")
	defer utils.Timer("Day 4 total")()

	rows := utils.ScanFileLinesToStrings("day4/input.txt", " ")

	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func isValidFields(fieldsMap map[string]bool) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for i := 0; i < len(requiredFields); i++ {
		if !fieldsMap[requiredFields[i]] {
			return false
		}
	}
	return true
}

func part1(rows [][]string) int {
	defer utils.Timer("Part 1")()
	fieldsMap := map[string]bool{}
	valid := 0

	for i := 0; i < len(rows); i++ {
		row := rows[i]

		if row[0] == "" {
			if isValidFields(fieldsMap) {
				valid++
			}
			fieldsMap = map[string]bool{}
			continue
		}

		for j := 0; j < len(row); j++ {
			split := strings.Split(row[j], ":")
			key := split[0]
			fieldsMap[key] = true
		}
	}

	if isValidFields(fieldsMap) {
		valid++
	}
	return valid
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 2")()
	passports := passport.Parse(rows)
	return len(passports)
}

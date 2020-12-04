package day4

import (
	"log"
	"regexp"
	"strconv"
	"strings"

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
			val := split[1]
			isValid := false
			switch key {
			case "byr":
				{
					year, err := strconv.Atoi(val)
					if err != nil || year < 1920 || year > 2002 {
						break
					}
					isValid = true
				}
			case "iyr":
				{
					year, err := strconv.Atoi(val)
					if err != nil || year < 2010 || year > 2020 {
						break
					}
					isValid = true
				}
			case "eyr":
				{
					year, err := strconv.Atoi(val)
					if err != nil || year < 2020 || year > 2030 {
						break
					}
					isValid = true
				}
			case "hgt":
				{
					scale := val[len(val)-2:]
					height, err := strconv.Atoi(val[:len(val)-2])
					if err != nil {
						break
					}
					switch scale {
					case "cm":
						{
							isValid = height >= 150 && height <= 193
						}
					case "in":
						{
							isValid = height >= 59 && height <= 76
						}
					}
				}
			case "hcl":
				{
					if len(val) != 7 || val[:1] != "#" {
						break
					}
					color := val[1:]
					re := regexp.MustCompile("^[a-zA-Z0-9]*$")
					if re.MatchString(color) {
						isValid = true
					}
				}
			case "ecl":
				{
					validColors := map[string]bool{
						"amb": true,
						"blu": true,
						"brn": true,
						"gry": true,
						"grn": true,
						"hzl": true,
						"oth": true,
					}
					if validColors[val] {
						isValid = true
					}
				}
			case "pid":
				{
					if len(val) != 9 {
						break
					}
					re := regexp.MustCompile("^[0-9]*$")
					if re.MatchString(val) {
						isValid = true
					}
				}
			}
			fieldsMap[key] = isValid
		}
	}

	if isValidFields(fieldsMap) {
		valid++
	}
	return valid
}

package day7

import (
	"log"
	"strconv"
	"strings"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {
	log.Println("Running day 7")
	defer utils.Timer("Day 7 total")()

	rows := utils.ScanFileLinesToStrings("day7/input.txt", " contain ")

	rules := parseRules(rows)

	log.Printf("Part 1 answer: %d", part1(rules))
	log.Printf("Part 2 answer: %d", part2(rules))

}

func parseRules(rows [][]string) map[string]map[string]int {
	rules := map[string]map[string]int{}
	for _, row := range rows {
		key := row[0][:len(row[0])-5]

		rules[key] = map[string]int{}
		if row[1] == "no other bags." {
			continue
		}
		contains := strings.Split(row[1][:len(row[1])-1], ", ")
		for _, c := range contains {
			split := strings.Split(c, " ")
			num, _ := strconv.Atoi(split[0])
			bag := split[1] + " " + split[2]
			rules[key][bag] = num
		}
	}
	return rules
}

func part1(rules map[string]map[string]int) int {
	defer utils.Timer("Part 1")()

	res := map[string]bool{}

	bags := []string{"shiny gold"}

	for len(bags) > 0 {
		newBags := make([]string, 0)
		for bag, rule := range rules {
			if res[bag] {
				continue
			}
			for _, r := range bags {
				if _, ok := rule[r]; ok {
					newBags = append(newBags, bag)
					res[bag] = true
				}
			}
		}
		bags = newBags
	}

	return len(res)
}

func part2(rules map[string]map[string]int) int {
	type Bag struct {
		num  int
		name string
	}

	count := 0
	bags := []Bag{{1, "shiny gold"}}
	for len(bags) > 0 {
		newBags := make([]Bag, 0)
		for _, bag := range bags {
			rule := rules[bag.name]

			for ruleBag, num := range rule {
				count += (num * bag.num)
				newBags = append(newBags, Bag{num * bag.num, ruleBag})
			}
		}
		bags = newBags
	}

	return count
}

package day14

import (
	"log"
	"strconv"

	"github.com/gpng/advent-of-code-2020/utils"
)

// Run code
func Run() {

	defer utils.Timer("Day 14 total")()

	rows := utils.ScanFileLinesToStrings("day14/input.txt", " = ")
	log.Printf("Part 1 answer: %d", part1(rows))
	log.Printf("Part 2 answer: %d", part2(rows))
}

func part1(rows [][]string) int64 {
	defer utils.Timer("Part 1")()

	memory := map[int]int64{}
	mask := map[int]string{}

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if row[0] == "mask" {
			mask = parseMask(row)
			continue
		}
		val, _ := strconv.Atoi(row[1])
		// convert to binary string
		bin := strconv.FormatInt(int64(val), 2)
		// append 0s to make it length 36
		bin36 := ""
		for j := 0; j < 36-len(bin); j++ {
			bin36 += "0"
		}
		bin36 += bin
		// apply mask
		for k, v := range mask {
			if v == "X" {
				continue
			}
			bin36 = replaceAt(bin36, k, v)
		}
		// parse back to int
		binVal, _ := strconv.ParseInt(bin36, 2, 64)
		// finally assign to memory
		pointer, _ := strconv.Atoi(row[0][4 : len(row[0])-1])
		memory[pointer] = binVal
	}

	sum := int64(0)
	for _, val := range memory {
		sum += val
	}
	return sum
}

func part2(rows [][]string) int {
	defer utils.Timer("Part 1")()

	memory := map[int64]int{}
	mask := map[int]string{}

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if row[0] == "mask" {
			mask = parseMask(row)
			continue
		}

		// get address string
		address, _ := strconv.Atoi(row[0][4 : len(row[0])-1])
		// convert to binary string
		binAddress := strconv.FormatInt(int64(address), 2)
		// append 0s to make it length 36
		binAddress36 := ""
		for j := 0; j < 36-len(binAddress); j++ {
			binAddress36 += "0"
		}
		binAddress36 += binAddress

		// get all the indexes of floating bits, and set value at those indexes to 0
		xIndexes := make([]int, 0)
		for k, v := range mask {
			if v == "0" {
				continue
			}
			if v == "X" {
				xIndexes = append(xIndexes, k)
				binAddress36 = replaceAt(binAddress36, k, "0")
			} else {
				binAddress36 = replaceAt(binAddress36, k, v)
			}
		}

		addresses := []string{binAddress36}
		addresses = append(addresses, replaceAt(binAddress36, xIndexes[0], "1"))

		// for every existing combination, add new address at floating bit index replaced with 1
		for x := 1; x < len(xIndexes); x++ {
			xIndex := xIndexes[x]
			newAddresses := make([]string, 0)
			for _, addr := range addresses {
				newAddresses = append(newAddresses, replaceAt(addr, xIndex, "1"))
			}
			addresses = append(addresses, newAddresses...)
		}

		// finally assign the values to each address
		val, _ := strconv.Atoi(row[1])
		for _, addr := range addresses {
			addrVal, _ := strconv.ParseInt(addr, 2, 64)
			memory[addrVal] = val
		}
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}
	return sum
}

func replaceAt(str string, index int, with string) string {
	out := []rune(str)
	out[index] = []rune(with)[0]
	return string(out)
}

func parseMask(row []string) map[int]string {
	maskRaw := row[1]
	mask := map[int]string{}
	for i := range maskRaw {
		char := maskRaw[i : i+1]
		mask[i] = char
	}
	return mask
}

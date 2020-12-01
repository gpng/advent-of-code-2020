package main

import (
	"flag"
	"log"
	"os"

	"github.com/gpng/advent-of-code-2020/day1"
)

func main() {
	day := flag.Int("d", 0, "Day to run")

	flag.Parse()

	runMap := map[int]func(){
		1: day1.Run,
	}

	// run all
	if *day == 0 {
		for i := 1; i <= len(runMap); i++ {
			runMap[i]()
		}
		os.Exit(0)
	}

	fn, ok := runMap[*day]
	if !ok {
		log.Fatalf("Invalid day %d", *day)
	}

	fn()

	os.Exit(0)
}

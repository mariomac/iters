package main

import (
	"fmt"
	"slices"

	"github.com/mariomac/iters"
)

func main_iterate() {
	numbers := iters.Iterate(1, double)
	sixNums := iters.Limit(6, numbers)
	words := iters.Map(sixNums, asWord)

	fmt.Println(slices.Collect(words))
}

func double(n int) int {
	return 2 * n
}

func asWord(n int) string {
	if n < 10 {
		return []string{
			"zero", "one", "two", "three", "four", "five",
			"six", "seven", "eight", "nine",
		}[n]
	} else {
		return "many"
	}
}

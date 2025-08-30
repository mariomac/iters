package main

import (
	"fmt"

	"github.com/mariomac/iters"
)

func main_reduce() {
	// create an infinite sequence of 1, 2, 3, ...
	seq := iters.Iterate(1, func(n int) int {
		return n + 1
	})
	// limit the sequence to 8 items
	seq8 := iters.Limit(8, seq)
	// multiply each item in the limited sequence
	fac8, _ := iters.Reduce(seq8, func(a, b int) int {
		return a * b
	})
	fmt.Println("The factorial of 8 is", fac8)
}

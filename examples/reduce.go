package main

import (
	"fmt"

	"github.com/mariomac/iters"
)

func main_reduce() {
	// create a sequence in range [1, 8]
	seq := iters.OfRange(1, 9)
	// multiply each item in the limited sequence
	fac8, _ := iters.Reduce(seq, func(a, b int) int {
		return a * b
	})
	fmt.Println("The factorial of 8 is", fac8)
}

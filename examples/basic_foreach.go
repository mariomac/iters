package main

import (
	"fmt"
	"slices"

	"github.com/mariomac/iters"
)

func main_basic_foreach() {
	numbers := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	prime := iters.Filter(numbers, isPrime)
	iters.ForEach(prime, func(n int) {
		fmt.Printf("%d is a prime number\n", n)
	})
}

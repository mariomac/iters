package main

import (
	"fmt"

	"github.com/mariomac/iters"
)

func main_basic_foreach() {
	numbers := iters.OfRange(1, 12)
	prime := iters.Filter(numbers, isPrime)
	iters.ForEach(prime, func(n int) {
		fmt.Printf("%d is a prime number\n", n)
	})
}

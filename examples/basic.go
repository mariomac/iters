package main

import (
	"fmt"
	"slices"

	"github.com/mariomac/iters"
)

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main_basic() {
	numbers := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})

	for n := range iters.Filter(numbers, isPrime) {
		fmt.Printf("%d is a prime number\n", n)
	}
}

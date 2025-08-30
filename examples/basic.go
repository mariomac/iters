package main

import (
	"fmt"

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
	numbers := iters.OfRange(1, 12)

	for n := range iters.Filter(numbers, isPrime) {
		fmt.Printf("%d is a prime number\n", n)
	}
}

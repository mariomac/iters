package main

import (
	"fmt"

	"github.com/mariomac/iters"
)

func main_comparable() {
	words := iters.Distinct(
		iters.Of("hello", "hello", "!", "ho", "ho", "ho", "!"),
	).ToSlice()

	fmt.Printf("Deduplicated words: %v\n", words)
}

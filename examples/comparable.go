package main

import (
	"fmt"
	"slices"

	"github.com/mariomac/iters"
)

func main_comparable() {
	words := iters.Distinct(
		slices.Values([]string{"hello", "hello", "!", "ho", "ho", "ho", "!"}),
	)

	fmt.Printf("Deduplicated words: %v\n", slices.Collect(words))
}

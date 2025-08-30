package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/mariomac/iters"
)

func main_dice() {
	rnd := rand.New(rand.NewSource(time.Now().UnixMilli()))
	fmt.Println("let me throw 5 times a dice for you")

	results := iters.Map(
		iters.Generate(rnd.Int),
		func(n int) int {
			return n%6 + 1
		},
	)
	takeFive := iters.Limit(5, results)

	fmt.Printf("results: %v\n", slices.Collect(takeFive))
}

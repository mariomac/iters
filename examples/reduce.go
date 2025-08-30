package main

import (
	"fmt"

	"github.com/mariomac/gostream/item"
	"github.com/mariomac/iters"
)

func main_reduce() {
	fac8, _ := iters.Iterate(1, item.Increment[int]).
		Limit(8).
		Reduce(item.Multiply[int])
	fmt.Println("The factorial of 8 is", fac8)
}

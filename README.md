# Go iterator helper suite

[![Go Reference](https://pkg.go.dev/badge/github.com/mariomac/iters.svg)](https://pkg.go.dev/github.com/mariomac/iters)
[![Go Report Card](https://goreportcard.com/badge/github.com/mariomac/iters)](https://goreportcard.com/report/github.com/mariomac/iters)
[![codecov](https://codecov.io/github/mariomac/iters/graph/badge.svg?token=D3C6Y3OCXZ)](https://codecov.io/github/mariomac/iters)

Type safe Stream processing library inspired in the
[Gosequence library](https://github.com/mariomac/gosequence), which at the same time is heavily inspired in
the [Java Streams API](https://docs.oracle.com/javase/8/docs/api/java/util/sequence/Stream.html).

## Table of contents

* [Table of contents](#table-of-contents)
* [Requirements](#requirements)
* [Usage examples](#usage-examples)
* [Extra credits](#extra-credits)

## Requirements

* Go 1.24 or higher

## Usage examples

### Example 1: basic creation, transformation and iteration

1. Creates a literal iter.Seq containing all the integers from 1 to 11.
2. From the sequence, selects all the integers that are prime
3. Iterates the sequence. For each filtered int, prints a message.

```go
import (
  "fmt"
  "github.com/mariomac/iters"
)

func main() {
    numbers := iters.OfRange(1, 12)
    
    for n := range iters.Filter(numbers, isPrime) {
        fmt.Printf("%d is a prime number\n", n)
    }
}

func isPrime(n int) bool {
  for i := 2; i <= n/2; i++ {
    if n%i == 0 {
      return false
    }
  }
  return true
}
```

Output: 
```
1 is a prime number
2 is a prime number
3 is a prime number
5 is a prime number
7 is a prime number
11 is a prime number
```

Alternatively, you can use the `ForEach` method to iterate the sequence in a functional way:

```go
iters.ForEach(prime, func(n int) {
    fmt.Printf("%d is a prime number\n", n)
})
```

### Example 2: generation, map, limit and slice conversion

1. Creates an **infinite** sequence of random integers (no problem, the generated sequences are evaluated lazily!)
2. Divides the random integer to get a number between 1 and 6
3. Limits the infinite sequence to 5 elements.
4. Collects the sequence items as a slice.

```go
rnd := rand.New(rand.NewSource(time.Now().UnixMilli()))
fmt.Println("let me throw 5 times a dice for you")

results := iters.Map(
    iters.Generate(rnd.Int),
    func(n int) int {
        return n%6 + 1
    },
)
takeFive := iters.Limit(5, results)

fmt.Printf("results: %v\n",
	slices.Collect(takeFive))
```

Output:
```
let me throw 5 times a dice for you
results: [3 5 2 1 3]
```

### Example 3: Generation from an iterator, Map to a different type

1. Generates an infinite sequence composed by `1`, `double(1)`, `double(double(1))`, etc...
   and cut it to 6 elements.
2. Maps the numbers' sequence to a strings' sequence.
3. Converts the words sequence to a slice and prints it.


```go
func main() {
    numbers := iters.Iterate(1, double)
    sixNums := iters.Limit(6, numbers)
    words := iters.Map(sixNums, asWord)
    
    fmt.Println(slices.Collect(words))
}

func double(n int) int {
    return 2 * n
}

func asWord(n int) string {
    if n < 10 {
        return []string{"zero", "one", "two", "three", "four", "five",
            "six", "seven", "eight", "nine"}[n]
    } else {
        return "many"
    }
}
```

Output:
```
[one two four eight many many]
```

### Example 4: deduplication of elements

Following example requires to compare the elements of the `iter.Seq`, so the `iter.Seq` needs to be
composed by `comparable` elements (this is, accepted by the the `==` and `!=` operators):

1. Instantiate an `iter.Seq` of `comparable` items.
2. Pass it to the `Distinct` method, that will return a copy of the original `iter.Seq` without
   duplicates
3. Operating as any other sequence.

```go
words := iters.Distinct(
    iters.Of("hello", "hello", "!", "ho", "ho", "ho", "!"),
)

fmt.Printf("Deduplicated words: %v\n", slices.Collect(words))
```

Output:

```
Deduplicated words: [hello ! ho]
```

### Example 5: Reduce

1. Generate an incremental sequence from 1 to 8, both included.
2. Reduce all the elements multiplying them

```go
// create a sequence in range [1, 8]
seq := iters.OfRange(1, 9)
// multiply each item in the limited sequence
fac8, _ := iters.Reduce(seq8, func(a, b int) int {
    return a * b
})
fmt.Println("The factorial of 8 is", fac8)
```

Output: 

```
The factorial of 8 is 40320
```

## Extra credits

This library is a port of the [Gostream library](https://github.com/mariomac/gostream) from the same
author, adopted to work directly with the Go 1.23 `iter.Seq` and `iter.Seq2` without
intermediate types.

The Gostream processing and aggregation functions, and thus most of the functions of this library,
are heavily inspired in the
[Java Stream Specification](https://docs.oracle.com/javase/8/docs/api/java/util/sequence/Stream.html).


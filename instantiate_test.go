package iters

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	gen := Iterate(2, func(n int) int {
		return n * n
	})
	assert.Equal(t,
		[]int{2, 4, 16, 256, 65536},
		slices.Collect(Limit(5, gen)))
	// test that iterating for the second time produces the same results
	assert.Equal(t,
		[]int{2, 4, 16, 256, 65536},
		slices.Collect(Limit(5, gen)))
}

func TestGenerate(t *testing.T) {
	cnt := 0
	gen := Generate(func() int {
		cnt++
		return cnt
	})
	assert.Equal(t,
		[]int{1, 2, 3, 4, 5},
		slices.Collect(Limit(5, gen)))
}

func TestConcat(t *testing.T) {
	concat := Concat[int](
		Of(1, 2, 3, 4, 5, 6),
		Of(7, 8, 9, 10),
	)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slices.Collect(concat))
	// test that iterating for the second time produces the same results
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slices.Collect(concat))
}

func TestEmpty(t *testing.T) {
	assert.Empty(t, slices.Collect(Empty[int]()))
}

func TestSeq2KeyValues(t *testing.T) {
	// Create an iter.Seq2 from key-value pairs
	seq2 := maps.All(map[string]int{"a": 1, "b": 2, "c": 3})

	assert.Equal(t, []string{"a", "b", "c"},
		slices.Sorted(Keys(seq2)))

	assert.Equal(t, []int{1, 2, 3},
		slices.Sorted(Values(seq2)))
}

func TestOfChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	result := slices.Collect(OfChannel(ch))
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestOfSlice(t *testing.T) {
	slice := []string{"a", "b", "c"}
	result := slices.Collect(OfSlice(slice))
	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func TestOfMap(t *testing.T) {
	m := map[string]int{"x": 10, "y": 20, "z": 30}
	keys := make([]string, 0)
	values := make([]int, 0)

	for k, v := range OfMap(m) {
		keys = append(keys, k)
		values = append(values, v)
	}

	slices.Sort(keys)
	slices.Sort(values)
	assert.Equal(t, []string{"x", "y", "z"}, keys)
	assert.Equal(t, []int{10, 20, 30}, values)
}

func TestOfMapKeys(t *testing.T) {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	result := slices.Collect(OfMapKeys(m))
	slices.Sort(result)
	assert.Equal(t, []string{"bar", "baz", "foo"}, result)
}

func TestOfMapValues(t *testing.T) {
	m := map[string]int{"x": 100, "y": 200, "z": 300}
	result := slices.Collect(OfMapValues(m))
	slices.Sort(result)
	assert.Equal(t, []int{100, 200, 300}, result)
}

func TestOf(t *testing.T) {
	result := slices.Collect(Of(1, 2, 3, 4, 5))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

func TestOfRange(t *testing.T) {
	result := slices.Collect(OfRange(0, 5))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, result)

	// Test with different types
	result2 := slices.Collect(OfRange(10, 13))
	assert.Equal(t, []int{10, 11, 12}, result2)
}

func TestZip(t *testing.T) {
	keys := Of("a", "b", "c", "d")
	values := Of(1, 2, 3)

	assert.Equal(t,
		map[string]int{"a": 1, "b": 2, "c": 3},
		maps.Collect(Zip(keys, values)))

	// Test equal length sequences
	keys2 := Of("x", "y", "z")
	values2 := Of(10, 20, 30)

	assert.Equal(t,
		map[string]int{"x": 10, "y": 20, "z": 30},
		maps.Collect(Zip(keys2, values2)))
}

func TestZipInfinites_Skip2_Limit2(t *testing.T) {
	seq2 := Zip(
		Iterate(1, func(n int) int { return n + 1 }),
		Iterate(1, func(n int) int { return n * 2 }))

	tail := Skip2(3, seq2)
	body := Limit2(3, tail)

	assert.Equal(t,
		map[int]int{4: 8, 5: 16, 6: 32},
		maps.Collect(body))
}

package iters

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty2(t *testing.T) {
	assert.Empty(t, maps.Collect(Empty2[int, int]()))
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

func TestConcat2(t *testing.T) {
	concat := Concat2[int, int](
		Zip(Of(1, 2, 3), Of(4, 5, 6)),
		Zip(Of(7, 8, 9), Of(10, 11, 12)),
		Zip(Of(13, 14, 15), Of(16, 17, 18)),
	)

	assert.Equal(t, []int{1, 2, 3, 7, 8, 9, 13, 14, 15},
		slices.Collect(Keys(concat)))
	assert.Equal(t, []int{4, 5, 6, 10, 11, 12, 16, 17, 18},
		slices.Collect(Values(concat)))
}

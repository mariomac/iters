package iters

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfinite_Zip_Skip2_Limit2(t *testing.T) {
	seq2 := Zip(
		Iterate(1, func(n int) int { return n + 1 }),
		Iterate(1, func(n int) int { return n * 2 }))

	tail := Skip2(3, seq2)
	body := Limit2(3, tail)

	assert.Equal(t,
		map[int]int{4: 8, 5: 16, 6: 32},
		maps.Collect(body))
}

func TestMap2Seq(t *testing.T) {
	seq2 := Zip(
		Of(1, 2, 3),
		Of(4, 5, 6))

	assert.Equal(t, []int{5, 7, 9},
		slices.Collect(Map2Seq(seq2, func(k, v int) int { return k + v })))
}

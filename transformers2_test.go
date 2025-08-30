package iters

import (
	"maps"
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

package iters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach2(t *testing.T) {
	var keys, vals []int
	ForEach2(func(yield func(a, b int) bool) {
		for _, i := range []int{1, 2, 3, 4} {
			if !yield(i, i*2) {
				return
			}
		}
	}, func(a, b int) {
		keys = append(keys, a)
		vals = append(vals, b)
	})
	assert.Equal(t, []int{1, 2, 3, 4}, keys)
	assert.Equal(t, []int{2, 4, 6, 8}, vals)
}

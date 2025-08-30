package iters

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func add[T integer](a, b T) T {
	return a + b
}

func isZero[T comparable](input T) bool {
	var zero T
	return input == zero
}

func increment[T integer](a T) T {
	return a + 1
}

func not[T any](condition func(i T) bool) func(i T) bool {
	return func(i T) bool {
		return !condition(i)
	}
}

func TestReduce(t *testing.T) {
	// test empty iter.Seq
	_, ok := Reduce(Empty[int](), add[int])
	assert.False(t, ok)

	// test one-element iter.Seq
	red, ok := Reduce(Of(8), add[int])
	assert.True(t, ok)
	assert.Equal(t, 8, red)

	// test multi-element iter.Seq
	red, ok = Reduce(Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), add[int])
	assert.True(t, ok)
	assert.Equal(t, 55, red)
}

func TestIterableStream_AllMatch(t *testing.T) {
	// for empty iter.Seq, following Java behavior as reference
	assert.True(t, AllMatch(Empty[string](), isZero[string]))
	assert.True(t, AllMatch(Of("hello", "world"), not(isZero[string])))
	assert.False(t, AllMatch(Of("", "world"), not(isZero[string])))
}

func TestIterableStream_AnyMatch(t *testing.T) {
	// for empty iter.Seq, following Java behavior as reference
	assert.False(t, AnyMatch(Empty[string](), isZero[string]))
	assert.True(t, AnyMatch(Of("hello", "world"), not(isZero[string])))
	assert.True(t, AnyMatch(Of("", "world"), not(isZero[string])))
	assert.False(t, AnyMatch(Of("", ""), not(isZero[string])))
}

func TestIterableStream_NoneMatch(t *testing.T) {
	// for empty iter.Seq, following Java behavior as reference
	assert.True(t, NoneMatch(Empty[string](), isZero[string]))
	assert.False(t, NoneMatch(Of("hello", "world"), not(isZero[string])))
	assert.False(t, NoneMatch(Of("", "world"), not(isZero[string])))
	assert.True(t, NoneMatch(Of("", ""), not(isZero[string])))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 0, Count(Empty[int]()))
	assert.Equal(t, 0, Count(Skip(3, Of(1, 2, 3))))
	assert.Equal(t, 3, Count(Of(1, 2, 3)))
	assert.Equal(t, 3, Count(Skip(3, Of(1, 2, 3, 4, 5, 6))))
	assert.Equal(t, 8, Count(Limit(8, Iterate[int](1, increment[int]))))
}

func TestFindFirst(t *testing.T) {
	_, ok := FindFirst(Empty[int]())
	require.False(t, ok)

	_, ok = FindFirst(Skip(3, Of(1, 2, 3)))
	require.False(t, ok)

	n, ok := FindFirst(Of(1, 2, 3))
	require.True(t, ok)
	assert.Equal(t, 1, n)

	n, ok = FindFirst(Skip(3, Of(1, 2, 3, 4, 5, 6)))
	require.True(t, ok)
	assert.Equal(t, 4, n)

	n, ok = FindFirst(Limit(8, Iterate[int](1, increment[int])))
	require.True(t, ok)
	assert.Equal(t, 1, n)
}

func TestMax(t *testing.T) {
	_, ok := Max(Empty[int]())
	require.False(t, ok)

	_, ok = Max(Skip(3, Of(1, 2, 3)))
	require.False(t, ok)

	n, ok := Max(Skip(2, Of(1, 2, 3)))
	require.True(t, ok)
	assert.Equal(t, 3, n)

	n, ok = Max(Of(1, 3, 2))
	require.True(t, ok)
	assert.Equal(t, 3, n)

	n, ok = Max(Skip(3, Of(1, 2, 3, 4, 5, 6)))
	require.True(t, ok)
	assert.Equal(t, 6, n)
}

func TestMin(t *testing.T) {
	_, ok := Min(Empty[int]())
	require.False(t, ok)

	n, ok := Min(Of(1, 2, 3))
	require.True(t, ok)
	assert.Equal(t, 1, n)
}

func TestMaxFunc(t *testing.T) {
	_, ok := MaxFunc(Empty[int](), cmp.Compare[int])
	require.False(t, ok)

	_, ok = MaxFunc(Skip(3, Of(1, 2, 3)), cmp.Compare[int])
	require.False(t, ok)

	n, ok := MaxFunc(Skip(2, Of(1, 2, 3)), cmp.Compare[int])
	require.True(t, ok)
	assert.Equal(t, 3, n)

	n, ok = MaxFunc(Of(1, 3, 2), cmp.Compare[int])
	require.True(t, ok)
	assert.Equal(t, 3, n)

	n, ok = MaxFunc(Skip(3, Of(1, 2, 3, 4, 5, 6)), cmp.Compare[int])
	require.True(t, ok)
	assert.Equal(t, 6, n)
}

func TestMinFunc(t *testing.T) {
	_, ok := MinFunc(Empty[int](), cmp.Compare[int])
	require.False(t, ok)

	n, ok := MinFunc(Of(1, 2, 3), cmp.Compare[int])
	require.True(t, ok)
	assert.Equal(t, 1, n)
}

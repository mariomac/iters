package iters

import "iter"

// ForEach2 invokes the consumer function for each pair of items of the iter.Seq2
func ForEach2[T1, T2 any](input iter.Seq2[T1, T2], consumer func(T1, T2)) {
	for v1, v2 := range input {
		consumer(v1, v2)
	}
}

package iters

import (
	"iter"
	"maps"
)

// Empty2 returns an empty iter.Seq2
func Empty2[T1, T2 any]() iter.Seq2[T1, T2] {
	return func(_ func(T1, T2) bool) {}
}

// OfMap returns an iter.Seq2 that iterates over all key-value pairs of the provided map.
// It is just an alias for maps.All.
func OfMap[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return maps.All(m)
}

// Zip joins the input iter.Seq[K] and iter.Seq[V] into an iter.Seq2[K, V].
// The resulting iter.Seq2 will have the same length as the shorter of the two input iter.Seq.
func Zip[K, V any](keys iter.Seq[K], vals iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		keyNext, keyStop := iter.Pull(keys)
		defer keyStop()
		valNext, valStop := iter.Pull(vals)
		defer valStop()

		for {
			k, kOk := keyNext()
			v, vOk := valNext()

			if !kOk || !vOk {
				return
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// Concat2 creates a lazily concatenated iter.Seq2 whose elements are all the elements of the first
// provided iter.Seq2 followed by all the elements of the second provided iter.Seq2, followed by the
// elements of the third iter.Seq2 (if any), and so on.
func Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

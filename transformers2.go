package iters

import "iter"

// Limit2 returns an iter.Seq2 consisting of the elements of the input iter.Seq2, truncated to
// be no longer than maxSize in length.
func Limit2[K, V any](maxSize int, input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		count := 0
		for k, v := range input {
			if count == maxSize {
				return
			}
			if !yield(k, v) {
				return
			}
			count++
		}
	}
}

// Skip2 returns an iter.Seq consisting of the remaining elements of the input iter.Seq2 after discarding
// the first n elements of the sequence.
func Skip2[K, V any](n int, input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull2(input)
		defer stop()
		skipped := 0
		for _, _, ok := next(); ok && skipped < n-1; _, _, ok = next() {
			skipped++
		}
		for k, v, ok := next(); ok; k, v, ok = next() {
			if !yield(k, v) {
				return
			}
		}
	}
}

package slices

import "golang.org/x/exp/constraints"

// IsSortedFunc returns true if the slice is sorted by the given predicate.
func IsSortedFunc[E any](source []E, compare func(E, E) int) bool {
	for i := 1; i < len(source); i++ {
		if compare(source[i-1], source[i]) > 0 {
			return false
		}
	}
	return true
}

// IsSorted returns true if the slice is sorted.
func IsSorted[E constraints.Ordered](source []E) bool {
	return IsSortedFunc(source, compareOrdered[E])
}

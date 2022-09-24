package slices

import "golang.org/x/exp/constraints"

// IsSortedFunc returns true if the slice is sorted by the given predicate.
func IsSortedFunc[E any](slice []E, compare func(E, E) int) bool {
	for i := 1; i < len(slice); i++ {
		if compare(slice[i-1], slice[i]) > 0 {
			return false
		}
	}
	return true
}

// IsSorted returns true if the slice is sorted.
func IsSorted[E constraints.Ordered](slice []E) bool {
	return IsSortedFunc(slice, compareOrdered[E])
}

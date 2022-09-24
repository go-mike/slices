package slices

import "golang.org/x/exp/constraints"

// IsSortedFunc returns true if the slice is sorted by the given predicate.
func IsSortedFunc[E any](slice []E, less func(E, E) bool) bool {
	for i := 1; i < len(slice); i++ {
		if less(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}

// IsSorted returns true if the slice is sorted.
func IsSorted[E constraints.Ordered](slice []E) bool {
	return IsSortedFunc(slice, func(a, b E) bool { return a < b })
}

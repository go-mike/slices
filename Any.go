package slices

// Any returns true if any item in the source slice matches the predicate.
// If the source slice is nil or empty, false is returned.
func Any[E any](source []E, predicate func(E) bool) bool {
	for _, item := range source {
		if predicate(item) {
			return true
		}
	}
	return false
}

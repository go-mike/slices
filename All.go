package slices

// All returns true if all items in the source slice match the predicate.
// If the source slice is nil or empty, true is returned.
func All[E any](source []E, predicate func(E) bool) bool {
	for _, item := range source {
		if !predicate(item) {
			return false
		}
	}
	return true
}

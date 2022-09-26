package slices

// AllIndexed returns true if all items in the source slice match the predicate.
func AllIndexed[E any](source []E, predicate func(E, int) bool) bool {
	for index, item := range source {
		if !predicate(item, index) {
			return false
		}
	}
	return true
}

// All returns true if all items in the source slice match the predicate.
// If the source slice is nil or empty, true is returned.
func All[E any](source []E, predicate func(E) bool) bool {
	return AllIndexed(source, func(item E, _ int) bool { return predicate(item) })
}

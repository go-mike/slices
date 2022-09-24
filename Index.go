package slices

// Index returns the index of the first element in source that is equal to the given value.
// If the value is not found, -1 is returned.
func Index[E comparable](source []E, item E) int {
	return IndexFunc(source, func(i E) bool { return i == item })
}

// IndexFunc returns the index of the first element in source that satisfies the given predicate.
// If the value is not found, -1 is returned.
func IndexFunc[E any](source []E, predicate func(E) bool) int {
	for i, v := range source {
		if predicate(v) {
			return i
		}
	}
	return -1
}

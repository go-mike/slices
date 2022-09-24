package slices

// LastIndex returns the index of the first element in source that is equal to the given value.
// If the value is not found, -1 is returned.
func LastIndex[E comparable](source []E, item E) int {
	return LastIndexFunc(source, func(i E) bool { return i == item })
}

// LastIndexFunc returns the index of the first element in source that satisfies the given predicate.
// If the value is not found, -1 is returned.
func LastIndexFunc[E any](source []E, predicate func(E) bool) int {
	for i := len(source) - 1; i >= 0; i-- {
		if predicate(source[i]) {
			return i
		}
	}
	return -1
}

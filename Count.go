package slices

// Count returns the number of items in the source slice that are equal to the specified item.
func Count[E comparable](source []E, item E) int {
	return CountFunc(source, func(i E) bool { return i == item })
}

// CountFunc returns the number of items in the source slice that match the specified predicate.
func CountFunc[E any](source []E, predicate func(E) bool) int {
	var result int
	for _, item := range source {
		if predicate(item) {
			result++
		}
	}
	return result
}

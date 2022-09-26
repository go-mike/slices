package slices

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

// Count returns the number of items in the source slice that are equal to the specified item.
func Count[E comparable](source []E, item E) int {
	return CountFunc(source, func(i E) bool { return i == item })
}

// CountIndexedFunc returns the number of items in the source slice that match the specified predicate.
func CountIndexedFunc[E any](source []E, predicate func(E, int) bool) int {
	var result int
	for index, item := range source {
		if predicate(item, index) {
			result++
		}
	}
	return result
}

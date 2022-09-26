package slices

// IndexIndexedFunc returns the index of the first element in source that satisfies the given predicate.
// If the item is not found, -1 is returned.
func IndexIndexedFunc[E any](source []E, predicate func(E, int) bool) int {
	for index, item := range source {
		if predicate(item, index) {
			return index
		}
	}
	return -1
}

// IndexFunc returns the index of the first element in source that satisfies the given predicate.
// If the item is not found, -1 is returned.
func IndexFunc[E any](source []E, predicate func(E) bool) int {
	return IndexIndexedFunc(source, func(item E, _ int) bool { return predicate(item) })
}

// Index returns the index of the first element in source that is equal to the given item.
// If the item is not found, -1 is returned.
func Index[E comparable](source []E, item E) int {
	return IndexFunc(source, func(i E) bool { return i == item })
}

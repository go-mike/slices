package slices

// LastIndexIndexedFunc returns the index of the first element in source that satisfies the given predicate.
// The function is invoked with the index of the element as the second argument.
// If the item is not found, -1 is returned.
func LastIndexIndexedFunc[E any](source []E, predicate func(E, int) bool) int {
	for index := len(source) - 1; index >= 0; index-- {
		if predicate(source[index], index) {
			return index
		}
	}
	return -1
}

// LastIndexFunc returns the index of the first element in source that satisfies the given predicate.
// If the item is not found, -1 is returned.
func LastIndexFunc[E any](source []E, predicate func(E) bool) int {
	return LastIndexIndexedFunc(source, func(item E, _ int) bool { return predicate(item) })
}

// LastIndex returns the index of the first element in source that is equal to the given item.
// If the item is not found, -1 is returned.
func LastIndex[E comparable](source []E, item E) int {
	return LastIndexFunc(source, func(i E) bool { return i == item })
}

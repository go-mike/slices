package slices

// IndicesIndexedFunc returns the indices of all elements in source that satisfy the given predicate.
// The function is invoked with the index of the element as the second argument.
func IndicesIndexedFunc[E any](source []E, predicate func(E, int) bool) []int {
	result := make([]int, 0)
	for index, item := range source {
		if predicate(item, index) {
			result = append(result, index)
		}
	}
	return result
}

// IndicesFunc returns the indices of all elements in source that satisfy the given predicate.
func IndicesFunc[E any](source []E, predicate func(E) bool) []int {
	return IndicesIndexedFunc(source, func(item E, _ int) bool { return predicate(item) })
}

// Indices returns the indices of all elements in source that are equal to the given item.
func Indices[E comparable](source []E, item E) []int {
	return IndicesFunc(source, func(i E) bool { return i == item })
}

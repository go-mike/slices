package slices

// IndicesFunc returns the indices of all elements in source that satisfy the given predicate.
func IndicesFunc[E any](source []E, predicate func(E) bool) []int {
	result := make([]int, 0)
	for i, v := range source {
		if predicate(v) {
			result = append(result, i)
		}
	}
	return result
}

// Indices returns the indices of all elements in source that are equal to the given value.
func Indices[E comparable](source []E, item E) []int {
	return IndicesFunc(source, func(i E) bool { return i == item })
}

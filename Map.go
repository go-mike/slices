package slices

// MapIndexed returns a new slice containing the results of applying the given function to each element of the source slice.
// The function is invoked with the index of the element as the second argument.
func MapIndexed[E, R any](source []E, mapper func(E, int) R) []R {
	var result []R
	for index, item := range source {
		result = append(result, mapper(item, index))
	}
	return result
}

// Map returns a new slice containing the results of applying the given function to each element of the source slice.
func Map[E, R any](source []E, mapper func(E) R) []R {
	return MapIndexed(source, func(item E, _ int) R { return mapper(item) })
}

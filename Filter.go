package slices

// FilterIndexed returns a new slice containing only the items from the source slice that match the specified predicate.
// The function is invoked with the index of the element as the second argument.
func FilterIndexed[E any](source []E, predicate func(E, int) bool) []E {
	var result []E
	for index, item := range source {
		if predicate(item, index) {
			result = append(result, item)
		}
	}
	return result
}

// Filter returns a new slice containing only the items from the source slice that match the specified predicate.
func Filter[E any](source []E, predicate func(E) bool) []E {
	return FilterIndexed(source, func(item E, _ int) bool { return predicate(item) })
}

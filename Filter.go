package slices

// Filter returns a new slice containing only the items from the source slice that match the specified predicate.
func Filter[E any](source []E, predicate func(E) bool) []E {
	var result []E
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

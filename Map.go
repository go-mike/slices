package slices

// Map returns a new slice containing the results of applying the given function to each element of the source slice.
func Map[E, R any](source []E, mapper func(E) R) []R {
	var result []R
	for _, item := range source {
		result = append(result, mapper(item))
	}
	return result
}

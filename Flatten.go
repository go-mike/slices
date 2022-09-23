package slices

// Flatten returns a new slice containing all the items from the source slice of slices.
func Flatten[E any](source [][]E) []E {
	var result []E
	for _, item := range source {
		result = append(result, item...)
	}
	return result
}

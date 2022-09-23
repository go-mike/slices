package slices

// Reverse returns a new slice with the elements of the source slice in reverse order.
func Reverse[E any](source []E) []E {
	result := make([]E, len(source))
	for i := 0; i < len(source); i++ {
		result[i] = source[len(source)-i-1]
	}
	return result
}

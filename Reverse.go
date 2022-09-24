package slices

// Reverse returns a new slice with the elements of the source slice in reverse order.
func Reverse[E any](source []E) []E {
	result := make([]E, len(source))
	for i := 0; i < len(source); i++ {
		result[i] = source[len(source)-i-1]
	}
	return result
}

// ReverseInPlace reverses the elements of the source slice in place.
func ReverseInPlace[E any](source []E) {
	for i := 0; i < len(source)/2; i++ {
		source[i], source[len(source)-i-1] = source[len(source)-i-1], source[i]
	}
}

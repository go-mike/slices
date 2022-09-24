package slices

// TrimCapacity returns a slice with all extra capacity removed.
func TrimCapacity[E any](source []E) []E {
	return source[:len(source):len(source)]
}

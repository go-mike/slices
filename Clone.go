package slices

// Clone returns a copy of the source slice.
func Clone[E any](source []E) []E {
	var result []E = make([]E, len(source))
	copy(result, source)
	return result
}

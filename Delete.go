package slices

// Delete deletes the elements of source, starting at start, with count elements.
// It is equivalent to replacing count elements of source with 0 elements, starting at start.
func Delete[E any](source []E, start int, count int) []E {
	return Replace(source, start, count, nil)
}

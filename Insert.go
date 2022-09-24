package slices

// Insert inserts the elements of source into the destination slice at the specified start.
// It is equivalent to replacing 0 elements of the destination slice with the elements of source, starting at start.
func Insert[E any](destination []E, start int, source []E) []E {
	return Replace(destination, start, 0, source)
}

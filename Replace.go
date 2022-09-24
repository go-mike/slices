package slices

// Replace replaces the elements of source with the elements of replacement, starting at start, with count elements.
// It is equivalent to first deleting count elements from source, starting at start, and then inserting replacement at start.
func Replace[E any](source []E, replacement []E, start int, count int) []E {
	return Insert(Delete(source, start, count), replacement, start)
}
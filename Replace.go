package slices

// Replace replaces the elements of source with the elements of replacement, starting at start, with count elements.
// It is equivalent to first deleting count elements from source, starting at start, and then inserting replacement at start.
func Replace[E any](source []E, start int, count int, replacement []E) []E {
	if start+count > len(source) {
		count = len(source) - start
	}
	if start < 0 {
		count += start
		start = 0
	}
	if start > len(source) {
		start = len(source)
		count = 0
	}
	if count < 0 {
		count = 0
	}
	result := make([]E, 0, len(source)-count+len(replacement))
	result = append(result, source[:start]...)
	result = append(result, replacement...)
	result = append(result, source[start+count:]...)
	return result
}
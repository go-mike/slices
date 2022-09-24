package slices

// Delete deletes the elements of source, starting at start, with count elements.
func Delete[E any](source []E, start int, count int) []E {
	if start+count > len(source) {
		count = len(source) - start
	}
	if start < 0 {
		count += start
		start = 0
	}
	if count <= 0 {
		return source
	}
	result := make([]E, 0, len(source)-count)
	result = append(result, source[:start]...)
	result = append(result, source[start+count:]...)
	return result
}

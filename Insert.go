package slices

// Insert inserts the elements of source into the destination slice at the specified position.
func Insert[E any](destination []E, source []E, position int) []E {
	if position < 0 {
		position = 0
	}
	if position > len(destination) {
		position = len(destination)
	}
	result := make([]E, 0, len(destination)+len(source))
	result = append(result, destination[:position]...)
	result = append(result, source...)
	result = append(result, destination[position:]...)
	return result
}

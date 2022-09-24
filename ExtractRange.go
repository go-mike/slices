package slices

// ExtractRangeStep returns a slice of elements from source, starting at start, with count elements, and with a step of step.
func ExtractRangeStep[E any](source []E, start int, count int, step int) []E {
	if count <= 0 {
		return []E{}
	}
	result := make([]E, 0, count)
	for i := 0; i < count; i++ {
		pos := start + i*step
		if pos >= 0 && pos < len(source) {
			result = append(result, source[pos])
		}
	}
	return result
}

// ExtractRange returns a slice of elements from source, starting at start, with count elements.
func ExtractRange[E any](source []E, start int, count int) []E {
	return ExtractRangeStep(source, start, count, 1)
}

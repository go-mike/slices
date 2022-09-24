package slices

// Grow grows the slice by the specified amount.
func Grow[E any](source []E, amount int) []E {
	if amount < 0 {
		return source
	}
	result := make([]E, len(source)+amount)
	copy(result, source)
	return result
}

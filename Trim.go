package slices

// TrimStartFunc returns a slice with all leading elements that satisfy the predicate removed.
func TrimStartFunc[E any](source []E, predicate func(E) bool) []E {
	result := []E{}
	start := true
	for _, item := range source {
		if start {
			if !predicate(item) {
				start = false
				result = append(result, item)
			}
		} else {
			result = append(result, item)
		}
	}
	return result
}

// TrimStart returns a slice with all leading elements that are equal to the specified element removed.
func TrimStart[E comparable](source []E, element E) []E {
	return TrimStartFunc(source, func(item E) bool { return item == element })
}

// TrimEndFunc returns a slice with all trailing elements that satisfy the predicate removed.
func TrimEndFunc[E any](source []E, predicate func(E) bool) []E {
	result := []E{}
	end := true
	for i := len(source) - 1; i >= 0; i-- {
		if end {
			if !predicate(source[i]) {
				end = false
				result = append([]E{source[i]}, result...)
			}
		} else {
			result = append([]E{source[i]}, result...)
		}
	}
	return result
}

// TrimEnd returns a slice with all trailing elements that are equal to the specified element removed.
func TrimEnd[E comparable](source []E, element E) []E {
	return TrimEndFunc(source, func(item E) bool { return item == element })
}

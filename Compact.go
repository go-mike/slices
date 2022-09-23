package slices

// Compact returns a new slice with all duplicate adjacent items removed.
func Compact[E comparable](source []E) []E {
	return CompactFunc(source, func(a, b E) bool { return a == b })
}

// CompactFunc returns a new slice with all duplicate adjacent items removed.
// The predicate is used to determine if two items are the same or equals.
func CompactFunc[E any](source []E, equals func(E, E) bool) []E {
	if len(source) <= 1 {
		return source
	}
	var result []E = []E{source[0]}
	for _, item := range source[1:] {
		if !equals(item, result[len(result)-1]) {
			result = append(result, item)
		}
	}
	return result
}

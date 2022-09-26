package slices

// ContainsFuncIndexed returns true if the source slice contains an item that matches the specified predicate.
func ContainsFuncIndexed[E any](source []E, predicate func(E, int) bool) bool {
	return IndexIndexedFunc(source, predicate) >= 0
}

// ContainsFunc returns true if the source slice contains an item that matches the specified predicate.
func ContainsFunc[E any](source []E, predicate func(E) bool) bool {
	return IndexFunc(source, predicate) >= 0
}

// Contains returns true if the source slice contains the specified item.
func Contains[E comparable](source []E, item E) bool {
	return ContainsFunc(source, func(i E) bool { return i == item })
}

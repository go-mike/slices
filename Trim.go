package slices

// TrimStartIndexedFunc returns a slice with all leading elements that satisfy the predicate removed.
// The function is invoked with the index of the element as the second argument.
func TrimStartIndexedFunc[E any](source []E, predicate func(E, int) bool) []E {
	if len(source) == 0 {
		return []E{}
	}
	pos := IndexIndexedFunc(source, func(item E, index int) bool { return !predicate(item, index) })
	if pos == 0 {
		return source
	}
	if pos < 0 {
		return []E{}
	}
	return source[pos:]
}

// TrimEndIndexedFunc returns a slice with all trailing elements that satisfy the predicate removed.
// The function is invoked with the index of the element as the second argument.
func TrimEndIndexedFunc[E any](source []E, predicate func(E, int) bool) []E {
	if len(source) == 0 {
		return []E{}
	}
	pos := LastIndexIndexedFunc(source, func(item E, index int) bool { return !predicate(item, index) })
	if pos == len(source)-1 {
		return source
	}
	if pos < 0 {
		return []E{}
	}
	return source[:pos+1]
}

// TrimIndexedFunc returns a slice with all leading and trailing elements that satisfy the predicate removed.
// The function is invoked with the index of the element as the second argument.
func TrimIndexedFunc[E any](source []E, predicate func(E, int) bool) []E {
	return TrimStartIndexedFunc(TrimEndIndexedFunc(source, predicate), predicate)
}

// TrimFunc returns a slice with all leading and trailing elements that satisfy the predicate removed.
func TrimFunc[E any](source []E, predicate func(E) bool) []E {
	indexed := func(item E, _ int) bool { return predicate(item) }
	return TrimIndexedFunc(source, indexed)
}

// Trim returns a slice with all leading and trailing elements that are equal to the specified element removed.
func Trim[E comparable](source []E, element E) []E {
	return TrimIndexedFunc(source, func(item E, _ int) bool { return item == element })
}

// TrimStartFunc returns a slice with all leading elements that satisfy the predicate removed.
func TrimStartFunc[E any](source []E, predicate func(E) bool) []E {
	return TrimStartIndexedFunc(source, func(item E, _ int) bool { return predicate(item) })
}

// TrimStart returns a slice with all leading elements that are equal to the specified element removed.
func TrimStart[E comparable](source []E, element E) []E {
	return TrimStartFunc(source, func(item E) bool { return item == element })
}

// TrimEndFunc returns a slice with all trailing elements that satisfy the predicate removed.
func TrimEndFunc[E any](source []E, predicate func(E) bool) []E {
	return TrimEndIndexedFunc(source, func(item E, _ int) bool { return predicate(item) })
}

// TrimEnd returns a slice with all trailing elements that are equal to the specified element removed.
func TrimEnd[E comparable](source []E, element E) []E {
	return TrimEndFunc(source, func(item E) bool { return item == element })
}

package slices

// GroupBy returns a map of slices, where each slice contains all the items from
// the source slice where the specified key selector returns the same key.
func GroupBy[E any, K comparable](source []E, keySelector func(E) K) map[K][]E {
	var result map[K][]E = make(map[K][]E)
	for _, item := range source {
		var key K = keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}

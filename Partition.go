package slices

// Partition returns two slices, the first containing the elements for which the predicate
// returns true, the second containing the elements for which the predicate returns false.
func Partition[E any](source []E, predicate func(E) bool) (trueValues []E, falseValues []E) {
	for _, item := range source {
		if predicate(item) {
			trueValues = append(trueValues, item)
		} else {
			falseValues = append(falseValues, item)
		}
	}
	return
}

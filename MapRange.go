package slices

// MapRange returns a new slice containing the results of applying the given function to each element of the range.
func MapRange[R any](start, count int, mapper func(int) R) []R {
	return MapRangeStep(start, count, 1, mapper)
}

// MapRangeStep returns a new slice containing the results of applying the given function to each element of the range.
func MapRangeStep[R any](start, count, step int, mapper func(int) R) []R {
	var result []R = make([]R, count)
	for i := 0; i < count; i++ {
		result[i] = mapper(start+i*step)
	}
	return result
}

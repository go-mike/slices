package slices

// ReduceIndexed returns the result of applying the reducer function to each element of the source, passing the result of the previous application to the next application.
// The function is invoked with the index of the element as the third argument.
func ReduceIndexed[E, R any](source []E, seed R, reducer func(R, E, int) R) R {
	for index, item := range source {
		seed = reducer(seed, item, index)
	}
	return seed
}

// ReduceRightIndexed returns the result of applying the reducer function to each element of the source, passing the result of the previous application to the next application, from the end of the slice to the beginning.
// The function is invoked with the index of the element as the third argument.
func ReduceRightIndexed[E, R any](source []E, seed R, reducer func(R, E, int) R) R {
	for index := len(source) - 1; index >= 0; index-- {
		seed = reducer(seed, source[index], index)
	}
	return seed
}

// Reduce returns the result of applying the reducer function to each element of the source, passing the result of the previous application to the next application.
func Reduce[E, R any](source []E, seed R, reducer func(R, E) R) R {
	for _, item := range source {
		seed = reducer(seed, item)
	}
	return seed
}

// ReduceRight returns the result of applying the reducer function to each element of the source, passing the result of the previous application to the next application, from the end of the slice to the beginning.
func ReduceRight[E, R any](source []E, seed R, reducer func(R, E) R) R {
	for i := len(source) - 1; i >= 0; i-- {
		seed = reducer(seed, source[i])
	}
	return seed
}

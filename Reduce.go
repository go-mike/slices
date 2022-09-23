package slices

func Reduce[E, R any](source []E, seed R, reducer func(R, E) R) R {
	for _, item := range source {
		seed = reducer(seed, item)
	}
	return seed
}

func ReduceRight[E, R any](source []E, seed R, reducer func(R, E) R) R {
	for i := len(source) - 1; i >= 0; i-- {
		seed = reducer(seed, source[i])
	}
	return seed
}

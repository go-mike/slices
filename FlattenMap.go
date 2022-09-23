package slices

func FlattenMap[E, R any](source []E, mapper func(E) []R) []R {
	return Flatten(Map(source, mapper))
}

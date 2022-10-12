package slices

import "golang.org/x/exp/constraints"

// QuickSortFunc sorts the given slice using the given comparison function.
func QuickSortFunc[E any](source []E, compare func(E, E) int) []E {
	if len(source) <= 1 {
		return source
	}

	pivot := source[0]
	left := make([]E, 0, len(source))
	right := make([]E, 0, len(source))
	for _, element := range source[1:] {
		if compare(element, pivot) <= 0 {
			left = append(left, element)
		} else {
			right = append(right, element)
		}
	}

	left = QuickSortFunc(left, compare)
	right = QuickSortFunc(right, compare)
	return append(append(left, pivot), right...)
}

// QuickSort sorts the given slice.
func QuickSort[E constraints.Ordered](source []E) []E {
	return QuickSortFunc(source, compareOrdered[E])
}

// QuickSortInPlaceFunc sorts the given slice in place using the given comparison function.
func QuickSortInPlaceFunc[E any](source []E, compare func(E, E) int) {
	sorted := QuickSortFunc(source, compare)
	copy(source, sorted)
}

// QuickSortInPlace sorts the given slice in place.
func QuickSortInPlace[E constraints.Ordered](source []E) {
	QuickSortInPlaceFunc(source, compareOrdered[E])
}

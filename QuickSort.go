package slices

import "golang.org/x/exp/constraints"

// QuickSortInPlaceFunc sorts the given slice in place using the given comparison function.
func QuickSortInPlaceFunc[E any](source []E, compare func(E, E) int) {
	if len(source) <= 1 {
		return
	}

	pivot := source[0]
	countLeft := CountFunc(source[1:], func(element E) bool { return compare(element, pivot) <= 0 })
	source[0], source[countLeft] = source[countLeft], source[0]
	rightIndex := countLeft + 1
	for i := 0; i < countLeft; i++ {
		if compare(source[i], pivot) > 0 {
			for compare(source[rightIndex], pivot) > 0 {
				rightIndex++
			}
			source[i], source[rightIndex] = source[rightIndex], source[i]
		}
	}
	QuickSortInPlaceFunc(source[:countLeft], compare)
	QuickSortInPlaceFunc(source[countLeft+1:], compare)
}

// QuickSortInPlace sorts the given slice in place.
func QuickSortInPlace[E constraints.Ordered](source []E) {
	QuickSortInPlaceFunc(source, compareOrdered[E])
}

// QuickSortFunc sorts the given slice using the given comparison function.
func QuickSortFunc[E any](source []E, compare func(E, E) int) []E {
	result := make([]E, len(source))
	copy(result, source)
	QuickSortInPlaceFunc(result, compare)
	return result
}

// QuickSort sorts the given slice.
func QuickSort[E constraints.Ordered](source []E) []E {
	return QuickSortFunc(source, compareOrdered[E])
}

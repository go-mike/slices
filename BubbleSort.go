package slices

import "golang.org/x/exp/constraints"

// BubbleSortInPlaceFunc is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortInPlaceFunc[E any](source []E, compare func(E, E) int) {
	for i := len(source) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if compare(source[j], source[j+1]) > 0 {
				source[j], source[j+1] = source[j+1], source[j]
			}
		}
	}
}

// BubbleSortInPlace is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortInPlace[E constraints.Ordered](source []E) {
	BubbleSortInPlaceFunc(source, compareOrdered[E])
}

// BubbleSortFunc is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortFunc[E any](source []E, compare func(E, E) int) []E {
	clone := Clone(source)
	BubbleSortInPlaceFunc(clone, compare)
	return clone
}

// BubbleSort is a function that sorts a slice using the bubble sort algorithm.
func BubbleSort[E constraints.Ordered](source []E) []E {
	return BubbleSortFunc(source, compareOrdered[E])
}

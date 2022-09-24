package slices

import "golang.org/x/exp/constraints"

// BubbleSortInPlaceFunc is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortInPlaceFunc[E any](slice []E, compare func(E, E) int) {
	for i := len(slice) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if compare(slice[j], slice[j+1]) > 0 {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// BubbleSortInPlace is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortInPlace[E constraints.Ordered](slice []E) {
	BubbleSortInPlaceFunc(slice, compareOrdered[E])
}

// BubbleSortFunc is a function that sorts a slice using the bubble sort algorithm.
func BubbleSortFunc[E any](slice []E, compare func(E, E) int) []E {
	clone := Clone(slice)
	BubbleSortInPlaceFunc(clone, compare)
	return clone
}

// BubbleSort is a function that sorts a slice using the bubble sort algorithm.
func BubbleSort[E constraints.Ordered](slice []E) []E {
	return BubbleSortFunc(slice, compareOrdered[E])
}

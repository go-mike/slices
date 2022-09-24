package slices

import "golang.org/x/exp/constraints"

// MergeSortedFunc is a function that merges two sorted slices into a single sorted slice.
func MergeSortedFunc[E any](left []E, right []E, compare func(E, E) int) []E {
	merged := make([]E, len(left)+len(right))
	i := 0
	j := 0
	for k := 0; k < len(merged); k++ {
		if i >= len(left) {
			merged[k] = right[j]
			j++
		} else if j >= len(right) {
			merged[k] = left[i]
			i++
		} else if compare(left[i], right[j]) <= 0 {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

// MergeSorted is a function that merges two sorted slices into a single sorted slice.
func MergeSorted[E constraints.Ordered](left []E, right []E) []E {
	return MergeSortedFunc(left, right, compareOrdered[E])
}

// MergeSortFunc is a function that sorts a slice using the merge sort algorithm.
func MergeSortFunc[E any](slice []E, compare func(E, E) int) []E {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	left := MergeSortFunc(slice[:mid], compare)
	right := MergeSortFunc(slice[mid:], compare)

	return MergeSortedFunc(left, right, compare)
}

// MergeSort is a function that sorts a slice using the merge sort algorithm.
func MergeSort[E constraints.Ordered](slice []E) []E {
	return MergeSortFunc(slice, compareOrdered[E])
}

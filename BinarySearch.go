package slices

import "golang.org/x/exp/constraints"

// BinarySearchFunc returns the index where an element is found, or would be found, in the given slice.
// The slice must be sorted.
// The given predicate must return 0 if the element is found, -1 if the element is less than the current element,
// and 1 if the element is greater than the current element.
// The function returns a second item which is true if the element was found.
func BinarySearchFunc[E any](source []E, compare func(E) int) (int, bool) {
	low := 0
	high := len(source) - 1
	for low <= high {
		mid := (low + high) / 2
		result := compare(source[mid])
		if result < 0 {
			high = mid - 1
		} else if result > 0 {
			low = mid + 1
		} else {
			return mid, true
		}
	}
	return low, false
}

// BinarySearch returns the index where an element is found, or would be found, in the given slice.
// The slice must be sorted.
// The function returns a second item which is true if the element was found.
func BinarySearch[E constraints.Ordered](source []E, element E) (int, bool) {
	return BinarySearchFunc(source, func(item E) int {
		if item < element {
			return 1
		}
		if item > element {
			return -1
		}
		return 0
	})
}

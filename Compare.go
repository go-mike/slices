package slices

import (
	"math"

	"golang.org/x/exp/constraints"
)

// CompareFunc returns 1 if the first slice is greater than the second,
// -1 if the first slice is less than the second, and 0 if the slices are equal,
// by the given comparison function
func CompareFunc[E any](a []E, b []E, compare func(E, E) int) int {
	minLen := int(math.Min(float64(len(a)), float64(len(b))))
	for i := 0; i < minLen; i++ {
		result := compare(a[i], b[i])
		if result != 0 {
			return result
		}
	}
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	return 0
}

func compareOrdered[E constraints.Ordered](a E, b E) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// Compare returns 1 if the first slice is greater than the second,
// -1 if the first slice is less than the second, and 0 if the slices are equal.
func Compare[E constraints.Ordered](s1, s2 []E) int {
	return CompareFunc(s1, s2, compareOrdered[E])
}

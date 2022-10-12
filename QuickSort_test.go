package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_QuickSortFunc tests the quick sort algorithm.
func Test_QuickSortFunc(t *testing.T) {
	slice := []int{45, 20, 5, 35, 15, 30, 25, 10}
	sorted := QuickSortFunc(slice, compareOrdered[int])
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, sorted)
}

// Test_QuickSort tests the quick sort algorithm.
func Test_QuickSort(t *testing.T) {
	slice := []int{45, 20, 5, 35, 15, 30, 25, 10}
	sorted := QuickSort(slice)
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, sorted)
}

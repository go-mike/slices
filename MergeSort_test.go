package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_MergeSortedFunc_LeftSmaller tests the case where the left slice is smaller than the right slice.
func Test_MergeSortedFunc_LeftSmaller(t *testing.T) {
	left := []int{10, 20, 30}
	right := []int{5, 15, 25, 35, 45}
	merged := MergeSortedFunc(left, right, compareOrdered[int])
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, merged)
}

// Test_MergeSortedFunc_RightSmaller tests the case where the right slice is smaller than the left slice.
func Test_MergeSortedFunc_RightSmaller(t *testing.T) {
	left := []int{5, 15, 25, 35, 45}
	right := []int{10, 20, 30}
	merged := MergeSortedFunc(left, right, compareOrdered[int])
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, merged)
}

// Fuzz_MergeSortFunc tests the merge sort algorithm.
func Fuzz_MergeSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, slice []byte) {
		sorted := MergeSortFunc(slice, func(a, b byte) int { return int(a) - int(b) })
		assert.True(t, IsSortedFunc(sorted, compareOrdered[byte]))
	})
}

// Test_MergeSorted_LeftSmaller tests the case where the left slice is smaller than the right slice.
func Test_MergeSorted_LeftSmaller(t *testing.T) {
	left := []int{10, 20, 30}
	right := []int{5, 15, 25, 35, 45}
	merged := MergeSorted(left, right)
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, merged)
}

// Test_MergeSorted_RightSmaller tests the case where the right slice is smaller than the left slice.
func Test_MergeSorted_RightSmaller(t *testing.T) {
	left := []int{5, 15, 25, 35, 45}
	right := []int{10, 20, 30}
	merged := MergeSorted(left, right)
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, merged)
}

// Test_MergeSortFunc tests the merge sort algorithm.
func Test_MergeSortFunc(t *testing.T) {
	slice := []int{45, 20, 5, 35, 15, 30, 25, 10}
	sorted := MergeSortFunc(slice, compareOrdered[int])
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, sorted)
}

// Test_MergeSort tests the merge sort algorithm.
func Test_MergeSort(t *testing.T) {
	slice := []int{45, 20, 5, 35, 15, 30, 25, 10}
	sorted := MergeSort(slice)
	assert.Equal(t, []int{5, 10, 15, 20, 25, 30, 35, 45}, sorted)
}

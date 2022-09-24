package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_IsSortedFunc_Empty tests the case where the slice is empty.
func Test_IsSortedFunc_Empty(t *testing.T) {
	slice := []int{}
	result := IsSortedFunc(slice, func(a, b int) int { return a - b })
	assert.True(t, result)
}

// Test_IsSortedFunc_Sorted tests the case where the slice is sorted.
func Test_IsSortedFunc_Sorted(t *testing.T) {
	slice := []int{1, 2, 3}
	result := IsSortedFunc(slice, func(a, b int) int { return a - b })
	assert.True(t, result)
}

// Test_IsSortedFunc_Equal tests the case where the slice contains equal elements.
func Test_IsSortedFunc_Equal(t *testing.T) {
	slice := []int{1, 2, 2, 3}
	result := IsSortedFunc(slice, func(a, b int) int { return a - b })
	assert.True(t, result)
}

// Test_IsSortedFunc_Unsorted tests the case where the slice is unsorted.
func Test_IsSortedFunc_Unsorted(t *testing.T) {
	slice := []int{1, 3, 2}
	result := IsSortedFunc(slice, func(a, b int) int { return a - b })
	assert.False(t, result)
}

// Test_IsSorted_Empty tests the case where the slice is empty.
func Test_IsSorted_Empty(t *testing.T) {
	slice := []int{}
	result := IsSorted(slice)
	assert.True(t, result)
}

// Test_IsSorted_Sorted tests the case where the slice is sorted.
func Test_IsSorted_Sorted(t *testing.T) {
	slice := []int{1, 2, 3}
	result := IsSorted(slice)
	assert.True(t, result)
}

// Test_IsSorted_Unsorted tests the case where the slice is unsorted.
func Test_IsSorted_Unsorted(t *testing.T) {
	slice := []int{1, 3, 2}
	result := IsSorted(slice)
	assert.False(t, result)
}

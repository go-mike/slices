package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_BinarySearchFunc_Empty tests the case where the slice is empty.
func Test_BinarySearchFunc_Empty(t *testing.T) {
	slice := []int{}
	index, found := BinarySearchFunc(slice, func(int) int { return 0 })
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearchFunc_SingleMatch tests the case where the slice contains a single element that matches.
func Test_BinarySearchFunc_SingleMatch(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearchFunc(slice, func(int) int { return 0 })
	assert.Equal(t, 0, index)
	assert.True(t, found)
}

// Test_BinarySearchFunc_SingleBelowTarget tests the case where the slice contains a single element that is below the target.
func Test_BinarySearchFunc_SingleBelowTarget(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearchFunc(slice, func(int) int { return 1 })
	assert.Equal(t, 1, index)
	assert.False(t, found)
}

// Test_BinarySearchFunc_SingleAboveTarget tests the case where the slice contains a single element that is above the target.
func Test_BinarySearchFunc_SingleAboveTarget(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearchFunc(slice, func(int) int { return -1 })
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearchFunc_Match tests the case where the target is in the middle of the slice, and is a match.
func Test_BinarySearchFunc_Match(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 40 - a })
	assert.Equal(t, 3, index)
	assert.True(t, found)
}

// Test_BinarySearchFunc_Middle tests the case where the target is in the middle of the slice, but not a match.
func Test_BinarySearchFunc_Middle(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 35 - a })
	assert.Equal(t, 3, index)
	assert.False(t, found)
}

// Test_BinarySearchFunc_First tests the case where the target is the first element of the slice.
func Test_BinarySearchFunc_First(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 10 - a })
	assert.Equal(t, 0, index)
	assert.True(t, found)
}

// Test_BinarySearchFunc_Last tests the case where the target is the last element of the slice.
func Test_BinarySearchFunc_Last(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 50 - a })
	assert.Equal(t, 4, index)
	assert.True(t, found)
}

// Test_BinarySearchFunc_BelowFirst tests the case where the target is below the first element of the slice.
func Test_BinarySearchFunc_BelowFirst(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 5 - a })
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearchFunc_AboveLast tests the case where the target is above the last element of the slice.
func Test_BinarySearchFunc_AboveLast(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearchFunc(slice, func(a int) int { return 55 - a })
	assert.Equal(t, 5, index)
	assert.False(t, found)
}

// Test_BinarySearch_Empty tests the case where the slice is empty.
func Test_BinarySearch_Empty(t *testing.T) {
	slice := []int{}
	index, found := BinarySearch(slice, 0)
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearch_SingleMatch tests the case where the slice contains a single element that matches.
func Test_BinarySearch_SingleMatch(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearch(slice, 1)
	assert.Equal(t, 0, index)
	assert.True(t, found)
}

// Test_BinarySearch_SingleBelowTarget tests the case where the slice contains a single element that is below the target.
func Test_BinarySearch_SingleBelowTarget(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearch(slice, 2)
	assert.Equal(t, 1, index)
	assert.False(t, found)
}

// Test_BinarySearch_SingleAboveTarget tests the case where the slice contains a single element that is above the target.
func Test_BinarySearch_SingleAboveTarget(t *testing.T) {
	slice := []int{1}
	index, found := BinarySearch(slice, 0)
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearch_Match tests the case where the target is in the middle of the slice, and is a match.
func Test_BinarySearch_Match(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 40)
	assert.Equal(t, 3, index)
	assert.True(t, found)
}

// Test_BinarySearch_Middle tests the case where the target is in the middle of the slice, but not a match.
func Test_BinarySearch_Middle(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 35)
	assert.Equal(t, 3, index)
	assert.False(t, found)
}

// Test_BinarySearch_First tests the case where the target is the first element of the slice.
func Test_BinarySearch_First(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 10)
	assert.Equal(t, 0, index)
	assert.True(t, found)
}

// Test_BinarySearch_Last tests the case where the target is the last element of the slice.
func Test_BinarySearch_Last(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 50)
	assert.Equal(t, 4, index)
	assert.True(t, found)
}

// Test_BinarySearch_BelowFirst tests the case where the target is below the first element of the slice.
func Test_BinarySearch_BelowFirst(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 5)
	assert.Equal(t, 0, index)
	assert.False(t, found)
}

// Test_BinarySearch_AboveLast tests the case where the target is above the last element of the slice.
func Test_BinarySearch_AboveLast(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	index, found := BinarySearch(slice, 55)
	assert.Equal(t, 5, index)
	assert.False(t, found)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_TrimStartFunc_EmptySource tests the case where the source slice is empty.
func Test_TrimStartFunc_EmptySource(t *testing.T) {
	var source []int
	result := TrimStartFunc(source, func(item int) bool { return item < 3 })
	assert.Equal(t, []int{}, result)
}

// Test_TrimStartFunc_NotFound tests the case where the predicate is not satisfied by any element at the beginning of the source slice.
func Test_TrimStartFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartFunc(source, func(item int) bool { return item > 3 })
	assert.Equal(t, source, result)
}

// Test_TrimStartFunc_Found tests the case where the predicate is satisfied by some elements at the beginning of the source slice.
func Test_TrimStartFunc_Found(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartFunc(source, func(item int) bool { return item < 3 })
	assert.Equal(t, []int{3, 4, 5}, result)
}

// Test_TrimStart_EmptySource tests the case where the source slice is empty.
func Test_TrimStart_EmptySource(t *testing.T) {
	var source []int
	result := TrimStart(source, 3)
	assert.Equal(t, []int{}, result)
}

// Test_TrimStart_NotFound tests the case where the element is not found at the beginning of the source slice.
func Test_TrimStart_NotFound(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStart(source, 6)
	assert.Equal(t, source, result)
}

// Test_TrimStart_Found tests the case where the element is found at the beginning of the source slice.
func Test_TrimStart_Found(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStart(source, 1)
	assert.Equal(t, []int{2, 3, 4, 5}, result)
}

// Test_TrimEndFunc_EmptySource tests the case where the source slice is empty.
func Test_TrimEndFunc_EmptySource(t *testing.T) {
	var source []int
	result := TrimEndFunc(source, func(item int) bool { return item > 3 })
	assert.Equal(t, []int{}, result)
}

// Test_TrimEndFunc_NotFound tests the case where the predicate is not satisfied by any element at the end of the source slice.
func Test_TrimEndFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndFunc(source, func(item int) bool { return item < 3 })
	assert.Equal(t, source, result)
}

// Test_TrimEndFunc_Found tests the case where the predicate is satisfied by some elements at the end of the source slice.
func Test_TrimEndFunc_Found(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndFunc(source, func(item int) bool { return item > 3 })
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_TrimEnd_EmptySource tests the case where the source slice is empty.
func Test_TrimEnd_EmptySource(t *testing.T) {
	var source []int
	result := TrimEnd(source, 3)
	assert.Equal(t, []int{}, result)
}

// Test_TrimEnd_NotFound tests the case where the element is not found at the end of the source slice.
func Test_TrimEnd_NotFound(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEnd(source, 6)
	assert.Equal(t, source, result)
}

// Test_TrimEnd_Found tests the case where the element is found at the end of the source slice.
func Test_TrimEnd_Found(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEnd(source, 5)
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Contains_EmptySource tests the case where the source slice is empty.
func Test_Contains_EmptySource(t *testing.T) {
	var source []int
	result := Contains(source, 1)
	assert.False(t, result)
}

// Test_Contains_NotFound tests the case where the item is not found.
func Test_Contains_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := Contains(source, 4)
	assert.False(t, result)
}

// Test_Contains_Found tests the case where the item is found.
func Test_Contains_Found(t *testing.T) {
	source := []int{1, 2, 3}
	result := Contains(source, 2)
	assert.True(t, result)
}

// Test_ContainsFunc_EmptySource tests the case where the source slice is empty.
func Test_ContainsFunc_EmptySource(t *testing.T) {
	var source []int
	result := ContainsFunc(source, func(int) bool { return true })
	assert.False(t, result)
}

// Test_ContainsFunc_NotFound tests the case where the item is not found.
func Test_ContainsFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := ContainsFunc(source, func(i int) bool { return i == 4 })
	assert.False(t, result)
}

// Test_ContainsFunc_Found tests the case where the item is found.
func Test_ContainsFunc_Found(t *testing.T) {
	source := []int{1, 2, 3}
	result := ContainsFunc(source, func(i int) bool { return i == 2 })
	assert.True(t, result)
}

// Test_ContainsFuncIndexed_EmptySource tests the case where the source slice is empty.
func Test_ContainsFuncIndexed_EmptySource(t *testing.T) {
	var source []int
	result := ContainsFuncIndexed(source, func(int, int) bool { return true })
	assert.False(t, result)
}

// Test_ContainsFuncIndexed_OnItem tests the case where the item is searched by item.
func Test_ContainsFuncIndexed_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := ContainsFuncIndexed(source, func(i int, index int) bool { return i == 20 })
	assert.True(t, result)
	result = ContainsFuncIndexed(source, func(i int, index int) bool { return i == 40 })
	assert.False(t, result)
}

// Test_ContainsFuncIndexed_OnIndex tests the case where the item is searched by index.
func Test_ContainsFuncIndexed_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := ContainsFuncIndexed(source, func(i int, index int) bool { return index == 1 })
	assert.True(t, result)
	result = ContainsFuncIndexed(source, func(i int, index int) bool { return index == 5 })
	assert.False(t, result)
}

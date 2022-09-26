package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Index_EmptySource tests the case where the source slice is empty.
func Test_Index_EmptySource(t *testing.T) {
	var source []int
	result := Index(source, 1)
	assert.Equal(t, -1, result)
}

// Test_Index_NotFound tests the case where the item is not found.
func Test_Index_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := Index(source, 4)
	assert.Equal(t, -1, result)
}

// Test_Index_Found tests the case where the item is found.
func Test_Index_Found(t *testing.T) {
	source := []int{1, 2, 3, 2, 5}
	result := Index(source, 2)
	assert.Equal(t, 1, result)
}

// Test_IndexFunc_EmptySource tests the case where the source slice is empty.
func Test_IndexFunc_EmptySource(t *testing.T) {
	var source []int
	result := IndexFunc(source, func(int) bool { return true })
	assert.Equal(t, -1, result)
}

// Test_IndexFunc_NotFound tests the case where the item is not found.
func Test_IndexFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := IndexFunc(source, func(int) bool { return false })
	assert.Equal(t, -1, result)
}

// Test_IndexFunc_Found tests the case where the item is found.
func Test_IndexFunc_Found(t *testing.T) {
	source := []int{1, 2, 3, 2, 5}
	result := IndexFunc(source, func(i int) bool { return i == 2 })
	assert.Equal(t, 1, result)
}

// Test_IndexIndexedFunc_EmptySource tests the case where the source slice is empty.
func Test_IndexIndexedFunc_EmptySource(t *testing.T) {
	var source []int
	result := IndexIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, -1, result)
}

// Test_IndexIndexedFunc_OnItem tests the case where the item is searched by item.
func Test_IndexIndexedFunc_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := IndexIndexedFunc(source, func(i int, _ int) bool { return i == 20 })
	assert.Equal(t, 1, result)
	result = IndexIndexedFunc(source, func(i int, _ int) bool { return i == 40 })
	assert.Equal(t, -1, result)
}

// Test_IndexIndexedFunc_OnIndex tests the case where the item is searched by index.
func Test_IndexIndexedFunc_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := IndexIndexedFunc(source, func(_ int, i int) bool { return i == 3 })
	assert.Equal(t, 3, result)
	result = IndexIndexedFunc(source, func(_ int, i int) bool { return i == 5 })
	assert.Equal(t, -1, result)
}

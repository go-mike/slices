package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_LastIndex_EmptySource tests the case where the source slice is empty.
func Test_LastIndex_EmptySource(t *testing.T) {
	var source []int
	result := LastIndex(source, 1)
	assert.Equal(t, -1, result)
}

// Test_LastIndex_NotFound tests the case where the item is not found.
func Test_LastIndex_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := LastIndex(source, 4)
	assert.Equal(t, -1, result)
}

// Test_LastIndex_Found tests the case where the item is found.
func Test_LastIndex_Found(t *testing.T) {
	source := []int{1, 2, 3, 2, 5}
	result := LastIndex(source, 2)
	assert.Equal(t, 3, result)
}

// Test_LastIndexFunc_EmptySource tests the case where the source slice is empty.
func Test_LastIndexFunc_EmptySource(t *testing.T) {
	var source []int
	result := LastIndexFunc(source, func(int) bool { return true })
	assert.Equal(t, -1, result)
}

// Test_LastIndexFunc_NotFound tests the case where the item is not found.
func Test_LastIndexFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := LastIndexFunc(source, func(int) bool { return false })
	assert.Equal(t, -1, result)
}

// Test_LastIndexFunc_Found tests the case where the item is found.
func Test_LastIndexFunc_Found(t *testing.T) {
	source := []int{1, 2, 3, 2, 5}
	result := LastIndexFunc(source, func(i int) bool { return i == 2 })
	assert.Equal(t, 3, result)
}

// Test_LastIndexIndexedFunc_EmptySource tests the case where the source slice is empty.
func Test_LastIndexIndexedFunc_EmptySource(t *testing.T) {
	var source []int
	result := LastIndexIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, -1, result)
}

// Test_LastIndexIndexedFunc_OnItem tests the case where the item is searched by item.
func Test_LastIndexIndexedFunc_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := LastIndexIndexedFunc(source, func(i int, _ int) bool { return i == 20 })
	assert.Equal(t, 3, result)
}

// Test_LastIndexIndexedFunc_OnIndex tests the case where the item is searched by index.
func Test_LastIndexIndexedFunc_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := LastIndexIndexedFunc(source, func(_ int, i int) bool { return i == 3 })
	assert.Equal(t, 3, result)
}

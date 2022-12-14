package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_IndicesFunc_EmptySource tests the case where the source slice is empty.
func Test_IndicesFunc_EmptySource(t *testing.T) {
	source := []int{}
	result := IndicesFunc(source, func(i int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_IndicesFunc_NoMatch tests the case where the source slice contains no matching elements.
func Test_IndicesFunc_NoMatch(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := IndicesFunc(source, func(i int) bool { return i == 6 })
	assert.Equal(t, []int{}, result)
}

// Test_IndicesFunc_Match tests the case where the source slice contains matching elements.
func Test_IndicesFunc_Match(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := IndicesFunc(source, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{1, 3}, result)
}

// Test_Indices_EmptySource tests the case where the source slice is empty.
func Test_Indices_EmptySource(t *testing.T) {
	source := []int{}
	result := Indices(source, 1)
	assert.Equal(t, []int{}, result)
}

// Test_Indices_NoMatch tests the case where the source slice contains no matching elements.
func Test_Indices_NoMatch(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Indices(source, 6)
	assert.Equal(t, []int{}, result)
}

// Test_Indices_Match tests the case where the source slice contains matching elements.
func Test_Indices_Match(t *testing.T) {
	source := []int{1, 2, 3, 2, 5}
	result := Indices(source, 2)
	assert.Equal(t, []int{1, 3}, result)
}

// Test_IndicesIndexedFunc_EmptySource tests the case where the source slice is empty.
func Test_IndicesIndexedFunc_EmptySource(t *testing.T) {
	source := []int{}
	result := IndicesIndexedFunc(source, func(i int, index int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_IndicesIndexedFunc_OnItem tests the case where the item is searched by item.
func Test_IndicesIndexedFunc_OnItem(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := IndicesIndexedFunc(source, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{1, 3}, result)
}

// Test_IndicesIndexedFunc_OnIndex tests the case where the item is searched by index.
func Test_IndicesIndexedFunc_OnIndex(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := IndicesIndexedFunc(source, func(_ int, index int) bool { return index%2 == 0 })
	assert.Equal(t, []int{0, 2, 4}, result)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Compact_EmptySource tests the case where the source slice is empty.
func Test_Compact_EmptySource(t *testing.T) {
	var source []int
	result := Compact(source)
	assert.Empty(t, result)
}

// Test_Compact_SingleItem tests the case where the source slice contains a single item.
func Test_Compact_SingleItem(t *testing.T) {
	source := []int{1}
	result := Compact(source)
	assert.Equal(t, source, result)
}

// Test_Compact_UniqueItems tests the case where the source slice contains multiple unique items.
func Test_Compact_UniqueItems(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Compact(source)
	assert.Equal(t, source, result)
}

// Test_Compact_DuplicateItems tests the case where the source slice contains multiple duplicate items.
func Test_Compact_DuplicateItems(t *testing.T) {
	source := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	result := Compact(source)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_CompactFunc_EmptySource tests the case where the source slice is empty.
func Test_CompactFunc_EmptySource(t *testing.T) {
	var source []int
	result := CompactFunc(source, func(int, int) bool { return true })
	assert.Empty(t, result)
}

// Test_CompactFunc_SingleItem tests the case where the source slice contains a single item.
func Test_CompactFunc_SingleItem(t *testing.T) {
	source := []int{1}
	result := CompactFunc(source, func(int, int) bool { return true })
	assert.Equal(t, source, result)
}

// Test_CompactFunc_UniqueItems tests the case where the source slice contains multiple unique items.
func Test_CompactFunc_UniqueItems(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := CompactFunc(source, func(a, b int) bool { return a == b })
	assert.Equal(t, source, result)
}

// Test_CompactFunc_DuplicateItems tests the case where the source slice contains multiple duplicate items.
func Test_CompactFunc_DuplicateItems(t *testing.T) {
	source := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	result := CompactFunc(source, func(a, b int) bool { return a == b })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

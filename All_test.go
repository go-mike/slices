package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_All_EmptySource tests the case where the source slice is empty.
func Test_All_EmptySource(t *testing.T) {
	var source []int
	result := All(source, func(int) bool { return true })
	assert.True(t, result)
}

// Test_All_AllFalse tests the case where the predicate returns false for all items.
func Test_All_AllFalse(t *testing.T) {
	source := []int{1, 2, 3}
	result := All(source, func(int) bool { return false })
	assert.False(t, result)
}

// Test_All_AllTrue tests the case where the predicate returns true for all items.
func Test_All_AllTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := All(source, func(int) bool { return true })
	assert.True(t, result)
}

// Test_All_SomeTrue tests the case where the predicate returns true for some, but not all, items.
func Test_All_SomeTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := All(source, func(i int) bool { return i < 3 })
	assert.False(t, result)
}

// Test_AllIndexed_EmptySource tests the case where the source slice is empty.
func Test_AllIndexed_EmptySource(t *testing.T) {
	var source []int
	result := AllIndexed(source, func(int, int) bool { return true })
	assert.True(t, result)
}

// Test_AllIndexed_AllFalse tests the case where the predicate returns false for all items.
func Test_AllIndexed_AllFalse(t *testing.T) {
	source := []int{1, 2, 3}
	result := AllIndexed(source, func(int, int) bool { return false })
	assert.False(t, result)
}

// Test_AllIndexed_AllTrue tests the case where the predicate returns true for all items.
func Test_AllIndexed_AllTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := AllIndexed(source, func(int, int) bool { return true })
	assert.True(t, result)
}

// Test_AllIndexed_OnItems tests the case where the predicate returns true for all items, based on the values.
func Test_AllIndexed_OnItems(t *testing.T) {
	source := []int{10, 20, 30}
	result := AllIndexed(source, func(item int, index int) bool { return item > 5 })
	assert.True(t, result)
	result = AllIndexed(source, func(item int, index int) bool { return item > 15 })
	assert.False(t, result)
}

// Test_AllIndexed_OnIndexes tests the case where the predicate returns true for all items, based on the indexes.
func Test_AllIndexed_OnIndexes(t *testing.T) {
	source := []int{10, 20, 30}
	result := AllIndexed(source, func(item int, index int) bool { return index < 3 })
	assert.True(t, result)
	result = AllIndexed(source, func(item int, index int) bool { return index < 2 })
	assert.False(t, result)
}

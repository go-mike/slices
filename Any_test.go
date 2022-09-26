package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Any_EmptySource tests the case where the source slice is empty.
func Test_Any_EmptySource(t *testing.T) {
	var source []int
	result := Any(source, func(int) bool { return true })
	assert.False(t, result)
}

// Test_Any_AllFalse tests the case where the predicate returns false for all items.
func Test_Any_AllFalse(t *testing.T) {
	source := []int{1, 2, 3}
	result := Any(source, func(int) bool { return false })
	assert.False(t, result)
}

// Test_Any_AllTrue tests the case where the predicate returns true for all items.
func Test_Any_AllTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := Any(source, func(int) bool { return true })
	assert.True(t, result)
}

// Test_Any_SomeTrue tests the case where the predicate returns true for some, but not all, items.
func Test_Any_SomeTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := Any(source, func(i int) bool { return i < 3 })
	assert.True(t, result)
}

// Test_AnyIndexed_EmptySource tests the case where the source slice is empty.
func Test_AnyIndexed_EmptySource(t *testing.T) {
	var source []int
	result := AnyIndexed(source, func(int, int) bool { return true })
	assert.False(t, result)
}

// Test_AnyIndexed_AllFalse tests the case where the predicate returns false for all items.
func Test_AnyIndexed_AllFalse(t *testing.T) {
	source := []int{1, 2, 3}
	result := AnyIndexed(source, func(int, int) bool { return false })
	assert.False(t, result)
}

// Test_AnyIndexed_AllTrue tests the case where the predicate returns true for all items.
func Test_AnyIndexed_AllTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := AnyIndexed(source, func(int, int) bool { return true })
	assert.True(t, result)
}

// Test_AnyIndexed_OnItems tests the case where the predicate returns true for all items, based on the values.
func Test_AnyIndexed_OnItems(t *testing.T) {
	source := []int{10, 20, 30}
	result := AnyIndexed(source, func(i int, index int) bool { return i == 20 })
	assert.True(t, result)
	result = AnyIndexed(source, func(i int, index int) bool { return i == 40 })
	assert.False(t, result)
}

// Test_AnyIndexed_SomeTrue tests the case where the predicate returns true for some, but not all, items.
func Test_AnyIndexed_SomeTrue(t *testing.T) {
	source := []int{1, 2, 3}
	result := AnyIndexed(source, func(i int, index int) bool { return index == 2 })
	assert.True(t, result)
	result = AnyIndexed(source, func(i int, index int) bool { return index == 4 })
	assert.False(t, result)
}

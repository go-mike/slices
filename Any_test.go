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

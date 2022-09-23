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

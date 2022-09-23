package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Filter_EmptySource tests the case where the source slice is empty.
func Test_Filter_EmptySource(t *testing.T) {
	var source []int
	result := Filter(source, func(int) bool { return true })
	assert.Empty(t, result)
}

// Test_Filter_NoneMatch tests the case where no items match the predicate.
func Test_Filter_NoneMatch(t *testing.T) {
	source := []int{1, 2, 3}
	result := Filter(source, func(int) bool { return false })
	assert.Empty(t, result)
}

// Test_Filter_AllMatch tests the case where all items match the predicate.
func Test_Filter_AllMatch(t *testing.T) {
	source := []int{1, 2, 3}
	result := Filter(source, func(int) bool { return true })
	assert.Equal(t, source, result)
}

// Test_Filter_SomeMatch tests the case where some items match the predicate.
func Test_Filter_SomeMatch(t *testing.T) {
	source := []int{1, 2, 3}
	result := Filter(source, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2}, result)
}

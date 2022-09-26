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

// Test_FilterIndexed_EmptySource tests the case where the source slice is empty.
func Test_FilterIndexed_EmptySource(t *testing.T) {
	var source []int
	result := FilterIndexed(source, func(int, int) bool { return true })
	assert.Empty(t, result)
}

// Test_FilterIndexed_OnItem tests the case where the item is searched by item.
func Test_FilterIndexed_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := FilterIndexed(source, func(i int, _ int) bool { return i%20 != 0 })
	assert.Equal(t, []int{10, 30, 50}, result)
}

// Test_FilterIndexed_OnIndex tests the case where the item is searched by index.
func Test_FilterIndexed_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 20, 50}
	result := FilterIndexed(source, func(_ int, index int) bool { return index%2 == 0 })
	assert.Equal(t, []int{10, 30, 50}, result)
}

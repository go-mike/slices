package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Replace_EmptySource tests the case where the source slice is empty.
func Test_Replace_EmptySource(t *testing.T) {
	source := []int{}
	result := Replace(source, []int{6, 7, 8}, 0, 0)
	assert.Equal(t, []int{6, 7, 8}, result)
}

// Test_Replace_EmptyReplacement tests the case where the replacement slice is empty.
func Test_Replace_EmptyReplacement(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Replace(source, []int{}, 0, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Replace_EmptySourceAndReplacement tests the case where both the source and replacement slices are empty.
func Test_Replace_EmptySourceAndReplacement(t *testing.T) {
	source := []int{}
	result := Replace(source, []int{}, 0, 0)
	assert.Equal(t, []int{}, result)
}

// Test_Replace_NegativeStart tests the case where the start is negative.
func Test_Replace_NegativeStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Replace(source, []int{6, 7, 8}, -1, 3)
	assert.Equal(t, []int{6, 7, 8, 3, 4, 5}, result)
}

// Test_Replace_NegativeStartAndCount tests the case where the start is negative and count is smaller than the start.
func Test_Replace_NegativeStartAndCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Replace(source, []int{6, 7, 8}, -3, 2)
	assert.Equal(t, []int{6, 7, 8, 1, 2, 3, 4, 5}, result)
}

// Test_Replace_PositiveStart tests the case where the start is positive.
func Test_Replace_PositiveStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Replace(source, []int{6, 7, 8}, 1, 2)
	assert.Equal(t, []int{1, 6, 7, 8, 4, 5}, result)
}

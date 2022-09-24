package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Delete_EmptySource tests the case where the source slice is empty.
func Test_Delete_EmptySource(t *testing.T) {
	source := []int{}
	result := Delete(source, 0, 1)
	assert.Equal(t, []int{}, result)
}

// Test_Delete_ZeroCount tests the case where the count is zero.
func Test_Delete_ZeroCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Delete(source, 0, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Delete_NegativeCount tests the case where the count is negative.
func Test_Delete_NegativeCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Delete(source, 0, -1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Delete_NegativeStart tests the case where the start is negative.
func Test_Delete_NegativeStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Delete(source, -1, 3)
	assert.Equal(t, []int{3, 4, 5}, result)
}

// Test_Delete_NegativeStartAndCount tests the case where the start is negative and count is smaller than the start.
func Test_Delete_NegativeStartAndCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Delete(source, -3, 2)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Delete_PositiveStart tests the case where the start is positive.
func Test_Delete_PositiveStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Delete(source, 1, 3)
	assert.Equal(t, []int{1, 5}, result)
}

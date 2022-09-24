package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Insert_EmptySource tests the case where the source slice is empty.
func Test_Insert_EmptySource(t *testing.T) {
	source := []int{}
	result := Insert([]int{1, 2, 3, 4, 5}, source, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Insert_EmptyDestination tests the case where the destination slice is empty.
func Test_Insert_EmptyDestination(t *testing.T) {
	destination := []int{}
	result := Insert(destination, []int{1, 2, 3, 4, 5}, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Insert_NegativePosition tests the case where the position is negative.
func Test_Insert_NegativePosition(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, []int{6, 7, 8}, -1)
	assert.Equal(t, []int{6, 7, 8, 1, 2, 3, 4, 5}, result)
}

// Test_Insert_PositionGreaterThanLength tests the case where the position is greater than the length of the destination slice.
func Test_Insert_PositionGreaterThanLength(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, []int{6, 7, 8}, 10)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, result)
}

// Test_Insert tests the case where the position is within the bounds of the destination slice.
func Test_Insert(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, []int{6, 7, 8}, 2)
	assert.Equal(t, []int{1, 2, 6, 7, 8, 3, 4, 5}, result)
}

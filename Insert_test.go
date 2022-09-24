package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Insert_EmptySource tests the case where the source slice is empty.
func Test_Insert_EmptySource(t *testing.T) {
	source := []int{}
	result := Insert(source, 0, []int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Insert_EmptyDestination tests the case where the destination slice is empty.
func Test_Insert_EmptyDestination(t *testing.T) {
	destination := []int{}
	result := Insert(destination, 0, []int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_Insert_NegativeStart tests the case where the position is negative.
func Test_Insert_NegativeStart(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, -1, []int{6, 7, 8})
	assert.Equal(t, []int{6, 7, 8, 1, 2, 3, 4, 5}, result)
}

// Test_Insert_StartGreaterThanLength tests the case where the position is greater than the length of the destination slice.
func Test_Insert_StartGreaterThanLength(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, 10, []int{6, 7, 8})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, result)
}

// Test_Insert tests the case where the position is within the bounds of the destination slice.
func Test_Insert(t *testing.T) {
	result := Insert([]int{1, 2, 3, 4, 5}, 2, []int{6, 7, 8})
	assert.Equal(t, []int{1, 2, 6, 7, 8, 3, 4, 5}, result)
}

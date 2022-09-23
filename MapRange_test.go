package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_MapRange_EmptySource tests the case where the source slice is empty.
func Test_MapRange_EmptySource(t *testing.T) {
	result := MapRange(10, 0, func(int) int { return 1 })
	assert.Empty(t, result)
}

// Test_MapRange tests the case where the mapper returns items.
func Test_MapRange(t *testing.T) {
	result := MapRange(10, 5, func(i int) int { return 3*i + 1 })
	assert.Equal(t, []int{31, 34, 37, 40, 43}, result)
}

// Test_MapRangeStep_EmptySource tests the case where the source slice is empty.
func Test_MapRangeStep_EmptySource(t *testing.T) {
	result := MapRangeStep(10, 0, 2, func(int) int { return 1 })
	assert.Empty(t, result)
}

// Test_MapRangeStep tests the case where the mapper returns items.
func Test_MapRangeStep(t *testing.T) {
	result := MapRangeStep(10, 5, 2, func(i int) int { return 3*i + 1 })
	assert.Equal(t, []int{31, 37, 43, 49, 55}, result)
}

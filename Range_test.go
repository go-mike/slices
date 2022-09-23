package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Range_Empty tests the case where the count is zero.
func Test_Range_Empty(t *testing.T) {
	result := Range(10, 0)
	assert.Empty(t, result)
}

// Test_Range tests the case where the count is positive.
func Test_Range(t *testing.T) {
	result := Range(10, 5)
	assert.Equal(t, []int{10, 11, 12, 13, 14}, result)
}

// Test_RangeStep_Empty tests the case where the count is zero.
func Test_RangeStep_Empty(t *testing.T) {
	result := RangeStep(10, 0, 2)
	assert.Empty(t, result)
}

// Test_RangeStep tests the case where the count is positive.
func Test_RangeStep(t *testing.T) {
	result := RangeStep(10, 5, 2)
	assert.Equal(t, []int{10, 12, 14, 16, 18}, result)
}

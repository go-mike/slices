package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_ExtractRangeStep_EmptySource tests the case where the source slice is empty.
func Test_ExtractRangeStep_EmptySource(t *testing.T) {
	source := []int{}
	result := ExtractRangeStep(source, 0, 1, 1)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_ZeroCount tests the case where the count is zero.
func Test_ExtractRangeStep_ZeroCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 0, 0, 1)
	assert.Empty(t, result)
}

// Test_ExtractRangeStep_NegativeCount tests the case where the count is negative.
func Test_ExtractRangeStep_NegativeCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 0, -1, 1)
	assert.Empty(t, result)
}

// Test_ExtractRangeStep_ZeroStep tests the case where the step is zero.
func Test_ExtractRangeStep_ZeroStep(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, 3, 0)
	assert.Equal(t, []int{2, 2, 2}, result)
}

// Test_ExtractRangeStep_NegativeStep tests the case where the step is negative.
func Test_ExtractRangeStep_NegativeStep(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 4, 3, -1)
	assert.Equal(t, []int{5, 4, 3}, result)
}

// Test_ExtractRangeStep_PositiveStep tests the case where the step is positive.
func Test_ExtractRangeStep_PositiveStep(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 0, 3, 1)
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_ExtractRangeStep_PositiveLargerStep tests the case where the step is positive and larger than 1.
func Test_ExtractRangeStep_PositiveLargerStep(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 0, 3, 2)
	assert.Equal(t, []int{1, 3, 5}, result)
}

// Test_ExtractRangeStep_NegativeLargerStep tests the case where the step is negative and larger than -1.
func Test_ExtractRangeStep_NegativeLargerStep(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 4, 3, -2)
	assert.Equal(t, []int{5, 3, 1}, result)
}

// Test_ExtractRangeStep_StartTooLarge tests the case where the start is too large.
func Test_ExtractRangeStep_StartTooLarge(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 5, 3, 1)
	assert.Empty(t, result)
}

// Test_ExtractRangeStep_StartTooSmall tests the case where the start is too small.
func Test_ExtractRangeStep_StartTooSmall(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -1, 3, 1)
	assert.Equal(t, []int{1, 2}, result)
}

// Test_ExtractRangeStep_CountTooLarge tests the case where the count is too large.
func Test_ExtractRangeStep_CountTooLarge(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 3, 3, 1)
	assert.Equal(t, []int{4, 5}, result)
}

// Test_ExtractRange_EmptySource tests the case where the source slice is empty.
func Test_ExtractRange_EmptySource(t *testing.T) {
	source := []int{}
	result := ExtractRange(source, 0, 1)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRange tests the case where the values are normal.
func Test_ExtractRange(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRange(source, 1, 3)
	assert.Equal(t, []int{2, 3, 4}, result)
}

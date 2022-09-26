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
	result := ExtractRangeStep(source, 1, 0, 3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_NegativeCount tests the case where the count is negative.
func Test_ExtractRangeStep_NegativeCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, -3, 3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_PositiveStep_BeyondEnd tests the case where the step is positive and the start index is beyond the end of the source slice.
func Test_ExtractRangeStep_PositiveStep_BeyondEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 5, 3, 3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_PositiveStep_BeforeZero tests the case where the step is positive and the start index plus the count is before zero.
func Test_ExtractRangeStep_PositiveStep_BeforeZero(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -7, 3, 3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_NegativeStep_BeyondStart tests the case where the step is negative and the start index is beyond the start of the source slice.
func Test_ExtractRangeStep_NegativeStep_BeyondStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -1, 3, -3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_NegativeStep_AfterEnd tests the case where the step is negative and the start index plus the count is after the end of the source slice.
func Test_ExtractRangeStep_NegativeStep_AfterEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 11, 3, -3)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_StepOne_InsideSlice tests the case where the step is one and the start index and count are inside the source slice.
func Test_ExtractRangeStep_StepOne_InsideSlice(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, 3, 1)
	assert.Equal(t, []int{2, 3, 4}, result)
}

// Test_ExtractRangeStep_StepOne_StartZero tests the case where the step is one and the start index is zero.
func Test_ExtractRangeStep_StepOne_StartZero(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 0, 3, 1)
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_ExtractRangeStep_StepOne_ClipStart tests the case where the step is one and the start index is negative.
func Test_ExtractRangeStep_StepOne_ClipStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -3, 5, 1)
	assert.Equal(t, []int{1, 2}, result)
}

// Test_ExtractRangeStep_StepOne_ClipEnd tests the case where the step is one and the start index plus the count is beyond the end of the source slice.
func Test_ExtractRangeStep_StepOne_ClipEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 3, 5, 1)
	assert.Equal(t, []int{4, 5}, result)
}

// Test_ExtractRangeStep_StepOne_ClipBoth tests the case where the step is one and the start index is negative and the start index plus the count is beyond the end of the source slice.
func Test_ExtractRangeStep_StepOne_ClipBoth(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -3, 10, 1)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_ExtractRangeStep_PositiveStep_InsideSlice tests the case where the step is positive and the start index and count are inside the source slice.
func Test_ExtractRangeStep_PositiveStep_InsideSlice(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, 2, 2)
	assert.Equal(t, []int{2, 4}, result)
}

// Test_ExtractRangeStep_PositiveStep_ClipStart tests the case where the step is positive and the start index is negative.
func Test_ExtractRangeStep_PositiveStep_ClipStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -1, 3, 2)
	assert.Equal(t, []int{2, 4}, result)
}

// Test_ExtractRangeStep_PositiveStep_ClipEnd tests the case where the step is positive and the start index plus the count is beyond the end of the source slice.
func Test_ExtractRangeStep_PositiveStep_ClipEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, 3, 2)
	assert.Equal(t, []int{2, 4}, result)
}

// Test_ExtractRangeStep_PositiveStep_ClipBoth tests the case where the step is positive and the start index is negative and the start index plus the count is beyond the end of the source slice.
func Test_ExtractRangeStep_PositiveStep_ClipBoth(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -1, 10, 2)
	assert.Equal(t, []int{2, 4}, result)
}

// Test_ExtractRangeStep_NegativeStep_InsideSlice tests the case where the step is negative and the start index and count are inside the source slice.
func Test_ExtractRangeStep_NegativeStep_InsideSlice(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 3, 2, -2)
	assert.Equal(t, []int{4, 2}, result)
}

// Test_ExtractRangeStep_NegativeStep_ClipStart tests the case where the step is negative and the start index is beyond the start of the source slice.
func Test_ExtractRangeStep_NegativeStep_ClipStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 5, 3, -2)
	assert.Equal(t, []int{4, 2}, result)
}

// Test_ExtractRangeStep_NegativeStep_ClipEnd tests the case where the step is negative and the start index plus the count is after the end of the source slice.
func Test_ExtractRangeStep_NegativeStep_ClipEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 3, 3, -2)
	assert.Equal(t, []int{4, 2}, result)
}

// Test_ExtractRangeStep_NegativeStep_ClipBoth tests the case where the step is negative and the start index is beyond the start of the source slice and the start index plus the count is after the end of the source slice.
func Test_ExtractRangeStep_NegativeStep_ClipBoth(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 5, 10, -2)
	assert.Equal(t, []int{4, 2}, result)
}

// Test_ExtractRangeStep_ZeroStep_BeforeStart tests the case where the step is zero and the start index is before the start of the source slice.
func Test_ExtractRangeStep_ZeroStep_BeforeStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, -1, 3, 0)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_ZeroStep_AfterEnd tests the case where the step is zero and the start index is after the end of the source slice.
func Test_ExtractRangeStep_ZeroStep_AfterEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 5, 3, 0)
	assert.Equal(t, []int{}, result)
}

// Test_ExtractRangeStep_ZeroStep_InsideSlice tests the case where the step is zero and the start index and count are inside the source slice.
func Test_ExtractRangeStep_ZeroStep_InsideSlice(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRangeStep(source, 1, 3, 0)
	assert.Equal(t, []int{2, 2, 2}, result)
}

// Test_ExtractRange tests the case where the start index and count are inside the source slice.
func Test_ExtractRange(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := ExtractRange(source, 1, 3)
	assert.Equal(t, []int{2, 3, 4}, result)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Count_EmptySource tests the case where the source slice is empty.
func Test_Count_EmptySource(t *testing.T) {
	var source []int
	result := Count(source, 1)
	assert.Equal(t, 0, result)
}

// Test_Count_NotFound tests the case where the item is not found.
func Test_Count_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := Count(source, 4)
	assert.Equal(t, 0, result)
}

// Test_Count_FoundOnce tests the case where the item is found once.
func Test_Count_FoundOnce(t *testing.T) {
	source := []int{1, 2, 3}
	result := Count(source, 2)
	assert.Equal(t, 1, result)
}

// Test_Count_FoundMultiple tests the case where the item is found multiple times.
func Test_Count_FoundMultiple(t *testing.T) {
	source := []int{1, 2, 3, 2}
	result := Count(source, 2)
	assert.Equal(t, 2, result)
}

// Test_CountFunc_EmptySource tests the case where the source slice is empty.
func Test_CountFunc_EmptySource(t *testing.T) {
	var source []int
	result := CountFunc(source, func(int) bool { return true })
	assert.Equal(t, 0, result)
}

// Test_CountFunc_NotFound tests the case where the item is not found.
func Test_CountFunc_NotFound(t *testing.T) {
	source := []int{1, 2, 3}
	result := CountFunc(source, func(i int) bool { return i == 4 })
	assert.Equal(t, 0, result)
}

// Test_CountFunc_FoundOnce tests the case where the item is found once.
func Test_CountFunc_FoundOnce(t *testing.T) {
	source := []int{1, 2, 3}
	result := CountFunc(source, func(i int) bool { return i == 2 })
	assert.Equal(t, 1, result)
}

// Test_CountFunc_FoundMultiple tests the case where the item is found multiple times.
func Test_CountFunc_FoundMultiple(t *testing.T) {
	source := []int{1, 2, 3, 2}
	result := CountFunc(source, func(i int) bool { return i == 2 })
	assert.Equal(t, 2, result)
}

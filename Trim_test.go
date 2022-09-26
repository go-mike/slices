package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_TrimStartIndexedFunc_EmptySource tests the case where the source slice is empty.
func Test_TrimStartIndexedFunc_EmptySource(t *testing.T) {
	source := []int{}
	result := TrimStartIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_TrimStartIndexedFunc_False tests the case where the predicate returns false.
func Test_TrimStartIndexedFunc_False(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartIndexedFunc(source, func(int, int) bool { return false })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_TrimStartIndexedFunc_True tests the case where the predicate returns true.
func Test_TrimStartIndexedFunc_True(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_TrimStartIndexedFunc_TrueAtStart_OnItem tests the case where the predicate returns true at the start of the source slice.
func Test_TrimStartIndexedFunc_TrueAtStart_OnItem(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartIndexedFunc(source, func(item int, index int) bool { return item < 3 })
	assert.Equal(t, []int{3, 4, 5}, result)
}

// Test_TrimStartIndexedFunc_TrueAtStart_OnIndex tests the case where the predicate returns true at the start of the source slice.
func Test_TrimStartIndexedFunc_TrueAtStart_OnIndex(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartIndexedFunc(source, func(item int, index int) bool { return index < 2 })
	assert.Equal(t, []int{3, 4, 5}, result)
}

// Test_TrimEndIndexedFunc_EmptySource tests the case where the source slice is empty.
func Test_TrimEndIndexedFunc_EmptySource(t *testing.T) {
	source := []int{}
	result := TrimEndIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_TrimEndIndexedFunc_False tests the case where the predicate returns false.
func Test_TrimEndIndexedFunc_False(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndIndexedFunc(source, func(int, int) bool { return false })
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

// Test_TrimEndIndexedFunc_True tests the case where the predicate returns true.
func Test_TrimEndIndexedFunc_True(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndIndexedFunc(source, func(int, int) bool { return true })
	assert.Equal(t, []int{}, result)
}

// Test_TrimEndIndexedFunc_TrueAtEnd_OnItem tests the case where the predicate returns true at the end of the source slice.
func Test_TrimEndIndexedFunc_TrueAtEnd_OnItem(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndIndexedFunc(source, func(item int, index int) bool { return item > 3 })
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_TrimEndIndexedFunc_TrueAtEnd_OnIndex tests the case where the predicate returns true at the end of the source slice.
func Test_TrimEndIndexedFunc_TrueAtEnd_OnIndex(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndIndexedFunc(source, func(item int, index int) bool { return index > 2 })
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_TrimIndexedFunc_OnItem tests the case where the predicate returns true at the start and end of the source slice.
func Test_TrimIndexedFunc(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimIndexedFunc(source, func(item int, index int) bool { return item < 2 || item > 4 })
	assert.Equal(t, []int{2, 3, 4}, result)
}

// Test_TrimIndexedFunc_OnIndex tests the case where the predicate returns true at the start and end of the source slice.
func Test_TrimIndexedFunc_OnIndex(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimIndexedFunc(source, func(item int, index int) bool { return index < 1 || index > 3 })
	assert.Equal(t, []int{2, 3, 4}, result)
}

// Test_TrimFunc tests the case where the predicate returns true at the start and end of the source slice.
func Test_TrimFunc(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimFunc(source, func(item int) bool { return item < 2 || item > 4 })
	assert.Equal(t, []int{2, 3, 4}, result)
}

// Test_Trim tests the case where the predicate returns true at the start and end of the source slice.
func Test_Trim(t *testing.T) {
	source := []int{1, 2, 3, 4, 3, 2, 1}
	result := Trim(source, 1)
	assert.Equal(t, []int{2, 3, 4, 3, 2}, result)
}

// Test_TrimStartFunc tests the case where the predicate returns true at the start of the source slice.
func Test_TrimStartFunc(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStartFunc(source, func(item int) bool { return item < 3 })
	assert.Equal(t, []int{3, 4, 5}, result)
}

// Test_TrimStart tests the case where the predicate returns true at the start of the source slice.
func Test_TrimStart(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimStart(source, 1)
	assert.Equal(t, []int{2, 3, 4, 5}, result)
}

// Test_TrimEndFunc tests the case where the predicate returns true at the end of the source slice.
func Test_TrimEndFunc(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEndFunc(source, func(item int) bool { return item > 3 })
	assert.Equal(t, []int{1, 2, 3}, result)
}

// Test_TrimEnd tests the case where the predicate returns true at the end of the source slice.
func Test_TrimEnd(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := TrimEnd(source, 5)
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

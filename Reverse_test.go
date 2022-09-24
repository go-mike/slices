package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Reverse_EmptySource tests the case where the source slice is empty.
func Test_Reverse_EmptySource(t *testing.T) {
	var source []int
	result := Reverse(source)
	assert.Equal(t, []int{}, result)
}

// Test_Reverse tests the case where the source slice is not empty.
func Test_Reverse(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Reverse(source)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, result)
}

// Test_ReverseInPlace_NilSource tests the case where the source slice is nil.
func Test_ReverseInPlace_NilSource(t *testing.T) {
	var source []int
	ReverseInPlace(source)
	assert.Nil(t, source)
}

// Test_ReverseInPlace_EmptySource tests the case where the source slice is empty.
func Test_ReverseInPlace_EmptySource(t *testing.T) {
	source := []int{}
	ReverseInPlace(source)
	assert.Equal(t, []int{}, source)
}

// Test_ReverseInPlace tests the case where the source slice is not empty.
func Test_ReverseInPlace(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	ReverseInPlace(source)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, source)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Grow_NegativeAmount tests the Grow function with a negative amount.
func Test_Grow_NegativeAmount(t *testing.T) {
	source := []int{1, 2, 3}
	result := Grow(source, -1)
	assert.Equal(t, source, result)
}

// Test_Grow_PositiveAmount tests the Grow function with a positive amount.
func Test_Grow_PositiveAmount(t *testing.T) {
	source := []int{1, 2, 3}
	result := Grow(source, 2)
	assert.Equal(t, []int{1, 2, 3, 0, 0}, result)
}

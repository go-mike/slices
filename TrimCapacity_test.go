package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_TrimCapacity_EmptySource tests the TrimCapacity function with an empty source.
func Test_TrimCapacity_EmptySource(t *testing.T) {
	source := []int{}
	actual := TrimCapacity(source)
	assert.Equal(t, source, actual)
	assert.Equal(t, 0, cap(actual))
}

// Test_Trim tests the TrimCapacity function with a non-empty source.
func Test_TrimCapacity(t *testing.T) {
	source := make([]int, 1, 10)
	source[0] = 1
	actual := TrimCapacity(source)
	assert.Equal(t, source, actual)
	assert.Equal(t, 1, cap(actual))
}

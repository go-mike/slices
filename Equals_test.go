package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_EqualsFunc_DifferentSizes tests the case where the two slices have different sizes.
func Test_EqualsFunc_DifferentSizes(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	result := EqualsFunc(a, b, func(a, b int) bool { return a == b })
	assert.False(t, result)
}

// Test_EqualsFunc_DifferentElements tests the case where the two slices have different elements.
func Test_EqualsFunc_DifferentElements(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	result := EqualsFunc(a, b, func(a, b int) bool { return a == b })
	assert.False(t, result)
}

// Test_EqualsFunc_Equal tests the case where the two slices are equal.
func Test_EqualsFunc_Equal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := EqualsFunc(a, b, func(a, b int) bool { return a == b })
	assert.True(t, result)
}


// Test_Equals_DifferentSizes tests the case where the two slices have different sizes.
func Test_Equals_DifferentSizes(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	result := Equals(a, b)
	assert.False(t, result)
}

// Test_Equals_DifferentElements tests the case where the two slices have different elements.
func Test_Equals_DifferentElements(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	result := Equals(a, b)
	assert.False(t, result)
}

// Test_Equals_Equal tests the case where the two slices are equal.
func Test_Equals_Equal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := Equals(a, b)
	assert.True(t, result)
}

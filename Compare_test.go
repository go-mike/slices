package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_CompareFunc_FirstLower tests the case where the first slice is lower than the second.
func Test_CompareFunc_FirstLower(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	result := CompareFunc(a, b, func(a, b int) int { return a - b })
	assert.Equal(t, -1, result)
}

// Test_CompareFunc_FirstGreater tests the case where the first slice is greater than the second.
func Test_CompareFunc_FirstGreater(t *testing.T) {
	a := []int{1, 2, 4}
	b := []int{1, 2, 3}
	result := CompareFunc(a, b, func(a, b int) int { return a - b })
	assert.Equal(t, 1, result)
}

// Test_CompareFunc_Equal tests the case where the two slices are equal.
func Test_CompareFunc_Equal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := CompareFunc(a, b, func(a, b int) int { return a - b })
	assert.Equal(t, 0, result)
}

// Test_CompareFunc_FirstShorter tests the case where the first slice is shorter than the second.
func Test_CompareFunc_FirstShorter(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	result := CompareFunc(a, b, func(a, b int) int { return a - b })
	assert.Equal(t, -1, result)
}

// Test_CompareFunc_FirstLonger tests the case where the first slice is longer than the second.
func Test_CompareFunc_FirstLonger(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3}
	result := CompareFunc(a, b, func(a, b int) int { return a - b })
	assert.Equal(t, 1, result)
}

// Test_Compare_FirstLower tests the case where the first slice is lower than the second.
func Test_Compare_FirstLower(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	result := Compare(a, b)
	assert.Equal(t, -1, result)
}

// Test_Compare_FirstGreater tests the case where the first slice is greater than the second.
func Test_Compare_FirstGreater(t *testing.T) {
	a := []int{1, 2, 4}
	b := []int{1, 2, 3}
	result := Compare(a, b)
	assert.Equal(t, 1, result)
}

// Test_Compare_Equal tests the case where the two slices are equal.
func Test_Compare_Equal(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := Compare(a, b)
	assert.Equal(t, 0, result)
}

// Test_Compare_FirstShorter tests the case where the first slice is shorter than the second.
func Test_Compare_FirstShorter(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	result := Compare(a, b)
	assert.Equal(t, -1, result)
}

// Test_Compare_FirstLonger tests the case where the first slice is longer than the second.
func Test_Compare_FirstLonger(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3}
	result := Compare(a, b)
	assert.Equal(t, 1, result)
}

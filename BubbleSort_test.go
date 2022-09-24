package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_BubbleSortInPlaceFunc_Empty tests the case where the slice is empty.
func Test_BubbleSortInPlaceFunc_Empty(t *testing.T) {
	slice := []int{}
	BubbleSortInPlaceFunc(slice, func(a, b int) int { return a - b })
	assert.Equal(t, []int{}, slice)
}

// Test_BubbleSortInPlaceFunc tests the case where the slice is not empty.
func Test_BubbleSortInPlaceFunc(t *testing.T) {
	slice := []int{30, 50, 20, 40, 10}
	BubbleSortInPlaceFunc(slice, func(a, b int) int { return a - b })
	assert.Equal(t, []int{10, 20, 30, 40, 50}, slice)
}

// Test_BubbleSortInPlace_Empty tests the case where the slice is empty.
func Test_BubbleSortInPlace_Empty(t *testing.T) {
	slice := []int{}
	BubbleSortInPlace(slice)
	assert.Equal(t, []int{}, slice)
}

// Test_BubbleSortInPlace tests the case where the slice is not empty.
func Test_BubbleSortInPlace(t *testing.T) {
	slice := []int{30, 50, 20, 40, 10}
	BubbleSortInPlace(slice)
	assert.Equal(t, []int{10, 20, 30, 40, 50}, slice)
}

// Test_BubbleSortFunc_Empty tests the case where the slice is empty.
func Test_BubbleSortFunc_Empty(t *testing.T) {
	slice := []int{}
	clone := BubbleSortFunc(slice, func(a, b int) int { return a - b })
	assert.Equal(t, []int{}, slice)
	assert.Equal(t, []int{}, clone)
}

// Test_BubbleSortFunc tests the case where the slice is not empty.
func Test_BubbleSortFunc(t *testing.T) {
	slice := []int{30, 50, 20, 40, 10}
	clone := BubbleSortFunc(slice, func(a, b int) int { return a - b })
	assert.Equal(t, []int{30, 50, 20, 40, 10}, slice)
	assert.Equal(t, []int{10, 20, 30, 40, 50}, clone)
}

// Test_BubbleSort_Empty tests the case where the slice is empty.
func Test_BubbleSort_Empty(t *testing.T) {
	slice := []int{}
	clone := BubbleSort(slice)
	assert.Equal(t, []int{}, slice)
	assert.Equal(t, []int{}, clone)
}

// Test_BubbleSort tests the case where the slice is not empty.
func Test_BubbleSort(t *testing.T) {
	slice := []int{30, 50, 20, 40, 10}
	clone := BubbleSort(slice)
	assert.Equal(t, []int{30, 50, 20, 40, 10}, slice)
	assert.Equal(t, []int{10, 20, 30, 40, 50}, clone)
}

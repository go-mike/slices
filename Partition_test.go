package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Partition_EmptySource tests the case where the source slice is empty.
func Test_Partition_EmptySource(t *testing.T) {
	var source []int
	trueValues, falseValues := Partition(source, func(int) bool { return true })
	assert.Empty(t, trueValues)
	assert.Empty(t, falseValues)
}

// Test_Partition_AllTrue tests the case where the predicate returns true for all items.
func Test_Partition_AllTrue(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	trueValues, falseValues := Partition(source, func(int) bool { return true })
	assert.Equal(t, source, trueValues)
	assert.Empty(t, falseValues)
}

// Test_Partition_AllFalse tests the case where the predicate returns false for all items.
func Test_Partition_AllFalse(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	trueValues, falseValues := Partition(source, func(int) bool { return false })
	assert.Empty(t, trueValues)
	assert.Equal(t, source, falseValues)
}

// Test_Partition tests the case where the predicate returns true for some items.
func Test_Partition(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	trueValues, falseValues := Partition(source, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, trueValues)
	assert.Equal(t, []int{1, 3, 5}, falseValues)
}

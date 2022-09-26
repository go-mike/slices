package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Map_EmptySource tests the case where the source slice is empty.
func Test_Map_EmptySource(t *testing.T) {
	var source []int
	result := Map(source, func(int) int { return 1 })
	assert.Empty(t, result)
}

// Test_Map tests the case where the mapper returns items.
func Test_Map(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := Map(source, func(i int) int { return 3*i + 1 })
	assert.Equal(t, []int{4, 7, 10, 13, 16}, result)
}

// Test_MapIndexed_EmptySource tests the case where the source slice is empty.
func Test_MapIndexed_EmptySource(t *testing.T) {
	var source []int
	result := MapIndexed(source, func(int, int) int { return 1 })
	assert.Empty(t, result)
}

// Test_MapIndexed tests the case where the mapper returns items.
func Test_MapIndexed(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	result := MapIndexed(source, func(item int, index int) int { return 3*item + index })
	assert.Equal(t, []int{3, 7, 11, 15, 19}, result)
}

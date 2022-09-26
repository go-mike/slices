package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_GroupBy_EmptySource tests the case where the source slice is empty.
func Test_GroupBy_EmptySource(t *testing.T) {
	var source []int
	result := GroupBy(source, func(int) int { return 1 })
	assert.Empty(t, result)
}

// Test_GroupBy_SameKey tests the case where all items have the same key.
func Test_GroupBy_SameKey(t *testing.T) {
	var source []int = []int{1, 2, 3}
	result := GroupBy(source, func(int) int { return 1 })
	assert.Equal(t, map[int][]int{1: {1, 2, 3}}, result)
}

// Test_GroupBy_DifferentKeys tests the case where all items have different keys.
func Test_GroupBy_DifferentKeys(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	result := GroupBy(source, func(i int) int { return i % 2 })
	assert.Equal(t, map[int][]int{0: {2, 4}, 1: {1, 3, 5}}, result)
}

// Test_GroupByIndexed_EmptySource tests the case where the source slice is empty.
func Test_GroupByIndexed_EmptySource(t *testing.T) {
	var source []int
	result := GroupByIndexed(source, func(int, int) int { return 1 })
	assert.Empty(t, result)
}

// Test_GroupByIndexed_OnItem tests the case where the item is grouped by item.
func Test_GroupByIndexed_OnItem(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	result := GroupByIndexed(source, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, map[int][]int{0: {2, 4}, 1: {1, 3, 5}}, result)
}

// Test_GroupByIndexed_OnIndex tests the case where the item is grouped by index.
func Test_GroupByIndexed_OnIndex(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	result := GroupByIndexed(source, func(_ int, index int) int { return index % 2 })
	assert.Equal(t, map[int][]int{0: {1, 3, 5}, 1: {2, 4}}, result)
}

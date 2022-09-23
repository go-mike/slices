package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GroupBy_EmptySource(t *testing.T) {
	var source []int
	result := GroupBy(source, func(int) int { return 1 })
	assert.Empty(t, result)
}

func Test_GroupBy_SameKey(t *testing.T) {
	var source []int = []int{1, 2, 3}
	result := GroupBy(source, func(int) int { return 1 })
	assert.Equal(t, map[int][]int{1: {1, 2, 3}}, result)
}

func Test_GroupBy_DifferentKeys(t *testing.T) {
	var source []int = []int{1, 2, 3, 4, 5}
	result := GroupBy(source, func(i int) int { return i%2 })
	assert.Equal(t, map[int][]int{0: {2, 4}, 1: {1, 3, 5}}, result)
}

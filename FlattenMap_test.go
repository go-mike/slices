package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_FlattenMap_EmptySource tests the case where the source map is empty.
func Test_FlattenMap_EmptySource(t *testing.T) {
	var source []int
	result := FlattenMap(source, func(int) []int { return []int{1} })
	assert.Empty(t, result)
}

// Test_FlattenMap_EmptyItems tests the case where the mapper returns empty items.
func Test_FlattenMap_EmptyItems(t *testing.T) {
	source := []int{1, 2}
	result := FlattenMap(source, func(int) []int { return []int{} })
	assert.Empty(t, result)
}

// Test_FlattenMap tests the case where the mapper returns items.
func Test_FlattenMap(t *testing.T) {
	source := []int{1, 2}
	result := FlattenMap(source, func(i int) []int { return []int{i, i} })
	assert.Equal(t, []int{1, 1, 2, 2}, result)
}

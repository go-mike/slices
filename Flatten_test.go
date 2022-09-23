package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Flatten_EmptySource tests the case where the source slice is empty.
func Test_Flatten_EmptySource(t *testing.T) {
	var source [][]int
	result := Flatten(source)
	assert.Empty(t, result)
}

// Test_Flatten_EmptyItems tests the case where the source slice contains empty items.
func Test_Flatten_EmptyItems(t *testing.T) {
	source := [][]int{{}, {}}
	result := Flatten(source)
	assert.Empty(t, result)
}

// Test_Flatten tests the case where the source slice contains items.
func Test_Flatten(t *testing.T) {
	source := [][]int{{1, 2}, {3, 4}}
	result := Flatten(source)
	assert.Equal(t, []int{1, 2, 3, 4}, result)
}

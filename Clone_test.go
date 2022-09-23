package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Clone_EmptySource tests the case where the source slice is empty.
func Test_Clone_EmptySource(t *testing.T) {
	var source []int
	result := Clone(source)
	assert.Empty(t, result)
}

// Test_Clone_SingleItem tests the case where the source slice contains a single item.
func Test_Clone_SingleItem(t *testing.T) {
	source := []int{1}
	result := Clone(source)
	assert.Equal(t, source, result)
}

// Test_Clone_MultipleItems tests the case where the source slice contains multiple items.
func Test_Clone_MultipleItems(t *testing.T) {
	source := []int{1, 2, 3}
	result := Clone(source)
	assert.Equal(t, source, result)
}

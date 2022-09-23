package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Reduce_EmptySource tests the case where the source slice is empty.
func Test_Reduce_EmptySource(t *testing.T) {
	var source []int
	sum := Reduce(source, "X:", func(sum string, item int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:", sum)
}

// Test_Reduce tests the case where the source slice is not empty.
func Test_Reduce(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	sum := Reduce(source, "X:", func(sum string, item int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:12345", sum)
}

// Test_ReduceRight_EmptySource tests the case where the source slice is empty.
func Test_ReduceRight_EmptySource(t *testing.T) {
	var source []int
	sum := ReduceRight(source, "X:", func(sum string, item int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:", sum)
}

// Test_ReduceRight tests the case where the source slice is not empty.
func Test_ReduceRight(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	sum := ReduceRight(source, "X:", func(sum string, item int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:54321", sum)
}

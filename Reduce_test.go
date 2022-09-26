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

// Test_ReduceIndexed_EmptySource tests the case where the source slice is empty.
func Test_ReduceIndexed_EmptySource(t *testing.T) {
	var source []int
	sum := ReduceIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:", sum)
}

// Test_ReduceIndexed_OnItem tests the case where the item is reduced by item.
func Test_ReduceIndexed_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 40, 50}
	sum := ReduceIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:1020304050", sum)
}

// Test_ReduceIndexed_OnIndex tests the case where the item is reduced by index.
func Test_ReduceIndexed_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 40, 50}
	sum := ReduceIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", index) })
	assert.Equal(t, "X:01234", sum)
}

// Test_ReduceRightIndexed_EmptySource tests the case where the source slice is empty.
func Test_ReduceRightIndexed_EmptySource(t *testing.T) {
	var source []int
	sum := ReduceRightIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:", sum)
}

// Test_ReduceRightIndexed_OnItem tests the case where the item is reduced by item.
func Test_ReduceRightIndexed_OnItem(t *testing.T) {
	source := []int{10, 20, 30, 40, 50}
	sum := ReduceRightIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", item) })
	assert.Equal(t, "X:5040302010", sum)
}

// Test_ReduceRightIndexed_OnIndex tests the case where the item is reduced by index.
func Test_ReduceRightIndexed_OnIndex(t *testing.T) {
	source := []int{10, 20, 30, 40, 50}
	sum := ReduceRightIndexed(source, "X:", func(sum string, item int, index int) string { return sum + fmt.Sprintf("%d", index) })
	assert.Equal(t, "X:43210", sum)
}

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

// Fuzz_QuickSortInPlaceFunc fuzz test for QuickSortInPlaceFunc.
func Fuzz_QuickSortInPlaceFunc(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{1, 2})
	f.Add([]byte{2, 3, 4})
	f.Add([]byte{2, 4, 3})
	f.Add([]byte{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add([]byte{10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	f.Add([]byte{10, 6, 8, 10, 9, 7, 5, 2, 4, 3, 1})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		QuickSortInPlaceFunc(slice, func(a, b int) int { return a - b })
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_QuickSortInPlace fuzz test for QuickSortInPlace.
func Fuzz_QuickSortInPlace(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{1, 2})
	f.Add([]byte{2, 3, 4})
	f.Add([]byte{2, 4, 3})
	f.Add([]byte{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add([]byte{10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	f.Add([]byte{10, 6, 8, 10, 9, 7, 5, 2, 4, 3, 1})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		QuickSortInPlace(slice)
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_QuickSortFunc fuzz test for QuickSortFunc.
func Fuzz_QuickSortFunc(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{1, 2})
	f.Add([]byte{2, 3, 4})
	f.Add([]byte{2, 4, 3})
	f.Add([]byte{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add([]byte{10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	f.Add([]byte{10, 6, 8, 10, 9, 7, 5, 2, 4, 3, 1})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		slice = QuickSortFunc(slice, func(a, b int) int { return a - b })
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_QuickSort fuzz test for QuickSort.
func Fuzz_QuickSort(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{1, 2})
	f.Add([]byte{2, 3, 4})
	f.Add([]byte{2, 4, 3})
	f.Add([]byte{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	f.Add([]byte{10, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	f.Add([]byte{10, 6, 8, 10, 9, 7, 5, 2, 4, 3, 1})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		slice = QuickSort(slice)
		assert.IsNonDecreasing(t, slice)
	})
}

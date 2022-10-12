package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fuzz "github.com/AdaLogics/go-fuzz-headers"
)

// Fuzz_BubbleSortInPlaceFunc fuzz test for BubbleSortInPlaceFunc.
func Fuzz_BubbleSortInPlaceFunc(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{1, 2})
	f.Add([]byte{2, 3, 4})
	f.Add([]byte{2, 4, 3})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		BubbleSortInPlaceFunc(slice, func(a, b int) int { return a - b })
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_BubbleSortInPlace fuzz test for BubbleSortInPlace.
func Fuzz_BubbleSortInPlace(f *testing.F) {
	f.Add([]byte{0})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		BubbleSortInPlace(slice)
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_BubbleSortFunc fuzz test for BubbleSortFunc.
func Fuzz_BubbleSortFunc(f *testing.F) {
	f.Add([]byte{0})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		slice = BubbleSortFunc(slice, func(a, b int) int { return a - b })
		assert.IsNonDecreasing(t, slice)
	})
}

// Fuzz_BubbleSort fuzz test for BubbleSort.
func Fuzz_BubbleSort(f *testing.F) {
	f.Add([]byte{0})
	f.Fuzz(func(t *testing.T, data []byte) {
		fuzzer := fuzz.NewConsumer(data)
		var slice []int
		if err := fuzzer.CreateSlice(&slice); err != nil {
			t.Skip(err)
		}
		slice = BubbleSort(slice)
		assert.IsNonDecreasing(t, slice)
	})
}

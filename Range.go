package slices

// Range returns a new slice containing the range of integers from start to start+count-1.
func Range(start, count int) []int {
	return MapRange(start, count, func(i int) int { return i })
}

// RangeStep returns a new slice containing the range of integers from start to start+step*count-1 with the given step.
func RangeStep(start, count, step int) []int {
	return MapRangeStep(start, count, step, func(i int) int { return i })
}

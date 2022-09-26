package slices

import "math"

// ExtractRangeStep returns a slice of elements from source, starting at start, with count elements, and with a step of step.
func ExtractRangeStep[E any](source []E, start int, count int, step int) []E {
	if count <= 0 || len(source) == 0 {
		return []E{}
	}
	if step > 0 {
		if start >= len(source) || start + (count-1) * step < 0 {
			return []E{}
		} else if start < 0 {
			diff := int(math.Ceil(float64(-start) / float64(step)))
			start += diff * step
			count -= diff
		}
		if start + (count-1) * step >= len(source) {
			count = (len(source) - start + step - 1) / step
		}
	} else if step < 0 {
		if start < 0 || start + (count-1) * step >= len(source) {
			return []E{}
		} else if start >= len(source) {
			diff := int(math.Ceil(float64(start - len(source) + 1) / float64(-step)))
			start += diff * step
			count -= diff
		}
		if start + (count-1) * step < 0 {
			count = (start + 1) / -step
		}
	} else {
		if start < 0 || start >= len(source) {
			return []E{}
		}
		result := make([]E, count)
		for i := 0; i < count; i++ {
			result[i] = source[start]
		}
		return result
	}
	if step == 1 {
		return source[start : start+count]
	}
	result := make([]E, 0, count)
	for i := 0; i < count; i++ {
		pos := start + i*step
		if pos >= 0 && pos < len(source) {
			result = append(result, source[pos])
		}
	}
	return result
}

// ExtractRange returns a slice of elements from source, starting at start, with count elements.
func ExtractRange[E any](source []E, start int, count int) []E {
	return ExtractRangeStep(source, start, count, 1)
}

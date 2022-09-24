package slices

// EqualsFunc returns true if the two slices are equal by the given predicate.
func EqualsFunc[E any](a []E, b []E, equals func(E, E) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equals(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Equals returns true if the two slices are equal.
func Equals[E comparable](a []E, b []E) bool {
	return EqualsFunc(a, b, func(a, b E) bool { return a == b })
}

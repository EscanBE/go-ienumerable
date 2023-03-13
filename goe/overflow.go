package goe

// add64p performs + operation on two int64 operands and panic if overflow
func add64p(a, b int64) int64 {
	c := a + b
	if (c > a) == (b > 0) {
		return c
	}
	panic("overflow")
}

package go_ienumerable

func (src *enumerable[T]) Average() float64 {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()

	sum := src.SumFloat64()
	if sum == 0.0 {
		return 0
	}

	return sum / float64(len(src.data))
}

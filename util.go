package go_ienumerable

func (src *enumerable[T]) assertSrcNonNil() {
	if src == nil {
		panic("src is nil")
	}
}

func (src *enumerable[T]) copyExceptData() IEnumerable[T] {
	return &enumerable[T]{
		data:             nil,
		equalsComparator: src.equalsComparator,
		lessComparator:   src.lessComparator,
	}
}

func (src *enumerable[T]) withData(data []T) IEnumerable[T] {
	src.data = data
	return src
}

func copySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	if src != nil {
		copy(dst, src)
	}
	return dst
}

package go_ienumerable

func (src *enumerable[T]) ChunkToHolder(size int) ChunkHolder[T] {
	src.assertSrcNonNil()
	src.assertSizeGt0(size)

	if len(src.data) < 1 {
		return NewChunkHolder[T]()
	}

	result := make([][]T, 0)
	tmpSlide := make([]T, 0)

	for _, ele := range src.data {
		if len(tmpSlide) < size {
			tmpSlide = append(tmpSlide, ele) // just append ele
		} else {
			result = append(result, tmpSlide) // persist previous
			tmpSlide = make([]T, 0)           // remake
			tmpSlide = append(tmpSlide, ele)  // append ele
		}
	}

	if len(tmpSlide) > 0 {
		result = append(result, tmpSlide)
	}

	return NewChunkHolder[T](result...)
}

func (src *enumerable[T]) ChunkToAny(size int) IEnumerable[[]any] {
	src.assertSrcNonNil()
	src.assertSizeGt0(size)

	if len(src.data) < 1 {
		return NewIEnumerable[[]any]()
	}

	result := make([][]any, 0)
	tmpSlide := make([]any, 0)

	for _, ele := range src.data {
		if len(tmpSlide) < size {
			tmpSlide = append(tmpSlide, ele) // just append ele
		} else {
			result = append(result, tmpSlide) // persist previous
			tmpSlide = make([]any, 0)         // remake
			tmpSlide = append(tmpSlide, ele)  // append ele
		}
	}

	if len(tmpSlide) > 0 {
		result = append(result, tmpSlide)
	}

	return NewIEnumerable[[]any](result...)
}

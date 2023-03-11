package go_ienumerable

func (src *enumerable[T]) Select(selector func(v T) any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertSelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	result := make([]any, len(src.data))

	for i, v := range src.data {
		result[i] = selector(v)
	}

	return NewIEnumerable[any](result...)
}

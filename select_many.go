package go_ienumerable

func (src *enumerable[T]) SelectMany(selector func(v T) []any) IEnumerable[any] {
	src.assertSrcNonNil()
	src.assertArraySelectorNonNil(selector)

	if len(src.data) < 1 {
		return NewIEnumerable[any]()
	}

	transformed := make([]any, 0)

	for _, v := range src.data {
		transformed = append(transformed, selector(v)...)
	}

	return NewIEnumerable[any](transformed...)
}

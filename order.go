package go_ienumerable

import "sort"

func (src *enumerable[T]) Order() IEnumerable[T] {
	src.assertSrcNonNil()

	if src.lessComparer == nil {
		panicRequire(requireLessComparer)
	}

	return src.OrderBy(src.lessComparer)
}

func (src *enumerable[T]) OrderBy(lessComparer func(left, right T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertComparerNonNil(lessComparer)

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	copied := copySlice(src.data)

	sort.SliceStable(copied, func(i, j int) bool {
		return lessComparer(copied[i], copied[j])
	})

	return src.copyExceptData().withData(copied)
}

func (src *enumerable[T]) OrderByDescending() IEnumerable[T] {
	src.assertSrcNonNil()

	if src.lessComparer == nil {
		panicRequire(requireLessComparer)
	}

	return src.OrderByDescendingBy(src.lessComparer)
}

func (src *enumerable[T]) OrderByDescendingBy(lessComparer func(left, right T) bool) IEnumerable[T] {
	src.assertSrcNonNil()
	src.assertComparerNonNil(lessComparer)

	if len(src.data) < 1 {
		return src.copyExceptData().withEmptyData()
	}

	copied := copySlice(src.data)

	sort.SliceStable(copied, func(i, j int) bool {
		return lessComparer(copied[j], copied[i])
	})

	return src.copyExceptData().withData(copied)
}

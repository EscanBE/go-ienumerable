package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
)

func (src *enumerable[T]) WithComparerFrom(another IEnumerable[T]) IEnumerable[T] {
	eAnother := another.(*enumerable[T])
	if eAnother.defaultComparer != nil {
		src.defaultComparer = eAnother.defaultComparer
	}
	return src
}

func (src *enumerable[T]) WithDefaultComparer(comparer comparers.IComparer[T]) IEnumerable[T] {
	if comparer == nil {
		src.defaultComparer = nil
	} else {
		src.defaultComparer = comparers.ConvertFromComparerIntoDefaultComparer[T](comparer)
	}
	return src
}

func (src *enumerable[T]) WithDefaultComparerAny(comparer comparers.IComparer[any]) IEnumerable[T] {
	if comparer == nil {
		src.defaultComparer = nil
	} else {
		src.defaultComparer = comparer
	}
	return src
}

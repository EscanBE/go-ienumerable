package goe

func (src *enumerable[T]) LongCount(optionalPredicate OptionalPredicate[T]) int64 {
	count := src.Count(optionalPredicate)
	return int64(count)
}

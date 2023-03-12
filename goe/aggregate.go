package goe

func (src *enumerable[T]) Aggregate(f func(pr, v T) T) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertAggregateFuncNonNil(f)

	result := src.data[0]
	for i := 1; i < len(src.data); i++ {
		result = f(result, src.data[i])
	}

	return result
}

func (src *enumerable[T]) AggregateWithSeed(seed T, f func(pr, v T) T) T {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertAggregateFuncNonNil(f)

	result := seed
	for i := 0; i < len(src.data); i++ {
		result = f(result, src.data[i])
	}

	return result
}

func (src *enumerable[T]) AggregateWithAnySeed(seed any, f func(pr any, v T) any) any {
	src.assertSrcNonNil()
	src.assertSrcNonEmpty()
	src.assertAggregateAnySeedFuncNonNil(f)

	result := seed
	for i := 0; i < len(src.data); i++ {
		result = f(result, src.data[i])
	}

	return result
}

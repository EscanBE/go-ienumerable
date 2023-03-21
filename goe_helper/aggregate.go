package goe_helper

import "github.com/EscanBE/go-ienumerable/goe"

// Aggregate applies an accumulator function over a sequence.
func Aggregate[TSource any](source goe.IEnumerable[TSource], f func(pr, v TSource) TSource) TSource {
	assertCollectionNotNil(source, "source")
	assertCollectionEmpty(source, "source")
	assertAccumulatorFunctionNotNil(f)

	result := source.ElementAt(0, false)
	for i := 1; i < source.Count(nil); i++ {
		result = f(result, source.ElementAt(i, false))
	}

	return result
}

// AggregateSeed applies an accumulator function over a sequence. The specified seed value is used as the initial accumulator value.
func AggregateSeed[TSource, TAccumulate any](source goe.IEnumerable[TSource], seed TAccumulate, fAccumulator func(pr TAccumulate, v TSource) TAccumulate) TAccumulate {
	assertCollectionNotNil(source, "source")
	assertCollectionEmpty(source, "source")
	assertAccumulatorFunctionNotNil(fAccumulator)

	result := seed
	for i := 0; i < source.Count(nil); i++ {
		result = fAccumulator(result, source.ElementAt(i, false))
	}

	return result
}

// AggregateSeedTransform applies an accumulator function over a sequence. The specified seed value is used as the initial accumulator value.
func AggregateSeedTransform[TSource, TAccumulate, TResult any](source goe.IEnumerable[TSource], seed TAccumulate, fAccumulator func(pr TAccumulate, v TSource) TAccumulate, fResultSelector func(fr TAccumulate) TResult) TResult {
	assertCollectionNotNil(source, "source")
	assertCollectionEmpty(source, "source")
	assertAccumulatorFunctionNotNil(fAccumulator)
	assertResultSelectorFunctionNotNil(fResultSelector)

	result := seed
	for i := 0; i < source.Count(nil); i++ {
		result = fAccumulator(result, source.ElementAt(i, false))
	}

	return fResultSelector(result)
}

package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
)

// ZipSelect applies a specified function to the corresponding elements of two sequences, producing a sequence of the results.
func ZipSelect[TFirst, TSecond, TResult any](first goe.IEnumerable[TFirst], second goe.IEnumerable[TSecond], resultSelector func(v1 TFirst, v2 TSecond) TResult) goe.IEnumerable[TResult] {
	if resultSelector == nil {
		panic("result selector is required")
	}

	assertCollectionNotNil(first, "first")
	assertCollectionNotNil(second, "second")

	minLength := min(first.Count(nil), second.Count(nil))

	results := make([]TResult, minLength)

	for i := 0; i < minLength; i++ {
		results[i] = resultSelector(first.ElementAt(i, false), second.ElementAt(i, false))
	}

	return goe.NewIEnumerable[TResult](results...)
}

// Zip2 produces a sequence of tuples with elements from the two specified sequences.
func Zip2[TFirst, TSecond any](first goe.IEnumerable[TFirst], second goe.IEnumerable[TSecond]) goe.IEnumerable[goe.ValueTuple2[TFirst, TSecond]] {
	assertCollectionNotNil(first, "first")
	assertCollectionNotNil(second, "second")

	minLength := min(first.Count(nil), second.Count(nil))

	results := make([]goe.ValueTuple2[TFirst, TSecond], minLength)

	for i := 0; i < minLength; i++ {
		results[i] = goe.ValueTuple2[TFirst, TSecond]{
			First:  first.ElementAt(i, false),
			Second: second.ElementAt(i, false),
		}
	}

	return goe.NewIEnumerable[goe.ValueTuple2[TFirst, TSecond]](results...)
}

// Zip3 produces a sequence of tuples with elements from the three specified sequences.
func Zip3[TFirst, TSecond, TThird any](first goe.IEnumerable[TFirst], second goe.IEnumerable[TSecond], third goe.IEnumerable[TThird]) goe.IEnumerable[goe.ValueTuple3[TFirst, TSecond, TThird]] {
	assertCollectionNotNil(first, "first")
	assertCollectionNotNil(second, "second")
	assertCollectionNotNil(third, "third")

	minLength := min(first.Count(nil), second.Count(nil), third.Count(nil))

	results := make([]goe.ValueTuple3[TFirst, TSecond, TThird], minLength)

	for i := 0; i < minLength; i++ {
		results[i] = goe.ValueTuple3[TFirst, TSecond, TThird]{
			First:  first.ElementAt(i, false),
			Second: second.ElementAt(i, false),
			Third:  third.ElementAt(i, false),
		}
	}

	return goe.NewIEnumerable[goe.ValueTuple3[TFirst, TSecond, TThird]](results...)
}

func min(numbers ...int) int {
	min := numbers[0]
	for i := 1; i < len(numbers); i++ {
		n := numbers[i]
		if n < min {
			min = n
		}
	}
	return min
}

package helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
)

func assertCollectionNotNil[T any](collection goe.IEnumerable[T], collectionName string) {
	if collection == nil {
		panic(fmt.Sprintf("%s collection is nil", collectionName))
	}
}

func assertCollectionEmpty[T any](collection goe.IEnumerable[T], collectionName string) {
	if collection.Count(nil) < 1 {
		panic(fmt.Sprintf("%s collection is empty", collectionName))
	}
}

func assertAccumulatorFunctionNotNil[TAccumulate, TSource any](f func(pr TAccumulate, v TSource) TAccumulate) {
	if f == nil {
		panic("accumulator function is nil")
	}
}

func assertResultSelectorFunctionNotNil[TSource, TResult any](f func(v TSource) TResult) {
	if f == nil {
		panic("result selector function is nil")
	}
}

func assertResultSelectorFunctionNotNil2[TSource, TCollection, TResult any](f func(v TSource, c TCollection) TResult) {
	if f == nil {
		panic("result selector function is nil")
	}
}

func assertManyResultSelectorFunctionNotNil[TSource, TResult any](f func(v TSource) []TResult) {
	if f == nil {
		panic("result selector function is nil")
	}
}

func assertCollectionSelectorFunctionNotNil[TSource, TResult any](f func(v TSource) []TResult) {
	if f == nil {
		panic("collection selector function is nil")
	}
}

package goe

import (
	"fmt"
	comparers "github.com/EscanBE/go-ienumerable/goe/comparers2"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"sort"
)

// ensure implementation
var _ IOrderedEnumerable[any] = &orderedEnumerable[any]{}

type orderedEnumerable[T any] struct {
	sourceIEnumerable  IEnumerable[T]
	chainableComparers []chainableComparer[T]
}

type chainableComparer[T any] struct {
	keySelector KeySelector[T]
	compareFunc CompareFunc[any]
	orderType   chainableComparerOrderType
}

type chainableComparerOrderType byte

//goland:noinspection GoSnakeCaseUsage
const (
	CLC_ASC  chainableComparerOrderType = 0
	CLC_DESC chainableComparerOrderType = 1
)

// newIOrderedEnumerable returns a new IOrderedEnumerable with the same type as data elements
func newIOrderedEnumerable[T any](src IEnumerable[T], keySelector KeySelector[T], compareFunc CompareFunc[any], orderType chainableComparerOrderType) IOrderedEnumerable[T] {
	if keySelector == nil {
		keySelector = func(src T) any {
			return src
		}
	}
	return (&orderedEnumerable[T]{
		sourceIEnumerable: src,
	}).chainMoreComparer(keySelector, compareFunc, orderType)
}

func (o *orderedEnumerable[T]) ThenBy(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	return o.chainMoreComparer(keySelector, compareFunc, CLC_ASC)
}

func (o *orderedEnumerable[T]) ThenByDescending(keySelector KeySelector[T], compareFunc CompareFunc[any]) IOrderedEnumerable[T] {
	return o.chainMoreComparer(keySelector, compareFunc, CLC_DESC)
}

func (o *orderedEnumerable[T]) GetOrderedEnumerable() IEnumerable[T] {
	o.assertSrcNonNil()

	e := o.sourceIEnumerable.(*enumerable[T])

	result := e.copyExceptData()

	if len(e.data) > 0 {
		copied := copySlice(e.data)

		sort.SliceStable(copied, func(i, j int) bool {
			for i2, comparer := range o.chainableComparers {
				var v1, v2 T

				if comparer.orderType == CLC_ASC {
					v1 = copied[i]
					v2 = copied[j]
				} else {
					v1 = copied[j]
					v2 = copied[i]
				}

				var k1, k2 any

				if comparer.keySelector != nil {
					k1 = comparer.keySelector(v1)
					k2 = comparer.keySelector(v2)
				} else {
					k1 = v1
					k2 = v2
				}

				if comparer.compareFunc == nil {
					isNil1 := k1 == nil
					isNil2 := k2 == nil

					if !isNil1 {
						_, isNil1 = reflection.RootValueExtractor(k1)
					}
					if !isNil2 {
						_, isNil2 = reflection.RootValueExtractor(k2)
					}

					if isNil1 && isNil2 {
						continue // next comparer
					}

					if isNil1 {
						k1 = nil
					}
					if isNil2 {
						k2 = nil
					}

					var defaultComparer comparers.IComparer[any]
					var foundDefaultCompare bool

					if k1 != nil {
						defaultComparer, foundDefaultCompare = comparers.TryGetDefaultComparerFromValue(k1)
					}

					if !foundDefaultCompare && k2 != nil {
						defaultComparer, foundDefaultCompare = comparers.TryGetDefaultComparerFromValue(k2)
					}

					if !foundDefaultCompare {
						panic(fmt.Sprintf("no default comparer found for %T", func() any {
							if k1 == nil {
								return k2
							} else {
								return k1
							}
						}()))
					}

					comparer.compareFunc = defaultComparer.CompareAny
					o.chainableComparers[i2] = comparer
					//fmt.Println("Cached")
				} else {
					// fmt.Println("Re-use")
				}

				compareResult := comparer.compareFunc(k1, k2)
				if compareResult < 0 {
					return true
				}
				if compareResult > 0 {
					return false
				}
				continue
			}

			return false
		})

		result = result.withData(copied)
	}

	return result
}

func (o *orderedEnumerable[T]) chainMoreComparer(keySelector KeySelector[T], compareFunc CompareFunc[any], orderType chainableComparerOrderType) *orderedEnumerable[T] {
	o.assertSrcNonNil()
	assertKeySelectorNonNil[T](keySelector)

	return &orderedEnumerable[T]{
		sourceIEnumerable: o.sourceIEnumerable,
		chainableComparers: append(copySlice(o.chainableComparers), chainableComparer[T]{
			keySelector: keySelector,
			compareFunc: compareFunc,
			orderType:   orderType,
		}),
	}
}

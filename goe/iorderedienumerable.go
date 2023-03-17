package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"sort"
)

// ensure implementation
var _ IOrderedEnumerable[any] = &orderedEnumerable[any]{}

type orderedEnumerable[T any] struct {
	sourceIEnumerable  IEnumerable[T]
	chainableComparers []chainableComparer[T]
}

type chainableComparer[T any] struct {
	compareFunc CompareFunc[T]
	orderType   chainableComparerOrderType
}

type chainableComparerOrderType byte

//goland:noinspection GoSnakeCaseUsage
const (
	CLC_ASC  chainableComparerOrderType = 0
	CLC_DESC chainableComparerOrderType = 1
)

// newIOrderedEnumerable returns a new IOrderedEnumerable with the same type as data elements
func newIOrderedEnumerable[T any](src IEnumerable[T], compareFuncOrComparer interface{}, orderType chainableComparerOrderType) IOrderedEnumerable[T] {
	return (&orderedEnumerable[T]{
		sourceIEnumerable: src,
	}).chainMoreComparer(compareFuncOrComparer, orderType)
}

func (o *orderedEnumerable[T]) ThenBy(compareFuncOrComparer interface{}) IOrderedEnumerable[T] {
	o.assertSrcNonNil()

	return o.chainMoreComparer(compareFuncOrComparer, CLC_ASC)
}

func (o *orderedEnumerable[T]) ThenByDescending(compareFuncOrComparer interface{}) IOrderedEnumerable[T] {
	o.assertSrcNonNil()

	return o.chainMoreComparer(compareFuncOrComparer, CLC_DESC)
}

func (o *orderedEnumerable[T]) GetOrderedEnumerable() IEnumerable[T] {
	o.assertSrcNonNil()

	e := o.sourceIEnumerable.(*enumerable[T])

	result := e.copyExceptData()

	if len(e.data) > 0 {
		copied := copySlice(e.data)

		sort.SliceStable(copied, func(i, j int) bool {
			for _, comparer := range o.chainableComparers {
				var v1, v2 T

				if comparer.orderType == CLC_ASC {
					v1 = copied[i]
					v2 = copied[j]
				} else {
					v1 = copied[j]
					v2 = copied[i]
				}

				compareResult := comparer.compareFunc(v1, v2)
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

func (o *orderedEnumerable[T]) chainMoreComparer(compareFuncOrComparer interface{}, orderType chainableComparerOrderType) *orderedEnumerable[T] {
	var compareFunc CompareFunc[T]

	if compareFuncOrComparer == nil {
		panic(getErrorNilComparer())
	} else if cff, okCff := compareFuncOrComparer.(func(v1, v2 T) int); okCff {
		if cff == nil {
			panic(getErrorNilComparer())
		}

		compareFunc = cff
	} else if cft, okCft := compareFuncOrComparer.(CompareFunc[T]); okCft {
		if cft == nil {
			panic(getErrorNilComparer())
		}

		compareFunc = cft
	} else if cpr, okCpr := compareFuncOrComparer.(comparers.IComparer[T]); okCpr {
		/* This will never reach since comparers.IComparer[T] is an interface and there is a nil check above
		if cpr == nil {
			panic(getErrorNilComparer())
		}
		*/
		compareFunc = func(v1, v2 T) int {
			return cpr.Compare(v1, v2)
		}
	} else if cprA, okCprA := compareFuncOrComparer.(comparers.IComparer[any]); okCprA {
		/* This will never reach since comparers.IComparer[T] is an interface and there is a nil check above
		if cpr == nil {
			panic(getErrorNilComparer())
		}
		*/
		compareFunc = func(v1, v2 T) int {
			return cprA.Compare(v1, v2)
		}
	} else {
		panic(getErrorComparerMustBeCompareFuncOrIComparer())
	}

	return &orderedEnumerable[T]{
		sourceIEnumerable: o.sourceIEnumerable,
		chainableComparers: append(copySlice(o.chainableComparers), chainableComparer[T]{
			compareFunc: compareFunc,
			orderType:   orderType,
		}),
	}
}

package goe

import "sort"

// ensure implementation
var _ IOrderedEnumerable[any] = &orderedEnumerable[any]{}

type orderedEnumerable[T any] struct {
	sourceIEnumerable  IEnumerable[T]
	chainableComparers []chainableComparer[T]
}

type chainableComparer[T any] struct {
	compare   func(v1, v2 T) int
	orderType chainableComparerOrderType
}

type chainableComparerOrderType byte

//goland:noinspection GoSnakeCaseUsage
const (
	CLC_ASC  chainableComparerOrderType = 0
	CLC_DESC chainableComparerOrderType = 1
)

// newIOrderedEnumerable returns a new IOrderedEnumerable with the same type as data elements
func newIOrderedEnumerable[T any](src IEnumerable[T], comparer func(v1, v2 T) int, orderType chainableComparerOrderType) IOrderedEnumerable[T] {
	return &orderedEnumerable[T]{
		sourceIEnumerable: src,
		chainableComparers: []chainableComparer[T]{
			{
				compare:   comparer,
				orderType: orderType,
			},
		},
	}
}

func (o *orderedEnumerable[T]) ThenBy(comparer func(v1 T, v2 T) int) IOrderedEnumerable[T] {
	o.assertSrcNonNil()
	o.assertComparerNonNil(comparer)

	return o.chainMoreLessComparer(comparer, CLC_ASC)
}

func (o *orderedEnumerable[T]) ThenByDescending(comparer func(v1 T, v2 T) int) IOrderedEnumerable[T] {
	o.assertSrcNonNil()
	o.assertComparerNonNil(comparer)

	return o.chainMoreLessComparer(comparer, CLC_DESC)
}

func (o *orderedEnumerable[T]) GetEnumerable() IEnumerable[T] {
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

				compareResult := comparer.compare(v1, v2)
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

func (o *orderedEnumerable[T]) chainMoreLessComparer(comparer func(v1 T, v2 T) int, orderType chainableComparerOrderType) *orderedEnumerable[T] {
	return &orderedEnumerable[T]{
		sourceIEnumerable: o.sourceIEnumerable,
		chainableComparers: append(copySlice(o.chainableComparers), chainableComparer[T]{
			compare:   comparer,
			orderType: orderType,
		}),
	}
}

func (o *orderedEnumerable[T]) assertSrcNonNil() {
	if o == nil {
		panic(getErrorSourceIsNil())
	}
}

func (_ *orderedEnumerable[T]) assertComparerNonNil(lessComparer func(v1 T, v2 T) int) {
	if lessComparer == nil {
		panic(getErrorNilComparer())
	}
}

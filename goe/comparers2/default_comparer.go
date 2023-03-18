package comparers

// ensure implementation
var _ IComparer[any] = wrappedAsDefaultComparer{}

type wrappedAsDefaultComparer struct {
	compareTypedFunc func(x, y any) int
	compareAnyFunc   func(x, y any) int
}

// ConvertFromComparerIntoDefaultComparer converts any comparer into a default comparer IComparer[any]
// that both CompareTyped and CompareAny receive any params for convenient purpose,
// still doing type check when calling CompareTyped as usual.
func ConvertFromComparerIntoDefaultComparer[T any](comparer IComparer[T]) IComparer[any] {
	return wrappedAsDefaultComparer{
		compareTypedFunc: func(x, y any) int {
			vx, okx := x.(T)
			if !okx {
				panic("can not cast x to typed value")
			}

			vy, oky := y.(T)
			if !oky {
				panic("can not cast y to typed value")
			}

			return comparer.CompareTyped(vx, vy)
		},
		compareAnyFunc: func(x, y any) int {
			return comparer.CompareAny(x, y)
		},
	}
}

func (w wrappedAsDefaultComparer) CompareTyped(x, y any) int {
	return w.compareTypedFunc(x, y)
}

func (w wrappedAsDefaultComparer) CompareAny(x, y any) int {
	return w.compareAnyFunc(x, y)
}

package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_TakeWhile(t *testing.T) {
	fPredicate := func(value int) bool {
		return value >= 5
	}
	tests := []struct {
		name      string
		src       IEnumerable[int]
		predicate func(value int) bool
		want      IEnumerable[int]
	}{
		{
			name:      "normal",
			src:       NewIEnumerable[int](6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 6, 7, 8, 9, 99),
			predicate: fPredicate,
			want:      NewIEnumerable[int](6, 7, 8, 9, 8, 7, 6, 5),
		},
		{
			name:      "false from beginning",
			src:       createIntEnumerable(1, 10),
			predicate: fPredicate,
			want:      createEmptyIntEnumerable(),
		},
		{
			name:      "empty",
			src:       createEmptyIntEnumerable(),
			predicate: fPredicate,
			want:      createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)

			// Predicate
			got := tt.src.TakeWhile(tt.predicate)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			got = tt.src.TakeWhile(Predicate[int](tt.predicate))

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			// PredicateWithIndex
			var predicateWithIndex func(v int, idx int) bool
			if tt.predicate == nil {
				predicateWithIndex = nil
			} else {
				predicateWithIndex = func(v int, idx int) bool {
					return tt.predicate(v)
				}
			}

			got = tt.src.TakeWhile(predicateWithIndex)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			got = tt.src.TakeWhile(PredicateWithIndex[int](predicateWithIndex))

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)
		})
	}

	t.Run("match by index", func(t *testing.T) {
		ieSrc := createIntEnumerable(100, 120)

		got := ieSrc.TakeWhile(func(_, index int) bool {
			return index < 5
		})

		want := createIntEnumerable(100, 104)
		if !assert.True(t, reflect.DeepEqual(want.ToArray(), got.ToArray())) {
			fmt.Printf("Want: [%v]\nGot:  [%v]", want.ToArray(), got.ToArray())
		}
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.TakeWhile(nil)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate func(v int) bool

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.TakeWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate Predicate[int]

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.TakeWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate func(v int, i int) bool

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.TakeWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate PredicateWithIndex[int]

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.TakeWhile(fPredicate)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc CompareFunc[int]
		ieSrc.TakeWhile(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc EqualsFunc[int]
		ieSrc.TakeWhile(badFunc)
	})
}

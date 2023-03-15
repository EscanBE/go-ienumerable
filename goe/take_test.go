package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func Test_enumerable_Take_TakeLast(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		count        int
		wantTake     IEnumerable[int]
		wantTakeLast IEnumerable[int]
	}{
		{
			name:         "partial",
			src:          createIntEnumerable(2, 11),
			count:        5,
			wantTake:     createIntEnumerable(2, 6),
			wantTakeLast: createIntEnumerable(7, 11),
		},
		{
			name:         "negative",
			src:          createIntEnumerable(2, 11),
			count:        -1 * (rand.Intn(100) + 1),
			wantTake:     createEmptyIntEnumerable(),
			wantTakeLast: createEmptyIntEnumerable(),
		},
		{
			name:         "all",
			src:          createIntEnumerable(1, 5),
			count:        5,
			wantTake:     createIntEnumerable(1, 5),
			wantTakeLast: createIntEnumerable(1, 5),
		},
		{
			name:         "all",
			src:          createIntEnumerable(1, 5),
			count:        6,
			wantTake:     createIntEnumerable(1, 5),
			wantTakeLast: createIntEnumerable(1, 5),
		},
		{
			name:         "empty",
			src:          createEmptyIntEnumerable(),
			count:        10,
			wantTake:     createEmptyIntEnumerable(),
			wantTakeLast: createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)
			gotTake := tt.src.Take(tt.count)
			gotTakeLast := tt.src.TakeLast(tt.count)

			assert.True(t, reflect.DeepEqual(tt.wantTake.ToArray(), gotTake.ToArray()))
			assert.True(t, reflect.DeepEqual(tt.wantTakeLast.ToArray(), gotTakeLast.ToArray()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, gotTake)
			bSrc.assertUnchangedIgnoreData(t, gotTakeLast)
		})
	}
}

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

		var badFunc LessFunc[int]
		ieSrc.TakeWhile(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc GreaterFunc[int]
		ieSrc.TakeWhile(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc EqualsFunc[int]
		ieSrc.TakeWhile(badFunc)
	})
}

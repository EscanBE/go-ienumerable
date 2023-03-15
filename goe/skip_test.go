package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func Test_enumerable_Skip_SkipLast(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		count        int
		wantSkip     IEnumerable[int]
		wantSkipLast IEnumerable[int]
	}{
		{
			name:         "partial",
			src:          createIntEnumerable(2, 11),
			count:        5,
			wantSkip:     createIntEnumerable(7, 11),
			wantSkipLast: createIntEnumerable(2, 6),
		},
		{
			name:         "negative",
			src:          createIntEnumerable(2, 11),
			count:        -1 * (rand.Intn(100) + 1),
			wantSkip:     createIntEnumerable(2, 11),
			wantSkipLast: createIntEnumerable(2, 11),
		},
		{
			name:         "all",
			src:          createIntEnumerable(1, 5),
			count:        5,
			wantSkip:     createEmptyIntEnumerable(),
			wantSkipLast: createEmptyIntEnumerable(),
		},
		{
			name:         "all",
			src:          createIntEnumerable(1, 5),
			count:        6,
			wantSkip:     createEmptyIntEnumerable(),
			wantSkipLast: createEmptyIntEnumerable(),
		},
		{
			name:         "empty",
			src:          createEmptyIntEnumerable(),
			count:        10,
			wantSkip:     createEmptyIntEnumerable(),
			wantSkipLast: createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)
			gotSkip := tt.src.Skip(tt.count)
			gotSkipLast := tt.src.SkipLast(tt.count)

			assert.True(t, reflect.DeepEqual(tt.wantSkip.ToArray(), gotSkip.ToArray()))
			assert.True(t, reflect.DeepEqual(tt.wantSkipLast.ToArray(), gotSkipLast.ToArray()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, gotSkip)
			bSrc.assertUnchangedIgnoreData(t, gotSkipLast)
		})
	}
}

func Test_enumerable_SkipWhile(t *testing.T) {
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
			want:      NewIEnumerable[int](4, 3, 2, 1, 6, 7, 8, 9, 99),
		},
		{
			name:      "false from beginning then take all",
			src:       createIntEnumerable(1, 10),
			predicate: fPredicate,
			want:      createIntEnumerable(1, 10),
		},
		{
			name:      "empty",
			src:       createEmptyIntEnumerable(),
			predicate: fPredicate,
			want:      createEmptyIntEnumerable(),
		},
		{
			name:      "no match at all",
			src:       createIntEnumerable(5, 20),
			predicate: fPredicate,
			want:      createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)

			// Predicate
			got := tt.src.SkipWhile(tt.predicate)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			got = tt.src.SkipWhile(Predicate[int](tt.predicate))

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

			got = tt.src.SkipWhile(predicateWithIndex)

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			got = tt.src.SkipWhile(PredicateWithIndex[int](predicateWithIndex))

			if !assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray())) {
				fmt.Printf("Want: [%v]\nGot:  [%v]", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)
		})
	}

	t.Run("match by index", func(t *testing.T) {
		ieSrc := createIntEnumerable(100, 120)

		got := ieSrc.SkipWhile(func(_, index int) bool {
			return index < 5
		})

		want := createIntEnumerable(105, 120)
		if !assert.True(t, reflect.DeepEqual(want.ToArray(), got.ToArray())) {
			fmt.Printf("Want: [%v]\nGot:  [%v]", want.ToArray(), got.ToArray())
		}
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.SkipWhile(nil)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate func(v int) bool

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.SkipWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate Predicate[int]

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.SkipWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate func(v int, i int) bool

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.SkipWhile(fPredicate)
	})

	t.Run("not accepting nil predicate", func(t *testing.T) {
		ieSrc := createIntEnumerable(1, 10)

		var fPredicate PredicateWithIndex[int]

		defer deferExpectPanicContains(t, getErrorNilPredicate().Error(), true)

		ieSrc.SkipWhile(fPredicate)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc LessFunc[int]
		ieSrc.SkipWhile(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc GreaterFunc[int]
		ieSrc.SkipWhile(badFunc)
	})

	t.Run("panic if not supported comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1)

		defer deferExpectPanicContains(t, "predicate must be", true)

		var badFunc EqualsFunc[int]
		ieSrc.SkipWhile(badFunc)
	})
}

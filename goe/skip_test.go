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
			wantSkip:     injectIntComparers(createEmptyIntEnumerable()),
			wantSkipLast: injectIntComparers(createEmptyIntEnumerable()),
		},
		{
			name:         "all",
			src:          createIntEnumerable(1, 5),
			count:        6,
			wantSkip:     injectIntComparers(createEmptyIntEnumerable()),
			wantSkipLast: injectIntComparers(createEmptyIntEnumerable()),
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

			assert.True(t, reflect.DeepEqual(tt.wantSkip.exposeData(), gotSkip.exposeData()))
			assert.True(t, reflect.DeepEqual(tt.wantSkipLast.exposeData(), gotSkipLast.exposeData()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, gotSkip)
			bSrc.assertUnchangedIgnoreData(t, gotSkipLast)
		})
	}
}

func Test_enumerable_SkipWhile(t *testing.T) {
	tests := []struct {
		name          string
		src           IEnumerable[int]
		predicate     func(value int) bool
		predicateWidx func(value int, index int) bool
		want          IEnumerable[int]
	}{
		{
			name: "normal",
			src:  NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1),
			predicate: func(v int) bool {
				return v <= 5
			},
			predicateWidx: func(v, i int) bool {
				return v <= 5
			},
			want: NewIEnumerable[int](6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1),
		},
		{
			name: "false from beginning",
			src:  createIntEnumerable(1, 10),
			predicate: func(v int) bool {
				return v >= 5
			},
			predicateWidx: func(v, i int) bool {
				return v >= 5
			},
			want: createIntEnumerable(1, 10),
		},
		{
			name: "empty",
			src:  createEmptyIntEnumerable(),
			predicate: func(v int) bool {
				return v >= 0
			},
			predicateWidx: func(v, i int) bool {
				return v >= 0
			},
			want: createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)

			// SkipWhile
			got := tt.src.SkipWhile(tt.predicate)

			assert.True(t, reflect.DeepEqual(tt.want.exposeData(), got.exposeData()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			// SkipWhileWidx
			got = tt.src.SkipWhileWidx(tt.predicateWidx)

			assert.True(t, reflect.DeepEqual(tt.want.exposeData(), got.exposeData()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)
		})
	}

	t.Run("nil predicate", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "predicate is nil")
		}()

		createRandomIntEnumerable(2).SkipWhile(nil)
	})

	t.Run("nil predicate", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "predicate is nil")
		}()

		createRandomIntEnumerable(2).SkipWhileWidx(nil)
	})
}

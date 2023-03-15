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
			want: createIntEnumerable(1, 5),
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
			want: createEmptyIntEnumerable(),
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

			// TakeWhile
			got := tt.src.TakeWhile(tt.predicate)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)

			// TakeWhileWidx
			got = tt.src.TakeWhileWidx(tt.predicateWidx)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

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

		createRandomIntEnumerable(2).TakeWhile(nil)
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

		createRandomIntEnumerable(2).TakeWhileWidx(nil)
	})
}

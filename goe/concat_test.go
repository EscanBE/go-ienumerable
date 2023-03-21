package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Concat(t *testing.T) {
	tests := []struct {
		name      string
		first     IEnumerable[int]
		second    IEnumerable[int]
		want      IEnumerable[int]
		wantPanic bool
	}{
		{
			name:      "first nil",
			first:     nil,
			second:    createEmptyIntEnumerable(),
			wantPanic: true,
		},
		{
			name:      "second nil",
			first:     createEmptyIntEnumerable(),
			second:    nil,
			wantPanic: true,
		},
		{
			name:   "both empty + bring along",
			first:  NewIEnumerable[int](),
			second: createEmptyIntEnumerable(),
			want:   createEmptyIntEnumerable(),
		},
		{
			name:   "both non empty",
			first:  NewIEnumerable[int](6, 5, 4),
			second: NewIEnumerable[int](9, 1, 7),
			want:   NewIEnumerable[int](6, 5, 4, 9, 1, 7),
		},
		{
			name:   "first empty, second non-empty",
			first:  NewIEnumerable[int](),
			second: NewIEnumerable[int](9, 1, 7),
			want:   NewIEnumerable[int](9, 1, 7),
		},
		{
			name:   "first non-empty, second empty",
			first:  NewIEnumerable[int](6, 5, 4),
			second: NewIEnumerable[int](),
			want:   NewIEnumerable[int](6, 5, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bFirst := backupForAssetUnchanged(tt.first)
			bSecond := backupForAssetUnchanged(tt.second)

			if tt.wantPanic && tt.first == nil {
				return
			}

			defer deferWantPanicDepends(t, tt.wantPanic)

			got := tt.first.Concat(tt.second)

			if assert.Len(t, got.ToArray(), tt.want.Count(nil)) {
				assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))
			}

			bFirst.assertUnchanged(t, tt.first)
			bSecond.assertUnchanged(t, tt.second)
			bFirst.assertUnchangedIgnoreData(t, got)
		})
	}
}

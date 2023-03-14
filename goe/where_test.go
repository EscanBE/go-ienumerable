package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Where(t *testing.T) {
	tests := []struct {
		name      string
		src       IEnumerable[int]
		predicate func(int) bool
		want      IEnumerable[int]
	}{
		{
			name: "partial",
			src:  createIntEnumerable(1, 10),
			predicate: func(i int) bool {
				return i >= 5
			},
			want: createIntEnumerable(5, 10),
		},
		{
			name: "none match",
			src:  createIntEnumerable(1, 5),
			predicate: func(i int) bool {
				return i > 5
			},
			want: injectIntComparers(createEmptyIntEnumerable()),
		},
		{
			name: "empty",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 0
			},
			want: createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)
			got := tt.src.Where(tt.predicate)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), got.ToArray()))

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)
		})
	}
}

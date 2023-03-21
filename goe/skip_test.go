package goe

import (
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

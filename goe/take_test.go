package goe

import (
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

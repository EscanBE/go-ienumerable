package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func Test_enumerable_Take(t *testing.T) {
	tests := []struct {
		name  string
		src   IEnumerable[int]
		count int
		want  IEnumerable[int]
	}{
		{
			name:  "partial",
			src:   createIntEnumerable(2, 11),
			count: 5,
			want:  createIntEnumerable(2, 6),
		},
		{
			name:  "negative",
			src:   createIntEnumerable(2, 11),
			count: -1 * (rand.Intn(100) + 1),
			want:  injectIntComparers(createEmptyIntEnumerable()),
		},
		{
			name:  "all",
			src:   createIntEnumerable(1, 5),
			count: 5,
			want:  createIntEnumerable(1, 5),
		},
		{
			name:  "all",
			src:   createIntEnumerable(1, 5),
			count: 6,
			want:  createIntEnumerable(1, 5),
		},
		{
			name:  "empty",
			src:   createEmptyIntEnumerable(),
			count: 10,
			want:  createEmptyIntEnumerable(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)
			got := tt.src.Take(tt.count)

			wantS := tt.want.exposeData()
			gotS := got.exposeData()

			if len(wantS) == 0 && len(wantS) == len(gotS) {

			} else if !assert.True(t, reflect.DeepEqual(wantS, gotS)) {
				fmt.Printf("Expect: %v\nActual: %v\n", tt.want.exposeData(), got.exposeData())
			}

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, got)
		})
	}
}

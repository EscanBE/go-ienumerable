package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_ToArray(t *testing.T) {
	tests := []struct {
		name string
		src  IEnumerable[int]
		want []int
	}{
		{
			name: "empty",
			src:  createEmptyIntEnumerable(),
			want: []int{},
		},
		{
			name: "with data",
			src:  createIntEnumerable(4, 8),
			want: []int{4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)
			got := tt.src.ToArray()

			if len(tt.want) == 0 && len(tt.want) == len(got) {

			} else if !assert.True(t, reflect.DeepEqual(tt.want, got)) {
				fmt.Printf("Expect: %v\nActual: %v\n", tt.want, got)
			}

			bSrc.assertUnchanged(t, tt.src)
		})
	}
}

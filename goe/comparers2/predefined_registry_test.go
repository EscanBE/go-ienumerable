package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_getDefaultComparerByKindAndType(t *testing.T) {
	tests := []struct {
		name        string
		key         DefaultComparerKey
		x           any
		y           any
		wantCompare int
	}{
		{
			name: "int64",
			key: DefaultComparerKey{
				Kind: reflect.Int64,
				Type: reflect.TypeOf(int64(2)),
			},
			x:           2.2,
			y:           1,
			wantCompare: 1,
		},
		{
			name: "pointer int64",
			key: DefaultComparerKey{
				Kind: reflect.String,
				Type: reflect.TypeOf("x"),
			},
			x:           "22",
			y:           "21",
			wantCompare: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantCompare, mappedDefaultComparers[tt.key].CompareTyped(tt.x, tt.y), "CompareTyped(%v,%v)", tt.x, tt.y)
		})
	}
}

func Test_getDefaultComparerKeyFromSampleValue(t *testing.T) {
	tests := []struct {
		name        string
		sampleValue any
		want        DefaultComparerKey
		wantPanic   bool
	}{
		{
			name:        "int64",
			sampleValue: int64(1),
			want: DefaultComparerKey{
				Kind: reflect.Int64,
				Type: reflect.TypeOf(int64(2)),
			},
			wantPanic: false,
		},
		{
			name: "pointer int64",
			sampleValue: func() any {
				var i int64 = 1
				return &i
			}(),
			want: DefaultComparerKey{
				Kind: reflect.Int64,
				Type: reflect.TypeOf(int64(2)),
			},
			wantPanic: false,
		},
		{
			name:        "nil",
			sampleValue: nil,
			wantPanic:   true,
		},
		{
			name: "nil string",
			sampleValue: func() any {
				var s *string
				return s
			}(),
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer deferExpectPanicContains(t, "sample value can not be nil or invalid", tt.wantPanic)
			assert.Equalf(t, tt.want, getDefaultComparerKeyFromSampleValue(tt.sampleValue), "getDefaultComparerKeyFromSampleValue(%v)", tt.sampleValue)
		})
	}
}

func TestTryGetDefaultComparer(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		gotComparer, gotOk := TryGetDefaultComparer[int64]()
		assert.True(t, gotOk)
		assert.NotNil(t, gotComparer)
		assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
	})
	t.Run("pointer int64", func(t *testing.T) {
		//vo := reflect.ValueOf(*new(*int64))
		//fmt.Println(vo.Kind(), vo.Type().Elem(), vo.Elem())
		key, err := tryGetDefaultComparerKeyFromSampleValue(*new(*int64))
		fmt.Println(key)
		assert.Nil(t, err)
		gotComparer, gotOk := TryGetDefaultComparer[*int64]()
		assert.True(t, gotOk)
		assert.NotNil(t, gotComparer)
		if gotOk {
			assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
		}
	})
}

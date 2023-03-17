package reflection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsNumericKind(t *testing.T) {
	i8 := int8(99)
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{
			value: nil,
			want:  false,
		},
		{
			value: &i8,
			want:  false,
		},
		{
			value: int8(1),
			want:  true,
		},
		{
			value: uint8(1),
			want:  true,
		},
		{
			value: int16(1),
			want:  true,
		},
		{
			value: uint16(1),
			want:  true,
		},
		{
			value: int32(1),
			want:  true,
		},
		{
			value: uint32(1),
			want:  true,
		},
		{
			value: int64(1),
			want:  true,
		},
		{
			value: uint64(1),
			want:  true,
		},
		{
			value: int(1),
			want:  true,
		},
		{
			value: uint(1),
			want:  true,
		},
		{
			value: float32(1),
			want:  true,
		},
		{
			value: float64(1),
			want:  true,
		},
		{
			value: complex(float32(1), float32(1)),
			want:  true,
		},
		{
			value: complex(float64(1), float64(1)),
			want:  true,
		},
		{
			value: "string",
			want:  false,
		},
		{
			value: true,
			want:  false,
		},
		{
			value: []int{},
			want:  false,
		},
		{
			value: [2]int{},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsNumericKind(reflect.ValueOf(tt.value)), "IsNumericKind(%v)", tt.value)
		})
	}
}

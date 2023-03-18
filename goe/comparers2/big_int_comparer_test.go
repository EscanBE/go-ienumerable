package comparers

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_bigIntComparer_CompareTyped(t *testing.T) {
	greater := new(big.Int).SetInt64(22)
	less := new(big.Int).SetInt64(11)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BigIntComparer.CompareTyped(less, greater))
		assert.Equal(t, -1, BigIntComparer.CompareTyped(nil, less))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BigIntComparer.CompareTyped(greater, less))
		assert.Equal(t, 1, BigIntComparer.CompareTyped(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BigIntComparer.CompareTyped(less, less))
		assert.Zero(t, BigIntComparer.CompareTyped(greater, greater))
		assert.Zero(t, BigIntComparer.CompareTyped(nil, nil))
		var x, y *big.Int
		assert.Zero(t, BigIntComparer.CompareTyped(x, y))
	})
}

func Test_bigIntComparer_CompareAny(t *testing.T) {
	greater := new(big.Int).SetInt64(22)
	less := new(big.Int).SetInt64(11)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BigIntComparer.CompareAny(less, greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, less))
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigIntComparer.CompareAny(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BigIntComparer.CompareAny(nil, nil))
		assert.Zero(t, BigIntComparer.CompareAny(less, less))
		assert.Zero(t, BigIntComparer.CompareAny(greater, greater))
	})

	t.Run("less pointer", func(t *testing.T) {
		assert.Equal(t, -1, BigIntComparer.CompareAny(&less, greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(&less, &greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(&less, &greater))
		var nilBigInt *big.Int
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nilBigInt, less))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, &less))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nilBigInt, &less))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nilBigInt, greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nil, &greater))
		assert.Equal(t, -1, BigIntComparer.CompareAny(nilBigInt, &greater))
	})

	t.Run("greater pointer", func(t *testing.T) {
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, &less))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, less))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, &less))
		var nilBigInt *big.Int
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, nilBigInt))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, nilBigInt))
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigIntComparer.CompareAny(greater, nilBigInt))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BigIntComparer.CompareAny(&greater, nilBigInt))
	})

	t.Run("equals pointer", func(t *testing.T) {
		var nilBigInt1, nilBigInt2 *big.Int
		assert.Zero(t, BigIntComparer.CompareAny(nilBigInt1, nilBigInt2))
		assert.Zero(t, BigIntComparer.CompareAny(nil, nilBigInt2))
		assert.Zero(t, BigIntComparer.CompareAny(nilBigInt1, nil))
		assert.Zero(t, BigIntComparer.CompareAny(less, &less))
		assert.Zero(t, BigIntComparer.CompareAny(&less, less))
		assert.Zero(t, BigIntComparer.CompareAny(&less, &less))
		assert.Zero(t, BigIntComparer.CompareAny(greater, &greater))
		assert.Zero(t, BigIntComparer.CompareAny(&greater, greater))
		assert.Zero(t, BigIntComparer.CompareAny(&greater, &greater))
	})

	tests := []struct {
		name      string
		x         any
		y         any
		want      int
		wantPanic bool
	}{
		{
			name: "normal",
			x:    new(big.Int).SetInt64(22),
			y:    new(big.Int).SetInt64(11),
			want: 1,
		},
		{
			name: "normal",
			x:    new(big.Int).SetInt64(11),
			y:    new(big.Int).SetInt64(22),
			want: -1,
		},
		{
			name: "normal",
			x:    nil,
			y:    new(big.Int).SetInt64(11),
			want: -1,
		},
		{
			name: "normal",
			x:    new(big.Int).SetInt64(11),
			y:    nil,
			want: 1,
		},
		{
			name: "normal",
			x:    nil,
			y:    nil,
			want: 0,
		},
		{
			name: "any",
			x:    any(new(big.Int).SetInt64(22)),
			y:    any(new(big.Int).SetInt64(11)),
			want: 1,
		},
		{
			name: "*any",
			x: func() any {
				a := new(big.Int).SetInt64(11)
				return &a
			}(),
			y: func() any {
				a := new(big.Int).SetInt64(22)
				return &a
			}(),
			want: -1,
		},
		{
			name:      "not big Int",
			x:         int64(1),
			y:         new(big.Int).SetInt64(22),
			wantPanic: true,
		},
		{
			name:      "not big Int",
			x:         new(big.Int).SetInt64(22),
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not big Int",
			x:         nil,
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not big Int",
			x:         int64(1),
			y:         nil,
			wantPanic: true,
		},
		{
			name: "nil big Int",
			x: func() any {
				var t *big.Int
				return t
			}(),
			y: func() any {
				var t *big.Int
				return t
			}(),
			wantPanic: false,
		},
		{
			name: "wrapped time",
			x: func() any {
				var t = new(big.Int).SetInt64(11)
				at1 := any(&t)
				at2 := &at1
				at3 := any(&at2)
				at4 := &at3
				at5 := any(at4)
				at6 := &at5
				at7 := any(&at6)
				return at7
			}(),
			y: func() any {
				var t = new(big.Int).SetInt64(22)
				at1 := any(&t)
				at2 := any(&at1)
				at3 := any(&at2)
				at4 := &at3
				at5 := &at4
				at6 := any(&at5)
				at7 := &at6
				return at7
			}(),
			want:      -1,
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			defer deferExpectPanicContains(t, "can not be cast to big.Int", tt.wantPanic)
			assert.Equalf(t1, tt.want, BigIntComparer.CompareAny(tt.x, tt.y), "CompareAny(%v, %v)", tt.x, tt.y)
		})
	}
}

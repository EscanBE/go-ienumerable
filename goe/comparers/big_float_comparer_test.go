package comparers

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_bigFloatComparer_CompareTyped(t *testing.T) {
	greater := new(big.Float).SetFloat64(22)
	less := new(big.Float).SetFloat64(11)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BigFloatComparer.CompareTyped(less, greater))
		assert.Equal(t, -1, BigFloatComparer.CompareTyped(nil, less))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BigFloatComparer.CompareTyped(greater, less))
		assert.Equal(t, 1, BigFloatComparer.CompareTyped(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BigFloatComparer.CompareTyped(less, less))
		assert.Zero(t, BigFloatComparer.CompareTyped(greater, greater))
		assert.Zero(t, BigFloatComparer.CompareTyped(nil, nil))
		var x, y *big.Float
		assert.Zero(t, BigFloatComparer.CompareTyped(x, y))
	})
}

func Test_bigFloatComparer_CompareAny(t *testing.T) {
	greater := new(big.Float).SetFloat64(22)
	less := new(big.Float).SetFloat64(11)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BigFloatComparer.CompareAny(less, greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, less))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BigFloatComparer.CompareAny(nil, nil))
		assert.Zero(t, BigFloatComparer.CompareAny(less, less))
		assert.Zero(t, BigFloatComparer.CompareAny(greater, greater))
	})

	t.Run("less pointer", func(t *testing.T) {
		assert.Equal(t, -1, BigFloatComparer.CompareAny(&less, greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(&less, &greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(&less, &greater))
		var nilBigFloat *big.Float
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nilBigFloat, less))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, &less))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nilBigFloat, &less))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nilBigFloat, greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nil, &greater))
		assert.Equal(t, -1, BigFloatComparer.CompareAny(nilBigFloat, &greater))
	})

	t.Run("greater pointer", func(t *testing.T) {
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, &less))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, less))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, &less))
		var nilBigFloat *big.Float
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, nilBigFloat))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, nilBigFloat))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(greater, nilBigFloat))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BigFloatComparer.CompareAny(&greater, nilBigFloat))
	})

	t.Run("equals pointer", func(t *testing.T) {
		var nilBigFloat1, nilBigFloat2 *big.Float
		assert.Zero(t, BigFloatComparer.CompareAny(nilBigFloat1, nilBigFloat2))
		assert.Zero(t, BigFloatComparer.CompareAny(nil, nilBigFloat2))
		assert.Zero(t, BigFloatComparer.CompareAny(nilBigFloat1, nil))
		assert.Zero(t, BigFloatComparer.CompareAny(less, &less))
		assert.Zero(t, BigFloatComparer.CompareAny(&less, less))
		assert.Zero(t, BigFloatComparer.CompareAny(&less, &less))
		assert.Zero(t, BigFloatComparer.CompareAny(greater, &greater))
		assert.Zero(t, BigFloatComparer.CompareAny(&greater, greater))
		assert.Zero(t, BigFloatComparer.CompareAny(&greater, &greater))
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
			x:    new(big.Float).SetFloat64(22),
			y:    new(big.Float).SetFloat64(11),
			want: 1,
		},
		{
			name: "normal",
			x:    new(big.Float).SetFloat64(11),
			y:    new(big.Float).SetFloat64(22),
			want: -1,
		},
		{
			name: "normal",
			x:    nil,
			y:    new(big.Float).SetFloat64(11),
			want: -1,
		},
		{
			name: "normal",
			x:    new(big.Float).SetFloat64(11),
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
			x:    any(new(big.Float).SetFloat64(22)),
			y:    any(new(big.Float).SetFloat64(11)),
			want: 1,
		},
		{
			name: "*any",
			x: func() any {
				a := new(big.Float).SetFloat64(11)
				return &a
			}(),
			y: func() any {
				a := new(big.Float).SetFloat64(22)
				return &a
			}(),
			want: -1,
		},
		{
			name:      "not big Float",
			x:         int64(1),
			y:         new(big.Float).SetFloat64(22),
			wantPanic: true,
		},
		{
			name:      "not big Float",
			x:         new(big.Float).SetFloat64(22),
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not big Float",
			x:         nil,
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not big Float",
			x:         int64(1),
			y:         nil,
			wantPanic: true,
		},
		{
			name: "nil big Float",
			x: func() any {
				var t *big.Float
				return t
			}(),
			y: func() any {
				var t *big.Float
				return t
			}(),
			wantPanic: false,
		},
		{
			name: "wrapped time",
			x: func() any {
				var t = new(big.Float).SetFloat64(11)
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
				var t = new(big.Float).SetFloat64(22)
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
			defer deferExpectPanicContains(t, "can not be cast to big.Float", tt.wantPanic)
			assert.Equalf(t1, tt.want, BigFloatComparer.CompareAny(tt.x, tt.y), "CompareAny(%v, %v)", tt.x, tt.y)
		})
	}
}

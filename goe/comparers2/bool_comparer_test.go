package comparers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_boolComparer_CompareTyped(t *testing.T) {
	greater := true
	less := false

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BoolComparer.CompareTyped(less, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BoolComparer.CompareTyped(greater, less))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BoolComparer.CompareTyped(less, less))
		assert.Zero(t, BoolComparer.CompareTyped(greater, greater))
	})
}

func Test_boolComparer_CompareAny(t *testing.T) {
	greater := true
	less := false

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, BoolComparer.CompareAny(less, greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, less))
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BoolComparer.CompareAny(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, BoolComparer.CompareAny(nil, nil))
		assert.Zero(t, BoolComparer.CompareAny(less, less))
		assert.Zero(t, BoolComparer.CompareAny(greater, greater))
	})

	t.Run("less pointer", func(t *testing.T) {
		assert.Equal(t, -1, BoolComparer.CompareAny(&less, greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(&less, &greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(&less, &greater))
		var nilBool *bool
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, less))
		assert.Equal(t, -1, BoolComparer.CompareAny(nilBool, less))
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, &less))
		assert.Equal(t, -1, BoolComparer.CompareAny(nilBool, &less))
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(nilBool, greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(nil, &greater))
		assert.Equal(t, -1, BoolComparer.CompareAny(nilBool, &greater))
	})

	t.Run("greater pointer", func(t *testing.T) {
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, &less))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, less))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, &less))
		var nilBool *bool
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, nilBool))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, nilBool))
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, BoolComparer.CompareAny(greater, nilBool))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, BoolComparer.CompareAny(&greater, nilBool))
	})

	t.Run("equals pointer", func(t *testing.T) {
		var nilBool1, nilBool2 *bool
		assert.Zero(t, BoolComparer.CompareAny(nilBool1, nilBool2))
		assert.Zero(t, BoolComparer.CompareAny(nil, nilBool2))
		assert.Zero(t, BoolComparer.CompareAny(nilBool1, nil))
		assert.Zero(t, BoolComparer.CompareAny(less, &less))
		assert.Zero(t, BoolComparer.CompareAny(&less, less))
		assert.Zero(t, BoolComparer.CompareAny(&less, &less))
		assert.Zero(t, BoolComparer.CompareAny(greater, &greater))
		assert.Zero(t, BoolComparer.CompareAny(&greater, greater))
		assert.Zero(t, BoolComparer.CompareAny(&greater, &greater))
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
			x:    true,
			y:    false,
			want: 1,
		},
		{
			name: "normal",
			x:    false,
			y:    true,
			want: -1,
		},
		{
			name: "any",
			x:    any(true),
			y:    any(false),
			want: 1,
		},
		{
			name: "*any",
			x: func() any {
				a := false
				return &a
			}(),
			y: func() any {
				a := true
				return &a
			}(),
			want: -1,
		},
		{
			name:      "not bool",
			x:         int64(1),
			y:         "1",
			wantPanic: true,
		},
		{
			name:      "not bool",
			x:         "1",
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not bool",
			x:         nil,
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not bool",
			x:         int64(1),
			y:         nil,
			wantPanic: true,
		},
		{
			name: "nil bool",
			x: func() any {
				var s *bool
				return s
			}(),
			y: func() any {
				var s *bool
				return s
			}(),
			wantPanic: false,
		},
		{
			name: "wrapped bool",
			x: func() any {
				var s = false
				at1 := any(&s)
				at2 := &at1
				at3 := any(&at2)
				at4 := &at3
				at5 := any(at4)
				at6 := &at5
				at7 := any(&at6)
				return at7
			}(),
			y: func() any {
				var s = true
				at1 := any(&s)
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
			defer deferExpectPanicContains(t, "can not be cast to bool", tt.wantPanic)
			assert.Equalf(t1, tt.want, BoolComparer.CompareAny(tt.x, tt.y), "CompareAny(%v, %v)", tt.x, tt.y)
		})
	}
}

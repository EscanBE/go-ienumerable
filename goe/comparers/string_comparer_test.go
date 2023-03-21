package comparers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringComparer_CompareTyped(t *testing.T) {
	greater := "22"
	less := "11"

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, StringComparer.CompareTyped(less, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, StringComparer.CompareTyped(greater, less))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, StringComparer.CompareTyped(less, less))
		assert.Zero(t, StringComparer.CompareTyped(greater, greater))
	})
}

func Test_stringComparer_CompareAny(t *testing.T) {
	greater := "22"
	less := "11"

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, StringComparer.CompareAny(less, greater))
		assert.Equal(t, -1, StringComparer.CompareAny(nil, less))
		assert.Equal(t, -1, StringComparer.CompareAny(nil, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, StringComparer.CompareAny(greater, less))
		assert.Equal(t, 1, StringComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, StringComparer.CompareAny(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, StringComparer.CompareAny(nil, nil))
		assert.Zero(t, StringComparer.CompareAny(less, less))
		assert.Zero(t, StringComparer.CompareAny(greater, greater))
	})

	t.Run("less pointer", func(t *testing.T) {
		assert.Equal(t, -1, StringComparer.CompareAny(&less, greater))
		assert.Equal(t, -1, StringComparer.CompareAny(&less, &greater))
		assert.Equal(t, -1, StringComparer.CompareAny(&less, &greater))
		var nilString *string
		assert.Equal(t, -1, StringComparer.CompareAny(nil, less))
		assert.Equal(t, -1, StringComparer.CompareAny(nilString, less))
		assert.Equal(t, -1, StringComparer.CompareAny(nil, &less))
		assert.Equal(t, -1, StringComparer.CompareAny(nilString, &less))
		assert.Equal(t, -1, StringComparer.CompareAny(nil, greater))
		assert.Equal(t, -1, StringComparer.CompareAny(nilString, greater))
		assert.Equal(t, -1, StringComparer.CompareAny(nil, &greater))
		assert.Equal(t, -1, StringComparer.CompareAny(nilString, &greater))
	})

	t.Run("greater pointer", func(t *testing.T) {
		assert.Equal(t, 1, StringComparer.CompareAny(greater, &less))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, less))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, &less))
		var nilString *string
		assert.Equal(t, 1, StringComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, StringComparer.CompareAny(greater, nilString))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, nilString))
		assert.Equal(t, 1, StringComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, StringComparer.CompareAny(greater, nilString))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, StringComparer.CompareAny(&greater, nilString))
	})

	t.Run("equals pointer", func(t *testing.T) {
		var nilString1, nilString2 *string
		assert.Zero(t, StringComparer.CompareAny(nilString1, nilString2))
		assert.Zero(t, StringComparer.CompareAny(nil, nilString2))
		assert.Zero(t, StringComparer.CompareAny(nilString1, nil))
		assert.Zero(t, StringComparer.CompareAny(less, &less))
		assert.Zero(t, StringComparer.CompareAny(&less, less))
		assert.Zero(t, StringComparer.CompareAny(&less, &less))
		assert.Zero(t, StringComparer.CompareAny(greater, &greater))
		assert.Zero(t, StringComparer.CompareAny(&greater, greater))
		assert.Zero(t, StringComparer.CompareAny(&greater, &greater))
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
			x:    "22",
			y:    "11",
			want: 1,
		},
		{
			name: "normal",
			x:    "21",
			y:    "22",
			want: -1,
		},
		{
			name: "any",
			x:    any("22"),
			y:    any("21"),
			want: 1,
		},
		{
			name: "*any",
			x: func() any {
				a := "21"
				return &a
			}(),
			y: func() any {
				a := "22"
				return &a
			}(),
			want: -1,
		},
		{
			name:      "not string",
			x:         int64(1),
			y:         "1",
			wantPanic: true,
		},
		{
			name:      "not string",
			x:         "1",
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not string",
			x:         nil,
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not string",
			x:         int64(1),
			y:         nil,
			wantPanic: true,
		},
		{
			name: "nil string",
			x: func() any {
				var s *string
				return s
			}(),
			y: func() any {
				var s *string
				return s
			}(),
			wantPanic: false,
		},
		{
			name: "wrapped string",
			x: func() any {
				var s = "21"
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
				var s = "22"
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
			defer deferExpectPanicContains(t, "can not be cast to string", tt.wantPanic)
			assert.Equalf(t1, tt.want, StringComparer.CompareAny(tt.x, tt.y), "CompareAny(%v, %v)", tt.x, tt.y)
		})
	}
}

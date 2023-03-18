package comparers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_timeTimeComparer_CompareTyped(t *testing.T) {
	greater := time.Now()
	less := greater.Add(-24 * time.Hour)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, TimeComparer.CompareTyped(less, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, TimeComparer.CompareTyped(greater, less))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, TimeComparer.CompareTyped(less, less))
		assert.Zero(t, TimeComparer.CompareTyped(greater, greater))
	})
}

func Test_timeTimeComparer_CompareAny(t *testing.T) {
	greater := time.Now()
	less := greater.Add(-24 * time.Hour)

	t.Run("less", func(t *testing.T) {
		assert.Equal(t, -1, TimeComparer.CompareAny(less, greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, less))
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, greater))
	})

	t.Run("greater", func(t *testing.T) {
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, less))
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, TimeComparer.CompareAny(less, nil))
	})

	t.Run("equals", func(t *testing.T) {
		assert.Zero(t, TimeComparer.CompareAny(nil, nil))
		assert.Zero(t, TimeComparer.CompareAny(less, less))
		assert.Zero(t, TimeComparer.CompareAny(greater, greater))
	})

	t.Run("less pointer", func(t *testing.T) {
		assert.Equal(t, -1, TimeComparer.CompareAny(&less, greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(&less, &greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(&less, &greater))
		var nilTime *time.Time
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, less))
		assert.Equal(t, -1, TimeComparer.CompareAny(nilTime, less))
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, &less))
		assert.Equal(t, -1, TimeComparer.CompareAny(nilTime, &less))
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(nilTime, greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(nil, &greater))
		assert.Equal(t, -1, TimeComparer.CompareAny(nilTime, &greater))
	})

	t.Run("greater pointer", func(t *testing.T) {
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, &less))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, less))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, &less))
		var nilTime *time.Time
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, nilTime))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, nilTime))
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, nil))
		assert.Equal(t, 1, TimeComparer.CompareAny(greater, nilTime))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, nil))
		assert.Equal(t, 1, TimeComparer.CompareAny(&greater, nilTime))
	})

	t.Run("equals pointer", func(t *testing.T) {
		var nilTime1, nilTime2 *time.Time
		assert.Zero(t, TimeComparer.CompareAny(nilTime1, nilTime2))
		assert.Zero(t, TimeComparer.CompareAny(nil, nilTime2))
		assert.Zero(t, TimeComparer.CompareAny(nilTime1, nil))
		assert.Zero(t, TimeComparer.CompareAny(less, &less))
		assert.Zero(t, TimeComparer.CompareAny(&less, less))
		assert.Zero(t, TimeComparer.CompareAny(&less, &less))
		assert.Zero(t, TimeComparer.CompareAny(greater, &greater))
		assert.Zero(t, TimeComparer.CompareAny(&greater, greater))
		assert.Zero(t, TimeComparer.CompareAny(&greater, &greater))
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
			x:    time.Now().Add(time.Minute),
			y:    time.Now(),
			want: 1,
		},
		{
			name: "normal",
			x:    time.Now().Add(-time.Minute),
			y:    time.Now(),
			want: -1,
		},
		{
			name: "any",
			x:    any(time.Now().Add(time.Minute)),
			y:    any(time.Now()),
			want: 1,
		},
		{
			name: "*any",
			x: func() any {
				a := time.Now().Add(-time.Minute)
				return &a
			}(),
			y: func() any {
				a := time.Now().Add(+time.Minute)
				return &a
			}(),
			want: -1,
		},
		{
			name:      "not time",
			x:         int64(1),
			y:         time.Now(),
			wantPanic: true,
		},
		{
			name:      "not time",
			x:         time.Now(),
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not time",
			x:         nil,
			y:         int64(1),
			wantPanic: true,
		},
		{
			name:      "not time",
			x:         int64(1),
			y:         nil,
			wantPanic: true,
		},
		{
			name: "nil time",
			x: func() any {
				var t *time.Time
				return t
			}(),
			y: func() any {
				var t *time.Time
				return t
			}(),
			wantPanic: false,
		},
		{
			name: "wrapped time",
			x: func() any {
				var t = time.Now()
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
				var t = time.Now().Add(time.Minute)
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
			defer deferExpectPanicContains(t, "can not be cast to time.Time", tt.wantPanic)
			assert.Equalf(t1, tt.want, TimeComparer.CompareAny(tt.x, tt.y), "CompareAny(%v, %v)", tt.x, tt.y)
		})
	}
}

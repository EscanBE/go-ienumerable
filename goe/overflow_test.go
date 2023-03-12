package goe

import (
	"math"
	"testing"
)

func Test_add64p(t *testing.T) {
	tests := []struct {
		name         string
		a            int64
		b            int64
		wantOverflow bool
	}{
		{
			name:         "within",
			a:            math.MaxInt64,
			b:            0,
			wantOverflow: false,
		},
		{
			name:         "within",
			a:            math.MaxInt64 - 1,
			b:            1,
			wantOverflow: false,
		},
		{
			name:         "overflow",
			a:            math.MaxInt64,
			b:            1,
			wantOverflow: true,
		},
		{
			name:         "within",
			a:            math.MinInt64,
			b:            0,
			wantOverflow: false,
		},
		{
			name:         "within",
			a:            math.MinInt64 + 1,
			b:            -1,
			wantOverflow: false,
		},
		{
			name:         "overflow",
			a:            math.MinInt64,
			b:            -1,
			wantOverflow: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer deferWantPanicDepends(t, tt.wantOverflow)
			_ = add64p(tt.a, tt.b)
		})
	}
}

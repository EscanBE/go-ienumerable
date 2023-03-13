package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_int8Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    int8
		y    int8
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MinInt8,
			y:    math.MinInt8,
			want: 0,
		},
		{
			x:    math.MaxInt8,
			y:    math.MaxInt8,
			want: 0,
		},
		{
			x:    math.MinInt8,
			y:    math.MaxInt8,
			want: -1,
		},
		{
			x:    math.MaxInt8,
			y:    math.MinInt8,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewInt8Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint8Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    uint8
		y    uint8
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MaxUint8,
			y:    math.MaxUint8,
			want: 0,
		},
		{
			x:    0,
			y:    math.MaxInt8,
			want: -1,
		},
		{
			x:    math.MaxInt8,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUint8Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int16Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    int16
		y    int16
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MinInt16,
			y:    math.MinInt16,
			want: 0,
		},
		{
			x:    math.MaxInt16,
			y:    math.MaxInt16,
			want: 0,
		},
		{
			x:    math.MinInt16,
			y:    math.MaxInt16,
			want: -1,
		},
		{
			x:    math.MaxInt16,
			y:    math.MinInt16,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewInt16Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint16Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    uint16
		y    uint16
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MaxUint16,
			y:    math.MaxUint16,
			want: 0,
		},
		{
			x:    0,
			y:    math.MaxInt16,
			want: -1,
		},
		{
			x:    math.MaxInt16,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUint16Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int32Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    int32
		y    int32
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MinInt32,
			y:    math.MinInt32,
			want: 0,
		},
		{
			x:    math.MaxInt32,
			y:    math.MaxInt32,
			want: 0,
		},
		{
			x:    math.MinInt32,
			y:    math.MaxInt32,
			want: -1,
		},
		{
			x:    math.MaxInt32,
			y:    math.MinInt32,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewInt32Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint32Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    uint32
		y    uint32
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MaxUint32,
			y:    math.MaxUint32,
			want: 0,
		},
		{
			x:    0,
			y:    math.MaxInt32,
			want: -1,
		},
		{
			x:    math.MaxInt32,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUint32Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int64Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    int64
		y    int64
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MinInt64,
			y:    math.MinInt64,
			want: 0,
		},
		{
			x:    math.MaxInt64,
			y:    math.MaxInt64,
			want: 0,
		},
		{
			x:    math.MinInt64,
			y:    math.MaxInt64,
			want: -1,
		},
		{
			x:    math.MaxInt64,
			y:    math.MinInt64,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewInt64Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint64Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    uint64
		y    uint64
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MaxUint64,
			y:    math.MaxUint64,
			want: 0,
		},
		{
			x:    0,
			y:    math.MaxInt64,
			want: -1,
		},
		{
			x:    math.MaxInt64,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUint64Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intComparer_Compare(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MinInt,
			y:    math.MinInt,
			want: 0,
		},
		{
			x:    math.MaxInt,
			y:    math.MaxInt,
			want: 0,
		},
		{
			x:    math.MinInt,
			y:    math.MaxInt,
			want: -1,
		},
		{
			x:    math.MaxInt,
			y:    math.MinInt,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewIntComparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintComparer_Compare(t *testing.T) {
	tests := []struct {
		x    uint
		y    uint
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    math.MaxUint,
			y:    math.MaxUint,
			want: 0,
		},
		{
			x:    0,
			y:    math.MaxInt,
			want: -1,
		},
		{
			x:    math.MaxInt,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUintComparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintptrComparer_Compare(t *testing.T) {
	tests := []struct {
		x       uintptr
		y       uintptr
		wantptr int
	}{
		{
			x:       0,
			y:       0,
			wantptr: 0,
		},
		{
			x:       1_000_000,
			y:       1_000_000,
			wantptr: 0,
		},
		{
			x:       0,
			y:       1_000_000,
			wantptr: -1,
		},
		{
			x:       1_000_000,
			y:       0,
			wantptr: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewUintptrComparer()
			if got := i.Compare(tt.x, tt.y); got != tt.wantptr {
				t.Errorf("Compare() = %v, want %v", got, tt.wantptr)
			}
		})
	}
}

func Test_float32Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    float32
		y    float32
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    -1 * math.MaxFloat32,
			y:    -1 * math.MaxFloat32,
			want: 0,
		},
		{
			x:    math.MaxFloat32,
			y:    math.MaxFloat32,
			want: 0,
		},
		{
			x:    -1 * math.MaxFloat32,
			y:    math.MaxFloat32,
			want: -1,
		},
		{
			x:    math.MaxFloat32,
			y:    -1 * math.MaxFloat32,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewFloat32Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_float64Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    float64
		y    float64
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    -1 * math.MaxFloat64,
			y:    -1 * math.MaxFloat64,
			want: 0,
		},
		{
			x:    math.MaxFloat64,
			y:    math.MaxFloat64,
			want: 0,
		},
		{
			x:    -1 * math.MaxFloat64,
			y:    math.MaxFloat64,
			want: -1,
		},
		{
			x:    math.MaxFloat64,
			y:    -1 * math.MaxFloat64,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewFloat64Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complex64Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    complex64
		y    complex64
		want int
	}{
		{
			x:    complex(0, 0),
			y:    0,
			want: 0,
		},
		{
			x:    0,
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(0, 0),
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(0, 0),
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(2, 2),
			y:    complex(1, 3),
			want: 1,
		},
		{
			x:    complex(1, 3),
			y:    complex(2, 2),
			want: -1,
		},
		{
			x:    complex(2, 3),
			y:    complex(2, 4),
			want: -1,
		},
		{
			x:    complex(2, 4),
			y:    complex(2, 3),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewComplex64Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complex128Comparer_Compare(t *testing.T) {
	tests := []struct {
		x    complex128
		y    complex128
		want int
	}{
		{
			x:    complex(0, 0),
			y:    0,
			want: 0,
		},
		{
			x:    0,
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(0, 0),
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(0, 0),
			y:    complex(0, 0),
			want: 0,
		},
		{
			x:    complex(2, 2),
			y:    complex(1, 3),
			want: 1,
		},
		{
			x:    complex(1, 3),
			y:    complex(2, 2),
			want: -1,
		},
		{
			x:    complex(2, 3),
			y:    complex(2, 4),
			want: -1,
		},
		{
			x:    complex(2, 4),
			y:    complex(2, 3),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewComplex128Comparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringComparer_Compare(t *testing.T) {
	tests := []struct {
		x    string
		y    string
		want int
	}{
		{
			x:    "",
			y:    "",
			want: 0,
		},
		{
			x:    "def",
			y:    "def",
			want: 0,
		},
		{
			x:    "def",
			y:    "deg",
			want: -1,
		},
		{
			x:    "def",
			y:    "dff",
			want: -1,
		},
		{
			x:    "def",
			y:    "eef",
			want: -1,
		},
		{
			x:    "eef",
			y:    "def",
			want: 1,
		},
		{
			x:    "dff",
			y:    "def",
			want: 1,
		},
		{
			x:    "deg",
			y:    "def",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewStringComparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boolComparer_Compare(t *testing.T) {
	tests := []struct {
		x    bool
		y    bool
		want int
	}{
		{
			x:    false,
			y:    false,
			want: 0,
		},
		{
			x:    false,
			y:    true,
			want: -1,
		},
		{
			x:    true,
			y:    true,
			want: 0,
		},
		{
			x:    true,
			y:    false,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			i := NewBoolComparer()
			if got := i.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partitionedComparer_fromEL(t *testing.T) {
	equals := func(x, y int8) bool {
		return x == y
	}

	less := func(x, y int8) bool {
		return x < y
	}

	tests := []struct {
		equals    func(x, y int8) bool
		less      func(x, y int8) bool
		x         int8
		y         int8
		want      int
		wantPanic bool
	}{
		{
			equals: equals,
			less:   less,
			x:      9,
			y:      9,
			want:   0,
		},
		{
			equals: equals,
			less:   less,
			x:      -9,
			y:      9,
			want:   -1,
		},
		{
			equals: equals,
			less:   less,
			x:      9,
			y:      -1,
			want:   1,
		},
		{
			equals:    equals,
			less:      nil,
			x:         99,
			y:         9,
			wantPanic: true,
		},
		{
			equals:    nil,
			less:      less,
			x:         99,
			y:         8,
			wantPanic: true,
		},
		{
			equals:    nil,
			less:      nil,
			x:         99,
			y:         7,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.wantPanic {
					if err == nil {
						t.Errorf("expect error")
						return
					}
					assert.Contains(t, fmt.Sprintf("%v", err), "comparer is nil")
				} else {
					if err != nil {
						t.Errorf("not want error but got %v", err)
					}
				}
			}()

			i := NewPartitionedComparerFromEL(tt.equals, tt.less)

			assert.Equal(t, tt.want, i.Compare(tt.x, tt.y))
		})
	}
}

func Test_partitionedComparer_fromComparer(t *testing.T) {
	defaultComparer := NewInt8Comparer()

	tests := []struct {
		comparer  IComparer[int8]
		x         int8
		y         int8
		want      int
		wantPanic bool
	}{
		{
			comparer: defaultComparer,
			x:        9,
			y:        9,
			want:     0,
		},
		{
			comparer: defaultComparer,
			x:        -9,
			y:        9,
			want:     -1,
		},
		{
			comparer: defaultComparer,
			x:        9,
			y:        -1,
			want:     1,
		},
		{
			comparer:  nil,
			x:         99,
			y:         9,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			defer func() {
				err := recover()
				if tt.wantPanic {
					if err == nil {
						t.Errorf("expect error")
						return
					}
					assert.Contains(t, fmt.Sprintf("%v", err), "comparer is nil")
				} else {
					if err != nil {
						t.Errorf("not want error but got %v", err)
					}
				}
			}()

			i := NewPartitionedComparerFromComparer[int8](tt.comparer)

			assert.Equal(t, tt.want, i.Compare(tt.x, tt.y))
		})
	}
}

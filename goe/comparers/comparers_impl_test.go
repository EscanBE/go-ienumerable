package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
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
			comparer := Int8Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Uint8Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Int16Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			y:    math.MaxUint16,
			want: -1,
		},
		{
			x:    math.MaxUint16,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := Uint16Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Int32Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			y:    math.MaxUint32,
			want: -1,
		},
		{
			x:    math.MaxUint32,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := Uint32Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Int64Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			y:    math.MaxUint64,
			want: -1,
		},
		{
			x:    math.MaxUint64,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := Uint64Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := IntComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			y:    math.MaxUint,
			want: -1,
		},
		{
			x:    math.MaxUint,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := UintComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uintptrComparer_Compare(t *testing.T) {
	tests := []struct {
		x    uintptr
		y    uintptr
		want int
	}{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    1_000_000,
			y:    1_000_000,
			want: 0,
		},
		{
			x:    0,
			y:    1_000_000,
			want: -1,
		},
		{
			x:    1_000_000,
			y:    0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := UintptrComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
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
			comparer := Float32Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Float64Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Complex64Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := Complex128Comparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := StringComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
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
			comparer := BoolComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeComparer_durationComparer_Compare(t *testing.T) {
	tests := []struct {
		x    time.Duration
		y    time.Duration
		want int
	}{
		{
			x:    time.Minute,
			y:    time.Minute,
			want: 0,
		},
		{
			x:    time.Second,
			y:    time.Minute,
			want: -1,
		},
		{
			x:    time.Minute,
			y:    time.Second,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Duration %v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := DurationComparer

			if got := comparer.Compare(tt.x, tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tt.x, &tt.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
		t.Run(fmt.Sprintf("Time %v-%v", tt.x, tt.y), func(t *testing.T) {
			comparer := TimeComparer

			now := time.Now()
			tX := now.Add(tt.x)
			tY := now.Add(tt.y)

			if got := comparer.Compare(tX, tY); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}

			if got := comparer.ComparePointerMode(&tX, &tY); got != tt.want {
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
			assert.Equal(t, tt.want, i.ComparePointerMode(&tt.x, &tt.y))
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
			assert.Equal(t, tt.want, i.ComparePointerMode(&tt.x, &tt.y))
		})
	}

	t.Run("pointer mode", func(t *testing.T) {
		comparer := NewPartitionedComparerFromEL[int](func(v1, v2 int) bool {
			return v1 == v2
		}, func(v1, v2 int) bool {
			return v1 < v2
		})

		i := 2
		assert.Equal(t, 0, comparer.Compare(i, i))
		assert.Equal(t, 0, comparer.ComparePointerMode(&i, &i))
		assert.Equal(t, 0, comparer.ComparePointerMode(nil, nil))
		assert.Equal(t, 1, comparer.ComparePointerMode(&i, nil))
		assert.Equal(t, -1, comparer.ComparePointerMode(nil, &i))
	})
}

func Test_wrappedComparer(t *testing.T) {
	t.Run("provide params of correct type", func(t *testing.T) {
		assert.Equal(t, 1, HideTypedComparer[int8](Int8Comparer).Compare(int8(3), int8(1)))
		assert.Equal(t, 1, HideTypedComparer[int](IntComparer).Compare(3, 1))
		assert.Equal(t, 1, HideTypedComparer[string](StringComparer).Compare("3", "1"))
		assert.Equal(t, 1, HideTypedComparer[bool](BoolComparer).Compare(true, false))
	})

	t.Run("provide params of correct type but wrapped to any", func(t *testing.T) {
		assert.Equal(t, 1, HideTypedComparer[int8](Int8Comparer).Compare(any(int8(3)), any(int8(1))))
		assert.Equal(t, 1, HideTypedComparer[int](IntComparer).Compare(any(3), any(1)))
		assert.Equal(t, 1, HideTypedComparer[string](StringComparer).Compare(any("3"), any("1")))
		assert.Equal(t, 1, HideTypedComparer[bool](BoolComparer).Compare(any(true), any(false)))
	})

	t.Run("resolve for both type or pointer value", func(t *testing.T) {
		hiddenComparer := HideTypedComparer[int8](Int8Comparer)
		x := any(int8(3))
		y := any(int8(1))
		px := &x
		py := &y
		assert.Equal(t, 1, hiddenComparer.Compare(x, y))
		assert.Equal(t, 1, hiddenComparer.ComparePointerMode(px, py))
	})

	t.Run("provide params of wrong type", func(t *testing.T) {
		hiddenInt8Comparer := HideTypedComparer[int8](Int8Comparer)
		params1 := int32(3)
		params2 := int32(1)

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
			}
		}()

		_ = hiddenInt8Comparer.Compare(params1, params2)
	})

	t.Run("provide params of wrong type (any)", func(t *testing.T) {
		hiddenInt8Comparer := HideTypedComparer[int8](Int8Comparer)
		params1 := int32(3)
		params2 := int32(1)

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
			}
		}()

		_ = hiddenInt8Comparer.Compare(any(params1), any(params2))
	})
}

func Test_compareBothMode(t *testing.T) {
	testCompareBothModeFor[int8](t, 1, 3)
	testCompareBothModeFor[uint8](t, 1, 3)
	testCompareBothModeFor[byte](t, 1, 3) // uint8 alias
	testCompareBothModeFor[int16](t, 1, 3)
	testCompareBothModeFor[uint16](t, 1, 3)
	testCompareBothModeFor[int32](t, 1, 3)
	testCompareBothModeFor[rune](t, 1, 3) // int32 alias
	testCompareBothModeFor[uint32](t, 1, 3)
	testCompareBothModeFor[int64](t, 1, 3)
	testCompareBothModeFor[uint64](t, 1, 3)
	testCompareBothModeFor[int](t, 1, 3)
	testCompareBothModeFor[uint](t, 1, 3)
	testCompareBothModeFor[uintptr](t, 1, 3)
	testCompareBothModeFor[float32](t, 1, 3)
	testCompareBothModeFor[float64](t, 1, 3)
	testCompareBothModeFor[string](t, "1", "3")
	testCompareBothModeFor[bool](t, false, true)
	testCompareBothModeFor[complex64](t, 1, 3)
	testCompareBothModeFor[complex128](t, 1, 3)
	testCompareBothModeFor[time.Time](t, time.Now(), time.Now().Add(time.Second))
	testCompareBothModeFor[time.Duration](t, time.Second, time.Minute)
}

func testCompareBothModeFor[T any](t *testing.T, small, big T) {
	comparer := GetDefaultComparer[T]()
	assert.Equal(t, 0, comparer.Compare(big, big))
	assert.Equal(t, 0, comparer.Compare(small, small))
	assert.Equal(t, -1, comparer.Compare(small, big))
	assert.Equal(t, 1, comparer.Compare(big, small))
	aSmall := any(small)
	aBig := any(big)
	assert.Equal(t, 0, comparer.ComparePointerMode(nil, nil))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &aBig))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &aSmall))
	assert.Equal(t, 1, comparer.ComparePointerMode(&aBig, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&aSmall, nil))
}
package comparers

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_getDefaultComparerByKindAndType(t *testing.T) {
	tests := []struct {
		name        string
		key         reflect.Type
		x           any
		y           any
		wantCompare int
	}{
		{
			name:        "int64",
			key:         reflect.TypeOf(int64(2)),
			x:           2.2,
			y:           1,
			wantCompare: 1,
		},
		{
			name:        "pointer int64",
			key:         reflect.TypeOf("x"),
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
		want        reflect.Type
		wantPanic   bool
	}{
		{
			name:        "int64",
			sampleValue: int64(1),
			want:        reflect.TypeOf(int64(2)),
			wantPanic:   false,
		},
		{
			name: "pointer int64",
			sampleValue: func() any {
				var i int64 = 1
				return &i
			}(),
			want:      reflect.TypeOf(int64(2)),
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
			want: reflect.TypeOf("string"),
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
		gotComparer, gotOk := TryGetDefaultComparer[*int64]()
		assert.True(t, gotOk)
		if gotOk {
			assert.NotNil(t, gotComparer)
			assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
		}
	})
	t.Run("pointer time.Time", func(t *testing.T) {
		gotComparer, gotOk := TryGetDefaultComparer[*time.Time]()
		assert.True(t, gotOk)
		if gotOk {
			assert.NotNil(t, gotComparer)
			assert.Equal(t, 1, gotComparer.CompareAny(time.Now().Add(time.Hour), time.Now()))
		}
	})
	t.Run("not supported type", func(t *testing.T) {
		type myType struct{}
		gotComparer, gotOk := TryGetDefaultComparer[myType]()
		assert.False(t, gotOk)
		assert.Nil(t, gotComparer)
	})
	t.Run("numeric same comparer", func(t *testing.T) {
		c1, _ := TryGetDefaultComparer[int]()
		assert.NotNil(t, c1)
		c2, _ := TryGetDefaultComparer[*uint16]()
		assert.NotNil(t, c2)
		c3, _ := TryGetDefaultComparer[float64]()
		assert.NotNil(t, c3)
		c4, _ := TryGetDefaultComparer[*complex128]()
		assert.NotNil(t, c4)
		assert.True(t, c1 == c2)
		assert.True(t, c1 == c3)
		assert.True(t, c1 == c4)
		c5, _ := TryGetDefaultComparer[time.Time]()
		assert.False(t, c1 == c5)
	})
}

func TestTryGetDefaultComparerFromValue(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		var i int64 = 5
		gotComparer, gotOk := TryGetDefaultComparerFromValue(i)
		assert.True(t, gotOk)
		assert.NotNil(t, gotComparer)
		assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
	})
	t.Run("pointer int64", func(t *testing.T) {
		var i int64
		i = 5
		gotComparer, gotOk := TryGetDefaultComparerFromValue(&i)
		assert.True(t, gotOk)
		if gotOk {
			assert.NotNil(t, gotComparer)
			assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
		}
	})
	t.Run("pointer time.Time", func(t *testing.T) {
		t2 := time.Now()
		gotComparer, gotOk := TryGetDefaultComparerFromValue(&t2)
		assert.True(t, gotOk)
		if gotOk {
			assert.NotNil(t, gotComparer)
			assert.Equal(t, 1, gotComparer.CompareAny(time.Now().Add(time.Hour), time.Now()))
		}
	})
	t.Run("not supported type", func(t *testing.T) {
		type myType struct{}
		gotComparer, gotOk := TryGetDefaultComparerFromValue(myType{})
		assert.False(t, gotOk)
		assert.Nil(t, gotComparer)

		gotComparer, gotOk = TryGetDefaultComparerFromValue(&myType{})
		assert.False(t, gotOk)
		assert.Nil(t, gotComparer)

		gotComparer, gotOk = TryGetDefaultComparerFromValue(nil)
		assert.False(t, gotOk)
		assert.Nil(t, gotComparer)
	})
	t.Run("numeric same comparer", func(t *testing.T) {
		var i int
		var ui16 uint16
		var f64 float64
		var cpl complex128
		c1, _ := TryGetDefaultComparerFromValue(&i)
		assert.NotNil(t, c1)
		c2, _ := TryGetDefaultComparerFromValue(&ui16)
		assert.NotNil(t, c2)
		c3, _ := TryGetDefaultComparerFromValue(&f64)
		assert.NotNil(t, c3)
		c4, _ := TryGetDefaultComparerFromValue(&cpl)
		assert.NotNil(t, c4)
		assert.True(t, c1 == c2)
		assert.True(t, c1 == c3)
		assert.True(t, c1 == c4)
		var t3 time.Time
		c5, _ := TryGetDefaultComparerFromValue(&t3)
		assert.False(t, c1 == c5)
	})
}

func TestGetDefaultComparer(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		gotComparer := GetDefaultComparer[int64]()
		assert.NotNil(t, gotComparer)
		assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
	})
	t.Run("pointer int64", func(t *testing.T) {
		gotComparer := GetDefaultComparer[*int64]()
		assert.NotNil(t, gotComparer)
		assert.Equal(t, 1, gotComparer.CompareAny(1.1, 1))
	})
	t.Run("pointer time.Time", func(t *testing.T) {
		gotComparer := GetDefaultComparer[*time.Time]()
		assert.NotNil(t, gotComparer)
		assert.Equal(t, 1, gotComparer.CompareAny(time.Now().Add(time.Hour), time.Now()))
	})
	t.Run("not supported type", func(t *testing.T) {
		type myType struct{}
		defer deferExpectPanicContains(t, "not found any default comparer for", true)
		_ = GetDefaultComparer[myType]()
	})
	t.Run("numeric same comparer", func(t *testing.T) {
		c1 := GetDefaultComparer[int]()
		assert.NotNil(t, c1)
		c2 := GetDefaultComparer[*uint16]()
		assert.NotNil(t, c2)
		c3 := GetDefaultComparer[float64]()
		assert.NotNil(t, c3)
		c4 := GetDefaultComparer[*complex128]()
		assert.NotNil(t, c4)
		assert.True(t, c1 == c2)
		assert.True(t, c1 == c3)
		assert.True(t, c1 == c4)
		c5 := GetDefaultComparer[time.Time]()
		assert.False(t, c1 == c5)
	})
}

type wrappedInt32 struct {
	value int32
}

var _ IComparer[wrappedInt32] = wrappedInt32Comparer{}

// Create a custom comparer for a type
type wrappedInt32Comparer struct {
}

func (i wrappedInt32Comparer) CompareTyped(x, y wrappedInt32) int {
	if x.value < y.value {
		return -1
	}

	if x.value > y.value {
		return 1
	}

	return 0
}

func (i wrappedInt32Comparer) CompareAny(any, any) int {
	panic("not implemented")
	return 0
}

func TestRegisterDefaultComparer(t *testing.T) {
	// This example shows how register and resolve default comparer for type wrappedInt32

	// register for auto resolve, run once at app boots up
	RegisterDefaultComparer[wrappedInt32](wrappedInt32Comparer{})

	// resolve
	comparer := GetDefaultComparer[wrappedInt32]()

	var big, small wrappedInt32
	big = wrappedInt32{
		value: 9,
	}
	small = wrappedInt32{
		value: 3,
	}

	assert.Equal(t, 0, comparer.CompareTyped(small, small))
	assert.Equal(t, 0, comparer.CompareTyped(big, big))
	assert.Equal(t, 1, comparer.CompareTyped(big, small))
	assert.Equal(t, -1, comparer.CompareTyped(small, big))
}

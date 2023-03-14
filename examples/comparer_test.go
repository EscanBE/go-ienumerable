package examples

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_comparer_1(t *testing.T) {
	comparer := comparers.Int32Comparer

	var big, small int32
	big = 9
	small = 3

	assert.Equal(t, 0, comparer.Compare(small, small))
	assert.Equal(t, 0, comparer.Compare(big, big))
	assert.Equal(t, 1, comparer.Compare(big, small))
	assert.Equal(t, -1, comparer.Compare(small, big))

	assert.Equal(t, 0, comparer.ComparePointerMode(&small, &small))
	assert.Equal(t, 0, comparer.ComparePointerMode(&big, &big))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, &small))
	assert.Equal(t, -1, comparer.ComparePointerMode(&small, &big))

	assert.Equal(t, 0, comparer.ComparePointerMode(nil, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&small, nil))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &big))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &small))
}

type wrappedInt32 struct {
	value int32
}

// Create a custom comparer for a type
type wrappedInt32Comparer struct {
}

func (i wrappedInt32Comparer) Compare(x, y wrappedInt32) int {
	if x.value < y.value {
		return -1
	}

	if x.value > y.value {
		return 1
	}

	return 0
}

func (i wrappedInt32Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(comparers.AnyPointerToType[wrappedInt32](x), comparers.AnyPointerToType[wrappedInt32](y))
}

// ensure implementation
var _ comparers.IComparer[wrappedInt32] = wrappedInt32Comparer{}

func Test_comparer_2(t *testing.T) {
	// This example shows how to use a custom comparer
	comparer := wrappedInt32Comparer{}

	var big, small wrappedInt32
	big = wrappedInt32{
		value: 9,
	}
	small = wrappedInt32{
		value: 3,
	}

	assert.Equal(t, 0, comparer.Compare(small, small))
	assert.Equal(t, 0, comparer.Compare(big, big))
	assert.Equal(t, 1, comparer.Compare(big, small))
	assert.Equal(t, -1, comparer.Compare(small, big))

	assert.Equal(t, 0, comparer.ComparePointerMode(&small, &small))
	assert.Equal(t, 0, comparer.ComparePointerMode(&big, &big))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, &small))
	assert.Equal(t, -1, comparer.ComparePointerMode(&small, &big))

	assert.Equal(t, 0, comparer.ComparePointerMode(nil, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&small, nil))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &big))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &small))
}

func Test_comparer_3(t *testing.T) {
	// This example shows how register and resolve default comparer for type wrappedInt32

	// register for auto resolve, run once at app boots up
	comparers.RegisterDefaultTypedComparer[wrappedInt32](wrappedInt32Comparer{}, false)

	// resolve
	comparer := comparers.GetDefaultComparer[wrappedInt32]()

	var big, small wrappedInt32
	big = wrappedInt32{
		value: 9,
	}
	small = wrappedInt32{
		value: 3,
	}

	assert.Equal(t, 0, comparer.Compare(small, small))
	assert.Equal(t, 0, comparer.Compare(big, big))
	assert.Equal(t, 1, comparer.Compare(big, small))
	assert.Equal(t, -1, comparer.Compare(small, big))

	assert.Equal(t, 0, comparer.ComparePointerMode(&small, &small))
	assert.Equal(t, 0, comparer.ComparePointerMode(&big, &big))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, &small))
	assert.Equal(t, -1, comparer.ComparePointerMode(&small, &big))

	assert.Equal(t, 0, comparer.ComparePointerMode(nil, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&big, nil))
	assert.Equal(t, 1, comparer.ComparePointerMode(&small, nil))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &big))
	assert.Equal(t, -1, comparer.ComparePointerMode(nil, &small))
}

package examples

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_comparer_1(t *testing.T) {
	comparer := comparers.NumericComparer

	var big, small int32
	big = 9
	small = 3

	assert.Equal(t, 0, comparer.CompareTyped(small, small))
	assert.Equal(t, 0, comparer.CompareTyped(big, big))
	assert.Equal(t, 1, comparer.CompareTyped(big, small))
	assert.Equal(t, -1, comparer.CompareTyped(small, big))

	assert.Equal(t, 0, comparer.CompareAny(&small, &small))
	assert.Equal(t, 0, comparer.CompareAny(&big, &big))
	assert.Equal(t, 1, comparer.CompareAny(&big, &small))
	assert.Equal(t, -1, comparer.CompareAny(&small, &big))

	assert.Equal(t, 0, comparer.CompareAny(nil, nil))
	assert.Equal(t, 1, comparer.CompareAny(&big, nil))
	assert.Equal(t, 1, comparer.CompareAny(&small, nil))
	assert.Equal(t, -1, comparer.CompareAny(nil, &big))
	assert.Equal(t, -1, comparer.CompareAny(nil, &small))
}

type wrappedInt32 struct {
	value int32
}

var _ comparers.IComparer[wrappedInt32] = wrappedInt32Comparer{}

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

func (i wrappedInt32Comparer) CompareAny(x, y any) int {
	return i.CompareAny(x, y)
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

	assert.Equal(t, 0, comparer.CompareTyped(small, small))
	assert.Equal(t, 0, comparer.CompareTyped(big, big))
	assert.Equal(t, 1, comparer.CompareTyped(big, small))
	assert.Equal(t, -1, comparer.CompareTyped(small, big))

	assert.Equal(t, 0, comparer.CompareAny(&small, &small))
	assert.Equal(t, 0, comparer.CompareAny(&big, &big))
	assert.Equal(t, 1, comparer.CompareAny(&big, &small))
	assert.Equal(t, -1, comparer.CompareAny(&small, &big))

	assert.Equal(t, 0, comparer.CompareAny(nil, nil))
	assert.Equal(t, 1, comparer.CompareAny(&big, nil))
	assert.Equal(t, 1, comparer.CompareAny(&small, nil))
	assert.Equal(t, -1, comparer.CompareAny(nil, &big))
	assert.Equal(t, -1, comparer.CompareAny(nil, &small))
}

func Test_comparer_3(t *testing.T) {
	// This example shows how register and resolve default comparer for type wrappedInt32

	// register for auto resolve, run once at app boots up
	comparers.RegisterDefaultComparer[wrappedInt32](wrappedInt32Comparer{})

	// resolve
	comparer := comparers.GetDefaultComparer[wrappedInt32]()

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

	assert.Equal(t, 0, comparer.CompareAny(&small, &small))
	assert.Equal(t, 0, comparer.CompareAny(&big, &big))
	assert.Equal(t, 1, comparer.CompareAny(&big, &small))
	assert.Equal(t, -1, comparer.CompareAny(&small, &big))

	assert.Equal(t, 0, comparer.CompareAny(nil, nil))
	assert.Equal(t, 1, comparer.CompareAny(&big, nil))
	assert.Equal(t, 1, comparer.CompareAny(&small, nil))
	assert.Equal(t, -1, comparer.CompareAny(nil, &big))
	assert.Equal(t, -1, comparer.CompareAny(nil, &small))
}

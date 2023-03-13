package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

var _ IComparer[*int8] = ptrInt8ComparerImpl{}

type ptrInt8ComparerImpl struct {
}

func (p ptrInt8ComparerImpl) Compare(x, y *int8) int {
	if x == nil && y == nil {
		return 0
	}
	if x == nil {
		return -1
	}
	if y == nil {
		return 1
	}
	return Int8Comparer.Compare(*x, *y)
}

var ptrInt8Comparer = ptrInt8ComparerImpl{}

func TestRegisterDefaultTypedComparer(t *testing.T) {
	t.Run("unable to detect type", func(t *testing.T) {
		defer deferExpectPanicContains(t, "empty or <nil> type name")
		//goland:noinspection GoVarAndConstTypeMayBeOmitted
		var comparerAny IComparer[any] = HideTypedComparer[*int8](ptrInt8Comparer)
		RegisterDefaultTypedComparer[any](comparerAny, false)
	})

	t.Run("comparer is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "comparer is nil")
		RegisterDefaultTypedComparer[byte](nil, false)
	})
}

func TestRegisterDefaultComparerForType(t *testing.T) {
	t.Run("empty type name", func(t *testing.T) {
		defer deferExpectPanicContains(t, "empty or <nil> type name")
		//goland:noinspection GoVarAndConstTypeMayBeOmitted
		var comparerAny IComparer[any] = HideTypedComparer[*int8](ptrInt8Comparer)
		RegisterDefaultComparerForType[any]("", comparerAny, false)
	})

	t.Run("nil type name", func(t *testing.T) {
		typeName := fmt.Sprintf("%T", *new(any))
		defer deferExpectPanicContains(t, "empty or <nil> type name")
		//goland:noinspection GoVarAndConstTypeMayBeOmitted
		var comparerAny IComparer[any] = HideTypedComparer[*int8](ptrInt8Comparer)
		RegisterDefaultComparerForType[any](typeName, comparerAny, false)
	})

	t.Run("comparer is nil", func(t *testing.T) {
		randTypeName := fmt.Sprintf("%d-test-byte", rand.Int63())
		defer deferExpectPanicContains(t, "comparer is nil")
		RegisterDefaultComparerForType[byte](randTypeName, nil, false)
	})

	t.Run("register twice", func(t *testing.T) {
		randTypeName := fmt.Sprintf("%d-test-*int8", rand.Int63())
		RegisterDefaultComparerForType[*int8](randTypeName, ptrInt8Comparer, false) // first ok

		RegisterDefaultComparerForType[*int8](randTypeName, ptrInt8Comparer, true) // second ok with override enabled

		defer deferExpectPanicContains(t, "had been registered before")
		RegisterDefaultComparerForType[*int8](randTypeName, ptrInt8Comparer, false) // third fails without override
	})
}

func Test_GetDefaultComparerByName(t *testing.T) {
	for i := 1; i <= 7; i++ {
		typeName := fmt.Sprintf("int%d", int(math.Pow(2, float64(i))))
		t.Run("can get correctly", func(t *testing.T) {
			c1, f1 := TryGetDefaultComparerByTypeName(typeName)
			if i >= 3 && i <= 6 {
				assert.NotNil(t, c1)
				assert.True(t, f1)
			} else {
				assert.Nil(t, c1)
				assert.False(t, f1)
			}
		})

		t.Run("panic correctly", func(t *testing.T) {
			if i >= 3 && i <= 6 {
				c1 := GetDefaultComparerByTypeName(typeName)
				assert.NotNil(t, c1)
			} else {
				defer deferExpectPanicContains(t, fmt.Sprintf("no default comparer registered for type [%s]", typeName))
				_ = GetDefaultComparerByTypeName(typeName)
			}
		})
	}
}

func Test_GetDefaultComparer(t *testing.T) {
	t.Run("can get correctly", func(t *testing.T) {
		c1, f1 := TryGetDefaultComparer[*int16]()
		assert.Nil(t, c1)
		assert.False(t, f1)

		c1, f1 = TryGetDefaultComparer[int16]()
		assert.NotNil(t, c1)
		assert.True(t, f1)

		c1, f1 = TryGetDefaultComparer[any]()
		assert.Nil(t, c1)
		assert.False(t, f1)

		c1 = GetDefaultComparer[int16]()
		assert.NotNil(t, c1)
	})

	t.Run("panic when not registered", func(t *testing.T) {
		defer deferExpectPanicContains(t, fmt.Sprintf("no default comparer registered for type [*int16]"))
		_ = GetDefaultComparer[*int16]()
	})

	t.Run("panic when can not detect type", func(t *testing.T) {
		defer deferExpectPanicContains(t, fmt.Sprintf("unable to detect type for provided type"))
		_ = GetDefaultComparer[any]()
	})
}

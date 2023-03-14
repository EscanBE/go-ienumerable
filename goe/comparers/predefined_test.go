package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

type testStruct struct {
}

var _ IComparer[testStruct] = testStructComparerImpl{}

type testStructComparerImpl struct {
}

func (p testStructComparerImpl) Compare(x, y testStruct) int {
	return 0
}

func (p testStructComparerImpl) ComparePointerMode(x, y any) int {
	return 0
}

var testStructComparer = testStructComparerImpl{}

func TestRegisterDefaultTypedComparer(t *testing.T) {
	t.Run("unable to detect type", func(t *testing.T) {
		defer deferExpectPanicContains(t, "empty or <nil> type name")
		//goland:noinspection GoVarAndConstTypeMayBeOmitted
		var comparerAny IComparer[any] = ConvertToDefaultComparer[testStruct](testStructComparer)
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
		var comparerAny IComparer[any] = ConvertToDefaultComparer[testStruct](testStructComparer)
		RegisterDefaultComparerForType[any]("", comparerAny, false)
	})

	t.Run("nil type name", func(t *testing.T) {
		typeName := fmt.Sprintf("%T", *new(any))
		defer deferExpectPanicContains(t, "empty or <nil> type name")
		//goland:noinspection GoVarAndConstTypeMayBeOmitted
		var comparerAny IComparer[any] = ConvertToDefaultComparer[testStruct](testStructComparer)
		RegisterDefaultComparerForType[any](typeName, comparerAny, false)
	})

	t.Run("comparer is nil", func(t *testing.T) {
		randTypeName := fmt.Sprintf("%d-test-byte", rand.Int63())
		defer deferExpectPanicContains(t, "comparer is nil")
		RegisterDefaultComparerForType[byte](randTypeName, nil, false)
	})

	t.Run("register twice", func(t *testing.T) {
		randTypeName := fmt.Sprintf("%d-test-testStruct", rand.Int63())
		RegisterDefaultComparerForType[testStruct](randTypeName, testStructComparer, false) // first ok

		RegisterDefaultComparerForType[testStruct](randTypeName, testStructComparer, true) // second ok with override enabled

		defer deferExpectPanicContains(t, "had been registered before")
		RegisterDefaultComparerForType[testStruct](randTypeName, testStructComparer, false) // third fails without override
	})

	t.Run("register for pointer", func(t *testing.T) {
		randTypeName := fmt.Sprintf("*%d-test-testStruct", rand.Int63())

		defer deferExpectPanicContains(t, "can not register for pointer")
		RegisterDefaultComparerForType[testStruct](randTypeName, testStructComparer, false)
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
		c1, f1 := TryGetDefaultComparer[int16]()
		assert.NotNil(t, c1)
		assert.True(t, f1)

		c1, f1 = TryGetDefaultComparer[any]()
		assert.Nil(t, c1)
		assert.False(t, f1)

		c1 = GetDefaultComparer[int16]()
		assert.NotNil(t, c1)
	})

	t.Run("can get correctly from pointer", func(t *testing.T) {
		c1, f1 := TryGetDefaultComparer[*int]()
		assert.NotNil(t, c1)
		assert.True(t, f1)

		assert.Equal(t, 01, c1.Compare(3, 1))
		assert.Equal(t, 00, c1.Compare(3, 3))
		assert.Equal(t, -1, c1.Compare(1, 3))
	})

	t.Run("panic when not registered", func(t *testing.T) {
		type notExisting struct{}
		defer deferExpectPanicContains(t, fmt.Sprintf("no default comparer registered for type [*comparers.notExisting]"))
		_ = GetDefaultComparer[*notExisting]()
	})

	t.Run("panic when not registered", func(t *testing.T) {
		type notExisting struct{}
		defer deferExpectPanicContains(t, fmt.Sprintf("no default comparer registered for type [comparers.notExisting]"))
		_ = GetDefaultComparer[notExisting]()
	})

	t.Run("panic when can not detect type", func(t *testing.T) {
		defer deferExpectPanicContains(t, fmt.Sprintf("unable to detect type for provided type"))
		_ = GetDefaultComparer[any]()
	})
}

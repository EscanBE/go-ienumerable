package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func createNilEnumerable() IEnumerable[any] {
	return &enumerable[any]{
		data: nil,
	}
}

func createEmptyEnumerable() IEnumerable[any] {
	return &enumerable[any]{
		data: make([]any, 0),
	}
}

func createEmptyIntEnumerable() IEnumerable[int] {
	return &enumerable[int]{
		data: make([]int, 0),
	}
}

func createIntEnumerable(from, to int) IEnumerable[int] {
	if from > to {
		panic(fmt.Errorf("createIntEnumerable from %d > to %d", from, to))
	}
	data := make([]int, 0)
	for i := from; i <= to; i++ {
		data = append(data, i)
	}
	return injectIntComparers(NewIEnumerable[int](data...))
}

func createRandomIntEnumerable(size int) IEnumerable[int] {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return injectIntComparers(NewIEnumerable[int](data...))
}

func injectIntComparers(e IEnumerable[int]) IEnumerable[int] {
	return e.
		WithLessComparer(func(i1, i2 int) bool {
			return i1 < i2
		}).
		WithEqualsComparer(func(i1, i2 int) bool {
			return i1 == i2
		})
}

type copiedOriginal[T comparable] struct {
	isNil              bool
	data               []T
	dataType           string
	hasEqualsComparer  bool
	hasLessComparer    bool
	hasDefaultComparer bool
}

func backupForAssetUnchanged[T comparable](ie IEnumerable[T]) copiedOriginal[T] {
	if ie == nil {
		return copiedOriginal[T]{
			isNil: true,
		}
	}
	cast := e[T](ie)
	return copiedOriginal[T]{
		data:               copySlice(cast.data),
		dataType:           cast.dataType,
		hasEqualsComparer:  cast.equalityComparer != nil,
		hasLessComparer:    cast.lessComparer != nil,
		hasDefaultComparer: cast.defaultComparer != nil,
	}
}

func (c copiedOriginal[T]) assertUnchanged(t *testing.T, e IEnumerable[T]) {
	if c.isNil {
		if e != nil {
			t.Errorf("copied of original is nil but got asset with non-nil %v", e)
		}
		return
	}

	eData := e.ToArray()
	if len(c.data) != len(eData) {
		assert.Lenf(t, eData, len(c.data), "data of source IEnumerable has been changed, expect len %d but changed to %d", len(c.data), len(eData))
	} else if len(c.data) > 0 {
		for i1, d1 := range c.data {
			d2 := eData[i1]
			assert.Equalf(t, d1, d2, "data of source IEnumerable has been changed, expect element at [%d] = %d but changed to %d", i1, d1, d2)
		}
	}

	c.assertUnchangedIgnoreData(t, e)
}

func (c copiedOriginal[T]) assertUnchangedIgnoreData(t *testing.T, ie IEnumerable[T]) {
	cast := e[T](ie)

	assert.Equalf(t, c.dataType, cast.dataType, "dataType has changed, expect %s but got %s", c.dataType, cast.dataType)

	exists := func(b bool) string {
		if b {
			return "exists"
		} else {
			return "not-exist"
		}
	}

	assert.Equalf(t, c.hasEqualsComparer, cast.equalityComparer != nil, "equality comparer state has changed, expect %s, but got %s", exists(c.hasEqualsComparer), exists(cast.equalityComparer != nil))
	assert.Equalf(t, c.hasLessComparer, cast.lessComparer != nil, "less comparer state has changed, expect %s, but got %s", exists(c.hasLessComparer), exists(cast.lessComparer != nil))
	assert.Equalf(t, c.hasDefaultComparer, cast.defaultComparer != nil, "default comparer state has changed, expect %s, but got %s", exists(c.hasDefaultComparer), exists(cast.defaultComparer != nil))
}

func deferWantPanicDepends(t *testing.T, wantPanic bool) {
	err := recover()
	if err == nil && wantPanic {
		t.Errorf("expect panic")
	} else if err != nil && !wantPanic {
		t.Errorf("expect not panic but got %v", err)
	}
}

func deferExpectPanicContains(t *testing.T, msgPart string) {
	err := recover()
	if err == nil {
		t.Errorf("expect error")
		return
	}

	assert.Contains(t, fmt.Sprintf("%v", err), msgPart)
}

func Test_enumerable_copyExceptData(t *testing.T) {
	t.Run("copy all except data", func(t *testing.T) {
		e := new(enumerable[int])
		e.data = []int{2, 3}
		e.dataType = "int"
		e.equalityComparer = func(v1, v2 int) bool {
			return v1 == v2
		}
		e.lessComparer = func(v1, v2 int) bool {
			return v1 < v2
		}

		copied := e.copyExceptData()
		assert.Len(t, copied.data, 0)
		assert.Equal(t, "int", copied.dataType)
		assert.NotNil(t, copied.equalityComparer)
		assert.NotNil(t, copied.lessComparer)
	})

	t.Run("copy nil yields nil", func(t *testing.T) {
		e := new(enumerable[int])
		e = nil

		copied := e.copyExceptData()

		assert.Nil(t, copied)
	})
}

func Test_enumerable_withData(t *testing.T) {
	t.Run("data copied", func(t *testing.T) {
		e := new(enumerable[int])
		e.data = []int{2, 3}

		copied := e.copyExceptData().withData(e.data)
		assert.Len(t, copied.data, 2)
	})

	t.Run("copy nil yields nil", func(t *testing.T) {
		e := new(enumerable[int])
		e = nil

		copied := e.copyExceptData().withData([]int{})

		assert.Nil(t, copied)
	})
}

func Test_enumerable_withEmptyData(t *testing.T) {
	t.Run("data copied", func(t *testing.T) {
		e := new(enumerable[int])
		e.data = []int{2, 3}

		copied := e.copyExceptData().withEmptyData()
		assert.Len(t, copied.data, 0)
	})

	t.Run("copy nil yields nil", func(t *testing.T) {
		e := new(enumerable[int])
		e = nil

		copied := e.copyExceptData().withEmptyData()

		assert.Nil(t, copied)
	})
}

func Test_enumerable_assertSrcNonNil(t *testing.T) {
	e := new(enumerable[int])
	e = nil

	defer deferWantPanicDepends(t, true)
	e.assertSrcNonNil()
}

func Test_enumerable_assertSrcNonEmpty(t *testing.T) {
	e := new(enumerable[int])

	defer deferWantPanicDepends(t, true)
	e.assertSrcNonEmpty()
}

func Test_enumerable_assertPredicateNonNil(t *testing.T) {
	e := new(enumerable[int])

	defer deferWantPanicDepends(t, true)
	e.assertPredicateNonNil(nil)
}

func Test_enumerable_assertSelectorNonNil(t *testing.T) {
	t.Run("selector", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferWantPanicDepends(t, true)

		e.assertSelectorNonNil(nil)
	})

	t.Run("array selector", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferWantPanicDepends(t, true)

		e.assertArraySelectorNonNil(nil)
	})
}

func Test_enumerable_assertAggregateFuncNonNil(t *testing.T) {
	t.Run("aggregate func", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferWantPanicDepends(t, true)

		e.assertAggregateFuncNonNil(nil)
	})

	t.Run("aggregate any seed func", func(t *testing.T) {
		e := new(enumerable[int])

		defer deferWantPanicDepends(t, true)

		e.assertAggregateAnySeedFuncNonNil(nil)
	})
}

func Test_enumerable_findDefaultComparer(t *testing.T) {
	t.Run("resolve registered", func(t *testing.T) {
		assert.NotNil(t, e[int32](NewIEnumerable[int32]()).findDefaultComparer())
	})

	t.Run("panic if type not registered", func(t *testing.T) {
		type MyInt64 int64
		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]")

		e[MyInt64](NewIEnumerable[MyInt64]()).findDefaultComparer()
	})
}

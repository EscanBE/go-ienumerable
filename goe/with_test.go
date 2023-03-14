package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_With(t *testing.T) {
	eSrc := NewIEnumerable[int](1, 2, 3)

	back := e[int](eSrc)
	assert.Nil(t, back.equalityComparer)
	assert.Nil(t, back.lessComparer)

	test_enumerable_With_addWiths(eSrc)

	assert.NotNil(t, back.equalityComparer)
	assert.NotNil(t, back.lessComparer)
}

//goland:noinspection GoSnakeCaseUsage
func test_enumerable_With_addWiths(e IEnumerable[int]) {
	e.WithEqualsComparer(func(i1, i2 int) bool {
		return i1 == i2
	}).WithLessComparer(func(i1, i2 int) bool {
		return i1 < i2
	})
}

func Test_enumerable_WithDefaultComparers(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int8](2, 1).WithDefaultComparers()

		back := e[int8](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int8](2))
	})

	t.Run("byte", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[byte](2, 1).WithDefaultComparers()

		back := e[byte](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[byte](2))
	})

	t.Run("uint8", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint8](2, 1).WithDefaultComparers()

		back := e[uint8](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint8](2))
	})

	t.Run("int16", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int16](2, 1).WithDefaultComparers()

		back := e[int16](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int16](2))
	})

	t.Run("uint16", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint16](2, 1).WithDefaultComparers()

		back := e[uint16](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint16](2))
	})

	t.Run("int32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int32](2, 1).WithDefaultComparers()

		back := e[int32](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int32](2))
	})

	t.Run("rune", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[rune](2, 1).WithDefaultComparers()

		back := e[rune](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[rune](2))
	})

	t.Run("uint32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint32](2, 1).WithDefaultComparers()

		back := e[uint32](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint32](2))
	})

	t.Run("int64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int64](2, 1).WithDefaultComparers()

		back := e[int64](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int64](2))
	})

	t.Run("uint64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint64](2, 1).WithDefaultComparers()

		back := e[uint64](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint64](2))
	})

	t.Run("int", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int](2, 1).WithDefaultComparers()

		back := e[int](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int](2))
	})

	t.Run("uint", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint](2, 1).WithDefaultComparers()

		back := e[uint](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint](2))
	})

	t.Run("uintptr", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uintptr](2, 1).WithDefaultComparers()

		back := e[uintptr](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uintptr](2))
	})

	t.Run("float32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[float32](2.0, 1.0).WithDefaultComparers()

		back := e[float32](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[float32](2.0))
	})

	t.Run("float64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[float64](2.0, 1.0).WithDefaultComparers()

		back := e[float64](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[float64](2.0))
	})

	t.Run("complex64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[complex64](2.0, 1.0).WithDefaultComparers()

		back := e[complex64](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.Nil(t, back.lessComparer)

		eSrc.Except(NewIEnumerable[complex64](2.0))
	})

	t.Run("complex128", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[complex128](2.0, 1.0).WithDefaultComparers()

		back := e[complex128](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.Nil(t, back.lessComparer)

		eSrc.Except(NewIEnumerable[complex128](2.0))
	})

	t.Run("string", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[string]("2", "1").WithDefaultComparers()

		back := e[string](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[string]("2"))
	})

	t.Run("bool", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[bool](true, false, true, false).WithDefaultComparers()

		back := e[bool](eSrc)

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		got := eSrc.Order().ToArray()
		assert.False(t, got[0])
		assert.False(t, got[1])
		assert.True(t, got[2])
		assert.True(t, got[3])

		got = eSrc.Except(NewIEnumerable[bool](false)).ToArray()
		assert.Len(t, got, 2)
		assert.True(t, got[0])
		assert.True(t, got[1])
	})

	t.Run("not supported", func(t *testing.T) {
		type x struct {
		}

		defer deferWantPanicDepends(t, true)

		_ = NewIEnumerable[x](x{}, x{}).WithDefaultComparers()
	})
}

func Test_enumerable_WithComparersFrom(t *testing.T) {
	t.Run("equality comparer", func(t *testing.T) {
		equalityComparer := func(v1, v2 int) bool {
			return v1 == v2
		}
		eDst := NewIEnumerable[int](1, 2, 3).WithEqualsComparer(equalityComparer)
		eSrc := NewIEnumerable[int]()

		eD := e[int](eDst)
		eS := e[int](eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.Nil(t, eS.equalityComparer)

		assert.True(t, eDst.Contains(2))

		//

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.Nil(t, eS.equalityComparer)

		assert.True(t, eDst.Contains(2))

		//

		_ = eSrc.WithEqualsComparer(func(v1, v2 int) bool {
			return false
		})

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.NotNil(t, eS.equalityComparer)

		assert.False(t, eDst.Contains(2))

		//

		eD.equalityComparer = nil
		assert.Nil(t, eD.equalityComparer)

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.NotNil(t, eS.equalityComparer)

		assert.False(t, eDst.Contains(2))

		//

		eD.equalityComparer = equalityComparer
		assert.True(t, eDst.Contains(2))

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.NotNil(t, eS.equalityComparer)

		assert.False(t, eDst.Contains(2)) // override
	})

	t.Run("less comparer", func(t *testing.T) {
		lessComparer := func(v1, v2 int) bool {
			return v1 < v2
		}
		eDst := NewIEnumerable[int](1, 2, 3).WithLessComparer(lessComparer)
		eSrc := NewIEnumerable[int]()

		eD := e[int](eDst)
		eS := e[int](eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.Nil(t, eS.lessComparer)

		assert.Equal(t, 1, eDst.Min())

		//

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.Nil(t, eS.lessComparer)

		assert.Equal(t, 1, eDst.Min())

		//

		_ = eSrc.WithLessComparer(func(v1, v2 int) bool {
			return v1 > v2
		})

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min())

		//

		eD.lessComparer = nil
		assert.Nil(t, eD.lessComparer)

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min())

		//

		eD.lessComparer = lessComparer
		assert.Equal(t, 1, eDst.Min())

		_ = eDst.WithComparerFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min()) // override
	})
}

func Test_enumerable_WithDefaultComparer(t *testing.T) {
	t.Run("inject and remove default comparer", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(5)
		eSrc.WithDefaultComparer(nil)

		e := e[int](eSrc)
		assert.Nil(t, e.defaultComparer)

		// replace
		eSrc.WithDefaultComparer(comparers.IntComparer)
		assert.NotNil(t, e.defaultComparer)

		// eraser if input nil
		eSrc.WithDefaultComparer(nil)
		assert.Nil(t, e.defaultComparer)
	})
}

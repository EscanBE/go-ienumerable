package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_With(t *testing.T) {
	eSrc := NewIEnumerable[int](1, 2, 3)

	back := eSrc.(*enumerable[int])
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
	/*
		t.Run("performance", func(t *testing.T) {
			defer deferWantPanicDepends(t, false)

			eSrc1 := createRandomIntEnumerable(2_000_000).WithDefaultComparers()
			eSrc2 := NewIEnumerable[int](eSrc1.exposeData()...).
				WithEqualsComparer(func(v1, v2 int) bool {
					return v1 == v2
				}).
				WithLessComparer(func(v1, v2 int) bool {
					return v1 < v2
				})

			now := func() int64 {
				return time.Now().UnixNano()
			}

			var sum1, sum2 int64
			var cnt1, cnt2 int

			for turn := 1; turn <= 10; turn++ {
				startSort1 := now()
				eSrc1.Order()
				stopSort1 := now()
				sum1 += stopSort1 - startSort1

				startSort2 := now()
				eSrc2.Order()
				stopSort2 := now()
				sum2 += stopSort2 - startSort2

				cnt1++
				cnt2++
			}

			avg1 := int64(math.Floor(float64(sum1) / float64(cnt1)))
			avg2 := int64(math.Floor(float64(sum2) / float64(cnt2)))
			fmt.Printf("Sort 1 avg: %d\n", avg1)
			fmt.Printf("Sort 2 avg: %d\n", avg2)
		})
	*/

	t.Run("int8", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int8](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[int8])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int8](2))
	})

	t.Run("byte", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[byte](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[byte])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[byte](2))
	})

	t.Run("uint8", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint8](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uint8])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint8](2))
	})

	t.Run("int16", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int16](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[int16])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int16](2))
	})

	t.Run("uint16", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint16](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uint16])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint16](2))
	})

	t.Run("int32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int32](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[int32])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int32](2))
	})

	t.Run("rune", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[rune](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[rune])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[rune](2))
	})

	t.Run("uint32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint32](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uint32])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint32](2))
	})

	t.Run("int64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int64](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[int64])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int64](2))
	})

	t.Run("uint64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint64](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uint64])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint64](2))
	})

	t.Run("int", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[int](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[int])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[int](2))
	})

	t.Run("uint", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uint](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uint])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uint](2))
	})

	t.Run("uintptr", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[uintptr](2, 1).WithDefaultComparers()

		back := eSrc.(*enumerable[uintptr])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[uintptr](2))
	})

	t.Run("float32", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[float32](2.0, 1.0).WithDefaultComparers()

		back := eSrc.(*enumerable[float32])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[float32](2.0))
	})

	t.Run("float64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[float64](2.0, 1.0).WithDefaultComparers()

		back := eSrc.(*enumerable[float64])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[float64](2.0))
	})

	t.Run("complex64", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[complex64](2.0, 1.0).WithDefaultComparers()

		back := eSrc.(*enumerable[complex64])

		assert.NotNil(t, back.equalityComparer)
		assert.Nil(t, back.lessComparer)

		eSrc.Except(NewIEnumerable[complex64](2.0))
	})

	t.Run("complex128", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[complex128](2.0, 1.0).WithDefaultComparers()

		back := eSrc.(*enumerable[complex128])

		assert.NotNil(t, back.equalityComparer)
		assert.Nil(t, back.lessComparer)

		eSrc.Except(NewIEnumerable[complex128](2.0))
	})

	t.Run("string", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[string]("2", "1").WithDefaultComparers()

		back := eSrc.(*enumerable[string])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		eSrc.Order()
		eSrc.Except(NewIEnumerable[string]("2"))
	})

	t.Run("bool", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		eSrc := NewIEnumerable[bool](true, false, true, false).WithDefaultComparers()

		back := eSrc.(*enumerable[bool])

		assert.NotNil(t, back.equalityComparer)
		assert.NotNil(t, back.lessComparer)

		got := eSrc.Order().exposeData()
		assert.False(t, got[0])
		assert.False(t, got[1])
		assert.True(t, got[2])
		assert.True(t, got[3])

		got = eSrc.Except(NewIEnumerable[bool](false)).exposeData()
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

		eD := eDst.(*enumerable[int])
		eS := eSrc.(*enumerable[int])

		assert.NotNil(t, eD.equalityComparer)
		assert.Nil(t, eS.equalityComparer)

		assert.True(t, eDst.Contains(2))

		//

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.Nil(t, eS.equalityComparer)

		assert.True(t, eDst.Contains(2))

		//

		_ = eSrc.WithEqualsComparer(func(v1, v2 int) bool {
			return false
		})

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.NotNil(t, eS.equalityComparer)

		assert.False(t, eDst.Contains(2))

		//

		eD.equalityComparer = nil
		assert.Nil(t, eD.equalityComparer)

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.equalityComparer)
		assert.NotNil(t, eS.equalityComparer)

		assert.False(t, eDst.Contains(2))

		//

		eD.equalityComparer = equalityComparer
		assert.True(t, eDst.Contains(2))

		_ = eDst.WithComparersFrom(eSrc)

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

		eD := eDst.(*enumerable[int])
		eS := eSrc.(*enumerable[int])

		assert.NotNil(t, eD.lessComparer)
		assert.Nil(t, eS.lessComparer)

		assert.Equal(t, 1, eDst.Min())

		//

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.Nil(t, eS.lessComparer)

		assert.Equal(t, 1, eDst.Min())

		//

		_ = eSrc.WithLessComparer(func(v1, v2 int) bool {
			return v1 > v2
		})

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min())

		//

		eD.lessComparer = nil
		assert.Nil(t, eD.lessComparer)

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min())

		//

		eD.lessComparer = lessComparer
		assert.Equal(t, 1, eDst.Min())

		_ = eDst.WithComparersFrom(eSrc)

		assert.NotNil(t, eD.lessComparer)
		assert.NotNil(t, eS.lessComparer)

		assert.Equal(t, 3, eDst.Min()) // override
	})
}

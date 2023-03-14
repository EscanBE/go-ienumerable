package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_SelectMany(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{2, 3}, []int8{4, 5})

		eGot := eSrc.SelectMany(func(a []int8) []any {
			return []any{a[0] * 2, a[1] * 2}
		})
		gotData := eGot.exposeData()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "", gotE.dataType)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{2, 3}, []int8{4, 5})

		eGot := eSrc.SelectMany(func(a []int8) []any {
			return []any{fmt.Sprintf("%d", a[0]+1), fmt.Sprintf("%d", a[1]+1)}
		})
		gotData := eGot.exposeData()
		assert.Len(t, gotData, 4)
		assert.Equal(t, "3", gotData[0])
		assert.Equal(t, "4", gotData[1])
		assert.Equal(t, "5", gotData[2])
		assert.Equal(t, "6", gotData[3])
		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "string", fmt.Sprintf("%T", gotE.data[0]))
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]()

		eGot := eSrc.SelectMany(func(a []int8) []any {
			return []any{int64(a[0]), int64(a[1])}
		})

		gotData := eGot.exposeData()
		assert.Len(t, gotData, 0)

		gotE := e[any](eGot)
		assert.Nil(t, gotE.equalityComparer)
		assert.Nil(t, gotE.lessComparer)
		assert.Equal(t, "", gotE.dataType)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]()

		defer deferWantPanicDepends(t, true)

		_ = eSrc.SelectMany(nil)
	})
}

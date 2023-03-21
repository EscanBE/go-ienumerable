package helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSelectMany(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]([]int8{2, 3}, []int8{6, 7}, []int8{4, 5})

		eGot := SelectMany(eSrc, func(a []int8) []any {
			return []any{a[0] * 2, a[1] * 2}
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 6)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(12), gotData[2])
		assert.Equal(t, int8(14), gotData[3])
		assert.Equal(t, int8(8), gotData[4])
		assert.Equal(t, int8(10), gotData[5])
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]()

		eGot := SelectMany(eSrc, func(i []int8) []any {
			return []any{int64(i[0])}
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)
	})

	t.Run("skip empty result from selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]([]int8{1, 2}, []int8{0, 2}, []int8{1, 0}, []int8{0, 0})

		eGot := SelectMany(eSrc, func(i []int8) []any {
			result := make([]any, 0)
			for _, iv := range i {
				if iv != 0 {
					result = append(result, iv)
				}
			}
			return result
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8]()

		defer deferExpectPanicContains(t, "result selector function is nil", true)

		_ = SelectMany[int8, int](eSrc, nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[[]int]([]int{3, 1})

		ieGot := SelectMany(ieSrc, func(i []int) []any {
			return []any{time.Duration(i[0]) * time.Minute, time.Duration(i[1]) * time.Minute}
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		gotArray = ieGot.Order().GetOrderedEnumerable().ToArray()

		assert.Equal(t, 1*time.Minute, gotArray[0])
		assert.Equal(t, 3*time.Minute, gotArray[1])
	})

	t.Run("panic nil value as result of selector", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[[]int]([]int{9, 3})

		defer deferExpectPanicContains(t, "result array can not be nil", true)

		_ = SelectMany(ieSrc, func(i []int) []any {
			return nil
		})
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := goe.NewIEnumerable[[]int]([]int{3, 1, 2, 6}, []int{})

		ieGot := SelectMany(ieSrc, func(i []int) []any {
			result := make([]any, len(i))
			for idx, iv := range i {
				if iv == 2 {
					result[idx] = nil
				} else {
					result[idx] = &MyInt{
						Value: iv,
					}
				}
			}
			return result
		})

		gotArray := ieGot.ToArray()

		assert.Len(t, gotArray, 4)
		assert.Equal(t, 3, gotArray[0].(*MyInt).Value)
		assert.Equal(t, 1, gotArray[1].(*MyInt).Value)
		assert.Nil(t, gotArray[2])
		assert.Equal(t, 6, gotArray[3].(*MyInt).Value)
	})
}

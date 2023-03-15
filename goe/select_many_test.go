package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_enumerable_SelectMany(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{2, 3}, []int8{6, 7}, []int8{4, 5})

		eGot := eSrc.SelectMany(func(a []int8) []any {
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
		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "int8", gotE.dataType)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]()

		eGot := eSrc.SelectMany(func(i []int8) []any {
			return []any{int64(i[0])}
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)

		gotE := e[any](eGot)
		// because no result to detect default comparer
		assert.Nil(t, gotE.defaultComparer)
		assert.Equal(t, "", gotE.dataType)
	})

	t.Run("skip empty result from selector", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{1, 2}, []int8{0, 2}, []int8{1, 0}, []int8{0, 0})

		eGot := eSrc.SelectMany(func(i []int8) []any {
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

		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", gotE.dataType)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()

		defer deferWantPanicDepends(t, true)

		_ = eSrc.SelectMany(nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{3, 1})

		ieGot := ieSrc.SelectMany(func(i []int) []any {
			return []any{time.Duration(i[0]) * time.Minute, time.Duration(i[1]) * time.Minute}
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		eGot := e[any](ieGot)
		assert.Equal(t, "time.Duration", eGot.dataType)
		assert.NotNil(t, eGot.defaultComparer)
		assert.Equal(t, 1, eGot.defaultComparer.Compare(gotArray[0], gotArray[1]))
		assert.Equal(t, -1, eGot.defaultComparer.Compare(gotArray[1], gotArray[0]))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[0], 3*time.Minute))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[1], 1*time.Minute))
	})

	t.Run("panic nil value as result of selector", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{9, 3})

		defer deferExpectPanicContains(t, "result array can not be nil", true)

		_ = ieSrc.SelectMany(func(i []int) []any {
			return nil
		})
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := NewIEnumerable[[]int]([]int{3, 1}, []int{})

		ieGot := ieSrc.SelectMany(func(i []int) []any {
			result := make([]any, len(i))
			for idx, iv := range i {
				result[idx] = MyInt{
					Value: iv,
				}
			}
			return result
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3, gotArray[0].(MyInt).Value)
		assert.Equal(t, 1, gotArray[1].(MyInt).Value)

		eGot := e[any](ieGot)
		assert.Equal(t, "goe.MyInt", eGot.dataType)
		assert.Nil(t, eGot.defaultComparer)
	})
}

func Test_enumerable_SelectManyWithSampleValueOfResult(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{2, 3, 4, 5})

		eGot := eSrc.SelectManyWithSampleValueOfResult(func(i []int8) []any {
			result := make([]any, len(i))
			for idx, iv := range i {
				result[idx] = iv * 2
			}
			return result
		}, int8(0))
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "int8", gotE.dataType)
	})

	t.Run("string", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{2, 3, 4, 5})

		eGot := eSrc.SelectManyWithSampleValueOfResult(func(i []int8) []any {
			result := make([]any, len(i))
			for idx, iv := range i {
				result[idx] = fmt.Sprintf("%d", iv+1)
			}
			return result
		}, "")
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, "3", gotData[0])
		assert.Equal(t, "4", gotData[1])
		assert.Equal(t, "5", gotData[2])
		assert.Equal(t, "6", gotData[3])
		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "string", fmt.Sprintf("%T", gotE.data[0]))
		assert.Equal(t, "string", gotE.dataType)
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]()

		eGot := eSrc.SelectManyWithSampleValueOfResult(func(i []int8) []any {
			return []any{int64(i[0]), int64(i[1])}
		}, int64(0))

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)

		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int64", gotE.dataType)
	})

	t.Run("skip empty result from selector", func(t *testing.T) {
		eSrc := NewIEnumerable[[]int8]([]int8{1, 2}, []int8{0, 2}, []int8{1, 0}, []int8{0, 0})

		eGot := eSrc.SelectManyWithSampleValueOfResult(func(i []int8) []any {
			result := make([]any, 0)
			for _, iv := range i {
				if iv != 0 {
					result = append(result, iv)
				}
			}
			return result
		}, int8(0))

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)

		gotE := e[any](eGot)
		assert.NotNil(t, gotE.defaultComparer)
		assert.Equal(t, "int8", gotE.dataType)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()

		defer deferExpectPanicContains(t, getErrorNilSelector().Error(), true)

		_ = eSrc.SelectManyWithSampleValueOfResult(nil, 0)
	})

	t.Run("sample result value can not be nil", func(t *testing.T) {
		eSrc := NewIEnumerable[int8]()

		defer deferExpectPanicContains(t, getErrorSampleValueIsNil().Error(), true)

		_ = eSrc.SelectManyWithSampleValueOfResult(func(i int8) []any {
			return []any{i}
		}, nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{3, 1})

		ieGot := ieSrc.SelectManyWithSampleValueOfResult(func(i []int) []any {
			return []any{time.Duration(i[0]) * time.Minute, time.Duration(i[1]) * time.Minute}
		}, time.Duration(0))

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		eGot := e[any](ieGot)
		assert.Equal(t, "time.Duration", eGot.dataType)
		assert.NotNil(t, eGot.defaultComparer)
		assert.Equal(t, 1, eGot.defaultComparer.Compare(gotArray[0], gotArray[1]))
		assert.Equal(t, -1, eGot.defaultComparer.Compare(gotArray[1], gotArray[0]))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[0], 3*time.Minute))
		assert.Zero(t, eGot.defaultComparer.Compare(gotArray[1], 1*time.Minute))
	})

	t.Run("panic nil value as result of selector", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{9, 3})

		defer deferExpectPanicContains(t, "result array can not be nil", true)

		_ = ieSrc.SelectManyWithSampleValueOfResult(func(i []int) []any {
			return nil
		}, time.Duration(0))
	})

	t.Run("nil value as element of result of selector is ok", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{9, 3})

		ieGot := ieSrc.SelectManyWithSampleValueOfResult(func(i []int) []any {
			return []any{nil, nil}
		}, time.Duration(0))

		gotArray := ieGot.ToArray()

		assert.Len(t, gotArray, 2)
		assert.Nil(t, gotArray[0])
		assert.Nil(t, gotArray[1])

		eGot := e[any](ieGot)
		assert.Equal(t, "time.Duration", eGot.dataType)
		assert.NotNil(t, eGot.defaultComparer)
		assert.Equal(t, 1, eGot.defaultComparer.Compare(time.Minute, time.Second))
		assert.Equal(t, -1, eGot.defaultComparer.Compare(time.Second, time.Minute))
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := NewIEnumerable[[]int]([]int{3, 1})

		ieGot := ieSrc.SelectManyWithSampleValueOfResult(func(i []int) []any {
			result := make([]any, len(i))
			for idx, iv := range i {
				result[idx] = MyInt{
					Value: iv,
				}
			}
			return result
		}, MyInt{})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3, gotArray[0].(MyInt).Value)
		assert.Equal(t, 1, gotArray[1].(MyInt).Value)

		eGot := e[any](ieGot)
		assert.Equal(t, "goe.MyInt", eGot.dataType)
		assert.Nil(t, eGot.defaultComparer)
	})

	t.Run("panic if result type not match type of sample result value", func(t *testing.T) {
		ieSrc := NewIEnumerable[[]int]([]int{3, 1})

		defer deferExpectPanicContains(t, "sample result at index 0 yields by element at index 0 is type [int] but got result 3 of type [int64]", true)

		_ = ieSrc.SelectManyWithSampleValueOfResult(func(i []int) []any {
			return []any{int64(i[0]), int64(i[1])}
		}, 0)
	})
}

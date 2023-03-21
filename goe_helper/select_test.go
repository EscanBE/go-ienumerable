package goe_helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8](2, 3, 4, 5)

		eGot := Select(eSrc, func(i int8) int8 {
			return i * 2
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(8), gotData[2])
		assert.Equal(t, int8(10), gotData[3])
	})

	t.Run("string", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8](2, 3, 4, 5)

		eGot := Select(eSrc, func(i int8) string {
			return fmt.Sprintf("%d", i+1)
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
		assert.Equal(t, "3", gotData[0])
		assert.Equal(t, "4", gotData[1])
		assert.Equal(t, "5", gotData[2])
		assert.Equal(t, "6", gotData[3])
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8]()

		eGot := Select(eSrc, func(i int8) int64 {
			return int64(i)
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8]()

		defer deferExpectPanicContains(t, "result selector function is nil", true)

		_ = Select[int8, string](eSrc, nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[int](3, 1)

		ieGot := Select(ieSrc, func(i int) time.Duration {
			return time.Duration(i) * time.Minute
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		gotArray = ieGot.Order().GetOrderedEnumerable().ToArray()

		assert.Equal(t, 1*time.Minute, gotArray[0])
		assert.Equal(t, 3*time.Minute, gotArray[1])
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := goe.NewIEnumerable[int](3, 1, 2, 6)

		ieGot := Select(ieSrc, func(i int) *MyInt {
			if i == 2 {
				return nil
			}
			return &MyInt{
				Value: i,
			}
		})

		gotArray := ieGot.ToArray()

		assert.Len(t, gotArray, 4)
		assert.Equal(t, 3, gotArray[0].Value)
		assert.Equal(t, 1, gotArray[1].Value)
		assert.Nil(t, gotArray[2])
		assert.Equal(t, 6, gotArray[3].Value)
	})
}

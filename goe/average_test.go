package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_enumerable_Average(t *testing.T) {
	t.Run("accept any int/float", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc1 := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			float32(11.0), float64(12.0),
			int8(-1), int16(-1), int32(-1), int64(-1), int(-1),
			float32(-1.1), float64(-1.1),
		)

		avg1 := eSrc1.Average()
		assert.Greater(t, 0.001, math.Abs(avg1-3.726))

		//goland:noinspection GoRedundantConversion
		eSrc2 := NewIEnumerable[float64](
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64), float64(math.MaxInt64),
		)

		avg2 := eSrc2.Average()
		assert.Greater(t, 0.001, math.Abs(avg2-math.MaxInt64))

		fmt.Printf("Avg 1: %f | Avg 2: %f\n", avg1, avg2)
	})

	t.Run("cover sum zero", func(t *testing.T) {
		//goland:noinspection ALL
		eSrc := NewIEnumerable[any](
			int8(0), uint8(0),
			int16(0), uint16(0),
			int32(0), uint32(0),
			int64(0), uint64(0),
			int(0), uint(0),
			float32(0), float64(0),
		)

		avg := eSrc.Average()
		assert.Zero(t, avg)
	})
}

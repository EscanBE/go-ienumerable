package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
	"time"
)

func Test_enumerable_SumInt32(t *testing.T) {
	t.Run("accept any int with value in range int32", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
		)
		assert.Equal(t, int32(55), eSrc.SumInt32())
	})

	t.Run("sum correct", func(t *testing.T) {
		halfMax := (math.MaxInt32 - 1) / 2
		max := math.MaxInt32
		min := math.MinInt32
		eSrc := NewIEnumerable[any](
			int32(max), int32(max), 2, int32(min), int32(min), int32(halfMax), int32(halfMax),
		)
		assert.Equal(t, int32(max-1), eSrc.SumInt32())
	})

	t.Run("empty of any integer type always returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[int8]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint8]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[int16]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint16]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[int32]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint32]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[int64]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint64]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[int]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*int8]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint8]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*int16]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint16]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*int32]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint32]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*int64]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint64]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*int]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint]().SumInt32()))
	})

	t.Run("empty of whatever type returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[any]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[string]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[time.Location]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*string]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]int]().SumInt32()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]string]().SumInt32()))
	})

	t.Run("panic if result is overflow int32", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[int64](int64(math.MaxInt32), int64(math.MaxInt32), int64(math.MinInt32), int64(math.MinInt32), int64(math.MaxInt32), 3).SumInt32()
		fmt.Printf("Result: %d", result)
	})

	t.Run("panic when sum overflow int32 (int)", func(t *testing.T) {
		if math.MaxInt == math.MaxInt32 { // skip x86
			return
		}
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 9223372036854775807 of type int is over range of int32") {
				// ok (x64)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](int(math.MaxInt))
		eSrc.SumInt32()
	})

	t.Run("panic when sum overflow int32 (int64)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 9223372036854775807 of type int64 is over range of int32") {
				// ok (x64)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](int64(math.MaxInt64))
		eSrc.SumInt32()
	})

	t.Run("panic when sum overflow int32 (uint32)", func(t *testing.T) {
		u := uint32(math.MaxUint32)
		assert.Equal(t, float64(math.MaxUint32), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "value 4294967295 of type uint32 is over range of int32")
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt32()
	})

	t.Run("panic when sum overflow int32 (uint64)", func(t *testing.T) {
		u := uint64(math.MaxUint64)
		assert.Equal(t, float64(math.MaxUint64), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "value 18446744073709551615 of type uint64 is over range of int32")
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt32()
	})

	t.Run("panic when sum overflow int32 (uint)", func(t *testing.T) {
		u := uint(math.MaxUint)
		assert.Equal(t, float64(math.MaxUint), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 18446744073709551615 of type uint is over range of int32") {
				// ok (x64)
			} else if strings.Contains(errStr, "value 4294967295 of type uint is over range of int32") {
				// ok (x86)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt32()
	})

	t.Run("panic when element is not integer", func(t *testing.T) {
		str := "Hello World!"
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int32", str))
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			string(str),
		)
		eSrc.SumInt32()
	})
}

func Test_enumerable_SumInt(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	t.Run("accept any int with value in range int", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
		)
		assert.Equal(t, 55, eSrc.SumInt())
	})

	t.Run("sum correct", func(t *testing.T) {
		halfMax := (math.MaxInt - 1) / 2
		max := math.MaxInt
		min := math.MinInt
		eSrc := NewIEnumerable[any](
			max, max, 2, min, min, halfMax, halfMax,
		)
		assert.Equal(t, max-1, eSrc.SumInt())
	})

	t.Run("empty of any integer type always returns 0", func(t *testing.T) {
		assert.Equal(t, 0, NewIEnumerable[int8]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[uint8]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[int16]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[uint16]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[int32]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[uint32]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[int64]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[uint64]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[int]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[uint]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*int8]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*uint8]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*int16]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*uint16]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*int32]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*uint32]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*int64]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*uint64]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*int]().SumInt())
		assert.Equal(t, 0, NewIEnumerable[*uint]().SumInt())
	})

	//goland:noinspection GoRedundantConversion
	t.Run("empty of whatever type returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[any]().SumInt()))
		assert.Equal(t, 0, int(NewIEnumerable[string]().SumInt()))
		assert.Equal(t, 0, int(NewIEnumerable[time.Location]().SumInt()))
		assert.Equal(t, 0, int(NewIEnumerable[*string]().SumInt()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]int]().SumInt()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]string]().SumInt()))
	})

	t.Run("panic if result is overflow int", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[int64](int64(math.MaxInt), int64(math.MaxInt), int64(math.MinInt), int64(math.MinInt), int64(math.MaxInt), 3).SumInt()
		fmt.Printf("Result: %d", result)
	})

	t.Run("panic when sum overflow int (int64)", func(t *testing.T) {
		if x64 {
			return
		}
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 9223372036854775807 of type int64 is over range of int") {
				// ok (x64)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](int64(math.MaxInt64))
		eSrc.SumInt()
	})

	t.Run("panic when sum overflow int (uint32)", func(t *testing.T) {
		if x64 {
			return
		}

		u := uint32(math.MaxUint32)
		assert.Equal(t, float64(math.MaxUint32), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "value 4294967295 of type uint32 is over range of int")
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt()
	})

	t.Run("panic when sum overflow int (uint64)", func(t *testing.T) {
		u := uint64(math.MaxUint64)
		assert.Equal(t, float64(math.MaxUint64), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "value 18446744073709551615 of type uint64 is over range of int")
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt()
	})

	t.Run("panic when sum overflow int (uint)", func(t *testing.T) {
		u := uint(math.MaxUint)
		assert.Equal(t, float64(math.MaxUint), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 18446744073709551615 of type uint is over range of int") {
				// ok (x64)
			} else if strings.Contains(errStr, "value 4294967295 of type uint is over range of int") {
				// ok (x86)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt()
	})

	t.Run("panic when element is not integer", func(t *testing.T) {
		str := "Hello World!"
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int", str))
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			string(str),
		)
		eSrc.SumInt()
	})
}

func Test_enumerable_SumInt64(t *testing.T) {
	t.Run("accept any int with value in range int64", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			int8(-1), int16(-1), int32(-1), int64(-1), int(-1),
		)
		start := time.Now().Nanosecond()
		sum := eSrc.SumInt64()
		fmt.Printf("Consumed: %d ns\n", time.Now().Nanosecond()-start)
		assert.Equal(t, int64(50), sum)
	})

	t.Run("sum correct", func(t *testing.T) {
		halfMax := (math.MaxInt64 - 1) / 2
		max := math.MaxInt64
		min := math.MinInt64
		eSrc := NewIEnumerable[any](
			int64(max), int64(max), 2, int64(min), int64(min), int64(halfMax), int64(halfMax),
		)
		start := time.Now().Nanosecond()
		sum := eSrc.SumInt64()
		fmt.Printf("Consumed: %d ns\n", time.Now().Nanosecond()-start)
		assert.Equal(t, int64(max-1), sum)
	})

	t.Run("empty of any integer type always returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[int8]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint8]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[int16]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint16]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[int32]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint32]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[int64]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint64]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[int]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int8]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint8]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int16]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint16]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int32]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint32]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int64]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint64]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint]().SumInt64()))
	})

	t.Run("empty of whatever type returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[any]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[string]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[time.Location]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*string]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]int]().SumInt64()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]string]().SumInt64()))
	})

	t.Run("empty of interface{} (aka any) returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[any]().SumInt64()))
	})

	t.Run("panic if result is overflow int64 (positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[int64](int64(math.MaxInt64), int64(math.MaxInt64), int64(math.MinInt64), int64(math.MinInt64), int64(math.MaxInt64), 3).SumInt64()
		fmt.Printf("Result: %d", result)
	})

	t.Run("panic if result is overflow int64 (negative)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[int64](int64(math.MaxInt64), int64(math.MaxInt64), int64(math.MinInt64), int64(math.MinInt64), int64(math.MinInt64), -3).SumInt64()
		fmt.Printf("Result: %d", result)
	})

	t.Run("panic when sum overflow int64 (uint64)", func(t *testing.T) {
		u := uint64(math.MaxUint64)
		assert.Equal(t, float64(math.MaxUint64), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "value 18446744073709551615 of type uint64 is over range of int64")
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt64()
	})

	t.Run("panic when sum overflow int64 (uint)", func(t *testing.T) {
		if math.MaxUint == math.MaxUint32 { // skip x86
			return
		}
		u := uint(math.MaxUint)
		assert.Equal(t, float64(math.MaxUint), float64(u))
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			errStr := fmt.Sprintf("%v", err)
			if strings.Contains(errStr, "value 18446744073709551615 of type uint is over range of int64") {
				// ok (x64)
			} else {
				t.Errorf("un-wanted error: %s", errStr)
			}
		}()
		eSrc := NewIEnumerable[any](u)
		eSrc.SumInt64()
	})

	t.Run("panic when element is not integer", func(t *testing.T) {
		str := "Hello World!"
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int64", str))
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			string(str),
		)
		eSrc.SumInt64()
	})
}

func Test_enumerable_SumFloat64(t *testing.T) {
	t.Run("accept any int/float", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			float32(11.0), float64(12.0),
			int8(-1), int16(-1), int32(-1), int64(-1), int(-1),
			float32(-1.1), float64(-1.1),
		)

		sum := eSrc.SumFloat64()
		//goland:noinspection GoRedundantConversion
		assert.Greater(t, 0.001, math.Abs(sum-float64(70.8)))
		//goland:noinspection GoRedundantConversion
		assert.Less(t, float64(70.7), sum)
		//goland:noinspection GoRedundantConversion
		assert.Greater(t, float64(70.9), sum)
		assert.Less(t, float64(70), sum)
		assert.Greater(t, float64(71), sum)
	})

	t.Run("no overflow if int64 elements sum over range", func(t *testing.T) {
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int64(math.MaxInt64), int64(math.MaxInt64),
			float64(math.MaxInt64), float64(math.MaxInt64),
		)

		sum := eSrc.SumFloat64()
		//goland:noinspection GoRedundantConversion
		assert.Greater(t, 0.001, math.Abs(sum-float64(math.MaxInt64)*4))
		//goland:noinspection GoRedundantConversion
	})

	t.Run("empty of any integer/float type always returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[int8]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint8]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[int16]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint16]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[int32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[int64]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint64]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[int]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[uint]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[float32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[float64]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int8]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint8]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int16]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint16]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int64]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint64]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*int]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*uint]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*float32]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*float64]().SumFloat64()))
	})

	t.Run("empty of whatever type returns 0", func(t *testing.T) {
		assert.Equal(t, 0, int(NewIEnumerable[any]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[string]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[time.Location]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*string]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]int]().SumFloat64()))
		assert.Equal(t, 0, int(NewIEnumerable[*[]string]().SumFloat64()))
	})

	t.Run("panic if result is overflow float64 (infinity positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](math.MaxFloat64, math.MaxFloat64).SumFloat64()

		fmt.Printf("Result: %v, Inf: %t\n", result, math.IsInf(result, 1))
	})

	t.Run("panic if result is overflow float64 (infinity negative)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](-1*math.MaxFloat64, -1*math.MaxFloat64).SumFloat64()
		fmt.Printf("Result: %v, Inf: %t\n", result, math.IsInf(result, -1))
	})

	t.Run("panic when element is not integer/float", func(t *testing.T) {
		str := "Hello World!"
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to float64", str))
		}()
		//goland:noinspection GoRedundantConversion
		eSrc := NewIEnumerable[any](
			int8(1), uint8(2),
			int16(3), uint16(4),
			int32(5), uint32(6),
			int64(7), uint64(8),
			int(9), uint(10),
			float32(11), float64(12),
			string(str),
		)
		eSrc.SumFloat64()
	})
}

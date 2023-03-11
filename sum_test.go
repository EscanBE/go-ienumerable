package go_ienumerable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
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

	t.Run("empty returns 0 for only integer (string)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type string cannot be tried to cast to int32")
		}()

		_ = NewIEnumerable[string]().SumInt32()
	})

	t.Run("empty returns 0 for only integer (float)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type float64 cannot be tried to cast to int32")
		}()

		_ = NewIEnumerable[float64]().SumInt32()
	})

	t.Run("empty of any integer type always returns 0", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

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
			if strings.Contains(errStr, "value 9223372036854775807 of type int cannot be casted to int32") {
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
			if strings.Contains(errStr, "value 9223372036854775807 of type int64 cannot be casted to int32") {
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
			assert.Contains(t, fmt.Sprintf("%v", err), "value 4294967295 of type uint32 cannot be casted to int32")
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
			assert.Contains(t, fmt.Sprintf("%v", err), "value 18446744073709551615 of type uint64 cannot be casted to int32")
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
			if strings.Contains(errStr, "value 18446744073709551615 of type uint cannot be casted to int32") {
				// ok (x64)
			} else if strings.Contains(errStr, "value 4294967295 of type uint cannot be casted to int32") {
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
		assert.Equal(t, int64(50), eSrc.SumInt64())
	})

	t.Run("empty returns 0 for only integer (string)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type string cannot be tried to cast to int64")
		}()

		_ = NewIEnumerable[string]().SumInt64()
	})

	t.Run("empty returns 0 for only integer (float)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type float64 cannot be tried to cast to int64")
		}()

		_ = NewIEnumerable[float64]().SumInt64()
	})

	t.Run("empty of any integer type always returns 0", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

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
			assert.Contains(t, fmt.Sprintf("%v", err), "value 18446744073709551615 of type uint64 cannot be casted to int64")
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
			if strings.Contains(errStr, "value 18446744073709551615 of type uint cannot be casted to int64") {
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

func Test_enumerable_SumFloat32(t *testing.T) {
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
			float32(-1.0), float64(-1.0),
		)
		assert.Equal(t, float32(71), eSrc.SumFloat32())
	})

	t.Run("empty returns 0 for only integer/float (string)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type string cannot be tried to cast to float32")
		}()

		_ = NewIEnumerable[string]().SumFloat32()
	})

	t.Run("empty of any integer/float type always returns 0", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		assert.Equal(t, 0, int(NewIEnumerable[int8]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint8]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[int16]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint16]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[int32]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint32]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[int64]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint64]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[int]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[uint]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[float32]().SumFloat32()))
		assert.Equal(t, 0, int(NewIEnumerable[float64]().SumFloat32()))
	})

	t.Run("panic if result is overflow float32 (positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float32](math.MaxFloat32, 1).SumFloat32()

		fmt.Printf("Result: %v\n", result)
	})

	t.Run("panic if result is overflow float32 (positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float32](math.MaxFloat32, math.MaxFloat32).SumFloat32()

		fmt.Printf("Result: %v\n", result)
	})

	t.Run("ok if during computation not overflow float64", func(t *testing.T) {
		result := NewIEnumerable[float32](math.MaxFloat32, math.MaxFloat32, -1*math.MaxFloat32, -1*math.MaxFloat32, 1).SumFloat32()

		assert.Equal(t, float32(1), result)
	})

	t.Run("panic if result is overflow float32 (negative)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float32](-1*math.MaxFloat32, -1).SumFloat32()
		fmt.Printf("Result: %v\n", result)
	})

	t.Run("panic if result is overflow float32 (negative)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float32](-1*math.MaxFloat32, -1*math.MaxFloat32).SumFloat32()
		fmt.Printf("Result: %v\n", result)
	})

	t.Run("panic if any element overflow float32", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "value 1.7976931348623157e+308 of type float64 cannot be casted to float32", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](math.MaxFloat64).SumFloat32()

		assert.Equal(t, float32(1), result)
	})

	t.Run("panic when element is not integer/float", func(t *testing.T) {
		str := "Hello World!"
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to float32", str))
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
		eSrc.SumFloat32()
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
			float32(-1.0), float64(-1.0),
		)
		assert.Equal(t, float64(71), eSrc.SumFloat64())
	})

	t.Run("empty returns 0 for only integer/float (string)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "type string cannot be tried to cast to float64")
		}()

		_ = NewIEnumerable[string]().SumFloat64()
	})

	t.Run("empty of any integer/float type always returns 0", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

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
	})

	t.Run("panic if result is overflow float64 (positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](math.MaxFloat64, 1).SumFloat64()

		fmt.Printf("Result: %v, Inf: %t\n", result, math.IsInf(result, 1))
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

	t.Run("panic if result is overflow float64 (infinity positive)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](math.MaxFloat64, math.MaxFloat64, -1*math.MaxFloat64, -1*math.MaxFloat64).SumFloat64()

		fmt.Printf("Result: %v, Inf: %t\n", result, math.IsInf(result, 1))
	})

	t.Run("panic if result is overflow float64 (negative)", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[float64](-1*math.MaxFloat64, -1).SumFloat64()
		fmt.Printf("Result: %v, Inf: %t\n", result, math.IsInf(result, -1))
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

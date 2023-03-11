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

	t.Run("empty returns 0 any integer type", func(t *testing.T) {
		defer deferWantPanicDepends(t, false)

		_ = NewIEnumerable[int8]().SumInt32()
		_ = NewIEnumerable[uint8]().SumInt32()
		_ = NewIEnumerable[int16]().SumInt32()
		_ = NewIEnumerable[uint16]().SumInt32()
		_ = NewIEnumerable[int32]().SumInt32()
		_ = NewIEnumerable[uint32]().SumInt32()
		_ = NewIEnumerable[int64]().SumInt32()
		_ = NewIEnumerable[uint64]().SumInt32()
		_ = NewIEnumerable[int]().SumInt32()
		_ = NewIEnumerable[uint]().SumInt32()
	})

	t.Run("panic if result is overflow int32", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Equal(t, "overflow int32", fmt.Sprintf("%v", err))
		}()

		result := NewIEnumerable[int64](int64(math.MaxInt32), int64(math.MinInt32), int64(math.MaxInt32), 2).SumInt32()
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

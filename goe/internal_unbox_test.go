package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func Test_enumerable_unboxAnyAsByte_any(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	ieAny := NewIEnumerable[any]()

	//goland:noinspection GoSnakeCaseUsage, SpellCheckingInspection
	eAny := e[any](ieAny)

	t.Run("in range", func(t *testing.T) {
		eAny.unboxAnyAsByte(int8(0))
		eAny.unboxAnyAsByte(int8(math.MaxInt8))
		eAny.unboxAnyAsByte(uint8(0))
		eAny.unboxAnyAsByte(uint8(math.MaxUint8))
		eAny.unboxAnyAsByte(int16(0))
		eAny.unboxAnyAsByte(int16(math.MaxUint8))
		eAny.unboxAnyAsByte(uint16(0))
		eAny.unboxAnyAsByte(uint16(math.MaxUint8))
		eAny.unboxAnyAsByte(int32(0))
		eAny.unboxAnyAsByte(int32(math.MaxUint8))
		eAny.unboxAnyAsByte(uint32(0))
		eAny.unboxAnyAsByte(uint32(math.MaxUint8))
		eAny.unboxAnyAsByte(int64(0))
		eAny.unboxAnyAsByte(int64(math.MaxUint8))
		eAny.unboxAnyAsByte(uint64(0))
		eAny.unboxAnyAsByte(uint64(math.MaxUint8))
		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsByte(int(0))
		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsByte(int(math.MaxUint8))
		eAny.unboxAnyAsByte(uint(0))
		eAny.unboxAnyAsByte(uint(math.MaxUint8))
	})

	t.Run("int8 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int8 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int8(-1))
	})

	t.Run("int16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int16(math.MaxUint8 + 1))
	})

	t.Run("int16 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int16(-1))
	})

	t.Run("uint16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint16 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(uint16(math.MaxUint8 + 1))
	})

	t.Run("int32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int32(math.MaxUint8 + 1))
	})

	t.Run("int32 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int32(-1))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(uint32(math.MaxUint8 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int64(math.MaxUint8 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(int64(-1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(uint64(math.MaxUint8 + 1))
	})

	t.Run("int over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsByte(int(math.MaxUint8 + 1))
	})

	t.Run("int under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsByte(int(-1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to byte")
		}()

		eAny.unboxAnyAsByte(uint(math.MaxUint8 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to byte", str))
		}()

		eAny.unboxAnyAsByte(any(str))
	})
}

//goland:noinspection GoSnakeCaseUsage
func Test_enumerable_unboxAnyAsByte_specific(t *testing.T) {
	ieI8 := NewIEnumerable[int8]()
	ieU8 := NewIEnumerable[uint8]()
	ieI16 := NewIEnumerable[int16]()
	ieU16 := NewIEnumerable[uint16]()
	ieI32 := NewIEnumerable[int32]()
	ieU32 := NewIEnumerable[uint32]()
	ieI64 := NewIEnumerable[int64]()
	ieU64 := NewIEnumerable[uint64]()
	ieI := NewIEnumerable[int]()
	ieU := NewIEnumerable[uint]()

	eI8 := e[int8](ieI8)
	eU8 := e[uint8](ieU8)
	eI16 := e[int16](ieI16)
	eU16 := e[uint16](ieU16)
	eI32 := e[int32](ieI32)
	eU32 := e[uint32](ieU32)
	eI64 := e[int64](ieI64)
	eU64 := e[uint64](ieU64)
	eI := e[int](ieI)
	eU := e[uint](ieU)

	t.Run("in range", func(t *testing.T) {
		eI8.unboxAnyAsByte(int8(0))
		eI8.unboxAnyAsByte(int8(math.MaxInt8))
		eU8.unboxAnyAsByte(uint8(0))
		eU8.unboxAnyAsByte(uint8(math.MaxUint8))
		eI16.unboxAnyAsByte(int16(0))
		eI16.unboxAnyAsByte(int16(math.MaxUint8))
		eU16.unboxAnyAsByte(uint16(0))
		eU16.unboxAnyAsByte(uint16(math.MaxUint8))
		eI32.unboxAnyAsByte(int32(0))
		eI32.unboxAnyAsByte(int32(math.MaxUint8))
		eU32.unboxAnyAsByte(uint32(0))
		eU32.unboxAnyAsByte(uint32(math.MaxUint8))
		eI64.unboxAnyAsByte(int64(0))
		eI64.unboxAnyAsByte(int64(math.MaxUint8))
		eU64.unboxAnyAsByte(uint64(0))
		eU64.unboxAnyAsByte(uint64(math.MaxUint8))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsByte(int(0))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsByte(int(math.MaxUint8))
		eU.unboxAnyAsByte(uint(0))
		eU.unboxAnyAsByte(uint(math.MaxUint8))
	})

	t.Run("int8 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int8 cannot be casted to byte")
		}()

		eI8.unboxAnyAsByte(int8(-1))
	})

	t.Run("int16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 cannot be casted to byte")
		}()

		eI16.unboxAnyAsByte(int16(math.MaxUint8 + 1))
	})

	t.Run("int16 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 cannot be casted to byte")
		}()

		eI16.unboxAnyAsByte(int16(-1))
	})

	t.Run("uint16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint16 cannot be casted to byte")
		}()

		eU16.unboxAnyAsByte(uint16(math.MaxUint8 + 1))
	})

	t.Run("int32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 cannot be casted to byte")
		}()

		eI32.unboxAnyAsByte(int32(math.MaxUint8 + 1))
	})

	t.Run("int32 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 cannot be casted to byte")
		}()

		eI32.unboxAnyAsByte(int32(-1))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 cannot be casted to byte")
		}()

		eU32.unboxAnyAsByte(uint32(math.MaxUint8 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to byte")
		}()

		eI64.unboxAnyAsByte(int64(math.MaxUint8 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to byte")
		}()

		eI64.unboxAnyAsByte(int64(-1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to byte")
		}()

		eU64.unboxAnyAsByte(uint64(math.MaxUint8 + 1))
	})

	t.Run("int over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsByte(int(math.MaxUint8 + 1))
	})

	t.Run("int under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsByte(int(-1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to byte")
		}()

		eU.unboxAnyAsByte(uint(math.MaxUint8 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to byte", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsByte(str)
	})
}

func Test_enumerable_unboxAnyAsInt32_any(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	eany := NewIEnumerable[any]()

	ea1 := e[any](eany)

	t.Run("in range", func(t *testing.T) {
		ea1.unboxAnyAsInt32(int8(math.MinInt8))
		ea1.unboxAnyAsInt32(int8(math.MaxInt8))
		ea1.unboxAnyAsInt32(uint8(0))
		ea1.unboxAnyAsInt32(uint8(math.MaxUint8))
		ea1.unboxAnyAsInt32(int16(math.MinInt16))
		ea1.unboxAnyAsInt32(int16(math.MaxInt16))
		ea1.unboxAnyAsInt32(uint16(0))
		ea1.unboxAnyAsInt32(uint16(math.MaxUint16))
		ea1.unboxAnyAsInt32(int32(math.MinInt32))
		ea1.unboxAnyAsInt32(int32(math.MaxInt32))
		ea1.unboxAnyAsInt32(uint32(0))
		ea1.unboxAnyAsInt32(uint32(math.MaxInt32))
		ea1.unboxAnyAsInt32(int64(math.MinInt32))
		ea1.unboxAnyAsInt32(int64(math.MaxInt32))
		ea1.unboxAnyAsInt32(uint64(0))
		ea1.unboxAnyAsInt32(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		ea1.unboxAnyAsInt32(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		ea1.unboxAnyAsInt32(int(math.MaxInt32))
		ea1.unboxAnyAsInt32(uint(0))
		ea1.unboxAnyAsInt32(uint(math.MaxInt32))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 cannot be casted to int32")
		}()

		ea1.unboxAnyAsInt32(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int32")
		}()

		ea1.unboxAnyAsInt32(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int32")
		}()

		ea1.unboxAnyAsInt32(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int32")
		}()

		ea1.unboxAnyAsInt32(uint64(math.MaxInt32 + 1))
	})

	x64 := math.MaxInt > math.MaxInt32

	t.Run("int under range", func(t *testing.T) {
		if !x64 {
			return
		}
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to int32")
		}()

		//goland:noinspection GoRedundantConversion
		ea1.unboxAnyAsInt32(int(math.MinInt32 - 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		if !x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int32")
		}()

		ea1.unboxAnyAsInt32(uint(math.MaxInt32 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int32", str))
		}()

		ea1.unboxAnyAsInt32(any(str))
	})
}

//goland:noinspection GoSnakeCaseUsage
func Test_enumerable_unboxAnyAsInt32_specific(t *testing.T) {
	ieI8 := NewIEnumerable[int8]()
	ieU8 := NewIEnumerable[uint8]()
	ieI16 := NewIEnumerable[int16]()
	ieU16 := NewIEnumerable[uint16]()
	ieI32 := NewIEnumerable[int32]()
	ieU32 := NewIEnumerable[uint32]()
	ieI64 := NewIEnumerable[int64]()
	ieU64 := NewIEnumerable[uint64]()
	ieI := NewIEnumerable[int]()
	ieU := NewIEnumerable[uint]()

	eI8 := e[int8](ieI8)
	eU8 := e[uint8](ieU8)
	eI16 := e[int16](ieI16)
	eU16 := e[uint16](ieU16)
	eI32 := e[int32](ieI32)
	eU32 := e[uint32](ieU32)
	eI64 := e[int64](ieI64)
	eU64 := e[uint64](ieU64)
	eI := e[int](ieI)
	eU := e[uint](ieU)

	t.Run("in range", func(t *testing.T) {
		eI8.unboxAnyAsInt32(int8(math.MinInt8))
		eI8.unboxAnyAsInt32(int8(math.MaxInt8))
		eU8.unboxAnyAsInt32(uint8(0))
		eU8.unboxAnyAsInt32(uint8(math.MaxUint8))
		eI16.unboxAnyAsInt32(int16(math.MinInt16))
		eI16.unboxAnyAsInt32(int16(math.MaxInt16))
		eU16.unboxAnyAsInt32(uint16(0))
		eU16.unboxAnyAsInt32(uint16(math.MaxUint16))
		eI32.unboxAnyAsInt32(int32(math.MinInt32))
		eI32.unboxAnyAsInt32(int32(math.MaxInt32))
		eU32.unboxAnyAsInt32(uint32(0))
		eU32.unboxAnyAsInt32(uint32(math.MaxInt32))
		eI64.unboxAnyAsInt32(int64(math.MinInt32))
		eI64.unboxAnyAsInt32(int64(math.MaxInt32))
		eU64.unboxAnyAsInt32(uint64(0))
		eU64.unboxAnyAsInt32(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsInt32(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsInt32(int(math.MaxInt32))
		eU.unboxAnyAsInt32(uint(0))
		eU.unboxAnyAsInt32(uint(math.MaxInt32))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 cannot be casted to int32")
		}()

		eU32.unboxAnyAsInt32(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int32")
		}()

		eI64.unboxAnyAsInt32(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int32")
		}()

		eI64.unboxAnyAsInt32(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int32")
		}()

		eU64.unboxAnyAsInt32(uint64(math.MaxInt32 + 1))
	})

	x64 := math.MaxInt > math.MaxInt32

	t.Run("int under range", func(t *testing.T) {
		if !x64 {
			return
		}
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int cannot be casted to int32")
		}()

		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsInt32(int(math.MinInt32 - 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		if !x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int32")
		}()

		eU.unboxAnyAsInt32(uint(math.MaxInt32 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int32", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsInt32(str)
	})
}

func Test_enumerable_unboxAnyAsInt64_any(t *testing.T) {

	//goland:noinspection SpellCheckingInspection
	ieAny := NewIEnumerable[any]()

	//goland:noinspection GoSnakeCaseUsage, SpellCheckingInspection
	eAny := e[any](ieAny)

	x64 := math.MaxInt > math.MaxInt32

	t.Run("in range", func(t *testing.T) {
		eAny.unboxAnyAsInt64(int8(math.MinInt8))
		eAny.unboxAnyAsInt64(int8(math.MaxInt8))
		eAny.unboxAnyAsInt64(uint8(0))
		eAny.unboxAnyAsInt64(uint8(math.MaxUint8))
		eAny.unboxAnyAsInt64(int16(math.MinInt16))
		eAny.unboxAnyAsInt64(int16(math.MaxInt16))
		eAny.unboxAnyAsInt64(uint16(0))
		eAny.unboxAnyAsInt64(uint16(math.MaxUint16))
		eAny.unboxAnyAsInt64(int32(math.MinInt32))
		eAny.unboxAnyAsInt64(int32(math.MaxInt32))
		eAny.unboxAnyAsInt64(uint32(0))
		eAny.unboxAnyAsInt64(uint32(math.MaxUint32))
		eAny.unboxAnyAsInt64(int64(math.MinInt64))
		eAny.unboxAnyAsInt64(int64(math.MaxInt64))
		eAny.unboxAnyAsInt64(uint64(0))
		eAny.unboxAnyAsInt64(uint64(math.MaxInt64))
		if x64 {
			//goland:noinspection GoRedundantConversion
			eAny.unboxAnyAsInt64(int(math.MinInt64))
			//goland:noinspection GoRedundantConversion
			eAny.unboxAnyAsInt64(int(math.MaxInt64))
			eAny.unboxAnyAsInt64(uint(0))
			eAny.unboxAnyAsInt64(uint(math.MaxInt64))
		}
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int64")
		}()

		eAny.unboxAnyAsInt64(uint64(math.MaxInt64 + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		if !x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int64")
		}()

		eAny.unboxAnyAsInt64(uint(math.MaxInt64 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int64", str))
		}()

		eAny.unboxAnyAsInt64(any(str))
	})
}

//goland:noinspection GoSnakeCaseUsage
func Test_enumerable_unboxAnyAsInt64_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ieI8 := NewIEnumerable[int8]()
	ieU8 := NewIEnumerable[uint8]()
	ieI16 := NewIEnumerable[int16]()
	ieU16 := NewIEnumerable[uint16]()
	ieI32 := NewIEnumerable[int32]()
	ieU32 := NewIEnumerable[uint32]()
	ieI64 := NewIEnumerable[int64]()
	ieU64 := NewIEnumerable[uint64]()
	ieI := NewIEnumerable[int]()
	ieU := NewIEnumerable[uint]()

	eI8 := e[int8](ieI8)
	eU8 := e[uint8](ieU8)
	eI16 := e[int16](ieI16)
	eU16 := e[uint16](ieU16)
	eI32 := e[int32](ieI32)
	eU32 := e[uint32](ieU32)
	eI64 := e[int64](ieI64)
	eU64 := e[uint64](ieU64)
	eI := e[int](ieI)
	eU := e[uint](ieU)

	t.Run("in range", func(t *testing.T) {
		eI8.unboxAnyAsInt64(int8(math.MinInt8))
		eI8.unboxAnyAsInt64(int8(math.MaxInt8))
		eU8.unboxAnyAsInt64(uint8(0))
		eU8.unboxAnyAsInt64(uint8(math.MaxUint8))
		eI16.unboxAnyAsInt64(int16(math.MinInt16))
		eI16.unboxAnyAsInt64(int16(math.MaxInt16))
		eU16.unboxAnyAsInt64(uint16(0))
		eU16.unboxAnyAsInt64(uint16(math.MaxUint16))
		eI32.unboxAnyAsInt64(int32(math.MinInt32))
		eI32.unboxAnyAsInt64(int32(math.MaxInt32))
		eU32.unboxAnyAsInt64(uint32(0))
		eU32.unboxAnyAsInt64(uint32(math.MaxInt32))
		eI64.unboxAnyAsInt64(int64(math.MinInt64))
		eI64.unboxAnyAsInt64(int64(math.MaxInt64))
		eU64.unboxAnyAsInt64(uint64(0))
		eU64.unboxAnyAsInt64(uint64(math.MaxInt32))
		if x64 {
			//goland:noinspection GoRedundantConversion
			eI.unboxAnyAsInt64(int(math.MinInt64))
			//goland:noinspection GoRedundantConversion
			eI.unboxAnyAsInt64(int(math.MaxInt64))
			eU.unboxAnyAsInt64(uint(0))
			eU.unboxAnyAsInt64(uint(math.MaxInt64))
		}
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int64")
		}()

		eU64.unboxAnyAsInt64(uint64(math.MaxInt64 + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		if !x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int64")
		}()

		eU.unboxAnyAsInt64(uint(math.MaxInt64 + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int64", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsInt64(str)
	})
}

func Test_enumerable_unboxAnyAsInt_any(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	//goland:noinspection SpellCheckingInspection
	ieAny := NewIEnumerable[any]()

	//goland:noinspection GoSnakeCaseUsage, SpellCheckingInspection
	eAny := e[any](ieAny)

	t.Run("in range", func(t *testing.T) {
		eAny.unboxAnyAsInt(int8(math.MinInt8))
		eAny.unboxAnyAsInt(int8(math.MaxInt8))
		eAny.unboxAnyAsInt(uint8(0))
		eAny.unboxAnyAsInt(uint8(math.MaxUint8))
		eAny.unboxAnyAsInt(int16(math.MinInt16))
		eAny.unboxAnyAsInt(int16(math.MaxInt16))
		eAny.unboxAnyAsInt(uint16(0))
		eAny.unboxAnyAsInt(uint16(math.MaxUint16))
		eAny.unboxAnyAsInt(int32(math.MinInt32))
		eAny.unboxAnyAsInt(int32(math.MaxInt32))
		eAny.unboxAnyAsInt(uint32(0))
		eAny.unboxAnyAsInt(uint32(math.MaxInt32))
		eAny.unboxAnyAsInt(int64(math.MinInt))
		eAny.unboxAnyAsInt(int64(math.MaxInt))
		eAny.unboxAnyAsInt(uint64(0))
		eAny.unboxAnyAsInt(uint64(math.MaxInt))
		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsInt(int(math.MinInt))
		//goland:noinspection GoRedundantConversion
		eAny.unboxAnyAsInt(int(math.MaxInt))
		eAny.unboxAnyAsInt(uint(0))
		eAny.unboxAnyAsInt(uint(math.MaxInt))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		if x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 cannot be casted to int")
		}()

		eAny.unboxAnyAsInt(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		if x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int")
		}()

		eAny.unboxAnyAsInt(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		if x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 cannot be casted to int")
		}()

		eAny.unboxAnyAsInt(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int")
		}()

		eAny.unboxAnyAsInt(uint64(math.MaxInt + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int")
		}()

		eAny.unboxAnyAsInt(uint(math.MaxInt + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int", str))
		}()

		eAny.unboxAnyAsInt(any(str))
	})
}

//goland:noinspection GoSnakeCaseUsage
func Test_enumerable_unboxAnyAsInt_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ieI8 := NewIEnumerable[int8]()
	ieU8 := NewIEnumerable[uint8]()
	ieI16 := NewIEnumerable[int16]()
	ieU16 := NewIEnumerable[uint16]()
	ieI32 := NewIEnumerable[int32]()
	ieU32 := NewIEnumerable[uint32]()
	ieI64 := NewIEnumerable[int64]()
	ieU64 := NewIEnumerable[uint64]()
	ieI := NewIEnumerable[int]()
	ieU := NewIEnumerable[uint]()

	eI8 := e[int8](ieI8)
	eU8 := e[uint8](ieU8)
	eI16 := e[int16](ieI16)
	eU16 := e[uint16](ieU16)
	eI32 := e[int32](ieI32)
	eU32 := e[uint32](ieU32)
	eI64 := e[int64](ieI64)
	eU64 := e[uint64](ieU64)
	eI := e[int](ieI)
	eU := e[uint](ieU)

	t.Run("in range", func(t *testing.T) {
		eI8.unboxAnyAsInt(int8(math.MinInt8))
		eI8.unboxAnyAsInt(int8(math.MaxInt8))
		eU8.unboxAnyAsInt(uint8(0))
		eU8.unboxAnyAsInt(uint8(math.MaxUint8))
		eI16.unboxAnyAsInt(int16(math.MinInt16))
		eI16.unboxAnyAsInt(int16(math.MaxInt16))
		eU16.unboxAnyAsInt(uint16(0))
		eU16.unboxAnyAsInt(uint16(math.MaxUint16))
		eI32.unboxAnyAsInt(int32(math.MinInt32))
		eI32.unboxAnyAsInt(int32(math.MaxInt32))
		eU32.unboxAnyAsInt(uint32(0))
		eU32.unboxAnyAsInt(uint32(math.MaxInt32))
		eI64.unboxAnyAsInt(int64(math.MinInt32))
		eI64.unboxAnyAsInt(int64(math.MaxInt32))
		eU64.unboxAnyAsInt(uint64(0))
		eU64.unboxAnyAsInt(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsInt(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		eI.unboxAnyAsInt(int(math.MaxInt32))
		eU.unboxAnyAsInt(uint(0))
		eU.unboxAnyAsInt(uint(math.MaxInt32))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		if x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 (expect: uint32) cannot be casted to int")
		}()

		eU32.unboxAnyAsInt(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		if x64 {
			return
		}
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to int")
		}()

		eI64.unboxAnyAsInt(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		if x64 {
			return
		}

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to int")
		}()

		eI64.unboxAnyAsInt(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 cannot be casted to int")
		}()

		eU64.unboxAnyAsInt(uint64(math.MaxInt + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint cannot be casted to int")
		}()

		eU.unboxAnyAsInt(uint(math.MaxInt + 1))
	})

	t.Run("not integer", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to int", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsInt(str)
	})
}

//goland:noinspection GoRedundantConversion
func Test_enumerable_unboxAnyAsFloat64OrInt64OrInt64_any(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	//goland:noinspection SpellCheckingInspection
	ieAny := NewIEnumerable[any]()

	//goland:noinspection GoSnakeCaseUsage, SpellCheckingInspection
	eAny := e[any](ieAny)

	t.Run("int64 or float64 depends value", func(t *testing.T) {
		var vf float64
		var vi int64
		var dt unboxFloat64DataType
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int8(math.MinInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int8(math.MaxInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint8(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint8(math.MaxUint8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int16(math.MinInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int16(math.MaxInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint16(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint16(math.MaxUint16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int32(math.MinInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int32(math.MaxInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint32(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint32(math.MaxUint32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int64(math.MinInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint64(0))
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint64(math.MaxUint64))
		assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		if x64 {
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int(math.MinInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MinInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(int(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint(0))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(uint(math.MaxUint64))
			assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		}
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(float32(3.3))
		assert.Greater(t, 0.0001, math.Abs(float64(3.3)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		vf, vi, dt = eAny.unboxAnyAsFloat64OrInt64(float64(9.9))
		assert.Greater(t, 0.0001, math.Abs(float64(9.9)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
	})

	t.Run("neither integer nor float", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to float64", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsFloat64OrInt64(str)
	})
}

//goland:noinspection GoSnakeCaseUsage,GoRedundantConversion
func Test_enumerable_unboxAnyAsFloat64OrInt64OrInt64_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ieI8 := NewIEnumerable[int8]()
	ieU8 := NewIEnumerable[uint8]()
	ieI16 := NewIEnumerable[int16]()
	ieU16 := NewIEnumerable[uint16]()
	ieI32 := NewIEnumerable[int32]()
	ieU32 := NewIEnumerable[uint32]()
	ieI64 := NewIEnumerable[int64]()
	ieU64 := NewIEnumerable[uint64]()
	ieI := NewIEnumerable[int]()
	ieU := NewIEnumerable[uint]()
	ieF32 := NewIEnumerable[float32]()
	ieF64 := NewIEnumerable[float64]()

	eI8 := e[int8](ieI8)
	eU8 := e[uint8](ieU8)
	eI16 := e[int16](ieI16)
	eU16 := e[uint16](ieU16)
	eI32 := e[int32](ieI32)
	eU32 := e[uint32](ieU32)
	eI64 := e[int64](ieI64)
	eU64 := e[uint64](ieU64)
	eI := e[int](ieI)
	eU := e[uint](ieU)
	eF32 := e[float32](ieF32)
	eF64 := e[float64](ieF64)

	t.Run("int64 or float64 depends value", func(t *testing.T) {
		var vf float64
		var vi int64
		var dt unboxFloat64DataType
		vf, vi, dt = eI8.unboxAnyAsFloat64OrInt64(int8(math.MinInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI8.unboxAnyAsFloat64OrInt64(int8(math.MaxInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU8.unboxAnyAsFloat64OrInt64(uint8(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU8.unboxAnyAsFloat64OrInt64(uint8(math.MaxUint8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI16.unboxAnyAsFloat64OrInt64(int16(math.MinInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI16.unboxAnyAsFloat64OrInt64(int16(math.MaxInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU16.unboxAnyAsFloat64OrInt64(uint16(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU16.unboxAnyAsFloat64OrInt64(uint16(math.MaxUint16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI32.unboxAnyAsFloat64OrInt64(int32(math.MinInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI32.unboxAnyAsFloat64OrInt64(int32(math.MaxInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU32.unboxAnyAsFloat64OrInt64(uint32(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU32.unboxAnyAsFloat64OrInt64(uint32(math.MaxUint32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI64.unboxAnyAsFloat64OrInt64(int64(math.MinInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eI64.unboxAnyAsFloat64OrInt64(int64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU64.unboxAnyAsFloat64OrInt64(uint64(0))
		vf, vi, dt = eU64.unboxAnyAsFloat64OrInt64(uint64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eU64.unboxAnyAsFloat64OrInt64(uint64(math.MaxUint64))
		assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		if x64 {
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eI.unboxAnyAsFloat64OrInt64(int(math.MinInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MinInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eI.unboxAnyAsFloat64OrInt64(int(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eU.unboxAnyAsFloat64OrInt64(uint(0))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eU.unboxAnyAsFloat64OrInt64(uint(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eU.unboxAnyAsFloat64OrInt64(uint(math.MaxUint64))
			assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		}
		vf, vi, dt = eF32.unboxAnyAsFloat64OrInt64(3.3)
		assert.Greater(t, 0.0001, math.Abs(float64(3.3)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		vf, vi, dt = eF64.unboxAnyAsFloat64OrInt64(9.9)
		assert.Greater(t, 0.0001, math.Abs(float64(9.9)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
	})

	t.Run("neither integer nor float", func(t *testing.T) {
		str := "1"

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), fmt.Sprintf("value %s of type string cannot be casted to float64", str))
		}()

		e[string](NewIEnumerable[string]()).unboxAnyAsFloat64OrInt64(str)
	})
}

func Test_enumerable_unboxAnyAsX(t *testing.T) {
	t.Run("unbox any byte & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int8 = 3
		var ab = any(b)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(b))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(ab))

		var pb *int8
		assert.Equal(t, 0, int(e.unboxAnyAsByte(pb)))

		assert.Equal(t, byte(b), e.unboxAnyAsByte(&b))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&ab))

		var pab any = any(&ab)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(pab))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&pab))

		var papab any = any(&pab)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(papab))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&papab))

		var pstr *string
		assert.Equal(t, 0, int(e.unboxAnyAsByte(pstr)))
	})

	t.Run("unbox any int32 & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int32 = 3
		var ab = any(b)
		assert.Equal(t, b, e.unboxAnyAsInt32(b))
		assert.Equal(t, b, e.unboxAnyAsInt32(ab))

		var pb *int32
		assert.Equal(t, 0, int(e.unboxAnyAsInt32(pb)))

		assert.Equal(t, b, e.unboxAnyAsInt32(&b))
		assert.Equal(t, b, e.unboxAnyAsInt32(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, e.unboxAnyAsInt32(pab))
		assert.Equal(t, b, e.unboxAnyAsInt32(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, e.unboxAnyAsInt32(papab))
		assert.Equal(t, b, e.unboxAnyAsInt32(&papab))
	})

	t.Run("unbox any int64 & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int64 = 3
		var ab = any(b)
		assert.Equal(t, b, e.unboxAnyAsInt64(b))
		assert.Equal(t, b, e.unboxAnyAsInt64(ab))

		var pb *int64
		assert.Equal(t, 0, int(e.unboxAnyAsInt64(pb)))

		assert.Equal(t, b, e.unboxAnyAsInt64(&b))
		assert.Equal(t, b, e.unboxAnyAsInt64(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, e.unboxAnyAsInt64(pab))
		assert.Equal(t, b, e.unboxAnyAsInt64(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, e.unboxAnyAsInt64(papab))
		assert.Equal(t, b, e.unboxAnyAsInt64(&papab))
	})

	t.Run("unbox any int & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int = 3
		var ab = any(b)
		assert.Equal(t, b, e.unboxAnyAsInt(b))
		assert.Equal(t, b, e.unboxAnyAsInt(ab))

		var pb *int
		assert.Equal(t, 0, int(e.unboxAnyAsInt(pb)))

		assert.Equal(t, b, e.unboxAnyAsInt(&b))
		assert.Equal(t, b, e.unboxAnyAsInt(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, e.unboxAnyAsInt(pab))
		assert.Equal(t, b, e.unboxAnyAsInt(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, e.unboxAnyAsInt(papab))
		assert.Equal(t, b, e.unboxAnyAsInt(&papab))
	})
}

func testEnumerable_unboxAnyAsX[T any](t *testing.T, fUnbox func(v any) T) {
	t.Run("unbox any byte & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int8 = 3
		var ab = any(b)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(b))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(ab))

		var pb *int8
		assert.Equal(t, 0, int(e.unboxAnyAsByte(pb)))

		assert.Equal(t, byte(b), e.unboxAnyAsByte(&b))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&ab))

		var pab any = any(&ab)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(pab))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&pab))

		var papab any = any(&pab)
		assert.Equal(t, byte(b), e.unboxAnyAsByte(papab))
		assert.Equal(t, byte(b), e.unboxAnyAsByte(&papab))
	})

	t.Run("unbox any int32 & pointer", func(t *testing.T) {
		e := e[any](NewIEnumerable[any]())
		var b int32 = 3
		var ab = any(b)
		assert.Equal(t, b, e.unboxAnyAsInt32(b))
		assert.Equal(t, b, e.unboxAnyAsInt32(ab))

		var pb *int32
		assert.Equal(t, 0, int(e.unboxAnyAsInt32(pb)))

		assert.Equal(t, b, e.unboxAnyAsInt32(&b))
		assert.Equal(t, b, e.unboxAnyAsInt32(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, e.unboxAnyAsInt32(pab))
		assert.Equal(t, b, e.unboxAnyAsInt32(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, e.unboxAnyAsInt32(papab))
		assert.Equal(t, b, e.unboxAnyAsInt32(&papab))
	})
}

func Test_enumerable_unboxUsingReflecting(t *testing.T) {
	var i8 int8 = 1
	infoOf[int8](i8)
	infoOf[*int8](&i8)
}

func infoOf[T any](v T) {
	vo := reflect.ValueOf(v)
	fmt.Printf("\n\n========================\n\nInfo of %s(%v) %v\n", vo.Type().String(), vo.Type().Kind(), v)
	//fmt.Printf("Nil: %t", vo.IsNil())
	fmt.Printf("Kind: %v\n", vo.Kind())
	fmt.Printf("\nIs Zero: %t", vo.IsZero())
	fmt.Printf("\nCan Int: %t", vo.CanInt())
	if vo.CanInt() {
		fmt.Printf(" = %d", vo.Int())
	}
	fmt.Printf("\nCan Uint: %t", vo.CanUint())
	if vo.CanUint() {
		fmt.Printf(" = %d", vo.Uint())
	}
	fmt.Printf("\nCan Float: %t", vo.CanFloat())
	if vo.CanUint() {
		fmt.Printf(" = %f", vo.Float())
	}
	fmt.Printf("\nCan Complex: %t", vo.CanComplex())
	if vo.CanUint() {
		fmt.Printf(" = %f", vo.Complex())
	}
}

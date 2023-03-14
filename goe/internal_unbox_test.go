package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_enumerable_unboxAnyAsByte_any(t *testing.T) {
	eany := NewIEnumerable[any]()
	t.Run("in range", func(t *testing.T) {
		eany.unboxAnyAsByte(int8(0))
		eany.unboxAnyAsByte(int8(math.MaxInt8))
		eany.unboxAnyAsByte(uint8(0))
		eany.unboxAnyAsByte(uint8(math.MaxUint8))
		eany.unboxAnyAsByte(int16(0))
		eany.unboxAnyAsByte(int16(math.MaxUint8))
		eany.unboxAnyAsByte(uint16(0))
		eany.unboxAnyAsByte(uint16(math.MaxUint8))
		eany.unboxAnyAsByte(int32(0))
		eany.unboxAnyAsByte(int32(math.MaxUint8))
		eany.unboxAnyAsByte(uint32(0))
		eany.unboxAnyAsByte(uint32(math.MaxUint8))
		eany.unboxAnyAsByte(int64(0))
		eany.unboxAnyAsByte(int64(math.MaxUint8))
		eany.unboxAnyAsByte(uint64(0))
		eany.unboxAnyAsByte(uint64(math.MaxUint8))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsByte(int(0))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsByte(int(math.MaxUint8))
		eany.unboxAnyAsByte(uint(0))
		eany.unboxAnyAsByte(uint(math.MaxUint8))
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

		eany.unboxAnyAsByte(int8(-1))
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

		eany.unboxAnyAsByte(int16(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(int16(-1))
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

		eany.unboxAnyAsByte(uint16(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(int32(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(int32(-1))
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

		eany.unboxAnyAsByte(uint32(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(int64(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(int64(-1))
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

		eany.unboxAnyAsByte(uint64(math.MaxUint8 + 1))
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
		eany.unboxAnyAsByte(int(math.MaxUint8 + 1))
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
		eany.unboxAnyAsByte(int(-1))
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

		eany.unboxAnyAsByte(uint(math.MaxUint8 + 1))
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

		eany.unboxAnyAsByte(any(str))
	})
}

func Test_enumerable_unboxAnyAsByte_specific(t *testing.T) {
	ei8 := NewIEnumerable[int8]()
	eu8 := NewIEnumerable[uint8]()
	ei16 := NewIEnumerable[int16]()
	eu16 := NewIEnumerable[uint16]()
	ei32 := NewIEnumerable[int32]()
	eu32 := NewIEnumerable[uint32]()
	ei64 := NewIEnumerable[int64]()
	eu64 := NewIEnumerable[uint64]()
	ei := NewIEnumerable[int]()
	eu := NewIEnumerable[uint]()
	t.Run("in range", func(t *testing.T) {
		ei8.unboxAnyAsByte(int8(0))
		ei8.unboxAnyAsByte(int8(math.MaxInt8))
		eu8.unboxAnyAsByte(uint8(0))
		eu8.unboxAnyAsByte(uint8(math.MaxUint8))
		ei16.unboxAnyAsByte(int16(0))
		ei16.unboxAnyAsByte(int16(math.MaxUint8))
		eu16.unboxAnyAsByte(uint16(0))
		eu16.unboxAnyAsByte(uint16(math.MaxUint8))
		ei32.unboxAnyAsByte(int32(0))
		ei32.unboxAnyAsByte(int32(math.MaxUint8))
		eu32.unboxAnyAsByte(uint32(0))
		eu32.unboxAnyAsByte(uint32(math.MaxUint8))
		ei64.unboxAnyAsByte(int64(0))
		ei64.unboxAnyAsByte(int64(math.MaxUint8))
		eu64.unboxAnyAsByte(uint64(0))
		eu64.unboxAnyAsByte(uint64(math.MaxUint8))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsByte(int(0))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsByte(int(math.MaxUint8))
		eu.unboxAnyAsByte(uint(0))
		eu.unboxAnyAsByte(uint(math.MaxUint8))
	})

	t.Run("int8 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int8 (expect: int8) cannot be casted to byte")
		}()

		ei8.unboxAnyAsByte(int8(-1))
	})

	t.Run("int16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 (expect: int16) cannot be casted to byte")
		}()

		ei16.unboxAnyAsByte(int16(math.MaxUint8 + 1))
	})

	t.Run("int16 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 (expect: int16) cannot be casted to byte")
		}()

		ei16.unboxAnyAsByte(int16(-1))
	})

	t.Run("uint16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint16 (expect: uint16) cannot be casted to byte")
		}()

		eu16.unboxAnyAsByte(uint16(math.MaxUint8 + 1))
	})

	t.Run("int32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 (expect: int32) cannot be casted to byte")
		}()

		ei32.unboxAnyAsByte(int32(math.MaxUint8 + 1))
	})

	t.Run("int32 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 (expect: int32) cannot be casted to byte")
		}()

		ei32.unboxAnyAsByte(int32(-1))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 (expect: uint32) cannot be casted to byte")
		}()

		eu32.unboxAnyAsByte(uint32(math.MaxUint8 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to byte")
		}()

		ei64.unboxAnyAsByte(int64(math.MaxUint8 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to byte")
		}()

		ei64.unboxAnyAsByte(int64(-1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 (expect: uint64) cannot be casted to byte")
		}()

		eu64.unboxAnyAsByte(uint64(math.MaxUint8 + 1))
	})

	t.Run("int over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int (expect: int) cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsByte(int(math.MaxUint8 + 1))
	})

	t.Run("int under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int (expect: int) cannot be casted to byte")
		}()

		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsByte(int(-1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint (expect: uint) cannot be casted to byte")
		}()

		eu.unboxAnyAsByte(uint(math.MaxUint8 + 1))
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

		NewIEnumerable[string]().unboxAnyAsByte(str)
	})
}

func Test_enumerable_unboxAnyAsInt32_any(t *testing.T) {
	eany := NewIEnumerable[any]()
	t.Run("in range", func(t *testing.T) {
		eany.unboxAnyAsInt32(int8(math.MinInt8))
		eany.unboxAnyAsInt32(int8(math.MaxInt8))
		eany.unboxAnyAsInt32(uint8(0))
		eany.unboxAnyAsInt32(uint8(math.MaxUint8))
		eany.unboxAnyAsInt32(int16(math.MinInt16))
		eany.unboxAnyAsInt32(int16(math.MaxInt16))
		eany.unboxAnyAsInt32(uint16(0))
		eany.unboxAnyAsInt32(uint16(math.MaxUint16))
		eany.unboxAnyAsInt32(int32(math.MinInt32))
		eany.unboxAnyAsInt32(int32(math.MaxInt32))
		eany.unboxAnyAsInt32(uint32(0))
		eany.unboxAnyAsInt32(uint32(math.MaxInt32))
		eany.unboxAnyAsInt32(int64(math.MinInt32))
		eany.unboxAnyAsInt32(int64(math.MaxInt32))
		eany.unboxAnyAsInt32(uint64(0))
		eany.unboxAnyAsInt32(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsInt32(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsInt32(int(math.MaxInt32))
		eany.unboxAnyAsInt32(uint(0))
		eany.unboxAnyAsInt32(uint(math.MaxInt32))
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

		eany.unboxAnyAsInt32(uint32(math.MaxInt32 + 1))
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

		eany.unboxAnyAsInt32(int64(math.MaxInt32 + 1))
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

		eany.unboxAnyAsInt32(int64(math.MinInt32 - 1))
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

		eany.unboxAnyAsInt32(uint64(math.MaxInt32 + 1))
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
		eany.unboxAnyAsInt32(int(math.MinInt32 - 1))
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

		eany.unboxAnyAsInt32(uint(math.MaxInt32 + 1))
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

		eany.unboxAnyAsInt32(any(str))
	})
}

func Test_enumerable_unboxAnyAsInt32_specific(t *testing.T) {
	ei8 := NewIEnumerable[int8]()
	eu8 := NewIEnumerable[uint8]()
	ei16 := NewIEnumerable[int16]()
	eu16 := NewIEnumerable[uint16]()
	ei32 := NewIEnumerable[int32]()
	eu32 := NewIEnumerable[uint32]()
	ei64 := NewIEnumerable[int64]()
	eu64 := NewIEnumerable[uint64]()
	ei := NewIEnumerable[int]()
	eu := NewIEnumerable[uint]()
	t.Run("in range", func(t *testing.T) {
		ei8.unboxAnyAsInt32(int8(math.MinInt8))
		ei8.unboxAnyAsInt32(int8(math.MaxInt8))
		eu8.unboxAnyAsInt32(uint8(0))
		eu8.unboxAnyAsInt32(uint8(math.MaxUint8))
		ei16.unboxAnyAsInt32(int16(math.MinInt16))
		ei16.unboxAnyAsInt32(int16(math.MaxInt16))
		eu16.unboxAnyAsInt32(uint16(0))
		eu16.unboxAnyAsInt32(uint16(math.MaxUint16))
		ei32.unboxAnyAsInt32(int32(math.MinInt32))
		ei32.unboxAnyAsInt32(int32(math.MaxInt32))
		eu32.unboxAnyAsInt32(uint32(0))
		eu32.unboxAnyAsInt32(uint32(math.MaxInt32))
		ei64.unboxAnyAsInt32(int64(math.MinInt32))
		ei64.unboxAnyAsInt32(int64(math.MaxInt32))
		eu64.unboxAnyAsInt32(uint64(0))
		eu64.unboxAnyAsInt32(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsInt32(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsInt32(int(math.MaxInt32))
		eu.unboxAnyAsInt32(uint(0))
		eu.unboxAnyAsInt32(uint(math.MaxInt32))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 (expect: uint32) cannot be casted to int32")
		}()

		eu32.unboxAnyAsInt32(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to int32")
		}()

		ei64.unboxAnyAsInt32(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 (expect: int64) cannot be casted to int32")
		}()

		ei64.unboxAnyAsInt32(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 (expect: uint64) cannot be casted to int32")
		}()

		eu64.unboxAnyAsInt32(uint64(math.MaxInt32 + 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int (expect: int) cannot be casted to int32")
		}()

		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsInt32(int(math.MinInt32 - 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint (expect: uint) cannot be casted to int32")
		}()

		eu.unboxAnyAsInt32(uint(math.MaxInt32 + 1))
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

		NewIEnumerable[string]().unboxAnyAsInt32(str)
	})
}

func Test_enumerable_unboxAnyAsInt64_any(t *testing.T) {
	eany := NewIEnumerable[any]()

	x64 := math.MaxInt > math.MaxInt32

	t.Run("in range", func(t *testing.T) {
		eany.unboxAnyAsInt64(int8(math.MinInt8))
		eany.unboxAnyAsInt64(int8(math.MaxInt8))
		eany.unboxAnyAsInt64(uint8(0))
		eany.unboxAnyAsInt64(uint8(math.MaxUint8))
		eany.unboxAnyAsInt64(int16(math.MinInt16))
		eany.unboxAnyAsInt64(int16(math.MaxInt16))
		eany.unboxAnyAsInt64(uint16(0))
		eany.unboxAnyAsInt64(uint16(math.MaxUint16))
		eany.unboxAnyAsInt64(int32(math.MinInt32))
		eany.unboxAnyAsInt64(int32(math.MaxInt32))
		eany.unboxAnyAsInt64(uint32(0))
		eany.unboxAnyAsInt64(uint32(math.MaxUint32))
		eany.unboxAnyAsInt64(int64(math.MinInt64))
		eany.unboxAnyAsInt64(int64(math.MaxInt64))
		eany.unboxAnyAsInt64(uint64(0))
		eany.unboxAnyAsInt64(uint64(math.MaxInt64))
		if x64 {
			//goland:noinspection GoRedundantConversion
			eany.unboxAnyAsInt64(int(math.MinInt64))
			//goland:noinspection GoRedundantConversion
			eany.unboxAnyAsInt64(int(math.MaxInt64))
			eany.unboxAnyAsInt64(uint(0))
			eany.unboxAnyAsInt64(uint(math.MaxInt64))
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

		eany.unboxAnyAsInt64(uint64(math.MaxInt64 + 1))
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

		eany.unboxAnyAsInt64(uint(math.MaxInt64 + 1))
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

		eany.unboxAnyAsInt64(any(str))
	})
}

func Test_enumerable_unboxAnyAsInt64_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ei8 := NewIEnumerable[int8]()
	eu8 := NewIEnumerable[uint8]()
	ei16 := NewIEnumerable[int16]()
	eu16 := NewIEnumerable[uint16]()
	ei32 := NewIEnumerable[int32]()
	eu32 := NewIEnumerable[uint32]()
	ei64 := NewIEnumerable[int64]()
	eu64 := NewIEnumerable[uint64]()
	ei := NewIEnumerable[int]()
	eu := NewIEnumerable[uint]()
	t.Run("in range", func(t *testing.T) {
		ei8.unboxAnyAsInt64(int8(math.MinInt8))
		ei8.unboxAnyAsInt64(int8(math.MaxInt8))
		eu8.unboxAnyAsInt64(uint8(0))
		eu8.unboxAnyAsInt64(uint8(math.MaxUint8))
		ei16.unboxAnyAsInt64(int16(math.MinInt16))
		ei16.unboxAnyAsInt64(int16(math.MaxInt16))
		eu16.unboxAnyAsInt64(uint16(0))
		eu16.unboxAnyAsInt64(uint16(math.MaxUint16))
		ei32.unboxAnyAsInt64(int32(math.MinInt32))
		ei32.unboxAnyAsInt64(int32(math.MaxInt32))
		eu32.unboxAnyAsInt64(uint32(0))
		eu32.unboxAnyAsInt64(uint32(math.MaxInt32))
		ei64.unboxAnyAsInt64(int64(math.MinInt64))
		ei64.unboxAnyAsInt64(int64(math.MaxInt64))
		eu64.unboxAnyAsInt64(uint64(0))
		eu64.unboxAnyAsInt64(uint64(math.MaxInt32))
		if x64 {
			//goland:noinspection GoRedundantConversion
			ei.unboxAnyAsInt64(int(math.MinInt64))
			//goland:noinspection GoRedundantConversion
			ei.unboxAnyAsInt64(int(math.MaxInt64))
			eu.unboxAnyAsInt64(uint(0))
			eu.unboxAnyAsInt64(uint(math.MaxInt64))
		}
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 (expect: uint64) cannot be casted to int64")
		}()

		eu64.unboxAnyAsInt64(uint64(math.MaxInt64 + 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint (expect: uint) cannot be casted to int64")
		}()

		eu.unboxAnyAsInt64(uint(math.MaxInt64 + 1))
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

		NewIEnumerable[string]().unboxAnyAsInt64(str)
	})
}

func Test_enumerable_unboxAnyAsInt_any(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	eany := NewIEnumerable[any]()
	t.Run("in range", func(t *testing.T) {
		eany.unboxAnyAsInt(int8(math.MinInt8))
		eany.unboxAnyAsInt(int8(math.MaxInt8))
		eany.unboxAnyAsInt(uint8(0))
		eany.unboxAnyAsInt(uint8(math.MaxUint8))
		eany.unboxAnyAsInt(int16(math.MinInt16))
		eany.unboxAnyAsInt(int16(math.MaxInt16))
		eany.unboxAnyAsInt(uint16(0))
		eany.unboxAnyAsInt(uint16(math.MaxUint16))
		eany.unboxAnyAsInt(int32(math.MinInt32))
		eany.unboxAnyAsInt(int32(math.MaxInt32))
		eany.unboxAnyAsInt(uint32(0))
		eany.unboxAnyAsInt(uint32(math.MaxInt32))
		eany.unboxAnyAsInt(int64(math.MinInt))
		eany.unboxAnyAsInt(int64(math.MaxInt))
		eany.unboxAnyAsInt(uint64(0))
		eany.unboxAnyAsInt(uint64(math.MaxInt))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsInt(int(math.MinInt))
		//goland:noinspection GoRedundantConversion
		eany.unboxAnyAsInt(int(math.MaxInt))
		eany.unboxAnyAsInt(uint(0))
		eany.unboxAnyAsInt(uint(math.MaxInt))
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

		eany.unboxAnyAsInt(uint32(math.MaxInt32 + 1))
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

		eany.unboxAnyAsInt(int64(math.MaxInt32 + 1))
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

		eany.unboxAnyAsInt(int64(math.MinInt32 - 1))
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

		eany.unboxAnyAsInt(uint64(math.MaxInt + 1))
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

		eany.unboxAnyAsInt(uint(math.MaxInt + 1))
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

		eany.unboxAnyAsInt(any(str))
	})
}

func Test_enumerable_unboxAnyAsInt_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ei8 := NewIEnumerable[int8]()
	eu8 := NewIEnumerable[uint8]()
	ei16 := NewIEnumerable[int16]()
	eu16 := NewIEnumerable[uint16]()
	ei32 := NewIEnumerable[int32]()
	eu32 := NewIEnumerable[uint32]()
	ei64 := NewIEnumerable[int64]()
	eu64 := NewIEnumerable[uint64]()
	ei := NewIEnumerable[int]()
	eu := NewIEnumerable[uint]()
	t.Run("in range", func(t *testing.T) {
		ei8.unboxAnyAsInt(int8(math.MinInt8))
		ei8.unboxAnyAsInt(int8(math.MaxInt8))
		eu8.unboxAnyAsInt(uint8(0))
		eu8.unboxAnyAsInt(uint8(math.MaxUint8))
		ei16.unboxAnyAsInt(int16(math.MinInt16))
		ei16.unboxAnyAsInt(int16(math.MaxInt16))
		eu16.unboxAnyAsInt(uint16(0))
		eu16.unboxAnyAsInt(uint16(math.MaxUint16))
		ei32.unboxAnyAsInt(int32(math.MinInt32))
		ei32.unboxAnyAsInt(int32(math.MaxInt32))
		eu32.unboxAnyAsInt(uint32(0))
		eu32.unboxAnyAsInt(uint32(math.MaxInt32))
		ei64.unboxAnyAsInt(int64(math.MinInt32))
		ei64.unboxAnyAsInt(int64(math.MaxInt32))
		eu64.unboxAnyAsInt(uint64(0))
		eu64.unboxAnyAsInt(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsInt(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		ei.unboxAnyAsInt(int(math.MaxInt32))
		eu.unboxAnyAsInt(uint(0))
		eu.unboxAnyAsInt(uint(math.MaxInt32))
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

		eu32.unboxAnyAsInt(uint32(math.MaxInt32 + 1))
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

		ei64.unboxAnyAsInt(int64(math.MaxInt32 + 1))
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

		ei64.unboxAnyAsInt(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 (expect: uint64) cannot be casted to int")
		}()

		eu64.unboxAnyAsInt(uint64(math.MaxInt + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint (expect: uint) cannot be casted to int")
		}()

		eu.unboxAnyAsInt(uint(math.MaxInt + 1))
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

		NewIEnumerable[string]().unboxAnyAsInt(str)
	})
}

func Test_enumerable_unboxAnyAsFloat64OrInt64OrInt64_any(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32
	eany := NewIEnumerable[any]()
	t.Run("int64 or float64 depends value", func(t *testing.T) {
		var vf float64
		var vi int64
		var dt unboxFloat64DataType
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int8(math.MinInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int8(math.MaxInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint8(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint8(math.MaxUint8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int16(math.MinInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int16(math.MaxInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint16(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint16(math.MaxUint16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int32(math.MinInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int32(math.MaxInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint32(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint32(math.MaxUint32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int64(math.MinInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint64(0))
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint64(math.MaxUint64))
		assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		if x64 {
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int(math.MinInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MinInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(int(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint(0))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(uint(math.MaxUint64))
			assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		}
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(float32(3.3))
		assert.Greater(t, 0.0001, math.Abs(float64(3.3)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		vf, vi, dt = eany.unboxAnyAsFloat64OrInt64(float64(9.9))
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

		NewIEnumerable[string]().unboxAnyAsFloat64OrInt64(str)
	})
}

func Test_enumerable_unboxAnyAsFloat64OrInt64OrInt64_specific(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	ei8 := NewIEnumerable[int8]()
	eu8 := NewIEnumerable[uint8]()
	ei16 := NewIEnumerable[int16]()
	eu16 := NewIEnumerable[uint16]()
	ei32 := NewIEnumerable[int32]()
	eu32 := NewIEnumerable[uint32]()
	ei64 := NewIEnumerable[int64]()
	eu64 := NewIEnumerable[uint64]()
	ei := NewIEnumerable[int]()
	eu := NewIEnumerable[uint]()
	ef32 := NewIEnumerable[float32]()
	ef64 := NewIEnumerable[float64]()
	t.Run("int64 or float64 depends value", func(t *testing.T) {
		var vf float64
		var vi int64
		var dt unboxFloat64DataType
		vf, vi, dt = ei8.unboxAnyAsFloat64OrInt64(int8(math.MinInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei8.unboxAnyAsFloat64OrInt64(int8(math.MaxInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu8.unboxAnyAsFloat64OrInt64(uint8(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu8.unboxAnyAsFloat64OrInt64(uint8(math.MaxUint8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei16.unboxAnyAsFloat64OrInt64(int16(math.MinInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei16.unboxAnyAsFloat64OrInt64(int16(math.MaxInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu16.unboxAnyAsFloat64OrInt64(uint16(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu16.unboxAnyAsFloat64OrInt64(uint16(math.MaxUint16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei32.unboxAnyAsFloat64OrInt64(int32(math.MinInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei32.unboxAnyAsFloat64OrInt64(int32(math.MaxInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu32.unboxAnyAsFloat64OrInt64(uint32(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu32.unboxAnyAsFloat64OrInt64(uint32(math.MaxUint32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei64.unboxAnyAsFloat64OrInt64(int64(math.MinInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = ei64.unboxAnyAsFloat64OrInt64(int64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu64.unboxAnyAsFloat64OrInt64(uint64(0))
		vf, vi, dt = eu64.unboxAnyAsFloat64OrInt64(uint64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vf, vi, dt = eu64.unboxAnyAsFloat64OrInt64(uint64(math.MaxUint64))
		assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		if x64 {
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = ei.unboxAnyAsFloat64OrInt64(int(math.MinInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MinInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			//goland:noinspection GoRedundantConversion
			vf, vi, dt = ei.unboxAnyAsFloat64OrInt64(int(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eu.unboxAnyAsFloat64OrInt64(uint(0))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eu.unboxAnyAsFloat64OrInt64(uint(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vf, vi, dt = eu.unboxAnyAsFloat64OrInt64(uint(math.MaxUint64))
			assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		}
		vf, vi, dt = ef32.unboxAnyAsFloat64OrInt64(3.3)
		assert.Greater(t, 0.0001, math.Abs(float64(3.3)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		vf, vi, dt = ef64.unboxAnyAsFloat64OrInt64(9.9)
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

		NewIEnumerable[string]().unboxAnyAsFloat64OrInt64(str)
	})
}

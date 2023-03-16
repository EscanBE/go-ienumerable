package reflection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestUnboxAnyAsByte(t *testing.T) {
	t.Run("in range", func(t *testing.T) {
		UnboxAnyAsByte(int8(0))
		UnboxAnyAsByte(int8(math.MaxInt8))
		UnboxAnyAsByte(uint8(0))
		UnboxAnyAsByte(uint8(math.MaxUint8))
		UnboxAnyAsByte(int16(0))
		UnboxAnyAsByte(int16(math.MaxUint8))
		UnboxAnyAsByte(uint16(0))
		UnboxAnyAsByte(uint16(math.MaxUint8))
		UnboxAnyAsByte(int32(0))
		UnboxAnyAsByte(int32(math.MaxUint8))
		UnboxAnyAsByte(uint32(0))
		UnboxAnyAsByte(uint32(math.MaxUint8))
		UnboxAnyAsByte(int64(0))
		UnboxAnyAsByte(int64(math.MaxUint8))
		UnboxAnyAsByte(uint64(0))
		UnboxAnyAsByte(uint64(math.MaxUint8))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsByte(int(0))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsByte(int(math.MaxUint8))
		UnboxAnyAsByte(uint(0))
		UnboxAnyAsByte(uint(math.MaxUint8))
	})

	t.Run("int8 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int8 is over range of byte")
		}()

		UnboxAnyAsByte(int8(-1))
	})

	t.Run("int16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 is over range of byte")
		}()

		UnboxAnyAsByte(int16(math.MaxUint8 + 1))
	})

	t.Run("int16 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int16 is over range of byte")
		}()

		UnboxAnyAsByte(int16(-1))
	})

	t.Run("uint16 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint16 is over range of byte")
		}()

		UnboxAnyAsByte(uint16(math.MaxUint8 + 1))
	})

	t.Run("int32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 is over range of byte")
		}()

		UnboxAnyAsByte(int32(math.MaxUint8 + 1))
	})

	t.Run("int32 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int32 is over range of byte")
		}()

		UnboxAnyAsByte(int32(-1))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 is over range of byte")
		}()

		UnboxAnyAsByte(uint32(math.MaxUint8 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 is over range of byte")
		}()

		UnboxAnyAsByte(int64(math.MaxUint8 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 is over range of byte")
		}()

		UnboxAnyAsByte(int64(-1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 is over range of byte")
		}()

		UnboxAnyAsByte(uint64(math.MaxUint8 + 1))
	})

	t.Run("int over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int is over range of byte")
		}()

		//goland:noinspection GoRedundantConversion
		UnboxAnyAsByte(int(math.MaxUint8 + 1))
	})

	t.Run("int under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int is over range of byte")
		}()

		//goland:noinspection GoRedundantConversion
		UnboxAnyAsByte(int(-1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint is over range of byte")
		}()

		UnboxAnyAsByte(uint(math.MaxUint8 + 1))
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

		UnboxAnyAsByte(any(str))
	})
}

func TestUnboxAnyAsInt32(t *testing.T) {
	t.Run("in range", func(t *testing.T) {
		UnboxAnyAsInt32(int8(math.MinInt8))
		UnboxAnyAsInt32(int8(math.MaxInt8))
		UnboxAnyAsInt32(uint8(0))
		UnboxAnyAsInt32(uint8(math.MaxUint8))
		UnboxAnyAsInt32(int16(math.MinInt16))
		UnboxAnyAsInt32(int16(math.MaxInt16))
		UnboxAnyAsInt32(uint16(0))
		UnboxAnyAsInt32(uint16(math.MaxUint16))
		UnboxAnyAsInt32(int32(math.MinInt32))
		UnboxAnyAsInt32(int32(math.MaxInt32))
		UnboxAnyAsInt32(uint32(0))
		UnboxAnyAsInt32(uint32(math.MaxInt32))
		UnboxAnyAsInt32(int64(math.MinInt32))
		UnboxAnyAsInt32(int64(math.MaxInt32))
		UnboxAnyAsInt32(uint64(0))
		UnboxAnyAsInt32(uint64(math.MaxInt32))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsInt32(int(math.MinInt32))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsInt32(int(math.MaxInt32))
		UnboxAnyAsInt32(uint(0))
		UnboxAnyAsInt32(uint(math.MaxInt32))
	})

	t.Run("uint32 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint32 is over range of int32")
		}()

		UnboxAnyAsInt32(uint32(math.MaxInt32 + 1))
	})

	t.Run("int64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 is over range of int32")
		}()

		UnboxAnyAsInt32(int64(math.MaxInt32 + 1))
	})

	t.Run("int64 under range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int64 is over range of int32")
		}()

		UnboxAnyAsInt32(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 is over range of int32")
		}()

		UnboxAnyAsInt32(uint64(math.MaxInt32 + 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type int is over range of int32")
		}()

		//goland:noinspection GoRedundantConversion
		UnboxAnyAsInt32(int(math.MinInt32 - 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint is over range of int32")
		}()

		UnboxAnyAsInt32(uint(math.MaxInt32 + 1))
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

		UnboxAnyAsInt32(any(str))
	})
}

func TestUnboxAnyAsInt64(t *testing.T) {

	x64 := math.MaxInt > math.MaxInt32

	t.Run("in range", func(t *testing.T) {
		UnboxAnyAsInt64(int8(math.MinInt8))
		UnboxAnyAsInt64(int8(math.MaxInt8))
		UnboxAnyAsInt64(uint8(0))
		UnboxAnyAsInt64(uint8(math.MaxUint8))
		UnboxAnyAsInt64(int16(math.MinInt16))
		UnboxAnyAsInt64(int16(math.MaxInt16))
		UnboxAnyAsInt64(uint16(0))
		UnboxAnyAsInt64(uint16(math.MaxUint16))
		UnboxAnyAsInt64(int32(math.MinInt32))
		UnboxAnyAsInt64(int32(math.MaxInt32))
		UnboxAnyAsInt64(uint32(0))
		UnboxAnyAsInt64(uint32(math.MaxUint32))
		UnboxAnyAsInt64(int64(math.MinInt64))
		UnboxAnyAsInt64(int64(math.MaxInt64))
		UnboxAnyAsInt64(uint64(0))
		UnboxAnyAsInt64(uint64(math.MaxInt64))
		if x64 {
			//goland:noinspection GoRedundantConversion
			UnboxAnyAsInt64(int(math.MinInt64))
			//goland:noinspection GoRedundantConversion
			UnboxAnyAsInt64(int(math.MaxInt64))
			UnboxAnyAsInt64(uint(0))
			UnboxAnyAsInt64(uint(math.MaxInt64))
		}
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 is over range of int64")
		}()

		UnboxAnyAsInt64(uint64(math.MaxInt64 + 1))
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

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint is over range of int64")
		}()

		UnboxAnyAsInt64(uint(math.MaxInt64 + 1))
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

		UnboxAnyAsInt64(any(str))
	})
}

func TestUnboxAnyAsInt(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	t.Run("in range", func(t *testing.T) {
		UnboxAnyAsInt(int8(math.MinInt8))
		UnboxAnyAsInt(int8(math.MaxInt8))
		UnboxAnyAsInt(uint8(0))
		UnboxAnyAsInt(uint8(math.MaxUint8))
		UnboxAnyAsInt(int16(math.MinInt16))
		UnboxAnyAsInt(int16(math.MaxInt16))
		UnboxAnyAsInt(uint16(0))
		UnboxAnyAsInt(uint16(math.MaxUint16))
		UnboxAnyAsInt(int32(math.MinInt32))
		UnboxAnyAsInt(int32(math.MaxInt32))
		UnboxAnyAsInt(uint32(0))
		UnboxAnyAsInt(uint32(math.MaxInt32))
		UnboxAnyAsInt(int64(math.MinInt))
		UnboxAnyAsInt(int64(math.MaxInt))
		UnboxAnyAsInt(uint64(0))
		UnboxAnyAsInt(uint64(math.MaxInt))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsInt(int(math.MinInt))
		//goland:noinspection GoRedundantConversion
		UnboxAnyAsInt(int(math.MaxInt))
		UnboxAnyAsInt(uint(0))
		UnboxAnyAsInt(uint(math.MaxInt))
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

		UnboxAnyAsInt(uint32(math.MaxInt32 + 1))
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

		UnboxAnyAsInt(int64(math.MaxInt32 + 1))
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

		UnboxAnyAsInt(int64(math.MinInt32 - 1))
	})

	t.Run("uint64 over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint64 is over range of int")
		}()

		UnboxAnyAsInt(uint64(math.MaxInt + 1))
	})

	t.Run("uint over range", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect panic")
				return
			}

			assert.Contains(t, fmt.Sprintf("%v", err), "of type uint is over range of int")
		}()

		UnboxAnyAsInt(uint(math.MaxInt + 1))
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

		UnboxAnyAsInt(any(str))
	})
}

//goland:noinspection GoRedundantConversion
func TestUnboxAnyAsFloat64OrInt64OrInt64(t *testing.T) {
	x64 := math.MaxInt > math.MaxInt32

	t.Run("int64 or float64 depends value", func(t *testing.T) {
		var vi int64
		var vf float64
		var dt UnboxFloat64DataType
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int8(math.MinInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int8(math.MaxInt8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint8(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint8(math.MaxUint8))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint8), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int16(math.MinInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int16(math.MaxInt16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint16(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint16(math.MaxUint16))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint16), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int32(math.MinInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int32(math.MaxInt32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint32(0))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint32(math.MaxUint32))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxUint32), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int64(math.MinInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MinInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(int64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint64(0))
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint64(math.MaxInt64))
		assert.Equal(t, float64(0.0), vf)
		assert.Equal(t, int64(math.MaxInt64), vi)
		assert.Equal(t, dt, UF64_TYPE_INT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint64(math.MaxUint64))
		assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		if x64 {
			//goland:noinspection GoRedundantConversion
			vi, vf, dt = UnboxAnyAsInt64OrFloat64(int(math.MinInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MinInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			//goland:noinspection GoRedundantConversion
			vi, vf, dt = UnboxAnyAsInt64OrFloat64(int(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint(0))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint(math.MaxInt64))
			assert.Equal(t, float64(0.0), vf)
			assert.Equal(t, int64(math.MaxInt64), vi)
			assert.Equal(t, dt, UF64_TYPE_INT64)
			vi, vf, dt = UnboxAnyAsInt64OrFloat64(uint(math.MaxUint64))
			assert.Greater(t, 0.0001, math.Abs(float64(math.MaxUint64)-vf))
			assert.Equal(t, int64(0), vi)
			assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		}
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(float32(3.3))
		assert.Greater(t, 0.0001, math.Abs(float64(3.3)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
		vi, vf, dt = UnboxAnyAsInt64OrFloat64(float64(9.9))
		assert.Greater(t, 0.0001, math.Abs(float64(9.9)-vf))
		assert.Equal(t, int64(0), vi)
		assert.Equal(t, dt, UF64_TYPE_FLOAT64)
	})

	t.Run("nil returns zero int64", func(t *testing.T) {
		var pi64 *float64
		vi, vf, dt := UnboxAnyAsInt64OrFloat64(pi64)
		assert.Zero(t, vi)
		assert.Zero(t, vf)
		assert.Equal(t, UF64_TYPE_INT64, dt)
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

		UnboxAnyAsInt64OrFloat64(str)
	})
}

func Test_enumerable_unboxAnyAsX(t *testing.T) {
	t.Run("unbox any byte & pointer", func(t *testing.T) {
		var b int8 = 3
		var ab = any(b)
		assert.Equal(t, byte(b), UnboxAnyAsByte(b))
		assert.Equal(t, byte(b), UnboxAnyAsByte(ab))

		var pb *int8
		assert.Equal(t, 0, int(UnboxAnyAsByte(pb)))

		assert.Equal(t, byte(b), UnboxAnyAsByte(&b))
		assert.Equal(t, byte(b), UnboxAnyAsByte(&ab))

		var pab any = any(&ab)
		assert.Equal(t, byte(b), UnboxAnyAsByte(pab))
		assert.Equal(t, byte(b), UnboxAnyAsByte(&pab))

		var papab any = any(&pab)
		assert.Equal(t, byte(b), UnboxAnyAsByte(papab))
		assert.Equal(t, byte(b), UnboxAnyAsByte(&papab))

		var pstr *string
		assert.Equal(t, 0, int(UnboxAnyAsByte(pstr)))
	})

	t.Run("unbox any int32 & pointer", func(t *testing.T) {
		var b int32 = 3
		var ab = any(b)
		assert.Equal(t, b, UnboxAnyAsInt32(b))
		assert.Equal(t, b, UnboxAnyAsInt32(ab))

		var pb *int32
		assert.Equal(t, 0, int(UnboxAnyAsInt32(pb)))

		assert.Equal(t, b, UnboxAnyAsInt32(&b))
		assert.Equal(t, b, UnboxAnyAsInt32(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, UnboxAnyAsInt32(pab))
		assert.Equal(t, b, UnboxAnyAsInt32(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, UnboxAnyAsInt32(papab))
		assert.Equal(t, b, UnboxAnyAsInt32(&papab))
	})

	t.Run("unbox any int64 & pointer", func(t *testing.T) {
		var b int64 = 3
		var ab = any(b)
		assert.Equal(t, b, UnboxAnyAsInt64(b))
		assert.Equal(t, b, UnboxAnyAsInt64(ab))

		var pb *int64
		assert.Equal(t, 0, int(UnboxAnyAsInt64(pb)))

		assert.Equal(t, b, UnboxAnyAsInt64(&b))
		assert.Equal(t, b, UnboxAnyAsInt64(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, UnboxAnyAsInt64(pab))
		assert.Equal(t, b, UnboxAnyAsInt64(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, UnboxAnyAsInt64(papab))
		assert.Equal(t, b, UnboxAnyAsInt64(&papab))
	})

	t.Run("unbox any int & pointer", func(t *testing.T) {
		var b int = 3
		var ab = any(b)
		assert.Equal(t, b, UnboxAnyAsInt(b))
		assert.Equal(t, b, UnboxAnyAsInt(ab))

		var pb *int
		assert.Equal(t, 0, int(UnboxAnyAsInt(pb)))

		assert.Equal(t, b, UnboxAnyAsInt(&b))
		assert.Equal(t, b, UnboxAnyAsInt(&ab))

		var pab any = any(&ab)
		assert.Equal(t, b, UnboxAnyAsInt(pab))
		assert.Equal(t, b, UnboxAnyAsInt(&pab))

		var papab any = any(&pab)
		assert.Equal(t, b, UnboxAnyAsInt(papab))
		assert.Equal(t, b, UnboxAnyAsInt(&papab))
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

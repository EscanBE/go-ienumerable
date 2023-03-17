package reflection

import (
	"fmt"
	"math"
	"reflect"
)

// UnboxAnyAsInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int64.
//
// Panic if value is over range or not integer (nil value of any type are treated as zero)
func UnboxAnyAsInt64(v any) int64 {
	result, state := TryUnboxAnyAsInt64(v)
	if state == UNBOX_SUCCESS {
		return result
	}

	if state == UNBOX_NIL {
		return 0
	}

	if state == UNBOX_OVERFLOW {
		panic(makeOverflowError(v, "int64"))
	}

	panic(makeCastError(v, "int64"))
}

// TryUnboxAnyAsInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int64.
func TryUnboxAnyAsInt64(v any) (result int64, state UnboxResultState) {
	state = UNBOX_FAILED

	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			state = UNBOX_NIL
			return
		}

		vo = vo.Elem()
		kind = vo.Kind()

		if kind == reflect.Interface {
			elem := vo.Elem()
			if elem.Kind() == reflect.Ptr {
				vo = elem
				kind = reflect.Ptr
			}
		}
	}

	if kind == reflect.Interface {
		if vo.IsNil() {
			state = UNBOX_NIL
			return
		}

		subElem := vo.Elem()
		subKind := subElem.Kind()
		switch subKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		default:
			return
		}

		vo = subElem
		kind = subKind
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = vo.Int()
		state = UNBOX_SUCCESS
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxInt64 {
			result = int64(uv)
			state = UNBOX_SUCCESS
		} else {
			state = UNBOX_OVERFLOW
		}
		break
	default:
		if !vo.IsValid() {
			state = UNBOX_NIL
			return
		}
		break
	}
	return
}

// UnboxAnyAsInt32 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int32.
//
// Panic if value is over range or not integer (nil value of any type are treated as zero)
func UnboxAnyAsInt32(v any) int32 {
	result, state := TryUnboxAnyAsInt32(v)
	if state == UNBOX_SUCCESS {
		return result
	}

	if state == UNBOX_NIL {
		return 0
	}

	if state == UNBOX_OVERFLOW {
		panic(makeOverflowError(v, "int32"))
	}

	panic(makeCastError(v, "int32"))
}

// TryUnboxAnyAsInt32 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int32.
func TryUnboxAnyAsInt32(v any) (result int32, state UnboxResultState) {
	result64, state64 := TryUnboxAnyAsInt64(v)

	if state64 != UNBOX_SUCCESS {
		state = state64
		return
	}

	if math.MinInt32 > result64 || result64 > math.MaxInt32 {
		state = UNBOX_OVERFLOW
		return
	}

	result = int32(result64)
	state = UNBOX_SUCCESS
	return
}

// UnboxAnyAsInt unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int.
//
// Panic if value is over range or not integer (nil value of any type are treated as zero)
func UnboxAnyAsInt(v any) int {
	result, state := TryUnboxAnyAsInt(v)
	if state == UNBOX_SUCCESS {
		return result
	}

	if state == UNBOX_NIL {
		return 0
	}

	if state == UNBOX_OVERFLOW {
		panic(makeOverflowError(v, "int"))
	}

	panic(makeCastError(v, "int"))
}

// TryUnboxAnyAsInt unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into int.
func TryUnboxAnyAsInt(v any) (result int, state UnboxResultState) {
	result64, state64 := TryUnboxAnyAsInt64(v)

	if state64 != UNBOX_SUCCESS {
		state = state64
		return
	}

	result = int(result64)
	state = UNBOX_SUCCESS
	return
}

// UnboxAnyAsByte unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into byte (uint8, range 0-255).
//
// Panic if value is over range or not integer (nil value of any type are treated as zero)
func UnboxAnyAsByte(v any) byte {
	result, state := TryUnboxAnyAsByte(v)
	if state == UNBOX_SUCCESS {
		return result
	}

	if state == UNBOX_NIL {
		return 0
	}

	if state == UNBOX_OVERFLOW {
		panic(makeOverflowError(v, "byte"))
	}

	panic(makeCastError(v, "byte"))
}

// TryUnboxAnyAsByte unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64), or their pointer,
// or wrapped into any, into byte (uint8, range 0-255).
func TryUnboxAnyAsByte(v any) (result byte, state UnboxResultState) {
	result64, state64 := TryUnboxAnyAsInt64(v)

	if state64 != UNBOX_SUCCESS {
		state = state64
		return
	}

	if 0 > result64 || result64 > math.MaxUint8 {
		state = UNBOX_OVERFLOW
		return
	}

	result = byte(result64)
	state = UNBOX_SUCCESS
	return
}

// UnboxAnyAsInt64OrFloat64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64, float32/64), or their pointer,
// or wrapped into any, into float64 or int64 (priority int64 if integer).
//
// Panic if not integer/float (nil value of any type are treated as zero)
func UnboxAnyAsInt64OrFloat64(v any) (resultInt64 int64, resultFloat64 float64, resultDataType UnboxFloat64DataType) {
	resultInt64, resultFloat64, resultDataType, state := TryUnboxAnyAsInt64OrFloat64(v)
	if state == UNBOX_SUCCESS {
		return
	}

	if state == UNBOX_NIL {
		resultDataType = UF64_TYPE_INT64
		return
	}

	panic(makeCastError(v, "float64"))
}

// TryUnboxAnyAsInt64OrFloat64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64, float32/64), or their pointer,
// or wrapped into any, into float64 or int64 (priority int64 if integer).
func TryUnboxAnyAsInt64OrFloat64(v any) (resultInt64 int64, resultFloat64 float64, resultDataType UnboxFloat64DataType, state UnboxResultState) {
	resultDataType = UF64_TYPE_FAILED
	state = UNBOX_FAILED

	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			resultDataType = UF64_TYPE_NIL
			state = UNBOX_NIL
			return
		}

		vo = vo.Elem()
		kind = vo.Kind()

		if kind == reflect.Interface {
			elem := vo.Elem()
			if elem.Kind() == reflect.Ptr {
				vo = elem
				kind = reflect.Ptr
			}
		}
	}

	if kind == reflect.Interface {
		if vo.IsNil() {
			resultDataType = UF64_TYPE_NIL
			state = UNBOX_NIL
			return
		}

		subElem := vo.Elem()
		subKind := subElem.Kind()
		switch subKind {
		case reflect.Float32, reflect.Float64:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		default:
			return
		}

		vo = subElem
		kind = subKind
	}

	switch kind {
	case reflect.Float32, reflect.Float64:
		resultFloat64 = vo.Float()
		resultDataType = UF64_TYPE_FLOAT64
		state = UNBOX_SUCCESS
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		resultInt64 = vo.Int()
		resultDataType = UF64_TYPE_INT64
		state = UNBOX_SUCCESS
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxInt64 {
			resultInt64 = int64(uv)
			resultDataType = UF64_TYPE_INT64
			state = UNBOX_SUCCESS
		} else {
			resultFloat64 = float64(uv)
			resultDataType = UF64_TYPE_FLOAT64
			state = UNBOX_SUCCESS
		}
		break
	default:
		if !vo.IsValid() {
			resultDataType = UF64_TYPE_NIL
			state = UNBOX_NIL
			return
		}
		break
	}
	return
}

func makeOverflowError(v any, t string) error {
	return fmt.Errorf("value %v of type %T is over range of %s", v, v, t)
}

func makeCastError(v any, t string) error {
	return fmt.Errorf("value %v of type %T cannot be casted to %s", v, v, t)
}

package goe

import (
	"math"
	"reflect"
)

// unboxAnyAsByte unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into byte.
//
// Panic if value is over range or not integer
func (src *enumerable[T]) unboxAnyAsByte(v T) byte {
	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			return 0
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
		subKind := vo.Elem().Kind()
		switch subKind {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		default:
			panic(makeCastError(v, "byte"))
		}

		vo = vo.Elem()
		kind = vo.Kind()
	}

	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxUint8 {
			return byte(uv)
		}
		panic(makeCastError(v, "byte"))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iv := vo.Int()
		if 0 <= iv && iv <= math.MaxUint8 {
			return byte(iv)
		}
		panic(makeCastError(v, "byte"))
	default:
		panic(makeCastError(v, "byte"))
	}
}

// unboxAnyAsInt32 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int32.
//
// Panic if value is over range or not integer
func (src *enumerable[T]) unboxAnyAsInt32(v T) int32 {
	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			return 0
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
		subKind := vo.Elem().Kind()
		switch subKind {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		default:
			panic(makeCastError(v, "int32"))
		}

		vo = vo.Elem()
		kind = vo.Kind()
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iv := vo.Int()
		if math.MinInt32 <= iv && iv <= math.MaxInt32 {
			return int32(iv)
		}
		panic(makeCastError(v, "int32"))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxInt32 {
			return int32(uv)
		}
		panic(makeCastError(v, "int32"))
	default:
		panic(makeCastError(v, "int32"))
	}
}

// unboxAnyAsInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int64.
//
// Panic if value is over range or not integer
func (src *enumerable[T]) unboxAnyAsInt64(v T) int64 {
	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			return 0
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
		subKind := vo.Elem().Kind()
		switch subKind {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		default:
			panic(makeCastError(v, "int64"))
		}

		vo = vo.Elem()
		kind = vo.Kind()
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vo.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxInt64 {
			return int64(uv)
		}
		panic(makeCastError(v, "int64"))
	default:
		panic(makeCastError(v, "int64"))
	}
}

// unboxAnyAsInt unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int.
//
// Panic if value is over range or not integer
func (src *enumerable[T]) unboxAnyAsInt(v T) int {
	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			return 0
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
		subKind := vo.Elem().Kind()
		switch subKind {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		default:
			panic(makeCastError(v, "int"))
		}

		vo = vo.Elem()
		kind = vo.Kind()
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(vo.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uv := vo.Uint()
		if uv <= math.MaxInt {
			return int(uv)
		}
		panic(makeCastError(v, "int"))
	default:
		panic(makeCastError(v, "int"))
	}
}

type unboxFloat64DataType byte

//goland:noinspection GoSnakeCaseUsage
const (
	UF64_TYPE_FLOAT64 unboxFloat64DataType = 1
	UF64_TYPE_INT64   unboxFloat64DataType = 2
)

// unboxAnyAsFloat64OrInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) or float32/64
// into either int64 or float64 value (priority int64), data type specified in result.
// This design is for sum accuracy.
//
// Panic if neither integer nor float
func (src *enumerable[T]) unboxAnyAsFloat64OrInt64(v T) (rf float64, ri int64, dt unboxFloat64DataType) {
	vo := reflect.ValueOf(v)

	kind := vo.Kind()
	for kind == reflect.Ptr {
		if vo.IsNil() {
			dt = UF64_TYPE_INT64
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
		subKind := vo.Elem().Kind()
		switch subKind {
		case reflect.Float32, reflect.Float64:
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			break
		default:
			panic(makeCastError(v, "float64"))
		}

		vo = vo.Elem()
		kind = vo.Kind()
	}

	switch kind {
	case reflect.Float32, reflect.Float64:
		rf = vo.Float()
		dt = UF64_TYPE_FLOAT64
		return
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ri = vo.Int()
		dt = UF64_TYPE_INT64
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		rui := vo.Uint()
		if rui <= math.MaxInt64 {
			ri = int64(rui)
			dt = UF64_TYPE_INT64
		} else {
			rf = float64(rui)
			dt = UF64_TYPE_FLOAT64
		}
		return
	default:
		panic(makeCastError(v, "float64"))
	}
}

package go_ienumerable

import "fmt"

func (src *enumerable[T]) UnboxInt8() IEnumerable[int8] {
	size := len(src.data)
	result := make([]int8, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(int8); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "int8"))
			}
		}
	}

	return NewIEnumerable[int8](result...)
}

func (src *enumerable[T]) UnboxUInt8() IEnumerable[uint8] {
	size := len(src.data)
	result := make([]uint8, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uint8); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uint8"))
			}
		}
	}

	return NewIEnumerable[uint8](result...)
}

func (src *enumerable[T]) UnboxInt16() IEnumerable[int16] {
	size := len(src.data)
	result := make([]int16, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(int16); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "int16"))
			}
		}
	}

	return NewIEnumerable[int16](result...)
}

func (src *enumerable[T]) UnboxUInt16() IEnumerable[uint16] {
	size := len(src.data)
	result := make([]uint16, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uint16); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uint16"))
			}
		}
	}

	return NewIEnumerable[uint16](result...)
}

func (src *enumerable[T]) UnboxInt32() IEnumerable[int32] {
	size := len(src.data)
	result := make([]int32, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(int32); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "int32"))
			}
		}
	}

	return NewIEnumerable[int32](result...)
}

func (src *enumerable[T]) UnboxUInt32() IEnumerable[uint32] {
	size := len(src.data)
	result := make([]uint32, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uint32); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uint32"))
			}
		}
	}

	return NewIEnumerable[uint32](result...)
}

func (src *enumerable[T]) UnboxInt64() IEnumerable[int64] {
	size := len(src.data)
	result := make([]int64, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(int64); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "int64"))
			}
		}
	}

	return NewIEnumerable[int64](result...)
}

func (src *enumerable[T]) UnboxUInt64() IEnumerable[uint64] {
	size := len(src.data)
	result := make([]uint64, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uint64); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uint64"))
			}
		}
	}

	return NewIEnumerable[uint64](result...)
}

func (src *enumerable[T]) UnboxInt() IEnumerable[int] {
	size := len(src.data)
	result := make([]int, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(int); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "int"))
			}
		}
	}

	return NewIEnumerable[int](result...)
}

func (src *enumerable[T]) UnboxUInt() IEnumerable[uint] {
	size := len(src.data)
	result := make([]uint, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uint); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uint"))
			}
		}
	}

	return NewIEnumerable[uint](result...)
}

func (src *enumerable[T]) UnboxUIntptr() IEnumerable[uintptr] {
	size := len(src.data)
	result := make([]uintptr, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(uintptr); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "uintptr"))
			}
		}
	}

	return NewIEnumerable[uintptr](result...)
}

func (src *enumerable[T]) UnboxFloat32() IEnumerable[float32] {
	size := len(src.data)
	result := make([]float32, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(float32); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "float32"))
			}
		}
	}

	return NewIEnumerable[float32](result...)
}

func (src *enumerable[T]) UnboxFloat64() IEnumerable[float64] {
	size := len(src.data)
	result := make([]float64, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(float64); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "float64"))
			}
		}
	}

	return NewIEnumerable[float64](result...)
}

func (src *enumerable[T]) UnboxComplex64() IEnumerable[complex64] {
	size := len(src.data)
	result := make([]complex64, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(complex64); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "complex64"))
			}
		}
	}

	return NewIEnumerable[complex64](result...)
}

func (src *enumerable[T]) UnboxComplex128() IEnumerable[complex128] {
	size := len(src.data)
	result := make([]complex128, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(complex128); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "complex128"))
			}
		}
	}

	return NewIEnumerable[complex128](result...)
}

func (src *enumerable[T]) UnboxString() IEnumerable[string] {
	size := len(src.data)
	result := make([]string, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(string); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "string"))
			}
		}
	}

	return NewIEnumerable[string](result...)
}

func (src *enumerable[T]) UnboxBool() IEnumerable[bool] {
	size := len(src.data)
	result := make([]bool, size)
	if size > 0 {
		for i, v := range src.data {
			if cast, ok := any(v).(bool); ok {
				result[i] = cast
			} else {
				panic(makeCastError(v, "bool"))
			}
		}
	}

	return NewIEnumerable[bool](result...)
}

func makeCastError(v any, t string) error {
	return fmt.Errorf("value %v of type %T cannot be casted to %s", v, v, t)
}

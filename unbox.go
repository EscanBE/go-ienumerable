package go_ienumerable

import "fmt"

func (src *enumerable[T]) UnboxInt8() IEnumerable[int8] {
	size := len(src.data)
	result := make([]int8, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(int8); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to int8", t))
			}
		}
	}

	return NewIEnumerable[int8](result...)
}

func (src *enumerable[T]) UnboxUInt8() IEnumerable[uint8] {
	size := len(src.data)
	result := make([]uint8, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uint8); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uint8", t))
			}
		}
	}

	return NewIEnumerable[uint8](result...)
}

func (src *enumerable[T]) UnboxInt16() IEnumerable[int16] {
	size := len(src.data)
	result := make([]int16, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(int16); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to int16", t))
			}
		}
	}

	return NewIEnumerable[int16](result...)
}

func (src *enumerable[T]) UnboxUInt16() IEnumerable[uint16] {
	size := len(src.data)
	result := make([]uint16, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uint16); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uint16", t))
			}
		}
	}

	return NewIEnumerable[uint16](result...)
}

func (src *enumerable[T]) UnboxInt32() IEnumerable[int32] {
	size := len(src.data)
	result := make([]int32, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(int32); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to int32", t))
			}
		}
	}

	return NewIEnumerable[int32](result...)
}

func (src *enumerable[T]) UnboxUInt32() IEnumerable[uint32] {
	size := len(src.data)
	result := make([]uint32, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uint32); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uint32", t))
			}
		}
	}

	return NewIEnumerable[uint32](result...)
}

func (src *enumerable[T]) UnboxInt64() IEnumerable[int64] {
	size := len(src.data)
	result := make([]int64, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(int64); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to int64", t))
			}
		}
	}

	return NewIEnumerable[int64](result...)
}

func (src *enumerable[T]) UnboxUInt64() IEnumerable[uint64] {
	size := len(src.data)
	result := make([]uint64, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uint64); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uint64", t))
			}
		}
	}

	return NewIEnumerable[uint64](result...)
}

func (src *enumerable[T]) UnboxInt() IEnumerable[int] {
	size := len(src.data)
	result := make([]int, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(int); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to int", t))
			}
		}
	}

	return NewIEnumerable[int](result...)
}

func (src *enumerable[T]) UnboxUInt() IEnumerable[uint] {
	size := len(src.data)
	result := make([]uint, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uint); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uint", t))
			}
		}
	}

	return NewIEnumerable[uint](result...)
}

func (src *enumerable[T]) UnboxUIntptr() IEnumerable[uintptr] {
	size := len(src.data)
	result := make([]uintptr, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(uintptr); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to uintptr", t))
			}
		}
	}

	return NewIEnumerable[uintptr](result...)
}

func (src *enumerable[T]) UnboxFloat32() IEnumerable[float32] {
	size := len(src.data)
	result := make([]float32, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(float32); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to float32", t))
			}
		}
	}

	return NewIEnumerable[float32](result...)
}

func (src *enumerable[T]) UnboxFloat64() IEnumerable[float64] {
	size := len(src.data)
	result := make([]float64, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(float64); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to float64", t))
			}
		}
	}

	return NewIEnumerable[float64](result...)
}

func (src *enumerable[T]) UnboxComplex64() IEnumerable[complex64] {
	size := len(src.data)
	result := make([]complex64, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(complex64); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to complex64", t))
			}
		}
	}

	return NewIEnumerable[complex64](result...)
}

func (src *enumerable[T]) UnboxComplex128() IEnumerable[complex128] {
	size := len(src.data)
	result := make([]complex128, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(complex128); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to complex128", t))
			}
		}
	}

	return NewIEnumerable[complex128](result...)
}

func (src *enumerable[T]) UnboxString() IEnumerable[string] {
	size := len(src.data)
	result := make([]string, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(string); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to string", t))
			}
		}
	}

	return NewIEnumerable[string](result...)
}

func (src *enumerable[T]) UnboxBool() IEnumerable[bool] {
	size := len(src.data)
	result := make([]bool, size)
	if size > 0 {
		for i, t := range src.data {
			if cast, ok := any(t).(bool); ok {
				result[i] = cast
			} else {
				panic(fmt.Errorf("%v can not be casted to bool", t))
			}
		}
	}

	return NewIEnumerable[bool](result...)
}

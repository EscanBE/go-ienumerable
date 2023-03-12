package go_ienumerable

import "fmt"

func (src *enumerable[T]) UnboxByte() IEnumerable[byte] {
	size := len(src.data)
	result := make([]byte, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = src.unboxAnyAsByte(v)
		}
	}

	return NewIEnumerable[byte](result...)
}

func (src *enumerable[T]) UnboxInt32() IEnumerable[int32] {
	size := len(src.data)
	result := make([]int32, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = src.unboxAnyAsInt32(v)
		}
	}

	return NewIEnumerable[int32](result...)
}

func (src *enumerable[T]) UnboxInt64() IEnumerable[int64] {
	size := len(src.data)
	result := make([]int64, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = src.unboxAnyAsInt64(v)
		}
	}

	return NewIEnumerable[int64](result...)
}

func (src *enumerable[T]) UnboxInt() IEnumerable[int] {
	size := len(src.data)
	result := make([]int, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = src.unboxAnyAsInt(v)
		}
	}

	return NewIEnumerable[int](result...)
}

func (src *enumerable[T]) UnboxFloat64() IEnumerable[float64] {
	size := len(src.data)
	result := make([]float64, size)
	if size > 0 {
		for i, v := range src.data {
			vf, vi, dt := src.unboxAnyAsFloat64OrInt64(v)
			if dt == UF64_TYPE_FLOAT64 {
				result[i] = vf
			} else if dt == UF64_TYPE_INT64 {
				result[i] = float64(vi)
			}
		}
	}

	return NewIEnumerable[float64](result...)
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

func makeCastError2(v any, t, et string) error {
	return fmt.Errorf("value %v of type %T (expect: %s) cannot be casted to %s", v, v, et, t)
}

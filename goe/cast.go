package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
)

func (src *enumerable[T]) CastByte() IEnumerable[byte] {
	size := len(src.data)
	result := make([]byte, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = reflection.UnboxAnyAsByte(v)
		}
	}

	return NewIEnumerable[byte](result...)
}

func (src *enumerable[T]) CastInt32() IEnumerable[int32] {
	size := len(src.data)
	result := make([]int32, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = reflection.UnboxAnyAsInt32(v)
		}
	}

	return NewIEnumerable[int32](result...)
}

func (src *enumerable[T]) CastInt64() IEnumerable[int64] {
	size := len(src.data)
	result := make([]int64, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = reflection.UnboxAnyAsInt64(v)
		}
	}

	return NewIEnumerable[int64](result...)
}

func (src *enumerable[T]) CastInt() IEnumerable[int] {
	size := len(src.data)
	result := make([]int, size)
	if size > 0 {
		for i, v := range src.data {
			result[i] = reflection.UnboxAnyAsInt(v)
		}
	}

	return NewIEnumerable[int](result...)
}

func (src *enumerable[T]) CastFloat64() IEnumerable[float64] {
	size := len(src.data)
	result := make([]float64, size)
	if size > 0 {
		for i, v := range src.data {
			vi, vf, dt := reflection.UnboxAnyAsInt64OrFloat64(v)
			if dt == reflection.UF64_TYPE_FLOAT64 {
				result[i] = vf
			} else if dt == reflection.UF64_TYPE_INT64 {
				result[i] = float64(vi)
			}
		}
	}

	return NewIEnumerable[float64](result...)
}
func (src *enumerable[T]) CastString() IEnumerable[string] {
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

func (src *enumerable[T]) CastBool() IEnumerable[bool] {
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

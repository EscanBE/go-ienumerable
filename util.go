package go_ienumerable

import (
	"fmt"
	"math"
	"math/big"
)

func (src *enumerable[T]) exposeData() []T {
	return src.data
}

func (src *enumerable[T]) exposeDataType() string {
	return src.dataType
}

func (src *enumerable[T]) len() int {
	return len(src.data)
}

func (src *enumerable[T]) assertSrcNonNil() {
	if src == nil {
		panic("source is nil")
	}
}

func (src *enumerable[T]) assertSrcNonEmpty() {
	if len(src.data) < 1 {
		panic(getErrorSrcContainsNoElement())
	}
}

func getErrorSrcContainsNoElement() error {
	return fmt.Errorf("source contains no element")
}

func (_ *enumerable[T]) assertSizeGt0(size int) {
	if size < 1 {
		panic("size is below 1")
	}
}

func (src *enumerable[T]) assertIndex(index int) {
	if 0 > index || index >= len(src.data) {
		panic("index out of bound")
	}
}

func (col *enumerator[T]) assertCollectionNonNil() {
	if col == nil {
		panic("collection is nil")
	}
}

func (_ *enumerable[T]) assertSecondIEnumerableNonNil(second IEnumerable[T]) {
	if second == nil {
		panic("second IEnumerable is nil")
	}
}

func (src *enumerable[T]) assertPredicateNonNil(predicate func(T) bool) {
	if predicate == nil {
		panic(getErrorNilPredicate())
	}
}

func getErrorNilPredicate() error {
	return fmt.Errorf("predicate is nil")
}

func (src *enumerable[T]) assertComparerNonNil(comparer func(T, T) bool) {
	if comparer == nil {
		panic(getErrorNilComparer())
	}
}

func getErrorNilComparer() error {
	return fmt.Errorf("comparer is nil")
}

func (src *enumerable[T]) assertSelectorNonNil(selector func(T) any) {
	if selector == nil {
		panic(getErrorNilSelector())
	}
}

func (src *enumerable[T]) assertArraySelectorNonNil(selector func(T) []any) {
	if selector == nil {
		panic(getErrorNilSelector())
	}
}

func getErrorNilSelector() error {
	return fmt.Errorf("selector is nil")
}

func (src *enumerable[T]) assertAggregateFuncNonNil(f func(T, T) T) {
	if f == nil {
		panic(getErrorNilAggregateFunc())
	}
}

func (src *enumerable[T]) assertAggregateAnySeedFuncNonNil(f func(any, T) any) {
	if f == nil {
		panic(getErrorNilAggregateFunc())
	}
}

func getErrorNilAggregateFunc() error {
	return fmt.Errorf("aggregate function is nil")
}

func getErrorMoreThanOne() error {
	return fmt.Errorf("more than one element")
}

func getErrorMoreThanOneMatch() error {
	return fmt.Errorf("more than one element satisfies the condition in predicate")
}

func getErrorNoMatch() error {
	return fmt.Errorf("no element satisfies the condition in predicate")
}

func (src *enumerable[T]) copyExceptData() *enumerable[T] {
	if src == nil {
		return nil
	}
	return &enumerable[T]{
		data:             nil,
		dataType:         src.dataType,
		equalityComparer: src.equalityComparer,
		lessComparer:     src.lessComparer,
	}
}

func (src *enumerable[T]) withData(data []T) *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = data
	return src
}

func (src *enumerable[T]) withEmptyData() *enumerable[T] {
	if src == nil {
		return nil
	}
	src.data = make([]T, 0)
	return src
}

func copySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	if src != nil {
		copy(dst, src)
	}
	return dst
}

func (src *enumerable[T]) unboxAnyAsByte(v T) byte {
	if src.dataType == "" {
		// fall through
	} else if src.dataType == "int" {
		if vi, oki := any(v).(int); oki && 0 <= vi && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "int64" {
		if vi, oki := any(v).(int64); oki && 0 <= vi && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "int8" {
		if vi, oki := any(v).(int8); oki && 0 <= vi {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "int32" {
		if vi, oki := any(v).(int32); oki && 0 <= vi && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "int16" {
		if vi, oki := any(v).(int16); oki && 0 <= vi && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "uint" {
		if vi, oki := any(v).(uint); oki && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "uint64" {
		if vi, oki := any(v).(uint64); oki && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "uint32" {
		if vi, oki := any(v).(uint32); oki && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else if src.dataType == "uint8" {
		return any(v).(uint8)
	} else if src.dataType == "uint16" {
		if vi, oki := any(v).(uint16); oki && vi <= math.MaxUint8 {
			return byte(vi)
		}
		panic(makeCastError2(v, "byte", src.dataType))
	} else {
		panic(makeCastError(v, "byte"))
	}

	if vi, oki := any(v).(int); oki {
		if 0 > vi || vi > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(vi)
	} else if v64, ok64 := any(v).(int64); ok64 {
		if 0 > v64 || v64 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(v64)
	} else if v8, ok8 := any(v).(int8); ok8 {
		if 0 > v8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(v8)
	} else if v32, ok32 := any(v).(int32); ok32 {
		if 0 > v32 || v32 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(v32)
	} else if v16, ok16 := any(v).(int16); ok16 {
		if 0 > v16 || v16 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(v16)
	} else if vui, okui := any(v).(uint); okui {
		if vui > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(vui)
	} else if vu64, oku64 := any(v).(uint64); oku64 {
		if vu64 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(vu64)
	} else if vu32, oku32 := any(v).(uint32); oku32 {
		if vu32 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(vu32)
	} else if vu8, oku8 := any(v).(uint8); oku8 {
		return vu8
	} else if vu16, oku16 := any(v).(uint16); oku16 {
		if vu16 > math.MaxUint8 {
			panic(makeCastError(v, "byte"))
		}
		return byte(vu16)
	} else {
		panic(makeCastError(v, "byte"))
	}
}

func (src *enumerable[T]) unboxAnyAsInt32(v T) int32 {
	if src.dataType == "" {
		// fall through
	} else if src.dataType == "int" {
		if vi, oki := any(v).(int); oki && math.MinInt32 <= vi && vi <= math.MaxInt32 {
			return int32(vi)
		}
		panic(makeCastError2(v, "int32", src.dataType))
	} else if src.dataType == "int64" {
		if vi, oki := any(v).(int64); oki && math.MinInt32 <= vi && vi <= math.MaxInt32 {
			return int32(vi)
		}
		panic(makeCastError2(v, "int32", src.dataType))
	} else if src.dataType == "int8" {
		return int32(any(v).(int8))
	} else if src.dataType == "int32" {
		return any(v).(int32)
	} else if src.dataType == "int16" {
		return int32(any(v).(int16))
	} else if src.dataType == "uint" {
		if vi, oki := any(v).(uint); oki && vi <= math.MaxInt32 {
			return int32(vi)
		}
		panic(makeCastError2(v, "int32", src.dataType))
	} else if src.dataType == "uint64" {
		if vi, oki := any(v).(uint64); oki && vi <= math.MaxInt32 {
			return int32(vi)
		}
		panic(makeCastError2(v, "int32", src.dataType))
	} else if src.dataType == "uint32" {
		if vi, oki := any(v).(uint32); oki && vi <= math.MaxInt32 {
			return int32(vi)
		}
		panic(makeCastError2(v, "int32", src.dataType))
	} else if src.dataType == "uint8" {
		return int32(any(v).(uint8))
	} else if src.dataType == "uint16" {
		return int32(any(v).(uint16))
	} else {
		panic(makeCastError(v, "int32"))
	}

	if vi, oki := any(v).(int); oki {
		if math.MinInt32 > vi || vi > math.MaxInt32 {
			panic(makeCastError(v, "int32"))
		}
		return int32(vi)
	} else if v64, ok64 := any(v).(int64); ok64 {
		if math.MinInt32 > v64 || v64 > math.MaxInt32 {
			panic(makeCastError(v, "int32"))
		}
		return int32(v64)
	} else if v8, ok8 := any(v).(int8); ok8 {
		return int32(v8)
	} else if v32, ok32 := any(v).(int32); ok32 {
		return v32
	} else if v16, ok16 := any(v).(int16); ok16 {
		return int32(v16)
	} else if vui, okui := any(v).(uint); okui {
		if vui > math.MaxInt32 {
			panic(makeCastError(v, "int32"))
		}
		return int32(vui)
	} else if vu64, oku64 := any(v).(uint64); oku64 {
		if vu64 > math.MaxInt32 {
			panic(makeCastError(v, "int32"))
		}
		return int32(vu64)
	} else if vu32, oku32 := any(v).(uint32); oku32 {
		if vu32 > math.MaxInt32 {
			panic(makeCastError(v, "int32"))
		}
		return int32(vu32)
	} else if vu8, oku8 := any(v).(uint8); oku8 {
		return int32(vu8)
	} else if vu16, oku16 := any(v).(uint16); oku16 {
		return int32(vu16)
	} else {
		panic(makeCastError(v, "int32"))
	}
}

func (src *enumerable[T]) unboxAnyAsInt64(v T) int64 {
	if src.dataType == "" {
		// fall through
	} else if src.dataType == "int" {
		return int64(any(v).(int))
	} else if src.dataType == "int64" {
		return any(v).(int64)
	} else if src.dataType == "int8" {
		return int64(any(v).(int8))
	} else if src.dataType == "int32" {
		return int64(any(v).(int32))
	} else if src.dataType == "int16" {
		return int64(any(v).(int16))
	} else if src.dataType == "uint" {
		if vi, oki := any(v).(uint); oki && vi <= math.MaxInt64 {
			return int64(vi)
		}
		panic(makeCastError2(v, "int64", src.dataType))
	} else if src.dataType == "uint64" {
		if vi, oki := any(v).(uint64); oki && vi <= math.MaxInt64 {
			return int64(vi)
		}
		panic(makeCastError2(v, "int64", src.dataType))
	} else if src.dataType == "uint32" {
		return int64(any(v).(uint32))
	} else if src.dataType == "uint8" {
		return int64(any(v).(uint8))
	} else if src.dataType == "uint16" {
		return int64(any(v).(uint16))
	} else {
		panic(makeCastError(v, "int64"))
	}

	if vi, oki := any(v).(int); oki {
		return int64(vi)
	} else if v64, ok64 := any(v).(int64); ok64 {
		return v64
	} else if v8, ok8 := any(v).(int8); ok8 {
		return int64(v8)
	} else if v32, ok32 := any(v).(int32); ok32 {
		return int64(v32)
	} else if v16, ok16 := any(v).(int16); ok16 {
		return int64(v16)
	} else if vui, okui := any(v).(uint); okui {
		if vui > math.MaxInt64 {
			panic(makeCastError(v, "int64"))
		}
		return int64(vui)
	} else if vu64, oku64 := any(v).(uint64); oku64 {
		if vu64 > math.MaxInt64 {
			panic(makeCastError(v, "int64"))
		}
		return int64(vu64)
	} else if vu32, oku32 := any(v).(uint32); oku32 {
		return int64(vu32)
	} else if vu8, oku8 := any(v).(uint8); oku8 {
		return int64(vu8)
	} else if vu16, oku16 := any(v).(uint16); oku16 {
		return int64(vu16)
	} else {
		panic(makeCastError(v, "int64"))
	}
}

func (src *enumerable[T]) unboxAnyAsInt(v T) int {
	if src.dataType == "" {
		// fall through
	} else if src.dataType == "int" {
		return any(v).(int)
	} else if src.dataType == "int64" {
		//if vi, oki := any(v).(int64); oki && math.MinInt <= vi && vi <= math.MaxInt {
		//	return int(vi)
		//}
		//panic(makeCastError2(v, "int", src.dataType))
		return int(any(v).(int64))
	} else if src.dataType == "int8" {
		return int(any(v).(int8))
	} else if src.dataType == "int32" {
		return int(any(v).(int32))
	} else if src.dataType == "int16" {
		return int(any(v).(int16))
	} else if src.dataType == "uint" {
		if vi, oki := any(v).(uint); oki && vi <= math.MaxInt {
			return int(vi)
		}
		panic(makeCastError2(v, "int", src.dataType))
	} else if src.dataType == "uint64" {
		if vi, oki := any(v).(uint64); oki && vi <= math.MaxInt {
			return int(vi)
		}
		panic(makeCastError2(v, "int", src.dataType))
	} else if src.dataType == "uint32" {
		//if vi, oki := any(v).(uint32); oki && uint64(vi) <= uint64(math.MaxInt) {
		//	return int(vi)
		//}
		//panic(makeCastError2(v, "int", src.dataType))
		return int(any(v).(uint32))
	} else if src.dataType == "uint8" {
		return int(any(v).(uint8))
	} else if src.dataType == "uint16" {
		return int(any(v).(uint16))
	} else {
		panic(makeCastError(v, "int"))
	}

	if vi, oki := any(v).(int); oki {
		return vi
	} else if v64, ok64 := any(v).(int64); ok64 {
		//if math.MinInt > v64 || v64 > math.MaxInt {
		//	panic(makeCastError(v, "int"))
		//}
		return int(v64)
	} else if v8, ok8 := any(v).(int8); ok8 {
		return int(v8)
	} else if v32, ok32 := any(v).(int32); ok32 {
		return int(v32)
	} else if v16, ok16 := any(v).(int16); ok16 {
		return int(v16)
	} else if vui, okui := any(v).(uint); okui {
		if vui > math.MaxInt {
			panic(makeCastError(v, "int"))
		}
		return int(vui)
	} else if vu64, oku64 := any(v).(uint64); oku64 {
		if vu64 > math.MaxInt {
			panic(makeCastError(v, "int"))
		}
		return int(vu64)
	} else if vu32, oku32 := any(v).(uint32); oku32 {
		//if uint64(vu32) > uint64(math.MaxInt) {
		//	panic(makeCastError(v, "int"))
		//}
		return int(vu32)
	} else if vu8, oku8 := any(v).(uint8); oku8 {
		return int(vu8)
	} else if vu16, oku16 := any(v).(uint16); oku16 {
		return int(vu16)
	} else {
		panic(makeCastError(v, "int"))
	}
}

type unboxFloat64DataType byte

//goland:noinspection GoSnakeCaseUsage
const (
	UF64_TYPE_FLOAT64 unboxFloat64DataType = 1
	UF64_TYPE_INT64   unboxFloat64DataType = 2
)

func (src *enumerable[T]) unboxAnyAsFloat64OrInt64(v T) (rf float64, ri int64, dt unboxFloat64DataType) {
	if src.dataType == "" {
		// fall through
	} else if src.dataType == "int" {
		ri = int64(any(v).(int))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "int64" {
		ri = any(v).(int64)
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "int8" {
		ri = int64(any(v).(int8))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "int32" {
		ri = int64(any(v).(int32))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "int16" {
		ri = int64(any(v).(int16))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "uint" {
		vi := any(v).(uint)
		if vi <= math.MaxInt64 {
			ri = int64(vi)
			dt = UF64_TYPE_INT64
			return
		} else {
			rf = float64(vi)
			dt = UF64_TYPE_FLOAT64
			return
		}
	} else if src.dataType == "uint64" {
		vi := any(v).(uint64)
		if vi <= math.MaxInt64 {
			ri = int64(vi)
			dt = UF64_TYPE_INT64
			return
		} else {
			rf = float64(vi)
			dt = UF64_TYPE_FLOAT64
			return
		}
	} else if src.dataType == "uint32" {
		ri = int64(any(v).(uint32))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "uint8" {
		ri = int64(any(v).(uint8))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "uint16" {
		ri = int64(any(v).(uint16))
		dt = UF64_TYPE_INT64
		return
	} else if src.dataType == "float32" {
		rf = float64(any(v).(float32))
		dt = UF64_TYPE_FLOAT64
		return
	} else if src.dataType == "float64" {
		rf = any(v).(float64)
		dt = UF64_TYPE_FLOAT64
		return
	} else {
		panic(makeCastError(v, "float64"))
	}

	if vf64, okf64 := any(v).(float64); okf64 {
		rf = vf64
		dt = UF64_TYPE_FLOAT64
		return
	} else if vi, oki := any(v).(int); oki {
		ri = int64(vi)
		dt = UF64_TYPE_INT64
		return
	} else if v64, ok64 := any(v).(int64); ok64 {
		ri = v64
		dt = UF64_TYPE_INT64
		return
	} else if v8, ok8 := any(v).(int8); ok8 {
		ri = int64(v8)
		dt = UF64_TYPE_INT64
		return
	} else if v32, ok32 := any(v).(int32); ok32 {
		ri = int64(v32)
		dt = UF64_TYPE_INT64
		return
	} else if vf32, okf32 := any(v).(float32); okf32 {
		rf = float64(vf32)
		dt = UF64_TYPE_FLOAT64
		return
	} else if v16, ok16 := any(v).(int16); ok16 {
		ri = int64(v16)
		dt = UF64_TYPE_INT64
		return
	} else if vui, okui := any(v).(uint); okui {
		if vui > math.MaxInt64 {
			bf := new(big.Float)
			bf = bf.SetUint64(uint64(vui))
			rf, _ = bf.Float64()
			dt = UF64_TYPE_FLOAT64
			return
		}
		ri = int64(vui)
		dt = UF64_TYPE_INT64
		return
	} else if vu64, oku64 := any(v).(uint64); oku64 {
		if vu64 > math.MaxInt64 {
			bf := new(big.Float)
			bf = bf.SetUint64(vu64)
			rf, _ = bf.Float64()
			dt = UF64_TYPE_FLOAT64
			return
		}
		ri = int64(vu64)
		dt = UF64_TYPE_INT64
		return
	} else if vu32, oku32 := any(v).(uint32); oku32 {
		ri = int64(vu32)
		dt = UF64_TYPE_INT64
		return
	} else if vu8, oku8 := any(v).(uint8); oku8 {
		ri = int64(vu8)
		dt = UF64_TYPE_INT64
		return
	} else if vu16, oku16 := any(v).(uint16); oku16 {
		ri = int64(vu16)
		dt = UF64_TYPE_INT64
		return
	} else {
		panic(makeCastError(v, "float64"))
	}
}

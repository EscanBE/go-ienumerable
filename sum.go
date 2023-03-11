package go_ienumerable

import (
	"fmt"
	"github.com/johncgriffin/overflow"
	"math"
)

func (src *enumerable[T]) SumInt32() int32 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		typeName := fmt.Sprintf("%T", *new(T))
		switch typeName {
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int32", typeName))
		}
	}

	var sum int64 = 0

	for _, d := range src.data {
		if vi, oki := any(d).(int); oki {
			if math.MinInt32 > vi || vi > math.MaxInt32 {
				panic(makeCastError(d, "int32"))
			}
			sum = overflow.Add64p(sum, int64(vi))
		} else if v64, ok64 := any(d).(int64); ok64 {
			if math.MinInt32 > v64 || v64 > math.MaxInt32 {
				panic(makeCastError(d, "int32"))
			}
			sum = overflow.Add64p(sum, v64)
		} else if v8, ok8 := any(d).(int8); ok8 {
			sum = overflow.Add64p(sum, int64(v8))
		} else if v32, ok32 := any(d).(int32); ok32 {
			sum = overflow.Add64p(sum, int64(v32))
		} else if v16, ok16 := any(d).(int16); ok16 {
			sum = overflow.Add64p(sum, int64(v16))
		} else if vui, okui := any(d).(uint); okui {
			if vui > math.MaxInt32 {
				panic(makeCastError(d, "int32"))
			}
			sum = overflow.Add64p(sum, int64(vui))
		} else if vu64, oku64 := any(d).(uint64); oku64 {
			if vu64 > math.MaxInt32 {
				panic(makeCastError(d, "int32"))
			}
			sum = overflow.Add64p(sum, int64(vu64))
		} else if vu32, oku32 := any(d).(uint32); oku32 {
			if vu32 > math.MaxInt32 {
				panic(makeCastError(d, "int32"))
			}
			sum = overflow.Add64p(sum, int64(vu32))
		} else if vu8, oku8 := any(d).(uint8); oku8 {
			sum = overflow.Add64p(sum, int64(vu8))
		} else if vu16, oku16 := any(d).(uint16); oku16 {
			sum = overflow.Add64p(sum, int64(vu16))
		} else {
			panic(makeCastError(d, "int32"))
		}
	}

	if math.MinInt32 > sum || sum > math.MaxInt32 {
		panic("overflow int32")
	}

	return int32(sum)
}

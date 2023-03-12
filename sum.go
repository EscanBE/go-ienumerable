package go_ienumerable

import (
	"fmt"
	"math"
	"math/big"
)

func (src *enumerable[T]) SumInt32() int32 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int32", src.dataType))
		}
	}

	var sum int64 = 0

	for _, d := range src.data {
		i32 := src.unboxAnyAsInt32(d)
		sum = add64p(sum, int64(i32))
	}

	if math.MinInt32 > sum || sum > math.MaxInt32 {
		panic("overflow")
	}

	return int32(sum)
}

func (src *enumerable[T]) SumInt64() int64 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int64", src.dataType))
		}
	}

	sum := new(big.Int)

	for _, d := range src.data {
		i64 := src.unboxAnyAsInt64(d)
		sum = sum.Add(sum, big.NewInt(i64))
	}

	if !sum.IsInt64() {
		panic("overflow")
	}

	return sum.Int64()
}

func (src *enumerable[T]) SumFloat32() float32 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "float32", "float64":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to float32", src.dataType))
		}
	}

	var sum float64 = 0
	const minFloat32 = float64(-1*math.MaxFloat32) - 1

	for _, d := range src.data {
		previousSum := sum
		sign := 0
		if vf64, okf64 := any(d).(float64); okf64 {
			if vf64 > math.MaxFloat32 || vf64 < minFloat32 {
				panic(makeCastError(d, "float32"))
			}
			sum += vf64
			if vf64 > 0 {
				sign = 1
			} else if vf64 < 0 {
				sign = -1
			}
		} else if vi, oki := any(d).(int); oki {
			sum += float64(vi)
			if vi > 0 {
				sign = 1
			} else if vi < 0 {
				sign = -1
			}
		} else if v64, ok64 := any(d).(int64); ok64 {
			sum += float64(v64)
			if v64 > 0 {
				sign = 1
			} else if v64 < 0 {
				sign = -1
			}
		} else if vf32, okf32 := any(d).(float32); okf32 {
			sum += float64(vf32)
			if vf32 > 0 {
				sign = 1
			} else if vf32 < 0 {
				sign = -1
			}
		} else if v8, ok8 := any(d).(int8); ok8 {
			sum += float64(v8)
			if v8 > 0 {
				sign = 1
			} else if v8 < 0 {
				sign = -1
			}
		} else if v32, ok32 := any(d).(int32); ok32 {
			sum += float64(v32)
			if v32 > 0 {
				sign = 1
			} else if v32 < 0 {
				sign = -1
			}
		} else if v16, ok16 := any(d).(int16); ok16 {
			sum += float64(v16)
			if v16 > 0 {
				sign = 1
			} else if v16 < 0 {
				sign = -1
			}
		} else if vui, okui := any(d).(uint); okui {
			sum += float64(vui)
			if vui > 0 {
				sign = 1
			}
		} else if vu64, oku64 := any(d).(uint64); oku64 {
			sum += float64(vu64)
			if vu64 > 0 {
				sign = 1
			}
		} else if vu32, oku32 := any(d).(uint32); oku32 {
			sum += float64(vu32)
			if vu32 > 0 {
				sign = 1
			}
		} else if vu8, oku8 := any(d).(uint8); oku8 {
			sum += float64(vu8)
			if vu8 > 0 {
				sign = 1
			}
		} else if vu16, oku16 := any(d).(uint16); oku16 {
			sum += float64(vu16)
			if vu16 > 0 {
				sign = 1
			}
		} else {
			panic(makeCastError(d, "float32"))
		}

		if sign != 0 {
			if sign > 0 && sum <= previousSum {
				panic("overflow")
			} else if sign < 0 && sum >= previousSum {
				panic("overflow")
			}
		}
	}

	if minFloat32 > sum || sum > math.MaxFloat32 {
		panic("overflow")
	}

	return float32(sum)
}

func (src *enumerable[T]) SumFloat64() float64 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "float32", "float64":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to float64", src.dataType))
		}
	}

	var sum float64 = 0

	for _, d := range src.data {
		previousSum := sum
		sign := 0
		if vf64, okf64 := any(d).(float64); okf64 {
			sum += vf64
			if vf64 > 0 {
				sign = 1
			} else if vf64 < 0 {
				sign = -1
			}
		} else if vi, oki := any(d).(int); oki {
			sum += float64(vi)
			if vi > 0 {
				sign = 1
			} else if vi < 0 {
				sign = -1
			}
		} else if v64, ok64 := any(d).(int64); ok64 {
			sum += float64(v64)
			if v64 > 0 {
				sign = 1
			} else if v64 < 0 {
				sign = -1
			}
		} else if vf32, okf32 := any(d).(float32); okf32 {
			sum += float64(vf32)
			if vf32 > 0 {
				sign = 1
			} else if vf32 < 0 {
				sign = -1
			}
		} else if v8, ok8 := any(d).(int8); ok8 {
			sum += float64(v8)
			if v8 > 0 {
				sign = 1
			} else if v8 < 0 {
				sign = -1
			}
		} else if v32, ok32 := any(d).(int32); ok32 {
			sum += float64(v32)
			if v32 > 0 {
				sign = 1
			} else if v32 < 0 {
				sign = -1
			}
		} else if v16, ok16 := any(d).(int16); ok16 {
			sum += float64(v16)
			if v16 > 0 {
				sign = 1
			} else if v16 < 0 {
				sign = -1
			}
		} else if vui, okui := any(d).(uint); okui {
			sum += float64(vui)
			if vui > 0 {
				sign = 1
			}
		} else if vu64, oku64 := any(d).(uint64); oku64 {
			sum += float64(vu64)
			if vu64 > 0 {
				sign = 1
			}
		} else if vu32, oku32 := any(d).(uint32); oku32 {
			sum += float64(vu32)
			if vu32 > 0 {
				sign = 1
			}
		} else if vu8, oku8 := any(d).(uint8); oku8 {
			sum += float64(vu8)
			if vu8 > 0 {
				sign = 1
			}
		} else if vu16, oku16 := any(d).(uint16); oku16 {
			sum += float64(vu16)
			if vu16 > 0 {
				sign = 1
			}
		} else {
			panic(makeCastError(d, "float64"))
		}

		if sign != 0 {
			if math.IsInf(sum, sign) {
				panic("overflow")
			} else if sign > 0 && sum <= previousSum {
				panic("overflow")
			} else if sign < 0 && sum >= previousSum {
				panic("overflow")
			}
		}
	}

	return sum
}

package goe

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
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
		case "*int8", "*uint8", "*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*int", "*uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int32", src.dataType))
		}
	}

	var sum int64 = 0

	for _, d := range src.data {
		i32 := reflection.UnboxAnyAsInt32(d)
		sum = add64p(sum, int64(i32))
	}

	if math.MinInt32 > sum || sum > math.MaxInt32 {
		panic("overflow")
	}

	return int32(sum)
}

func (src *enumerable[T]) SumInt() int {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint":
			return 0
		case "*int8", "*uint8", "*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*int", "*uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int", src.dataType))
		}
	}

	sum := new(big.Int)

	for _, d := range src.data {
		i64 := int64(reflection.UnboxAnyAsInt(d))
		sum = sum.Add(sum, big.NewInt(i64))
	}

	if !sum.IsInt64() {
		panic("overflow")
	}

	vSum := sum.Int64()

	//if math.MinInt > vSum || vSum > math.MaxInt {
	//	panic("overflow")
	//}

	return int(vSum)
}

func (src *enumerable[T]) SumInt64() int64 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint":
			return 0
		case "*int8", "*uint8", "*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*int", "*uint":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to int64", src.dataType))
		}
	}

	sum := new(big.Int)

	for _, d := range src.data {
		i64 := reflection.UnboxAnyAsInt64(d)
		sum = sum.Add(sum, big.NewInt(i64))
	}

	if !sum.IsInt64() {
		panic("overflow")
	}

	return sum.Int64()
}

func (src *enumerable[T]) SumFloat64() float64 {
	src.assertSrcNonNil()

	if len(src.data) < 1 {
		switch src.dataType {
		case "":
			return 0
		case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "int", "uint", "float32", "float64":
			return 0
		case "*int8", "*uint8", "*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*int", "*uint", "*float32", "*float64":
			return 0
		default:
			panic(fmt.Errorf("type %s cannot be tried to cast to float64", src.dataType))
		}
	}

	sumBi := new(big.Int)
	sumBf := new(big.Float)

	for _, d := range src.data {
		ri, rf, dt := reflection.UnboxAnyAsInt64OrFloat64(d)
		if dt == reflection.UF64_TYPE_INT64 {
			sumBi = sumBi.Add(sumBi, big.NewInt(ri))
		} else if dt == reflection.UF64_TYPE_FLOAT64 {
			sumBf = sumBf.Add(sumBf, big.NewFloat(rf))
		}
	}

	sumBf = sumBf.Add(sumBf, new(big.Float).SetInt(sumBi))

	sum, _ := sumBf.Float64()

	if math.IsInf(sum, sumBf.Sign()) {
		panic("overflow")
	}

	return sum
}

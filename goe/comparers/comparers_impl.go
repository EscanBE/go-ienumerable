package comparers

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

// ensure implementation
var (
	_ IComparer[int8]       = int8Comparer{}
	_ IComparer[uint8]      = uint8Comparer{}
	_ IComparer[int16]      = int16Comparer{}
	_ IComparer[uint16]     = uint16Comparer{}
	_ IComparer[int32]      = int32Comparer{}
	_ IComparer[uint32]     = uint32Comparer{}
	_ IComparer[int64]      = int64Comparer{}
	_ IComparer[uint64]     = uint64Comparer{}
	_ IComparer[int]        = intComparer{}
	_ IComparer[uint]       = uintComparer{}
	_ IComparer[uintptr]    = uintptrComparer{}
	_ IComparer[float32]    = float32Comparer{}
	_ IComparer[float64]    = float64Comparer{}
	_ IComparer[complex64]  = complex64Comparer{}
	_ IComparer[complex128] = complex128Comparer{}
	_ IComparer[string]     = stringComparer{}
	_ IComparer[bool]       = boolComparer{}
	_ IComparer[any]        = partitionedComparer[any]{}
)

type int8Comparer struct {
}

// NewInt8Comparer returns IComparer for int8 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewInt8Comparer() IComparer[int8] {
	return int8Comparer{}
}

func (i int8Comparer) Compare(x, y int8) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i int8Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[int8](x), AnyPointerToType[int8](y))
}

// AnyPointerToType convert *T or *any (*interface{}) into T, otherwise panic
func AnyPointerToType[T any](p any) T {
	if pT, ok1 := p.(*T); ok1 {
		return *pT
	}

	if pA, ok2 := p.(*any); ok2 {
		return (*pA).(T)
	}

	return p.(T)
}

/*
type numberComparer struct {
}

func (i numberComparer) Compare(x, y any) int {
	xi64, xf64, tx64, sx64 := reflection.TryUnboxAnyAsInt64OrFloat64(x)
	if sx64 == reflection.UNBOX_FAILED {
		panic(fmt.Errorf("%v of type %T can not be casted to number", x, x))
	}

	yi64, yf64, ty64, sy64 := reflection.TryUnboxAnyAsInt64OrFloat64(y)
	if sy64 == reflection.UNBOX_FAILED {
		panic(fmt.Errorf("%v of type %T can not be casted to number", y, y))
	}

	if sx64 == reflection.UNBOX_NIL && sy64 == reflection.UNBOX_NIL {
		return 0
	}

	if sx64 == reflection.UNBOX_NIL {
		return -1
	}

	if sy64 == reflection.UNBOX_NIL {
		return 1
	}

	sx64.AssertIsSuccess()
	sy64.AssertIsSuccess()
	tx64.AssertHasNonNilResult()
	ty64.AssertHasNonNilResult()

	if tx64 == reflection.UF64_TYPE_INT64 && ty64 == reflection.UF64_TYPE_INT64 {
		if xi64 < yi64 {
			return -1
		} else if xi64 > yi64 {
			return 1
		} else {
			return 0
		}
	}

	if tx64 == reflection.UF64_TYPE_INT64 {
		xf64 = float64(xi64)
	}

	if ty64 == reflection.UF64_TYPE_INT64 {
		yf64 = float64(yi64)
	}

	if xf64 < yf64 {
		return -1
	} else if xf64 > yf64 {
		return 1
	} else {
		return 0
	}
}
*/

type uint8Comparer struct {
}

// NewUint8Comparer returns IComparer for uint8 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewUint8Comparer() IComparer[uint8] {
	return uint8Comparer{}
}

func (i uint8Comparer) Compare(x, y uint8) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uint8Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uint8](x), AnyPointerToType[uint8](y))
}

type int16Comparer struct {
}

// NewInt16Comparer returns IComparer for int16 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewInt16Comparer() IComparer[int16] {
	return int16Comparer{}
}

func (i int16Comparer) Compare(x, y int16) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i int16Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[int16](x), AnyPointerToType[int16](y))
}

type uint16Comparer struct {
}

// NewUint16Comparer returns IComparer for uint16 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewUint16Comparer() IComparer[uint16] {
	return uint16Comparer{}
}

func (i uint16Comparer) Compare(x, y uint16) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uint16Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uint16](x), AnyPointerToType[uint16](y))
}

type int32Comparer struct {
}

// NewInt32Comparer returns IComparer for int32 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewInt32Comparer() IComparer[int32] {
	return int32Comparer{}
}

func (i int32Comparer) Compare(x, y int32) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i int32Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[int32](x), AnyPointerToType[int32](y))
}

type uint32Comparer struct {
}

// NewUint32Comparer returns IComparer for uint32 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewUint32Comparer() IComparer[uint32] {
	return uint32Comparer{}
}

func (i uint32Comparer) Compare(x, y uint32) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uint32Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uint32](x), AnyPointerToType[uint32](y))
}

type int64Comparer struct {
}

// NewInt64Comparer returns IComparer for int64 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewInt64Comparer() IComparer[int64] {
	return int64Comparer{}
}

func (i int64Comparer) Compare(x, y int64) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i int64Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[int64](x), AnyPointerToType[int64](y))
}

type uint64Comparer struct {
}

// NewUint64Comparer returns IComparer for uint64 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewUint64Comparer() IComparer[uint64] {
	return uint64Comparer{}
}

func (i uint64Comparer) Compare(x, y uint64) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uint64Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uint64](x), AnyPointerToType[uint64](y))
}

type intComparer struct {
}

// NewIntComparer returns IComparer for int with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewIntComparer() IComparer[int] {
	return intComparer{}
}

func (i intComparer) Compare(x, y int) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i intComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[int](x), AnyPointerToType[int](y))
}

type uintComparer struct {
}

// NewUintComparer returns IComparer for uint with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewUintComparer() IComparer[uint] {
	return uintComparer{}
}

func (i uintComparer) Compare(x, y uint) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uintComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uint](x), AnyPointerToType[uint](y))
}

type uintptrComparer struct {
}

// NewUintptrComparer returns IComparer for uintptr with comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
//
// Ps: Don't know if this is comparison algorithm is correct
func NewUintptrComparer() IComparer[uintptr] {
	return uintptrComparer{}
}

func (i uintptrComparer) Compare(x, y uintptr) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i uintptrComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[uintptr](x), AnyPointerToType[uintptr](y))
}

type float32Comparer struct {
}

// NewFloat32Comparer returns IComparer for float32 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewFloat32Comparer() IComparer[float32] {
	return float32Comparer{}
}

func (i float32Comparer) Compare(x, y float32) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i float32Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[float32](x), AnyPointerToType[float32](y))
}

type float64Comparer struct {
}

// NewFloat64Comparer returns IComparer for float64 with default comparison:
//
// x < y ? -1 (x == y ? 0 : 1)
func NewFloat64Comparer() IComparer[float64] {
	return float64Comparer{}
}

func (i float64Comparer) Compare(x, y float64) int {
	if x < y {
		return -1
	}

	if x > y {
		return 1
	}

	return 0
}

func (i float64Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[float64](x), AnyPointerToType[float64](y))
}

type bigIntComparer struct {
}

// NewBigIntComparer returns IComparer for *big.Int with default comparison.
func NewBigIntComparer() IComparer[*big.Int] {
	return bigIntComparer{}
}

func (i bigIntComparer) Compare(x, y *big.Int) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return x.Cmp(y)
}

func (i bigIntComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return AnyPointerToType[*big.Int](x).Cmp(AnyPointerToType[*big.Int](y))
}

type complex64Comparer struct {
}

/*
NewComplex64Comparer returns IComparer for complex64 with comparison:

real(x) < real(y)
? -1
: (

	real(x) > real(y)
	? 1
	: (
		imag(x) < imag(y)
		? - 1
		: imag(x) > imag(y) ? 1 : 0
	)

)

Ps: Don't know if this is comparison algorithm is correct
*/
func NewComplex64Comparer() IComparer[complex64] {
	return complex64Comparer{}
}

func (i complex64Comparer) Compare(x, y complex64) int {
	var xrf = real(x)
	var yrf = real(y)

	if xrf < yrf {
		return -1
	}

	if xrf > yrf {
		return 1
	}

	var xif = imag(x)
	var yif = imag(y)

	if xif < yif {
		return -1
	}

	if xif > yif {
		return 1
	}

	return 0
}

func (i complex64Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[complex64](x), AnyPointerToType[complex64](y))
}

type complex128Comparer struct {
}

/*
NewComplex128Comparer returns IComparer for complex128 with comparison:

real(x) < real(y)
? -1
: (

	real(x) > real(y)
	? 1
	: (
		imag(x) < imag(y)
		? - 1
		: imag(x) > imag(y) ? 1 : 0
	)

)

Ps: Don't know if this is comparison algorithm is correct
*/
func NewComplex128Comparer() IComparer[complex128] {
	return complex128Comparer{}
}

func (i complex128Comparer) Compare(x, y complex128) int {
	var xrf = real(x)
	var yrf = real(y)

	if xrf < yrf {
		return -1
	}

	if xrf > yrf {
		return 1
	}

	var xif = imag(x)
	var yif = imag(y)

	if xif < yif {
		return -1
	}

	if xif > yif {
		return 1
	}

	return 0
}

func (i complex128Comparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[complex128](x), AnyPointerToType[complex128](y))
}

type stringComparer struct {
}

// NewStringComparer returns IComparer for string with default comparison.
func NewStringComparer() IComparer[string] {
	return stringComparer{}
}

func (i stringComparer) Compare(x, y string) int {
	return strings.Compare(x, y)
}

func (i stringComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[string](x), AnyPointerToType[string](y))
}

type boolComparer struct {
}

// NewBoolComparer returns IComparer for bool with comparison:
// x == y ? 0 : (!x ? -1 : 1)
func NewBoolComparer() IComparer[bool] {
	return boolComparer{}
}

func (i boolComparer) Compare(x, y bool) int {
	if x == y {
		return 0
	}

	if !x {
		return -1
	}

	return 1
}

func (i boolComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[bool](x), AnyPointerToType[bool](y))
}

type timeComparer struct {
}

// NewTimeComparer returns IComparer for time.Time with default comparison.
func NewTimeComparer() IComparer[time.Time] {
	return timeComparer{}
}

func (i timeComparer) Compare(x, y time.Time) int {
	if x == y {
		return 0
	}

	if x.Before(y) {
		return -1
	}

	return 1
}

func (i timeComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[time.Time](x), AnyPointerToType[time.Time](y))
}

type durationComparer struct {
}

// NewDurationComparer returns IComparer for time.Duration with default comparison.
func NewDurationComparer() IComparer[time.Duration] {
	return durationComparer{}
}

func (i durationComparer) Compare(x, y time.Duration) int {
	if x == y {
		return 0
	}

	if x < y {
		return -1
	}

	return 1
}

func (i durationComparer) ComparePointerMode(x, y any) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[time.Duration](x), AnyPointerToType[time.Duration](y))
}

// partitionedComparer can run in 2 options, if comparer is provided, use it, otherwise use the pair equals and less
type partitionedComparer[T any] struct {
	equals   func(v1, v2 T) bool
	less     func(v1, v2 T) bool
	comparer IComparer[T]
}

// NewPartitionedComparerFromEL creates an IComparer that uses
// both of the 'equals' and 'less' methods for underlying comparison.
func NewPartitionedComparerFromEL[T any](equals func(v1, v2 T) bool, less func(v1, v2 T) bool) IComparer[T] {
	if equals == nil {
		panic("equals comparer is nil")
	}
	if less == nil {
		panic("less comparer is nil")
	}
	return partitionedComparer[T]{
		equals: equals,
		less:   less,
	}
}

// NewPartitionedComparerFromComparer creates an IComparer that uses
// another IComparer for underlying comparison.
func NewPartitionedComparerFromComparer[T any](comparer IComparer[T]) IComparer[T] {
	if comparer == nil {
		panic("comparer is nil")
	}
	return partitionedComparer[T]{
		comparer: comparer,
	}
}

func (i partitionedComparer[T]) Compare(x, y T) int {
	if i.comparer != nil {
		return i.comparer.Compare(x, y)
	}

	if i.equals(x, y) {
		return 0
	}

	if i.less(x, y) {
		return -1
	}

	return 1
}

func (i partitionedComparer[T]) ComparePointerMode(x, y any) int {
	if i.comparer != nil {
		return i.comparer.ComparePointerMode(x, y)
	}

	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return i.Compare(AnyPointerToType[T](x), AnyPointerToType[T](y))
}

// defaultComparer will be used as default comparer instance for types.
// It shall wrap an IComparer[T] inside and present as IComparer[any],
// and its functions Compare and ComparePointerMode can accept
// interface{} as parameter and auto-forward to corresponding compare method.
//
// Eg: a default comparer that wraps a IComparer[int]:
//
// - pass int & int to any Compare* method will forward to Compare,
//
// - pass *int & *int to any Compare* method will forward to ComparePointerMode,
//
// - pass *int & int to any Compare* will panic (not same representation),
//
// - pass int & int32 to any Compare* also panic (int32 is not int),
type defaultComparer struct {
	compareFunc func(v1, v2 any) int
}

// ConvertToDefaultComparer wraps the typed comparer IComparer[T] into IComparer[any].
// With ability to receive both *T and T to compare, but both params must have the same representation: (T & T) or (*T & *T)
// This is used for auto-resolve comparer technique.
//
// Eg: a default comparer that wraps a IComparer[int]:
//
// - pass int & int to any Compare* method will forward to Compare,
//
// - pass *int & *int to any Compare* method will forward to ComparePointerMode,
//
// - pass *int & int to any Compare* will panic (not same representation),
//
// - pass int & int32 to any Compare* also panic (int32 is not int),
func ConvertToDefaultComparer[T any](comparer IComparer[T]) IComparer[any] {
	if comparer == nil {
		panic("comparer is nil")
	}

	if dc, ok := any(comparer).(*defaultComparer); ok {
		return dc
	}

	return &defaultComparer{
		compareFunc: func(unknown1, unknown2 any) int {
			if unknown1 == nil && unknown2 == nil {
				return comparer.ComparePointerMode(nil, nil)
			}

			if unknown1 == nil {
				pointer2, okPointer2 := unknown2.(*T)

				if okPointer2 {
					return comparer.ComparePointerMode(nil, pointer2)
				}

				pointerAny2, okPointerAny2 := unknown2.(*any)
				if okPointerAny2 {
					//goland:noinspection GoSnakeCaseUsage
					v2_2, okv2_2 := (*pointerAny2).(T)
					if okv2_2 {
						return comparer.ComparePointerMode(nil, &v2_2)
					}
				}

				panic(fmt.Sprintf("first param is nil but second param neither value or pointer. Found [%T]", unknown2))
			}

			if unknown2 == nil {
				pointer1, okPointer1 := unknown1.(*T)

				if okPointer1 {
					return comparer.ComparePointerMode(pointer1, nil)
				}

				pointerAny1, okPointerAny1 := unknown1.(*any)
				if okPointerAny1 {
					value1, okValue1 := (*pointerAny1).(T)
					if okValue1 {
						return comparer.ComparePointerMode(&value1, nil)
					}
				}

				panic(fmt.Sprintf("second param is nil but first param neither value or pointer. Found [%T]", unknown1))
			}

			value1, okValue1 := unknown1.(T)
			value2, okValue2 := unknown2.(T)
			if okValue1 && okValue2 {
				return comparer.Compare(value1, value2)
			}

			pointer1, okPointer1 := unknown1.(*T)
			pointer2, okPointer2 := unknown2.(*T)

			if okPointer1 && okPointer2 {
				return comparer.ComparePointerMode(pointer1, pointer2)
			}

			pointerAny1, okPointerAny1 := unknown1.(*any)
			pointerAny2, okPointerAny2 := unknown2.(*any)

			if okPointerAny1 && okPointerAny2 {
				v1, ok1 := (*pointerAny1).(T)
				v2, ok2 := (*pointerAny2).(T)

				if ok1 && ok2 {
					return comparer.Compare(v1, v2)
				}
			}

			if (!okValue1 && !okPointer1) || (!okValue2 && !okPointer2) {
				panic(fmt.Sprintf("first or second params neither value or pointer. Found [%T] and [%T]", unknown1, unknown2))
			}

			panic(fmt.Sprintf("both params must have the same presentation, value or pointer. Found [%T] and [%T]", unknown1, unknown2))
		},
	}
}

func (i *defaultComparer) Compare(x, y any) int {
	return i.compareFunc(x, y)
}

func (i defaultComparer) ComparePointerMode(x, y any) int {
	return i.compareFunc(x, y)
}

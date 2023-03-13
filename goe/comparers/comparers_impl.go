package comparers

import (
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
	if pT, ok := p.(*T); ok {
		return *pT
	}
	return (*p.(*any)).(T)
}

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

type wrappedComparer struct {
	compareFunc            func(v1, v2 any) int
	comparePointerModeFunc func(v1, v2 any) int
}

// HideTypedComparer wraps the typed comparer IComparer[T] into IComparer[any].
// This is used for auto-resolve comparer technique.
//
// Beware: input parameter still have to have correct type.
// For example when wraps a IComparer[int8],
// parameter passed to the Compare function still have to be int8,
// otherwise panic
func HideTypedComparer[T any](comparer IComparer[T]) IComparer[any] {
	return &wrappedComparer{
		compareFunc: func(v1, v2 any) int {
			return comparer.Compare(v1.(T), v2.(T))
		},
		comparePointerModeFunc: func(v1, v2 any) int {
			return comparer.ComparePointerMode(v1, v2)
		},
	}
}

func (i *wrappedComparer) Compare(x, y any) int {
	return i.compareFunc(x, y)
}

func (i wrappedComparer) ComparePointerMode(x, y any) int {
	return i.comparePointerModeFunc(x, y)
}

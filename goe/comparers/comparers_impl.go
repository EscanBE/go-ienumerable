package comparers

import "strings"

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

type uint8Comparer struct {
}

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

type int16Comparer struct {
}

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

type uint16Comparer struct {
}

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

type int32Comparer struct {
}

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

type uint32Comparer struct {
}

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

type int64Comparer struct {
}

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

type uint64Comparer struct {
}

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

type intComparer struct {
}

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

type uintComparer struct {
}

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

type uintptrComparer struct {
}

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

type float32Comparer struct {
}

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

type float64Comparer struct {
}

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

type complex64Comparer struct {
}

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

type complex128Comparer struct {
}

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

type stringComparer struct {
}

func NewStringComparer() IComparer[string] {
	return stringComparer{}
}

func (i stringComparer) Compare(x, y string) int {
	return strings.Compare(x, y)
}

type boolComparer struct {
}

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

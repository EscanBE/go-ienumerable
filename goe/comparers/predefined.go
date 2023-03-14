package comparers

import (
	"fmt"
	"strings"
	"time"
)

// TODO: implement default comparers for big.Int, big.Float

var (
	Int8Comparer       = NewInt8Comparer()
	Uint8Comparer      = NewUint8Comparer()
	Int16Comparer      = NewInt16Comparer()
	Uint16Comparer     = NewUint16Comparer()
	Int32Comparer      = NewInt32Comparer()
	Uint32Comparer     = NewUint32Comparer()
	Int64Comparer      = NewInt64Comparer()
	Uint64Comparer     = NewUint64Comparer()
	IntComparer        = NewIntComparer()
	UintComparer       = NewUintComparer()
	UintptrComparer    = NewUintptrComparer()
	Float32Comparer    = NewFloat32Comparer()
	Float64Comparer    = NewFloat64Comparer()
	Complex64Comparer  = NewComplex64Comparer()
	Complex128Comparer = NewComplex128Comparer()
	StringComparer     = NewStringComparer()
	BoolComparer       = NewBoolComparer()
	TimeComparer       = NewTimeComparer()
	DurationComparer   = NewDurationComparer()
)

var mappedDefaultComparers = map[string]IComparer[any]{
	getNormalizeTypeName[int8]():          ConvertToDefaultComparer[int8](Int8Comparer),
	getNormalizeTypeName[uint8]():         ConvertToDefaultComparer[uint8](Uint8Comparer),
	getNormalizeTypeName[int16]():         ConvertToDefaultComparer[int16](Int16Comparer),
	getNormalizeTypeName[uint16]():        ConvertToDefaultComparer[uint16](Uint16Comparer),
	getNormalizeTypeName[int32]():         ConvertToDefaultComparer[int32](Int32Comparer),
	getNormalizeTypeName[uint32]():        ConvertToDefaultComparer[uint32](Uint32Comparer),
	getNormalizeTypeName[int64]():         ConvertToDefaultComparer[int64](Int64Comparer),
	getNormalizeTypeName[uint64]():        ConvertToDefaultComparer[uint64](Uint64Comparer),
	getNormalizeTypeName[int]():           ConvertToDefaultComparer[int](IntComparer),
	getNormalizeTypeName[uint]():          ConvertToDefaultComparer[uint](UintComparer),
	getNormalizeTypeName[uintptr]():       ConvertToDefaultComparer[uintptr](UintptrComparer),
	getNormalizeTypeName[float32]():       ConvertToDefaultComparer[float32](Float32Comparer),
	getNormalizeTypeName[float64]():       ConvertToDefaultComparer[float64](Float64Comparer),
	getNormalizeTypeName[complex64]():     ConvertToDefaultComparer[complex64](Complex64Comparer),
	getNormalizeTypeName[complex128]():    ConvertToDefaultComparer[complex128](Complex128Comparer),
	getNormalizeTypeName[string]():        ConvertToDefaultComparer[string](StringComparer),
	getNormalizeTypeName[bool]():          ConvertToDefaultComparer[bool](BoolComparer),
	getNormalizeTypeName[time.Time]():     ConvertToDefaultComparer[time.Time](TimeComparer),
	getNormalizeTypeName[time.Duration](): ConvertToDefaultComparer[time.Duration](DurationComparer),
}

// GetDefaultComparer attempts to get IComparer for corresponding type and returns as IComparer[any].
// Panic if no default comparer registered for this type or unable to detect type,
// can specify via GetDefaultComparerByTypeName or TryGetDefaultComparerByTypeName
func GetDefaultComparer[T any]() IComparer[any] {
	typeName := fmt.Sprintf("%T", *new(T))

	if len(typeName) < 1 || typeName == "<nil>" {
		panic(fmt.Sprintf("unable to detect type for provided type"))
	}

	if comparer, found := TryGetDefaultComparerByTypeName(typeName); found {
		return comparer
	}

	panic(fmt.Sprintf("no default comparer registered for type [%s]", typeName))
}

// TryGetDefaultComparer attempts to get IComparer for corresponding type and returns as IComparer[any].
func TryGetDefaultComparer[T any]() (comparer IComparer[any], found bool) {
	typeName := fmt.Sprintf("%T", *new(T))
	if len(typeName) < 1 || typeName == "<nil>" {
		return
	}

	return TryGetDefaultComparerByTypeName(typeName)
}

// GetDefaultComparerByTypeName attempts to get IComparer for specified type and returns as IComparer[any].
// Panic if no default comparer registered for this type
func GetDefaultComparerByTypeName(typeName string) IComparer[any] {
	if comparer, found := TryGetDefaultComparerByTypeName(typeName); found {
		return comparer
	}

	panic(fmt.Sprintf("no default comparer registered for type [%s]", typeName))
}

// TryGetDefaultComparerByTypeName attempts to get IComparer for specified type and returns as IComparer[any].
func TryGetDefaultComparerByTypeName(typeName string) (comparer IComparer[any], found bool) {
	typeName = normalizeTypeName(typeName)
	comparer, found = mappedDefaultComparers[typeName]
	return
}

// RegisterDefaultComparerForType register a comparer for specific type into registry.
//
// Panic if duplicated for a type or comparer is nil or type name is empty/<nil>
func RegisterDefaultComparerForType[T any](typeName string, comparer IComparer[T], allowOverride bool) {
	if strings.HasPrefix(typeName, "*") {
		panic("can not register for pointer. Pointer is automatically registered when register for normal")
	}

	typeName = normalizeTypeName(typeName)
	if len(typeName) < 1 || typeName == "<nil>" || typeName == "*" {
		panic("empty or <nil> type name")
	}

	if comparer == nil {
		panic("comparer is nil")
	}

	if !allowOverride {
		if existing, found := mappedDefaultComparers[typeName]; found && existing != nil {
			panic(fmt.Sprintf("default comparer for type [%s] had been registered before", typeName))
		}
	}

	mappedDefaultComparers[typeName] = ConvertToDefaultComparer[T](comparer)
	fmt.Printf("Registered default comparer for type [%s]\n", typeName)
}

// RegisterDefaultTypedComparer register a comparer for specific type into registry.
// Will not able to detect type if T is nil-able, use RegisterDefaultCompilerForType with specified type name instead
//
// Panic if duplicated for a type or comparer is nil or type name is empty/<nil>
func RegisterDefaultTypedComparer[T any](comparer IComparer[T], allowOverride bool) {
	RegisterDefaultComparerForType[T](fmt.Sprintf("%T", *new(T)), comparer, allowOverride)
}

func normalizeTypeName(typeName string) string {
	return strings.TrimPrefix(strings.ToLower(strings.TrimSpace(typeName)), "*")
}

func getNormalizeTypeName[T any]() string {
	return normalizeTypeName(fmt.Sprintf("%T", *new(T)))
}

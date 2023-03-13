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

var mappedWrappedComparers = map[string]IComparer[any]{
	"int8":                              HideTypedComparer[int8](Int8Comparer),
	"uint8":                             HideTypedComparer[uint8](Uint8Comparer),
	"int16":                             HideTypedComparer[int16](Int16Comparer),
	"uint16":                            HideTypedComparer[uint16](Uint16Comparer),
	"int32":                             HideTypedComparer[int32](Int32Comparer),
	"uint32":                            HideTypedComparer[uint32](Uint32Comparer),
	"int64":                             HideTypedComparer[int64](Int64Comparer),
	"uint64":                            HideTypedComparer[uint64](Uint64Comparer),
	"int":                               HideTypedComparer[int](IntComparer),
	"uint":                              HideTypedComparer[uint](UintComparer),
	"uintptr":                           HideTypedComparer[uintptr](UintptrComparer),
	"float32":                           HideTypedComparer[float32](Float32Comparer),
	"float64":                           HideTypedComparer[float64](Float64Comparer),
	"complex64":                         HideTypedComparer[complex64](Complex64Comparer),
	"complex128":                        HideTypedComparer[complex128](Complex128Comparer),
	"string":                            HideTypedComparer[string](StringComparer),
	"bool":                              HideTypedComparer[bool](BoolComparer),
	fmt.Sprintf("%T", time.Time{}):      HideTypedComparer[time.Time](TimeComparer),
	fmt.Sprintf("%T", time.Duration(0)): HideTypedComparer[time.Duration](DurationComparer),
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
	comparer, found = mappedWrappedComparers[normalizeTypeName(typeName)]
	return
}

// RegisterDefaultComparerForType register a comparer for specific type into registry.
//
// Panic if duplicated for a type or comparer is nil or type name is empty/<nil>
func RegisterDefaultComparerForType[T any](typeName string, comparer IComparer[T], allowOverride bool) {
	typeName = normalizeTypeName(typeName)
	if len(typeName) < 1 || typeName == "<nil>" {
		panic("empty or <nil> type name")
	}

	if comparer == nil {
		panic("comparer is nil")
	}

	if !allowOverride {
		if existing, found := mappedWrappedComparers[typeName]; found && existing != nil {
			panic(fmt.Sprintf("default comparer for type [%s] had been registered before", typeName))
		}
	}

	mappedWrappedComparers[typeName] = HideTypedComparer[T](comparer)
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
	return strings.ToLower(strings.TrimSpace(typeName))
}

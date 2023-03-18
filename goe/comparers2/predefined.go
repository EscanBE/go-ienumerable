package comparers

import (
	"time"
)

// NumericComparer implements IComparer[any], this compares the real value of 2 input numeric,
// no matter it is int, float, complex or even types which defined on top of numeric like time.Duration (int64)
//
// > Eg 1: CompareTyped(int8(1), float64(1.1)) will results -1.
//
// _______________
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*int8(1), float64(1.1)), CompareAny(int8(1), *any(*float64(1.1))) all results -1.
//
// Panic if the final value detected is not numeric.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var NumericComparer IComparer[any] = NewNumericComparer()

// TimeComparer implements IComparer[time.Time], this compares the real value of 2 time.Time struct.
//
// If x before y, returns -1. If x after y, returns 1. Otherwise, returns 0.
//
// _______________
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*Time vs Time), CompareAny(Time, *any(*Time)).
//
// Panic if the final value detected is not time.Time struct.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var TimeComparer IComparer[time.Time] = NewTimeComparer()

// StringComparer implements IComparer[string], this compares the real value of 2 strings.
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*string vs string), CompareAny(string, *any(*string)).
//
// Panic if the final value detected is not a string.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var StringComparer IComparer[string] = NewStringComparer()

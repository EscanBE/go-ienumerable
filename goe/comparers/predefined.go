package comparers

import (
	"math/big"
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

// BoolComparer implements IComparer[bool], this compares the real value of 2 boolean with contract: false < true.
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*bool vs bool), CompareAny(bool, *any(*bool)).
//
// Panic if the final value detected is not a boolean.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var BoolComparer IComparer[bool] = NewBoolComparer()

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

// BigIntComparer implements IComparer[*big.Int], this compares the real value of 2 *big.Int struct.
//
// If x before y, returns -1. If x after y, returns 1. Otherwise, returns 0.
//
// _______________
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*big.Int vs *(big.Int)), CompareAny(*big.Int, *any(*big.Int)).
//
// Panic if the final value detected is not big.Int struct.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var BigIntComparer IComparer[*big.Int] = NewBigIntComparer()

// BigFloatComparer implements IComparer[*big.Float], this compares the real value of 2 *big.Float struct.
//
// If x before y, returns -1. If x after y, returns 1. Otherwise, returns 0.
//
// _______________
//
// Wrapped as interface or pointer are all accepted by compare using `CompareAny` method.
//
// > Eg 2: CompareAny(*big.Float vs *(big.Float)), CompareAny(*big.Float, *any(*big.Float)).
//
// Panic if the final value detected is not big.Float struct.
//
//goland:noinspection GoVarAndConstTypeMayBeOmitted
var BigFloatComparer IComparer[*big.Float] = NewBigFloatComparer()

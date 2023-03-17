package comparers

// NumericComparer implements IComparer, this compares the real value of 2 input numeric, no matter it is int, float or complex.
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
var NumericComparer IComparer[any] = numericComparer{}

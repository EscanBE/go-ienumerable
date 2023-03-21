package goe

// EqualsFunc is function returns true when and only when 'left' is equals to 'right'
type EqualsFunc[T any] func(left, right T) bool

// OptionalEqualsFunc is function returns true when and only when 'left' is equals to 'right'.
// This equality function is optional and will be resolved at runtime or using defined default comparer if not provided.
type OptionalEqualsFunc[T any] EqualsFunc[T]

// RequiredEqualsFunc is function returns true when and only when 'left' is equals to 'right'.
// This equality function is required and will cause panic if not provided.
type RequiredEqualsFunc[T any] EqualsFunc[T]

// CompareFunc is function:
//
// returns -1 when left < right
//
// returns 0 when left == right
//
// returns 1 when left > right
type CompareFunc[T any] func(left, right T) int

// OptionalCompareFunc is function:
//
// returns -1 when left < right
//
// returns 0 when left == right
//
// returns 1 when left > right
//
// This compare function is optional and will be resolved at runtime or using defined default comparer if not provided.
type OptionalCompareFunc[T any] func(left, right T) int

// Predicate is function that receives a value as input and returns boolean as output.
// Usually used as filters.
type Predicate[T any] func(value T) bool

// OptionalPredicate is function that receives a value as input and returns boolean as output.
// Usually used as filters.
// This predicate function is optional and will be resolved at runtime or using defined default comparer if not provided.
type OptionalPredicate[T any] Predicate[T]

// PredicateWithIndex is function that receives a value as input, along with an index value and returns boolean as output.
// Usually used as filters.
type PredicateWithIndex[T any] func(value T, index int) bool

// KeySelector is function that specify key that to be used for comparing elements within a collection.
type KeySelector[T any] func(value T) any

type ValueTuple2[T1, T2 any] struct {
	First  T1
	Second T2
}

type ValueTuple3[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

type Group[TKey, TElement any] struct {
	Key      TKey
	Elements TElement
}

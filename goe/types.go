package goe

// LessFunc is function returns true when and only when 'left' is less than 'right'
type LessFunc[T any] func(left, right T) bool

// GreaterFunc is function returns true when and only when 'left' is greater than 'right'
type GreaterFunc[T any] func(left, right T) bool

// EqualsFunc is function returns true when and only when 'left' is equals to 'right'
type EqualsFunc[T any] func(left, right T) bool

// CompareFunc is function:
//
// returns 0 when left == right
//
// returns -1 when left < right
//
// returns 1 when left > right
type CompareFunc[T any] func(left, right T) int

// Predicate is function that receives a value as input and returns boolean as output.
// Usually used as filters.
type Predicate[T any] func(value T) bool

// PredicateWithIndex is function that receives a value as input, along with an index value and returns boolean as output.
// Usually used as filters.
type PredicateWithIndex[T any] func(value T, index int) bool

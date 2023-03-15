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

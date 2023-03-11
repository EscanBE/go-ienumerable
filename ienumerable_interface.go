package go_ienumerable

type IEnumerable[T any] interface {
	// C#

	//Aggregate(fn interface{}, memo interface{}) IEnumerable

	// All returns true if all elements matches with predicate, also true when empty
	All(predicate func(T) bool) bool

	// Any determines whether a sequence contains any elements.
	Any() bool

	// AnyBy determines whether any element of a sequence satisfies a condition.
	AnyBy(predicate func(T) bool) bool

	// Count returns the number of elements in a sequence.
	Count() int

	// CountBy returns a number that represents how many elements in the specified sequence satisfy a condition.
	CountBy(predicate func(T) bool) int

	// Distinct returns distinct elements from a sequence.
	//
	// Require: equality comparer provided via WithEqualsComparer
	Distinct() IEnumerable[T]

	// DistinctBy returns distinct elements from a sequence by using the
	// specified equality comparer to compare values.
	DistinctBy(equalsComparer func(v1, v2 T) bool) IEnumerable[T]

	// Except produces the set difference of two sequences.
	//
	// Require: equality comparer provided via WithEqualsComparer
	Except(second IEnumerable[T]) IEnumerable[T]

	// ExceptBy produces the set difference of two sequences by using the
	// specified equality comparer to compare values.
	ExceptBy(second IEnumerable[T], equalsComparer func(v1, v2 T) bool) IEnumerable[T]

	// First returns the first element of a sequence
	First() T

	// FirstBy returns the first element in a sequence that satisfies a specified condition
	FirstBy(predicate func(T) bool) T

	// FirstOrDefault returns the first element of a sequence, or a specified default value if the sequence contains no elements.
	FirstOrDefault(defaultValue T) T

	// FirstOrDefaultBy returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found
	FirstOrDefaultBy(predicate func(T) bool, defaultValue T) T

	// GetEnumerator returns an enumerator that iterates through a collection.
	GetEnumerator() IEnumerator[T]

	//GroupBy(fieldName string) enumerable

	// Last returns the last element of a sequence
	Last() T

	// LastBy returns the last element in a sequence that satisfies a specified condition
	LastBy(predicate func(T) bool) T

	// LastOrDefault returns the last element of a sequence, or a specified default value if the sequence contains no elements.
	LastOrDefault(defaultValue T) T

	// LastOrDefaultBy returns the last element of the sequence that satisfies a condition, or a specified default value if no such element is found
	LastOrDefaultBy(predicate func(T) bool, defaultValue T) T

	// Order sorts the elements of a sequence in ascending order.
	//
	// Require: less comparer provided via WithLessComparer
	Order() IEnumerable[T]

	// OrderBy sorts the elements of a sequence in ascending order.
	// specified less comparer to compare values.
	OrderBy(lessComparer func(left, right T) bool) IEnumerable[T]

	// OrderByDescending sorts the elements of a sequence in descending order.
	//
	// Require: less comparer provided via WithLessComparer
	OrderByDescending() IEnumerable[T]

	// OrderByDescendingBy sorts the elements of a sequence in descending order.
	// specified less comparer to compare values.
	OrderByDescendingBy(lessComparer func(left, right T) bool) IEnumerable[T]

	// Reverse inverts the order of the elements in a sequence.
	Reverse() IEnumerable[T]

	//Select(selector interface{}) IEnumerable
	//SelectMany(selector interface{}) IEnumerable

	// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
	Skip(count int) IEnumerable[T]

	// Take returns a specified number of contiguous elements from the start of a sequence.
	Take(count int) IEnumerable[T]

	// ToArray creates an array from a IEnumerable[T].
	ToArray() []T

	// Union produces the set union of two sequences.
	//
	// Require: equality comparer provided via WithEqualsComparer
	Union(second IEnumerable[T]) IEnumerable[T]

	// UnionBy produces the set union of two sequences by using the
	// specified equality comparer to compare values.
	UnionBy(second IEnumerable[T], equalsComparer func(v1, v2 T) bool) IEnumerable[T]

	// Where filters a sequence of values based on a predicate.
	Where(predicate func(T) bool) IEnumerable[T]

	// Extra

	// FirstSafe returns the first element of a sequence, with error if sequence contains no element
	FirstSafe() (T, error)

	// FirstSafeBy returns the first element in a sequence that satisfies a specified condition, with error if sequence contains no element or predicate is nil
	FirstSafeBy(predicate func(T) bool) (T, error)

	// FirstOrDefaultSafeBy returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found, with error if predicate is nil
	FirstOrDefaultSafeBy(predicate func(T) bool, defaultValue T) (T, error)

	// LastOrDefaultSafeBy returns the last element of the sequence that satisfies a condition, or a specified default value if no such element is found, with error if predicate is nil
	LastOrDefaultSafeBy(predicate func(T) bool, defaultValue T) (T, error)

	// LastSafe returns the last element of a sequence, with error if sequence contains no element
	LastSafe() (T, error)

	// LastSafeBy returns the last element in a sequence that satisfies a specified condition, with error if sequence contains no element or predicate is nil
	LastSafeBy(predicate func(T) bool) (T, error)

	// Extra: the following methods are used to inject comparer into IEnumerable[T] instance
	// and those comparers is going to be used for methods like: Distinct, Order, etc

	// WithEqualsComparer the equality comparer, to indicate if 2 source values are equals, will be embedded
	// into this IEnumerable which automatically be used for the following methods:
	// Except, Distinct, Union
	WithEqualsComparer(equalsComparer func(v1, v2 T) bool) IEnumerable[T]

	// WithLessComparer the less comparer, to indicate if left source is lower than right source, will be embedded
	// into this IEnumerable which automatically be used for the following methods:
	// Order, OrderByDescending
	WithLessComparer(lessComparer func(left, right T) bool) IEnumerable[T]

	// WithDefaultComparers automatically detect type of T and inject comparers of the corresponding type,
	// equals to call WithEqualsComparer and WithLessComparer with predefined comparer.
	//
	// If the type of T is not a supported type, a panic will be raised.
	//
	// Supported types: int8/16/32/64, uint8/16/32/64, int, uint, uintptr, float32/64, complex64/128 (equality comparer only), string
	WithDefaultComparers() IEnumerable[T]

	// The following methods are internal APIs

	exposeData() []T
	len() int
	copy() IEnumerable[T]
}

package go_ienumerable

type IEnumerable[T any] interface {
	// C#

	// Aggregate applies an accumulator function over a sequence.
	Aggregate(f func(pr, v T) T) T

	// AggregateWithSeed applies an accumulator function over a sequence.
	// The specified seed value is used as the initial accumulator value.
	AggregateWithSeed(seed T, f func(pr, v T) T) T

	// AggregateWithAnySeed applies an accumulator function over a sequence.
	// The specified seed value is used as the initial accumulator value.
	//
	// Notice, the type (specified as 'any') of the seed and the aggregate function `f` param and result,
	// must be the same type
	AggregateWithAnySeed(seed any, f func(pr any, v T) any) any

	// All returns true if all elements matches with predicate, also true when empty
	All(predicate func(T) bool) bool

	// Any determines whether a sequence contains any elements.
	Any() bool

	// AnyBy determines whether any element of a sequence satisfies a condition.
	AnyBy(predicate func(T) bool) bool

	// Append appends a value to the end of the sequence and return a new sequence ends with input `element`
	Append(element T) IEnumerable[T]

	// AsEnumerable() IEnumerable[T] <= will not be implemented because no inheritance in Go thus never use

	// Average computes the average of a sequence of integer/float values.
	Average() float64

	// CastByte casts the source IEnumerable[T] into IEnumerable[byte]
	// if the source data type is byte (uint8), otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[byte], even default
	CastByte() IEnumerable[byte]

	// CastInt32 casts the source IEnumerable[T] into IEnumerable[int32]
	// if the source data type is int32, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int32]
	CastInt32() IEnumerable[int32]

	// CastInt64 casts the source IEnumerable[T] into IEnumerable[int64]
	// if the source data type is int64, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int64]
	CastInt64() IEnumerable[int64]

	// CastInt casts the source IEnumerable[T] into IEnumerable[int]
	// if the source data type is int, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int]
	CastInt() IEnumerable[int]

	// CastFloat64 casts the source IEnumerable[T] into IEnumerable[float64]
	// if the source data type is float64, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[float64]
	CastFloat64() IEnumerable[float64]

	// CastString casts the source IEnumerable[T] into IEnumerable[string]
	// if the source data type is string, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[string]
	CastString() IEnumerable[string]

	// CastBool casts the source IEnumerable[T] into IEnumerable[bool]
	// if the source data type is bool, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[bool]
	CastBool() IEnumerable[bool]

	// ChunkToHolder (Chunk) supposed to split the elements of a sequence into chunks of size at most size.
	// Use method GetChunkedIEnumeratorFromHolder to convert from ChunkHolder[T] back to IEnumerable[[]T].
	//
	// Due to limitation of Golang that can not define a method signature like
	//
	// `func (src *enumerable[T]) Chunk(size int) IEnumerable[[]T]`
	//
	// so the method Chunk is temporary renamed to ChunkToHolder and ChunkToAny, the name Chunk is reserved for future
	// for implementation when Go supports the above method signature.
	//
	// ChunkToHolder with result ChunkHolder designed as a stepping stone
	// then can use it to convert it back to IEnumerable[[]T] via GetChunkedIEnumeratorFromHolder function
	ChunkToHolder(size int) ChunkHolder[T]

	// ChunkToAny (Chunk) splits the elements of a sequence into chunks of size at most size.
	//
	// Due to limitation of Golang that can not define a method signature like
	//
	// `func (src *enumerable[T]) Chunk(size int) IEnumerable[[]T]`
	//
	// so the method Chunk is temporary renamed to ChunkToAny and ChunkToHolder, the name Chunk is reserved for future
	// for implementation when Go supports the above method signature.
	//
	// ChunkToAny with result IEnumerable[[]any] is not really a nice way since we have to convert it back to original type of it.
	ChunkToAny(size int) IEnumerable[[]any]

	// Concat concatenates two sequences.
	Concat(second IEnumerable[T]) IEnumerable[T]

	// Contains determines whether a sequence contains a specified element.
	//
	// Beware of compare numeric when IEnumerable[any] because int8(1) is not equals to int16(1), int32(1)...
	//
	// Require: equality comparer provided via WithEqualsComparer
	Contains(value T) bool

	// ContainsBy determines whether a sequence contains a specified element by using the specified equality comparer.
	//
	// Beware of compare numeric when IEnumerable[any] because int8(1) is not equals to int16(1), int32(1)...
	ContainsBy(value T, equalityComparer func(v1, v2 T) bool) bool

	// Count returns the number of elements in a sequence.
	Count() int

	// CountBy returns a number that represents how many elements in the specified sequence satisfy a condition.
	CountBy(predicate func(T) bool) int

	// DefaultIfEmpty returns the elements of the specified sequence
	// or the type parameter's default value in a singleton collection if the sequence is empty.
	DefaultIfEmpty() IEnumerable[T]

	// DefaultIfEmptyUsing returns the elements of the specified sequence
	// or the specified value in a singleton collection if the sequence is empty.
	DefaultIfEmptyUsing(defaultValue T) IEnumerable[T]

	// Distinct returns distinct elements from a sequence.
	//
	// Require: equality comparer provided via WithEqualsComparer
	Distinct() IEnumerable[T]

	// DistinctBy returns distinct elements from a sequence by using the
	// specified equality comparer to compare values.
	DistinctBy(equalsComparer func(v1, v2 T) bool) IEnumerable[T]

	// ElementAt returns the element at a specified index (0 based, from head) in a sequence.
	//
	// Panic if index is less than 0 or greater than or equal to the number of elements in source
	ElementAt(index int) T

	// ElementAtReverse returns the element at a specified reverse index (0 based, from tail) in a sequence.
	//
	// Panic if index is less than 0 or greater than or equal to the number of elements in source
	ElementAtReverse(reverseIndex int) T

	// ElementAtOrDefault returns the element at a specified index (0 based, from head) in a sequence.
	// If index is out of range, return default value of type.
	//
	// Beware of IEnumerable[any|interface{}], you will get nil no matter real type of underlying data is
	ElementAtOrDefault(index int) T

	// ElementAtReverseOrDefault returns the element at a specified reverse index (0 based, from tail) in a sequence.
	// If index is out of range, return default value of type.
	//
	// Beware of IEnumerable[any|interface{}], you will get nil no matter real type of underlying data is
	ElementAtReverseOrDefault(reverseIndex int) T

	// Empty returns an empty IEnumerable[T] that has the specified type argument.
	Empty() IEnumerable[T]

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

	// FirstOrDefault returns the first element of a sequence, or a default value of type if the sequence contains no elements.
	FirstOrDefault() T

	// FirstOrDefaultBy returns the first element of the sequence that satisfies a condition, or a default value of type if no such element is found
	FirstOrDefaultBy(predicate func(T) bool) T

	// FirstOrDefaultUsing returns the first element of a sequence, or a specified default value if the sequence contains no elements.
	FirstOrDefaultUsing(defaultValue T) T

	// FirstOrDefaultByUsing returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found
	FirstOrDefaultByUsing(predicate func(T) bool, defaultValue T) T

	// GetEnumerator returns an enumerator that iterates through a collection.
	GetEnumerator() IEnumerator[T]

	//GroupBy(fieldName string) enumerable

	// Last returns the last element of a sequence
	Last() T

	// LastBy returns the last element in a sequence that satisfies a specified condition
	LastBy(predicate func(T) bool) T

	// LastOrDefault returns the last element of a sequence, or a default value of type if the sequence contains no elements.
	LastOrDefault() T

	// LastOrDefaultBy returns the last element of the sequence that satisfies a condition, or a default value of type if no such element is found
	LastOrDefaultBy(predicate func(T) bool) T

	// LastOrDefaultUsing returns the last element of a sequence, or a specified default value if the sequence contains no elements.
	LastOrDefaultUsing(defaultValue T) T

	// LastOrDefaultByUsing returns the last element of the sequence that satisfies a condition, or a specified default value if no such element is found
	LastOrDefaultByUsing(predicate func(T) bool, defaultValue T) T

	// LongCount returns an int64 that represents the number of elements in a sequence.
	//
	// This method is meaning-less in Go because the Count method returns an int already has max value of int64 in x64 machines
	LongCount() int64

	// LongCountBy returns an int64 that represents how many elements in the specified sequence satisfy a condition.
	//
	// This method is meaning-less in Go because the CountBy method returns an int already has max value of int64 in x64 machines
	LongCountBy(predicate func(T) bool) int64

	// Min returns the minimum value in a sequence.
	//
	// Require: less comparer provided via WithLessComparer
	Min() T

	// MinBy returns the minimum value in a sequence according to provided comparer.
	MinBy(lessComparer func(left, right T) bool) T

	// Max returns the minimum value in a sequence.
	//
	// Require: less comparer provided via WithLessComparer
	Max() T

	// MaxBy returns the minimum value in a sequence according to provided LESS comparer.
	MaxBy(lessComparer func(left, right T) bool) T

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

	// Prepend adds a value to the beginning of the sequence and return a new sequence starts with input `element`
	Prepend(element T) IEnumerable[T]

	// Reverse inverts the order of the elements in a sequence.
	Reverse() IEnumerable[T]

	// Select projects each element of a sequence into a new form.
	//
	// Due to limitation of current Go, there is no way to directly cast into target type
	// in just one command, so additional transform from 'any' to target types is required.
	//
	// There are some Cast* methods
	// CastByte, CastInt, CastString, ... so can do combo like example:
	//
	// IEnumerable[int](src).Select(x => x + 1).CastInt() and will result IEnumerable[int]
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[any]
	Select(selector func(v T) any) IEnumerable[any]

	// SelectMany projects each element of a sequence to an IEnumerable[T]
	// and flattens the resulting sequences into one sequence.
	//
	// Due to limitation of current Go, there is no way to directly cast into target type
	// in just one command, so additional transform from 'any' to target types is required.
	//
	// There are some Cast* methods
	// CastByte, CastInt, CastString, ... so can do combo like example:
	//
	// IEnumerable[[]int](src).SelectMany([]int{x,y} => []any{x * 2, y * 2}).CastInt() and will result IEnumerable[int]
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[any]
	SelectMany(selector func(v T) []any) IEnumerable[any]

	// Single returns the only element of a sequence,
	// and panic if there is not exactly one element in the sequence.
	Single() T

	// SingleBy returns the only element of a sequence that satisfies a specified condition,
	// and panic if more than one such element exists.
	SingleBy(predicate func(T) bool) T

	// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
	Skip(count int) IEnumerable[T]

	// SumInt32 computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum result is overflow int32
	//
	// Notice 2: will panic if during sum, value is overflow int64
	//
	// Notice 3: will panic if element in array is not integer or is integer but overflow int32
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64)
	SumInt32() int32

	// SumInt computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum result is overflow int
	//
	// Notice 2: will panic if during sum, value is overflow int64
	//
	// Notice 3: will panic if element in array is not integer or is integer but overflow int
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64)
	SumInt() int

	// SumInt64 computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum is overflow int64
	//
	// Notice 2: will panic if element in array is not integer
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64)
	SumInt64() int64

	// SumFloat64 computes the sum of a sequence of integer/float values.
	//
	// Notice 1: will panic if sum is overflow float64
	//
	// Notice 2: will panic if element in array is not integer/float or is integer/float but overflow float64
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64 + accepted floats: float32/64)
	SumFloat64() float64

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

	// From this part, extra methods are defined to provide more utilities and/or to workaround
	// limitation of Golang compares to C#

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
	exposeDataType() string
	len() int

	// unboxAnyAsByte unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into byte.
	//
	// Panic if value is over range or not integer
	unboxAnyAsByte(v T) byte

	// unboxAnyAsInt32 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int32.
	//
	// Panic if value is over range or not integer
	unboxAnyAsInt32(v T) int32

	// unboxAnyAsInt unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int.
	//
	// Panic if value is over range or not integer
	unboxAnyAsInt(v T) int

	// unboxAnyAsInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) into int64.
	//
	// Panic if value is over range or not integer
	unboxAnyAsInt64(v T) int64

	// unboxAnyAsFloat64OrInt64 unbox any integer (int, int8/16/32/64, uint, uint8/16/32/64) or float32/64
	// into either int64 or float64 value (priority int64), data type specified in result.
	// This design is for sum accuracy.
	//
	// Panic if neither integer nor float
	unboxAnyAsFloat64OrInt64(v T) (rf float64, ri int64, dt unboxFloat64DataType)
}

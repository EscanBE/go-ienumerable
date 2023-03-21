package goe

import "github.com/EscanBE/go-ienumerable/goe/comparers"

// IEnumerable from C#, brought to Golang by VictorTrustyDev
//
//goland:noinspection GoSnakeCaseUsage
type IEnumerable[T any] interface {
	// Aggregate applies an accumulator function over a sequence.
	Aggregate(f func(previousValue, value T) T) T

	// AggregateSeed applies an accumulator function over a sequence.
	// The specified seed value is used as the initial accumulator value.
	AggregateSeed(seed T, f func(previousValue, value T) T) T

	// AggregateAnySeed applies an accumulator function over a sequence.
	// The specified seed value is used as the initial accumulator value.
	//
	// Contract: the type of the seed and the aggregate function `f` param and result,
	// must be the same type
	AggregateAnySeed(seed any, f func(previousValue any, value T) any) any

	// Aggregate_ImplementedInHelper aggregate methods are also implemented as helper, use the Aggregate methods from the helper package for method signature more likely C#.
	Aggregate_ImplementedInHelper()

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
	//
	// Panic if element can not be cast to number.
	Average() float64

	// CastByte casts the source IEnumerable[T] into IEnumerable[byte]
	// if the source data type is byte (uint8), otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[byte],
	// a default comparer will be assigned automatically if able to resolve.
	CastByte() IEnumerable[byte]

	// CastInt32 casts the source IEnumerable[T] into IEnumerable[int32]
	// if the source data type is int32, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int32],
	// a default comparer will be assigned automatically if able to resolve.
	CastInt32() IEnumerable[int32]

	// CastInt64 casts the source IEnumerable[T] into IEnumerable[int64]
	// if the source data type is int64, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int64],
	// a default comparer will be assigned automatically if able to resolve.
	CastInt64() IEnumerable[int64]

	// CastInt casts the source IEnumerable[T] into IEnumerable[int]
	// if the source data type is int, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[int],
	// a default comparer will be assigned automatically if able to resolve.
	CastInt() IEnumerable[int]

	// CastFloat64 casts the source IEnumerable[T] into IEnumerable[float64]
	// if the source data type is float64, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[float64],
	// a default comparer will be assigned automatically if able to resolve.
	CastFloat64() IEnumerable[float64]

	// CastString casts the source IEnumerable[T] into IEnumerable[string]
	// if the source data type is string, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[string],
	// a default comparer will be assigned automatically if able to resolve.
	CastString() IEnumerable[string]

	// CastBool casts the source IEnumerable[T] into IEnumerable[bool]
	// if the source data type is bool, otherwise panic.
	//
	// Notice: no comparer from source will be brought along with new IEnumerable[bool],
	// a default comparer will be assigned automatically if able to resolve.
	CastBool() IEnumerable[bool]

	// ChunkToHolder (known as Chunk) supposed to split the elements of a sequence into chunks of size at most size.
	// Use method GetChunkedIEnumeratorFromHolder to convert from ChunkHolder[T] back to IEnumerable[[]T].
	//
	// Suggestion: use helper function goe_helper.Chunk from helper package.
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
	// Suggestion: use helper function goe_helper.Chunk from helper package.
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

	// Chunk_ImplementedInHelper the Chunk method is also implemented as helper, use the Chunk method from the helper package for method signature more likely C#.
	Chunk_ImplementedInHelper()

	// Concat concatenates two sequences.
	Concat(second IEnumerable[T]) IEnumerable[T]

	// Contains determines whether a sequence contains a specified element.
	//
	// If passing nil as comparer function, the default comparer will be used or panic if no default comparer found.
	Contains(value T, optionalEqualsFunc OptionalEqualsFunc[T]) bool

	// Count returns a number that represents how many elements in the specified sequence satisfy a condition.
	Count(optionalPredicate OptionalPredicate[T]) int

	// DefaultIfEmpty returns the elements of the specified sequence
	// or the type parameter's default value in a singleton collection if the sequence is empty.
	DefaultIfEmpty() IEnumerable[T]

	// DefaultIfEmptyUsing returns the elements of the specified sequence
	// or the specified value in a singleton collection if the sequence is empty.
	DefaultIfEmptyUsing(defaultValue T) IEnumerable[T]

	// Distinct returns distinct elements from a sequence.
	//
	// If passing nil as comparer function, the default comparer will be used or panic if no default comparer found.
	Distinct(optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T]

	// DistinctBy returns distinct elements from a sequence according to a specified key selector function.
	//
	// If passing nil as comparer function, the default comparer will be used or panic if no default comparer found.
	DistinctBy(keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T]

	// ElementAt returns the element at a specified index (0 based, from head) in a sequence.
	//
	// When setting reverse to true, index is reverse index (0based, from tail).
	//
	// Panic if index is less than 0 or greater than or equal to the number of elements in source
	ElementAt(index int, reverse bool) T

	// ElementAtOrDefault returns the element at a specified index (0 based, from head) in a sequence.
	// If index is out of range, return default value of type.
	//
	// When setting reverse to true, index is reverse index (0based, from tail).
	//
	// Beware of IEnumerable[any|interface{}], you will get nil no matter real type of underlying data is
	ElementAtOrDefault(index int, reverse bool) T

	// Empty returns a new empty IEnumerable[T] that has the specified type argument.
	// Comparers will be copied into the new IEnumerable[T].
	Empty() IEnumerable[T]

	// Except produces the set difference of two sequences.
	//
	// If passing nil as comparer function, the default comparer will be used or panic if no default comparer found.
	Except(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T]

	// ExceptBy produces the set difference of two sequences according to a specified key selector function.
	//
	// _______________
	//
	// Contract: type of elements in second IEnumerable must be the same type with value returns by keySelector,
	// otherwise panic.
	//
	// _______________
	//
	// If passing nil as comparer function, the default comparer will be used or panic if no default comparer found.
	ExceptBy(second IEnumerable[any], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T]

	// First returns the first element of a sequence that satisfies a specified condition.
	//
	// If omitted the optional predicate, the first element will be returned.
	First(optionalPredicate OptionalPredicate[T]) T

	// FirstOrDefault returns the first element of a sequence, or a default value of type if the sequence contains no elements.
	//
	// If omitted the optional predicate, the first element will be returned.
	//
	// If omitted the optional default value param, default value of T will be returned.
	FirstOrDefault(optionalPredicate OptionalPredicate[T], optionalDefaultValue *T) T

	// GroupBy_ImplementedInHelper the GroupBy method is implemented as helper, use the GroupBy method from the helper package for method signature more likely C#.
	GroupBy_ImplementedInHelper()

	// GroupJoin_ImplementedInHelper the GroupJoin method is implemented as helper, use the GroupJoin method from the helper package for method signature more likely C#.
	GroupJoin_ImplementedInHelper()

	// GetEnumerator returns an enumerator that iterates through a collection.
	GetEnumerator() IEnumerator[T]

	//GroupBy(fieldName string) enumerable

	// Intersect produces the set intersection of two sequences.
	//
	// If omitted the optional equality comparer function, the default comparer will be used or panic if no default comparer found.
	Intersect(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T]

	// IntersectBy produces the set intersection of two sequences according to a specified key selector function and using the
	// optional equality-comparer to compare keys.
	//
	// If passing nil as equality comparer function, the default comparer will be used or panic if no default comparer found.
	IntersectBy(second IEnumerable[T], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T]

	// Join_ImplementedInHelper Join method are also implemented as helper, use the Join methods from the helper package for method signature more likely C#.
	Join_ImplementedInHelper()

	// Last returns the last element of a sequence that satisfies a specified condition.
	//
	// If omitted predicate, the last element will be returned.
	Last(optionalPredicate OptionalPredicate[T]) T

	// LastOrDefault returns the last element of a sequence, or a default value of type if the sequence contains no elements.
	//
	// If omitted predicate, the last element will be returned.
	//
	// If omitted the optional default value param, default value of T will be returned.
	LastOrDefault(optionalPredicate OptionalPredicate[T], optionalDefaultValue *T) T

	// LongCount returns an int64 that represents how many elements in the specified sequence satisfy an optional condition.
	//
	// This method is meaning-less in Go because the Count method returns an int already has max value of int64 in x64 machines
	LongCount(optionalPredicate OptionalPredicate[T]) int64

	// Min returns the minimum value in a sequence.
	//
	// Require: type must be registered for default comparer
	// or already set via WithDefaultComparer / WithDefaultComparerAny / WithComparerFrom,
	// otherwise panic.
	Min() T

	// MinBy returns the minimum value in a generic sequence according to a specified key selector function
	// and key comparer.
	//
	// If omitted the compareFunc, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	MinBy(keySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) T

	// Max returns the greatest value in a sequence.
	//
	// Require: type must be registered for default comparer
	// or already set via WithDefaultComparer / WithDefaultComparerAny / WithComparerFrom,
	// otherwise panic.
	Max() T

	// MaxBy returns the maximum value in a generic sequence according to a specified key selector function
	// and key comparer.
	//
	// If omitted the compare func, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	MaxBy(keySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) T

	// OfType_ImplementedInHelper OfType method are also implemented as helper, use the OfType methods from the helper package for method signature more likely C#.
	OfType_ImplementedInHelper()

	// Order sorts the elements of a sequence in ascending order.
	//
	// This method is implemented by using deferred execution,
	// that means you have to call `GetOrderedEnumerable` method
	// of the IOrderedEnumerable to invoke sorting and get the sorted IEnumerable.
	Order() IOrderedEnumerable[T]

	// OrderBy sorts the elements of a sequence in ascending order
	// according to the selected key.
	//
	// ________
	//
	// keySelector is required, compare function is optional.
	//
	// If omitted the compareFunc, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	//
	// This method is implemented by using deferred execution,
	// that means you have to call `GetOrderedEnumerable` method
	// of the IOrderedEnumerable to invoke sorting and get the sorted IEnumerable.
	OrderBy(keySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) IOrderedEnumerable[T]

	// OrderDescending sorts the elements of a sequence in descending order.
	//
	// This method is implemented by using deferred execution,
	// that means you have to call `GetOrderedEnumerable` method
	// of the IOrderedEnumerable to invoke sorting and get the sorted IEnumerable.
	OrderDescending() IOrderedEnumerable[T]

	// OrderByDescending sorts the elements of a sequence in descending order
	// according to the selected key.
	//
	// ________
	//
	// keySelector is required, compare function is optional.
	//
	// If omitted the optional compare function, the default comparer for corresponding type will be used,
	// or panic if no default compare found.
	//
	// This method is implemented by using deferred execution,
	// that means you have to call `GetOrderedEnumerable` method
	// of the IOrderedEnumerable to invoke sorting and get the sorted IEnumerable.
	OrderByDescending(keySelector KeySelector[T], optionalCompareFunc OptionalCompareFunc[any]) IOrderedEnumerable[T]

	// Prepend adds a value to the beginning of the sequence and return a new sequence starts with input `element`
	Prepend(element T) IEnumerable[T]

	// Range_ImplementedInHelper Range method are also implemented as helper, use the Range methods from the helper package for method signature more likely C#.
	Range_ImplementedInHelper()

	// Repeat generates a new sequence that contains one repeated value.
	//
	// Panic if count is less than 0
	Repeat(element T, count int) IEnumerable[T]

	// Reverse inverts the order of the elements in a sequence.
	Reverse() IEnumerable[T]

	// Select projects each element of a sequence into a new form,
	// if want to keep original data type, use SelectNewValue instead.
	//
	// Due to limitation of current Go, there is no way to directly cast into target type
	// in just one command, so additional transform from 'any' to target types might be required.
	//
	// There are some Cast* methods
	// CastByte, CastInt, CastString, ... so can do combo like example:
	//
	// IEnumerable[int](src).Select(x => x + 1).CastInt() and will result IEnumerable[int]
	//
	// Notice:
	//
	// - Comparer from source will NOT be brought along with new IEnumerable[any]
	//
	// - A default comparer will be assigned automatically if able to resolve.
	//
	// - It is not able to resolve default comparer for result type if sequence contains no element.
	//
	// - If not able to auto-resolve a default comparer for type of result,
	// you might need to specify comparers.IComparer[any] manually via WithDefaultComparer / WithDefaultComparerAny,
	// otherwise there is panic when you call methods where comparer is needed, like Distinct, Order,...
	Select(selector func(v T) any) IEnumerable[any]

	// SelectNewValue projects each element of a sequence into a new value , keep the original type.
	//
	// Notice: Existing comparer from source will be brought along with the new IEnumerable[T].
	SelectNewValue(selector func(v T) T) IEnumerable[T]

	// SelectMany projects each element of a sequence to an array of interface
	// and flattens the resulting sequences into one sequence: IEnumerable[any]
	//
	// Due to limitation of current Go, there is no way to directly cast into target type
	// in just one command, so additional transform from 'any' to target types is required.
	//
	// There are some Cast* methods
	// CastByte, CastInt, CastString, ... so can do combo like example:
	//
	// IEnumerable[[]int](src).SelectMany([]int{x,y} => []any{x * 2, y * 2}).CastInt() and will result IEnumerable[int]
	//
	// Notice:
	//
	// - Comparer from source will NOT be brought along with new IEnumerable[any]
	//
	// - A default comparer will be assigned automatically if able to resolve.
	//
	// - It is not able to resolve default comparer for result type if sequence contains no element.
	//
	// - If not able to auto-resolve a default comparer for type of result,
	// you might need to specify comparers.IComparer[any] manually via WithDefaultComparer / WithDefaultComparerAny,
	// otherwise there is panic when you call methods where comparer is needed, like Distinct, Order,...
	//
	// - Panic if selector returns nil
	SelectMany(selector func(v T) []any) IEnumerable[any]

	// Select_ImplementedInHelper select methods are also implemented as helper, use the Select* methods from the helper package for method signature more likely C#.
	Select_ImplementedInHelper()

	// SequenceEqual determines whether two sequences are equal according to an equality comparer.
	//
	// If passing nil as equality comparer function, the default comparer will be used or panic if no default comparer found.
	SequenceEqual(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) bool

	// Single returns the only element of a sequence that satisfies an optional condition,
	// and panic if more than one such element exists.
	Single(optionalPredicate OptionalPredicate[T]) T

	// SingleOrDefault returns the only element of a sequence that satisfies an optional condition
	// or a default value of type if no such element exists;
	// this method panics if more than one element satisfies the condition.
	//
	// If omitted the optional predicate, the only one element will be returned,
	// or panic if more than one,
	// or default value (provided by optional default value params or default of T) if no element,
	//
	//
	// If omitted the optional default value param, default value of T will be returned.
	SingleOrDefault(optionalPredicate OptionalPredicate[T], optionalDefaultValue *T) T

	// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
	Skip(count int) IEnumerable[T]

	// SkipLast returns a new enumerable collection that contains the elements from source
	// with the last count elements of the source collection omitted.
	SkipLast(count int) IEnumerable[T]

	// SkipWhile bypasses elements in a sequence as long as a specified condition is true
	//
	// The predicate param is required, must be either: Predicate[T] or PredicateWithIndex[T].
	SkipWhile(predicate interface{}) IEnumerable[T]

	// SumInt32 computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum result is overflow int32
	//
	// Notice 2: will panic if during sum, value is overflow int64
	//
	// Notice 3: will panic if element in sequence is not integer or is integer but overflow int32
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64).
	// But if sequence is empty, returns 0.
	SumInt32() int32

	// SumInt computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum result is overflow int
	//
	// Notice 2: will panic if during sum, value is overflow int64
	//
	// Notice 3: will panic if element in sequence is not integer or is integer but overflow int
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64).
	// But if sequence is empty, returns 0.
	SumInt() int

	// SumInt64 computes the sum of a sequence of integer values.
	//
	// Notice 1: will panic if sum is overflow int64
	//
	// Notice 2: will panic if element in sequence is not integer
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64).
	// But if sequence is empty, returns 0.
	SumInt64() int64

	// SumFloat64 computes the sum of a sequence of integer/float values.
	//
	// Notice 1: will panic if sum is overflow float64
	//
	// Notice 2: will panic if element in sequence is not integer/float or is integer/float but overflow float64
	// (accepted integers: int or int8/16/32/64, uint or uint/8/16/32/64 + accepted floats: float32/64).
	// But if sequence is empty, returns 0.
	SumFloat64() float64

	// Take returns a specified number of contiguous elements from the start of a sequence.
	Take(count int) IEnumerable[T]

	// TakeLast returns a new enumerable collection that contains the last count elements from source.
	TakeLast(count int) IEnumerable[T]

	// TakeWhile returns elements from a sequence as long as a specified condition is true.
	//
	// The predicate param is required, must be either: Predicate[T] or PredicateWithIndex[T].
	TakeWhile(predicate interface{}) IEnumerable[T]

	// ToArray creates an array from a IEnumerable[T].
	ToArray() []T

	// ToDictionary_ImplementedInHelper the ToDictionary method is also implemented as helper, use the ToDictionary method from the helper package for method signature more likely C#.
	ToDictionary_ImplementedInHelper()

	// TODO ToHashSet

	// TODO ToList

	// TODO ToLookUp

	// Union produces the set union of two sequences by using an optional equality function to compare values.
	//
	// If passing nil as equality comparer function, the default comparer will be used or panic if no default comparer found.
	Union(second IEnumerable[T], optionalEqualsFunc OptionalEqualsFunc[T]) IEnumerable[T]

	// UnionBy produces the set union of two sequences according to a specified key selector function and using the
	// optional equality-comparer to compare keys.
	//
	// If passing nil as equality comparer function, the default comparer will be used or panic if no default comparer found.
	UnionBy(second IEnumerable[T], keySelector KeySelector[T], optionalEqualsFunc OptionalEqualsFunc[any]) IEnumerable[T]

	// Where filters a sequence of values based on a predicate.
	Where(predicate func(T) bool) IEnumerable[T]

	// Zip_ImplementedInHelper Zip methods are implemented as helper, use the Zip methods from the helper package
	Zip_ImplementedInHelper()

	// From this part, extra methods are defined to provide more utilities and/or to workaround
	// limitation of Golang compares to C#

	// Extra: the following methods are used to inject comparers.IComparer into IEnumerable[T] instance
	// and those comparers is going to be used for methods like: Distinct, Order, etc...

	// WithComparerFrom copies existing comparer from the other IEnumerable[T] specified as parameter
	WithComparerFrom(copyFrom IEnumerable[T]) IEnumerable[T]

	// WithDefaultComparer setting default comparer to be used in this IEnumerable[T].
	//
	// If any existing (previously set or automatically detected) will be overridden.
	//
	// Setting to nil will remove existing if any.
	WithDefaultComparer(comparer comparers.IComparer[T]) IEnumerable[T]

	// WithDefaultComparerAny setting default comparer to be used in this IEnumerable[T].
	//
	// If any existing (previously set or automatically detected) will be overridden.
	//
	// Setting to nil will remove existing if any.
	WithDefaultComparerAny(comparer comparers.IComparer[any]) IEnumerable[T]
}

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

	//Distinct(selector interface{}) IEnumerable
	//DistinctBy(fieldName string) IEnumerable
	//Each(action interface{})
	//Except(predicate interface{}) IEnumerable
	//ExceptBy(fields map[string]interface{}) IEnumerable
	//Filter(predicate interface{}) IEnumerable
	//FilterBy(fields map[string]interface{}) IEnumerable
	//Find(predicate interface{}) IEnumerable
	//FindBy(fields map[string]interface{}) IEnumerable
	//FindIndex(predicate interface{}) int
	//FindIndexBy(fields map[string]interface{}) int

	// First returns the first element of a sequence
	First() T

	// FirstBy returns the first element in a sequence that satisfies a specified condition
	FirstBy(predicate func(T) bool) T

	// FirstOrDefault returns the first element of a sequence, or a specified default value if the sequence contains no elements.
	FirstOrDefault(defaultValue T) T

	// FirstOrDefaultBy returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found
	FirstOrDefaultBy(predicate func(T) bool, defaultValue T) T

	//GetEnumerator() IEnumerator
	//Group(keySelector interface{}) enumerable
	//GroupBy(fieldName string) enumerable
	//Index(keySelector interface{}) IEnumerable
	//IndexBy(fieldName string) IEnumerable
	//Keys() IEnumerable

	// Last returns the last element of a sequence
	Last() T

	// LastBy returns the last element in a sequence that satisfies a specified condition
	LastBy(predicate func(T) bool) T

	//Map(selector interface{}) IEnumerable
	//MapBy(fieldName string) IEnumerable
	//MapMany(selector interface{}) IEnumerable
	//MapManyBy(fieldName string) IEnumerable
	//Object() IEnumerable
	//Order(selector interface{}) IEnumerable
	//OrderBy(fieldName string) IEnumerable
	//Reduce(fn interface{}, memo interface{}) IEnumerable
	//Reject(predicate interface{}) IEnumerable
	//RejectBy(fields map[string]interface{}) IEnumerable
	//Reverse(selector interface{}) IEnumerable
	//ReverseBy(fieldName string) IEnumerable
	//Select(selector interface{}) IEnumerable
	//SelectBy(fieldName string) IEnumerable
	//SelectMany(selector interface{}) IEnumerable
	//SelectManyBy(fieldName string) IEnumerable
	//Sort(selector interface{}) IEnumerable
	//SortBy(fieldName string) IEnumerable

	// Skip bypasses a specified number of elements in a sequence and then returns the remaining elements.
	Skip(count int) IEnumerable[T]

	// Take returns a specified number of contiguous elements from the start of a sequence.
	Take(count int) IEnumerable[T]

	//Uniq(selector interface{}) IEnumerable
	//UniqBy(fieldName string) IEnumerable
	//Value(res interface{})
	//Values() IEnumerable

	// Where filters a sequence of values based on a predicate.
	Where(predicate func(T) bool) IEnumerable[T]

	// ToArray creates an array from a IEnumerable[T].
	ToArray() []T

	// Extra

	// FirstSafe returns the first element of a sequence, with error if sequence contains no element
	FirstSafe() (T, error)

	// FirstSafeBy returns the first element in a sequence that satisfies a specified condition, with error if sequence contains no element or predicate is nil
	FirstSafeBy(predicate func(T) bool) (T, error)

	// FirstOrDefaultSafeBy returns the first element of the sequence that satisfies a condition, or a specified default value if no such element is found, with error if predicate is nil
	FirstOrDefaultSafeBy(predicate func(T) bool, defaultValue T) (T, error)

	// LastSafe returns the last element of a sequence, with error if sequence contains no element
	LastSafe() (T, error)

	// LastSafeBy returns the last element in a sequence that satisfies a specified condition, with error if sequence contains no element or predicate is nil
	LastSafeBy(predicate func(T) bool) (T, error)

	// Extra comparators

	// WithEqualsComparator the equalsComparator will be embedded
	// into this IEnumerable which automatically serves for methods like...
	// TODO
	WithEqualsComparator(equalsComparable func(d1, d2 T) bool) IEnumerable[T]

	// WithLessComparator the lessComparator will be embedded
	// into this IEnumerable which automatically serves for methods like...
	// TODO
	WithLessComparator(less func(d1, d2 T) bool) IEnumerable[T]

	// internal APIs

	exposeData() []T
}

package comparers

// IComparer use method Compare to compare value of 2 input values.
//
// If left is less than right, returns -1.
//
// If left is equals to right, returns 0.
//
// If left is greater than right, returns 1.
type IComparer[T any] interface {
	// Compare compares value from params.
	//
	// If x is less than y, returns -1.
	//
	// If x is equals to y, returns 0.
	//
	// If x is greater than y, returns 1.
	Compare(x, y T) int

	// ComparePointerMode compares two params but in pointer mode (like *int vs *int).
	//
	// If both x and y are nil, return 0.
	//
	// If x is nil and y is not nil, return -1.
	//
	// If x is not nil and y is nil, return 1.
	//
	// If x is less than y, returns -1.
	//
	// If x is equals to y, returns 0.
	//
	// If x is greater than y, returns 1.
	ComparePointerMode(x, y *T) int
}

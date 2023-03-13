package comparers

// IComparer use method Compare to compare value of 2 input values.
//
// If left is less than right, returns -1.
//
// If left is equals to right, returns 0.
//
// If left is greater than right, returns 1.
type IComparer[T any] interface {
	Compare(x, y T) int
}

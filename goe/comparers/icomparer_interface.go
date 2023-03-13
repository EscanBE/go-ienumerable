package comparers

type IComparer[T any] interface {
	Compare(x, y T) int
}

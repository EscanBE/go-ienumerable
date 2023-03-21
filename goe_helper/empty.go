package goe_helper

import "github.com/EscanBE/go-ienumerable/goe"

// Empty returns an empty IEnumerable[T] that has the specified type argument.
func Empty[T any]() goe.IEnumerable[T] {
	return goe.NewIEnumerable[T]()
}

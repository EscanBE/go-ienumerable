package goe

import "fmt"

var _ ChunkHolder[any] = &chunkHolderImpl[any]{}

type ChunkHolder[T any] interface {
}

type chunkHolderImpl[T any] struct {
	data     [][]T
	dataType string
}

// NewChunkHolder returns an ChunkHolder with the same type as data elements
func NewChunkHolder[T any](data ...[]T) ChunkHolder[T] {
	dataType := fmt.Sprintf("%T", *new(T))
	if dataType == "<nil>" {
		dataType = ""
	}
	return &chunkHolderImpl[T]{
		data:     data,
		dataType: dataType,
	}
}

// GetChunkedIEnumeratorFromHolder moves the inner result into IEnumerable[[]T].
//
// Will panic if type of provided T not exactly matches with stored result
func GetChunkedIEnumeratorFromHolder[T any](ch ChunkHolder[T]) IEnumerable[[]T] {
	inputType := fmt.Sprintf("%T", *new(T))
	c := c[T](ch)
	resultType := c.dataType
	if inputType != resultType {
		if inputType == "<nil>" {
			inputType = "interface {}"
		}
		if resultType == "" {
			resultType = "interface {}"
		}
		panic(fmt.Sprintf("chunked has data with type of []%s, can not be casted into []%s", c.dataType, inputType))
	}
	return NewIEnumerable[[]T](c.data...)
}

// cast ChunkHolder back to *chunkHolderImpl for accessing private fields.
func c[T any](ch ChunkHolder[T]) *chunkHolderImpl[T] {
	return ch.(*chunkHolderImpl[T])
}

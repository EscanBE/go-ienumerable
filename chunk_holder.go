package go_ienumerable

import "fmt"

var _ ChunkHolder[any] = &chunkHolderImpl[any]{}

type ChunkHolder[T any] interface {
	exposeData() [][]T
	exposeDateType() string
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
func GetChunkedIEnumeratorFromHolder[T any](c ChunkHolder[T]) IEnumerable[[]T] {
	inputType := fmt.Sprintf("%T", *new(T))
	resultType := c.exposeDateType()
	if inputType != resultType {
		if inputType == "<nil>" {
			inputType = "interface {}"
		}
		if resultType == "" {
			resultType = "interface {}"
		}
		panic(fmt.Sprintf("chunked has data with type of []%s, can not be casted into []%s", c.exposeDateType(), inputType))
	}
	return NewIEnumerable[[]T](c.exposeData()...)
}

func (h *chunkHolderImpl[T]) exposeData() [][]T {
	return h.data
}

func (h *chunkHolderImpl[T]) exposeDateType() string {
	return h.dataType
}

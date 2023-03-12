package helper

import "github.com/EscanBE/go-ienumerable/goe"

// Chunk supposed to split the elements of source sequence into chunks of size at most size.
func Chunk[T any](source goe.IEnumerable[T], size int) goe.IEnumerable[[]T] {
	return goe.GetChunkedIEnumeratorFromHolder[T](source.ChunkToHolder(size))
}

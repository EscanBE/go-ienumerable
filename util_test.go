package go_ienumerable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func createNilEnumerable() IEnumerable[any] {
	return &enumerable[any]{
		data: nil,
	}
}

func createEmptyEnumerable() IEnumerable[any] {
	return &enumerable[any]{
		data: make([]any, 0),
	}
}

func createIntEnumerable(from, to int) IEnumerable[int] {
	if from > to {
		panic(fmt.Errorf("createIntEnumerable from %d > to %d", from, to))
	}
	data := make([]int, 0)
	for i := from; i <= to; i++ {
		data = append(data, i)
	}
	return NewIEnumerable[int](data...)
}

func createRandomIntEnumerable(size int) IEnumerable[int] {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return NewIEnumerable[int](data...)
}

type copiedOriginal[T comparable] struct {
	data                []T
	hasEqualsComparator bool
	hasLessComparator   bool
}

func backupForAssetUnchanged[T comparable](e IEnumerable[T]) copiedOriginal[T] {
	cast := e.(*enumerable[T])
	return copiedOriginal[T]{
		data:                copySlice(cast.data),
		hasEqualsComparator: cast.equalsComparator != nil,
		hasLessComparator:   cast.lessComparator != nil,
	}
}

func (c copiedOriginal[T]) assertUnchanged(t *testing.T, e IEnumerable[T]) {
	cast := e.(*enumerable[T])
	if len(c.data) != len(cast.data) {
		assert.Lenf(t, cast.data, len(c.data), "data of source IEnumerable has been changed, expect len %d but changed to %d", len(c.data), len(cast.data))
	} else if len(c.data) > 0 {
		for i1, d1 := range c.data {
			d2 := cast.data[i1]
			assert.Equalf(t, d1, d2, "data of source IEnumerable has been changed, expect element at [%d] = %d but changed to %d", i1, d1, d2)
		}
	}

	if c.hasLessComparator != (cast.lessComparator != nil) {

	}

	exists := func(b bool) string {
		if b {
			return "exists"
		} else {
			return "not-exist"
		}
	}

	assert.Equalf(t, c.hasEqualsComparator, cast.equalsComparator != nil, "equals comparator state has changed, expect %s, but got %s", exists(c.hasEqualsComparator), exists(cast.equalsComparator != nil))
	assert.Equalf(t, c.hasLessComparator, cast.lessComparator != nil, "less comparator state has changed, expect %s, but got %s", exists(c.hasLessComparator), exists(cast.lessComparator != nil))
}

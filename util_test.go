package go_ienumerable

import (
	"fmt"
	"math/rand"
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
		data = append(data, rand.Int())
	}
	return NewIEnumerable[int](data...)
}

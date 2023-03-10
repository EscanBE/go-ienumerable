package go_ienumerable

func createNilEnumerable() *enumerable[any] {
	return &enumerable[any]{
		data: nil,
	}
}

func createEmptyEnumerable() *enumerable[any] {
	return &enumerable[any]{
		data: make([]any, 0),
	}
}

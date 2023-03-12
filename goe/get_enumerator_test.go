package goe

import (
	"testing"
)

func Test_enumerable_GetEnumerator(t *testing.T) {
	t.Run("enumerator", func(t *testing.T) {
		testEnumerator[int](t, createIntEnumerable(2, 6).GetEnumerator())
	})

	t.Run("panic when nil enumerator", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)
		x := new(enumerator[int])
		x = nil
		x.MoveNext()
	})
}

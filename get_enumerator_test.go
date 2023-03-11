package go_ienumerable

import (
	"testing"
)

func Test_enumerable_GetEnumerator(t *testing.T) {
	testEnumerator[int](t, createIntEnumerable(2, 6).GetEnumerator())
}

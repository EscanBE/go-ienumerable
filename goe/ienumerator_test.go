package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IEnumerator(t *testing.T) {
	assert.Empty(t, Empty[string]().ToArray())

	testEnumerator[int](t, NewIEnumerator(2, 3, 4, 5, 6))
}

func testEnumerator[T any](t *testing.T, enumerator IEnumerator[T]) {
	assert.True(t, enumerator.MoveNext())
	assert.Equal(t, 2, enumerator.Current())
	assert.True(t, enumerator.MoveNext())
	assert.Equal(t, 3, enumerator.Current())
	assert.True(t, enumerator.MoveNext())
	assert.Equal(t, 4, enumerator.Current())
	assert.True(t, enumerator.MoveNext())
	assert.Equal(t, 5, enumerator.Current())
	assert.True(t, enumerator.MoveNext())
	assert.Equal(t, 6, enumerator.Current())
	assert.False(t, enumerator.MoveNext())

	enumerator.Reset()
	_, err := enumerator.CurrentSafe()
	assert.NotNil(t, err)

	defer deferWantPanicDepends(t, true)
	enumerator.Current()
}

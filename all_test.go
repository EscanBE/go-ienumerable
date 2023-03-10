package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_All(t *testing.T) {
	assert.True(t, createNilEnumerable().All(func(t any) bool {
		return true
	}))
	assert.True(t, createEmptyEnumerable().All(func(t any) bool {
		return true
	}))
	i1 := createIntEnumerable(1, 3)
	assert.True(t, i1.All(func(t int) bool {
		return t < 4
	}))
	assert.False(t, i1.All(func(t int) bool {
		return t >= 2
	}))
	i2 := createRandomIntEnumerable(10)
	assert.True(t, i2.All(func(t int) bool {
		return t >= 0
	}))
	assert.False(t, i2.All(func(t int) bool {
		return t < 0
	}))
}

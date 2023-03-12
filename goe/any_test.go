package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Any(t *testing.T) {
	assert.False(t, createNilEnumerable().Any())
	assert.False(t, createEmptyEnumerable().Any())

	i1 := createIntEnumerable(1, 3)
	bi1 := backupForAssetUnchanged(i1)
	assert.True(t, i1.Any())
	bi1.assertUnchanged(t, i1)
}

func Test_enumerable_AnyBy(t *testing.T) {
	assert.False(t, createNilEnumerable().AnyBy(func(t any) bool {
		return true
	}))
	assert.False(t, createEmptyEnumerable().AnyBy(func(t any) bool {
		return true
	}))

	i1 := createIntEnumerable(1, 3)
	bi1 := backupForAssetUnchanged(i1)
	assert.True(t, i1.AnyBy(func(t int) bool {
		return t < 2
	}))
	assert.False(t, i1.AnyBy(func(t int) bool {
		return t >= 4
	}))
	bi1.assertUnchanged(t, i1)

	i2 := createRandomIntEnumerable(10)
	assert.True(t, i2.AnyBy(func(t int) bool {
		return t >= 0
	}))
	assert.False(t, i2.AnyBy(func(t int) bool {
		return t < 0
	}))
}

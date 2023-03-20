package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_WithComparersFrom(t *testing.T) {
	t.Run("default comparer", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 3)
		ieDes := NewIEnumerable[int](5)

		eS := e[int](ieSrc)
		eD := e[int](ieDes)
		eD.defaultComparer = nil

		assert.NotNil(t, eS.defaultComparer)
		assert.Nil(t, eD.defaultComparer)

		assert.True(t, ieSrc.Contains(2, nil))

		//

		_ = ieSrc.WithComparerFrom(ieDes)

		assert.NotNil(t, eS.defaultComparer)
		assert.Nil(t, eD.defaultComparer)

		assert.True(t, ieSrc.Contains(2, nil)) // not changed

		//

		_ = ieDes.WithComparerFrom(ieSrc)

		assert.NotNil(t, eS.defaultComparer)
		assert.NotNil(t, eD.defaultComparer)

		assert.True(t, ieSrc.Contains(2, nil))
		assert.True(t, ieDes.Contains(5, nil))
	})
}

func Test_enumerable_WithDefaultComparer(t *testing.T) {
	t.Run("inject and remove default comparer", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(5)
		eSrc.WithDefaultComparer(nil)

		e := e[int](eSrc)
		assert.Nil(t, e.defaultComparer)

		// replace

		// eSrc.WithDefaultComparer(comparers.NumericComparer)
		panic("re-implement")

		assert.NotNil(t, e.defaultComparer)

		// eraser if input nil
		eSrc.WithDefaultComparer(nil)
		assert.Nil(t, e.defaultComparer)
	})
}

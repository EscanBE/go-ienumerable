package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
		eSrc := NewIEnumerable[time.Time]()
		eSrc.WithDefaultComparer(nil)

		e := e[time.Time](eSrc)
		assert.Nil(t, e.defaultComparer)

		// replace

		eSrc.WithDefaultComparer(comparers.TimeComparer)

		assert.NotNil(t, e.defaultComparer)

		// eraser if input nil
		eSrc.WithDefaultComparer(nil)
		assert.Nil(t, e.defaultComparer)
	})
}

func Test_enumerable_WithDefaultComparerAny(t *testing.T) {
	t.Run("inject and remove default comparer", func(t *testing.T) {
		eSrc := NewIEnumerable[time.Time]()
		eSrc.WithDefaultComparerAny(nil)

		e := e[time.Time](eSrc)
		assert.Nil(t, e.defaultComparer)

		// replace

		eSrc.WithDefaultComparerAny(comparers.ConvertFromComparerIntoDefaultComparer(comparers.TimeComparer))

		assert.NotNil(t, e.defaultComparer)

		// eraser if input nil
		eSrc.WithDefaultComparerAny(nil)
		assert.Nil(t, e.defaultComparer)
	})

	t.Run("Comparer any for exact type within IEnumerable[any]", func(t *testing.T) {
		now := time.Now()
		eSrc := NewIEnumerable[any](now.Add(time.Minute), now.Add(time.Hour), now, now.Add(time.Second))
		eSrc.WithDefaultComparerAny(nil)

		e := e[any](eSrc)
		assert.Nil(t, e.defaultComparer)

		// replace

		eSrc.WithDefaultComparerAny(comparers.ConvertFromComparerIntoDefaultComparer(comparers.TimeComparer))

		assert.NotNil(t, e.defaultComparer)

		min := eSrc.Min()
		assert.Equal(t, now, min)
		max := eSrc.Max()
		assert.Equal(t, now.Add(time.Hour), max)
	})
}

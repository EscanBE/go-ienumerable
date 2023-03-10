package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_With(t *testing.T) {
	eSrc := NewIEnumerable[int](1, 2, 3)

	back := eSrc.(*enumerable[int])
	assert.Nil(t, back.equalsComparator)
	assert.Nil(t, back.lessComparator)

	test_enumerable_With_addWiths(eSrc)

	assert.NotNil(t, back.equalsComparator)
	assert.NotNil(t, back.lessComparator)
}

//goland:noinspection GoSnakeCaseUsage
func test_enumerable_With_addWiths(e IEnumerable[int]) {
	e.WithEqualsComparator(func(i1, i2 int) bool {
		return i1 == i2
	}).WithLessComparator(func(i1, i2 int) bool {
		return i1 < i2
	})
}

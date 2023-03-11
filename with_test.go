package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_With(t *testing.T) {
	eSrc := NewIEnumerable[int](1, 2, 3)

	back := eSrc.(*enumerable[int])
	assert.Nil(t, back.equalityComparer)
	assert.Nil(t, back.lessComparer)

	test_enumerable_With_addWiths(eSrc)

	assert.NotNil(t, back.equalityComparer)
	assert.NotNil(t, back.lessComparer)
}

//goland:noinspection GoSnakeCaseUsage
func test_enumerable_With_addWiths(e IEnumerable[int]) {
	e.WithEqualsComparer(func(i1, i2 int) bool {
		return i1 == i2
	}).WithLessComparer(func(i1, i2 int) bool {
		return i1 < i2
	})
}

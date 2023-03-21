package comparers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_defaultComparer(t *testing.T) {
	test_defaultComparer[any](t, NumericComparer, 1, 1, 0, false)
	test_defaultComparer[any](t, NumericComparer, 2, 1, 1, false)
	test_defaultComparer[string](t, StringComparer, "2", "1", 1, false)
	test_defaultComparer[string](t, StringComparer, "2", 1, nil, true)
	test_defaultComparer[string](t, StringComparer, 2, "1", nil, true)
}

//goland:noinspection GoSnakeCaseUsage
func test_defaultComparer[T any](t *testing.T, comparer IComparer[T], x, y any, expect any, wantPanic bool) {
	defaultComparer := ConvertFromComparerIntoDefaultComparer[T](comparer)

	defer deferExpectPanicContains(t, "to typed value", wantPanic)

	assert.Equal(t, expect, defaultComparer.CompareTyped(x, y))
	assert.Equal(t, expect, defaultComparer.CompareAny(x, y))
}

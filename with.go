package go_ienumerable

import (
	"fmt"
	"strings"
)

func (src *enumerable[T]) WithEqualsComparer(equalsComparable func(d1 T, d2 T) bool) IEnumerable[T] {
	src.equalityComparer = equalsComparable
	return src
}

func (src *enumerable[T]) WithLessComparer(less func(d1 T, d2 T) bool) IEnumerable[T] {
	src.lessComparer = less
	return src
}

type requireWithExtraFunc byte

const (
	requireEqualityComparer requireWithExtraFunc = 0b01
	requireLessComparer     requireWithExtraFunc = 0b10
)

func panicRequire(require requireWithExtraFunc) {
	requiresName := getRequireName(require)
	panic(fmt.Errorf("the following comparer must be set: [%s]", strings.Join(requiresName, ",")))
}

func panicRequireEither(require requireWithExtraFunc) {
	requiresName := getRequireName(require)
	panic(fmt.Errorf("either of the following comparers must be set: [%s]", strings.Join(requiresName, ",")))
}

func getRequireName(require requireWithExtraFunc) []string {
	result := make([]string, 0)

	if require&requireEqualityComparer == requireEqualityComparer {
		result = append(result, "Equals Comparer (can be set using WithEqualsComparer)")
	}

	if require&requireLessComparer == requireLessComparer {
		result = append(result, "Less Comparer (can be set using WithLessComparer)")
	}

	return result
}

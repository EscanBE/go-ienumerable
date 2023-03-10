package go_ienumerable

import (
	"fmt"
	"strings"
)

func (src *enumerable[T]) WithEqualsComparator(equalsComparable func(d1 T, d2 T) bool) IEnumerable[T] {
	src.equalsComparator = equalsComparable
	return src
}

func (src *enumerable[T]) WithLessComparator(less func(d1 T, d2 T) bool) IEnumerable[T] {
	src.lessComparator = less
	return src
}

type requireWithExtraFunc byte

const (
	requireEqualsComparator requireWithExtraFunc = 0b01
	requireLessComparator   requireWithExtraFunc = 0b10
)

func panicRequireEither(require requireWithExtraFunc) {
	requiresName := getRequireName(require)
	panic(fmt.Errorf("either of the following comparators must be set: [%s]", strings.Join(requiresName, ",")))
}

func getRequireName(require requireWithExtraFunc) []string {
	result := make([]string, 0)

	if require&requireEqualsComparator == requireEqualsComparator {
		result = append(result, "Equals Comparator (can be set using WithEqualsComparator)")
	}

	if require&requireLessComparator == requireLessComparator {
		result = append(result, "Less Comparator (can be set using WithLessComparator)")
	}

	return result
}

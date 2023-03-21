package goe_helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type Student struct {
		Class string
		Name  string
	}

	ieSrc := goe.NewIEnumerable[Student](
		Student{
			Class: "A",
			Name:  "John",
		},
		Student{
			Class: "A",
			Name:  "Camila",
		},
		Student{
			Class: "A",
			Name:  "Stephen",
		},
		Student{
			Class: "B",
			Name:  "Hope",
		},
		Student{
			Class: "B",
			Name:  "Paul",
		},
	)

	t.Run("normal", func(t *testing.T) {
		var compareClassNameFunc goe.OptionalEqualsFunc[string] = func(key1, key2 string) bool {
			return key1 == key2
		}

		groups := GroupBy(ieSrc, func(src Student) string {
			return src.Class
		}, func(src Student) string {
			return src.Name
		}, compareClassNameFunc)

		assert.Equal(t, 2, groups.Count(nil))

		group1 := groups.ElementAt(0, false)
		group2 := groups.ElementAt(1, false)

		assert.Equal(t, "A", group1.Key)
		assert.Equal(t, 3, group1.Elements.Count(nil))
		assert.Equal(t, "B", group2.Key)
		assert.Equal(t, 2, group2.Elements.Count(nil))
	})

	keySelector := func(src Student) string {
		return src.Class
	}

	elementSelector := func(src Student) string {
		return src.Name
	}

	var keyEqualityFunc goe.OptionalEqualsFunc[string] = func(key1, key2 string) bool {
		return key1 == key2
	}

	t.Run("automatically resolve default comparer", func(t *testing.T) {
		groups := GroupBy(ieSrc, keySelector, elementSelector, nil)

		assert.Equal(t, 2, groups.Count(nil))

		group1 := groups.ElementAt(0, false)
		group2 := groups.ElementAt(1, false)

		assert.Equal(t, "A", group1.Key)
		assert.Equal(t, 3, group1.Elements.Count(nil))
		assert.Equal(t, "B", group2.Key)
		assert.Equal(t, 2, group2.Elements.Count(nil))
	})

	t.Run("panic source collection nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source collection is nil", true)

		_ = GroupBy[Student, string, string](nil, keySelector, elementSelector, keyEqualityFunc)
	})

	t.Run("panic key selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "key selector is nil", true)

		_ = GroupBy[Student, string, string](ieSrc, nil, elementSelector, keyEqualityFunc)
	})

	t.Run("panic element selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "element selector is nil", true)

		_ = GroupBy[Student, string, string](ieSrc, keySelector, nil, keyEqualityFunc)
	})

	t.Run("panic no default comparer", func(t *testing.T) {
		defer deferExpectPanicContains(t, "no default comparer registered for key type", true)

		_ = GroupBy[Student, Student, string](ieSrc, func(src Student) Student {
			return src
		}, elementSelector, nil)
	})
}

func TestGroupByTransform(t *testing.T) {
	type Student struct {
		Class string
		Name  string
	}

	ieSrc := goe.NewIEnumerable[Student](
		Student{
			Class: "A",
			Name:  "John",
		},
		Student{
			Class: "A",
			Name:  "Camila",
		},
		Student{
			Class: "A",
			Name:  "Stephen",
		},
		Student{
			Class: "B",
			Name:  "Hope",
		},
		Student{
			Class: "B",
			Name:  "Paul",
		},
	)

	t.Run("normal", func(t *testing.T) {
		var compareClassNameFunc goe.OptionalEqualsFunc[string] = func(key1, key2 string) bool {
			return key1 == key2
		}

		ieGot := GroupByTransform(
			ieSrc, func(src Student) string {
				return src.Class
			}, func(src Student) string {
				return src.Name
			}, func(key string, group goe.IEnumerable[string]) string {
				return fmt.Sprintf("%s: %s", key, strings.Join(group.ToArray(), "+"))
			}, compareClassNameFunc,
		)

		assert.Equal(t, 2, ieGot.Count(nil))

		assert.Equal(t, "A: John+Camila+Stephen", ieGot.ElementAt(0, false))
		assert.Equal(t, "B: Hope+Paul", ieGot.ElementAt(1, false))
	})

	keySelector := func(src Student) string {
		return src.Class
	}

	elementSelector := func(src Student) string {
		return src.Name
	}

	resultTransform := func(key string, group goe.IEnumerable[string]) string {
		return fmt.Sprintf("%s: %s", key, strings.Join(group.ToArray(), "+"))
	}

	var keyEqualityFunc goe.OptionalEqualsFunc[string] = func(key1, key2 string) bool {
		return key1 == key2
	}

	t.Run("automatically resolve default comparer", func(t *testing.T) {
		ieGot := GroupByTransform(ieSrc, keySelector, elementSelector, resultTransform, nil)

		assert.Equal(t, 2, ieGot.Count(nil))

		assert.Equal(t, "A: John+Camila+Stephen", ieGot.ElementAt(0, false))
		assert.Equal(t, "B: Hope+Paul", ieGot.ElementAt(1, false))
	})

	t.Run("panic source collection nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "source collection is nil", true)

		_ = GroupByTransform[Student, string, string, string](nil, keySelector, elementSelector, resultTransform, keyEqualityFunc)
	})

	t.Run("panic key selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "key selector is nil", true)

		_ = GroupByTransform[Student, string, string, string](ieSrc, nil, elementSelector, resultTransform, keyEqualityFunc)
	})

	t.Run("panic element selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "element selector is nil", true)

		_ = GroupByTransform[Student, string, string, string](ieSrc, keySelector, nil, resultTransform, keyEqualityFunc)
	})

	t.Run("panic result selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "result selector is nil", true)

		_ = GroupByTransform[Student, string, string, string](ieSrc, keySelector, elementSelector, nil, keyEqualityFunc)
	})

	t.Run("panic no default comparer", func(t *testing.T) {
		defer deferExpectPanicContains(t, "no default comparer registered for key type", true)

		_ = GroupByTransform[Student, Student, string, string](ieSrc, func(src Student) Student {
			return src
		}, elementSelector, func(key Student, group goe.IEnumerable[string]) string {
			return ""
		}, nil)
	})
}

package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
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

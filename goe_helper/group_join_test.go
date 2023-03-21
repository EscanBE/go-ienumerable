package goe_helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGroupJoin(t *testing.T) {
	type Person struct {
		Name string
	}
	type Pet struct {
		Name  string
		Owner Person
	}
	type CollectionInfo struct {
		OwnerName string
		Pets      []string
	}

	magnus := Person{
		Name: "Hedlund, Magnus",
	}
	terry := Person{
		Name: "Adams, Terry",
	}
	charlotte := Person{
		Name: "Weiss, Charlotte",
	}

	barley := Pet{
		Name:  "Barley",
		Owner: terry,
	}

	boots := Pet{
		Name:  "Boots",
		Owner: terry,
	}

	whiskers := Pet{
		Name:  "Whiskers",
		Owner: charlotte,
	}

	daisy := Pet{
		Name:  "Daisy",
		Owner: magnus,
	}

	iePeople := goe.NewIEnumerable(magnus, terry, charlotte)
	iePets := goe.NewIEnumerable(barley, boots, whiskers, daisy)

	t.Run("C# example", func(t *testing.T) {
		var compareOwnerFunc goe.OptionalEqualsFunc[Person] = func(person1, person2 Person) bool {
			return person1.Name == person2.Name
		}

		ieGot := GroupJoin(iePeople, iePets, func(person Person) Person {
			return person
		}, func(pet Pet) Person {
			return pet.Owner
		}, func(person Person, pets goe.IEnumerable[Pet]) CollectionInfo {
			return CollectionInfo{
				OwnerName: person.Name,
				Pets: pets.Select(func(p Pet) any {
					return p.Name
				}).CastString().ToArray(),
			}
		}, compareOwnerFunc)

		got := ieGot.ToArray()
		assert.Len(t, got, 3)

		getString := func(collection CollectionInfo) string {
			return fmt.Sprintf("%s - %s", collection.OwnerName, strings.Join(collection.Pets, "+"))
		}

		assert.Equal(t, "Hedlund, Magnus - Daisy", getString(got[0]))
		assert.Equal(t, "Adams, Terry - Barley+Boots", getString(got[1]))
		assert.Equal(t, "Weiss, Charlotte - Whiskers", getString(got[2]))
	})

	outerKeySelector := func(person Person) Person {
		return person
	}
	innerKeySelector := func(pet Pet) Person {
		return pet.Owner
	}
	resultKeySelector := func(person Person, pets goe.IEnumerable[Pet]) CollectionInfo {
		return CollectionInfo{
			OwnerName: person.Name,
			Pets: pets.Select(func(p Pet) any {
				return p.Name
			}).CastString().ToArray(),
		}
	}
	var compareOwnerFunc goe.OptionalEqualsFunc[Person] = func(person1, person2 Person) bool {
		return person1.Name == person2.Name
	}

	t.Run("automatically resolve default comparer", func(t *testing.T) {
		ieGot := GroupJoin[Person, Pet, string, CollectionInfo](iePeople, iePets, func(person Person) string {
			return person.Name
		}, func(pet Pet) string {
			return pet.Owner.Name
		}, resultKeySelector, nil)

		got := ieGot.ToArray()
		assert.Len(t, got, 3)

		getString := func(collection CollectionInfo) string {
			return fmt.Sprintf("%s - %s", collection.OwnerName, strings.Join(collection.Pets, "+"))
		}

		assert.Equal(t, "Hedlund, Magnus - Daisy", getString(got[0]))
		assert.Equal(t, "Adams, Terry - Barley+Boots", getString(got[1]))
		assert.Equal(t, "Weiss, Charlotte - Whiskers", getString(got[2]))
	})

	t.Run("panic outer collection nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "outer collection is nil", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](nil, iePets, outerKeySelector, innerKeySelector, resultKeySelector, compareOwnerFunc)
	})

	t.Run("panic inner collection nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "inner collection is nil", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](iePeople, nil, outerKeySelector, innerKeySelector, resultKeySelector, compareOwnerFunc)
	})

	t.Run("panic outer key selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "outer key selector is nil", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](iePeople, iePets, nil, innerKeySelector, resultKeySelector, compareOwnerFunc)
	})

	t.Run("panic inner key selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "inner key selector is nil", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](iePeople, iePets, outerKeySelector, nil, resultKeySelector, compareOwnerFunc)
	})

	t.Run("panic result selector is nil", func(t *testing.T) {
		defer deferExpectPanicContains(t, "result selector is nil", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](iePeople, iePets, outerKeySelector, innerKeySelector, nil, compareOwnerFunc)
	})

	t.Run("panic no default comparer", func(t *testing.T) {
		defer deferExpectPanicContains(t, "no default comparer registered for key type", true)

		_ = GroupJoin[Person, Pet, Person, CollectionInfo](iePeople, iePets, outerKeySelector, innerKeySelector, resultKeySelector, nil)
	})
}

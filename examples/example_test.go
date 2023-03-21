package examples

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/EscanBE/go-ienumerable/goe_helper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_example_1(t *testing.T) {

	// This is example contains Where, OrderByDescending, Reverse and FirstOrDefault

	got := goe.NewIEnumerable[string]("Hello", "World").
		Where(func(v string) bool {
			return len(v) < 3
		}).OrderDescending().
		GetOrderedEnumerable().
		Reverse().
		FirstOrDefault(nil, goe.Ptr("Oops"))
	fmt.Println(got)

	assert.Equal(t, "Oops", got)
}

func Test_example_2(t *testing.T) {

	// This is example contains Skip, Take, Select, Cast, Append and Aggregate

	transform := func(v byte) any {
		return v + 2
	}

	aggregate := func(str any, v int32) any {
		return fmt.Sprintf("%s%c", str, v)
	}

	// H 	e 	l 	l 	o 	_ 	W 	o 	r 	l	d
	// 72	101	108	108	111	32	87	111	114	108	100
	array := []byte{0, 70, 99, 106, 106, 109, 30, 85, 109, 112, 106, 98, 99, 66, 88, 69}
	got := goe.NewIEnumerable[byte](array...).
		Skip(1).
		Take(11).
		Select(transform).
		CastInt32().
		Append('"').
		AggregateAnySeed("\"", aggregate)

	fmt.Println(got)

	assert.Equal(t, "\"Hello World\"", got)
}

func Test_example_3(t *testing.T) {
	eSrc := goe.NewIEnumerable[string]("v2430", "v1530", "v3530", "v4530", "v2420", "v2160", "v3990")

	comparatorLevel1 := func(l, r any) int {
		leftString := l.(string)
		rightString := r.(string)
		return comparers.StringComparer.CompareTyped(string(leftString[1]), string(rightString[1]))
	}

	comparatorLevel2 := func(l, r any) int {
		leftString := l.(string)
		rightString := r.(string)
		return comparers.StringComparer.CompareTyped(string(leftString[2]), string(rightString[2]))
	}

	comparatorLevel3 := func(l, r any) int {
		leftString := l.(string)
		rightString := r.(string)
		return comparers.StringComparer.CompareTyped(string(leftString[3]), string(rightString[3]))
	}

	got := eSrc.OrderByDescending(goe.SelfSelector[string](), comparatorLevel1).
		ThenBy(goe.SelfSelector[string](), comparatorLevel2).
		ThenByDescending(goe.SelfSelector[string](), comparatorLevel3).
		GetOrderedEnumerable()

	fmt.Println(got)
	// v4530 v3530 v3990 v2160 v2430 v2420 v1530
}

func Test_example_4(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		type PetOwner struct {
			Name string
			Pets []string
		}
		eSrc := goe.NewIEnumerable[PetOwner](
			PetOwner{
				Name: "Higa",
				Pets: []string{"Scruffy", "Sam"},
			},
			PetOwner{
				Name: "Ashkenazi",
				Pets: []string{"Walker", "Sugar"},
			},
			PetOwner{
				Name: "Price",
				Pets: []string{"Scratches", "Diesel"},
			},
			PetOwner{
				Name: "Hines",
				Pets: []string{"Dusty"},
			},
		)

		eGot := goe_helper.SelectManyTransform(eSrc, func(petOwner PetOwner) []string {
			return petOwner.Pets
		}, func(petOwner PetOwner, petName string) goe.ValueTuple2[PetOwner, string] {
			return goe.ValueTuple2[PetOwner, string]{
				First:  petOwner,
				Second: petName,
			}
		}).Where(func(ownerAndPet goe.ValueTuple2[PetOwner, string]) bool {
			return strings.HasPrefix(ownerAndPet.Second, "S")
		}).Select(func(ownerAndPet goe.ValueTuple2[PetOwner, string]) any {
			return fmt.Sprintf("{Owner=%s, Pet=%s}", ownerAndPet.First.Name, ownerAndPet.Second)
		})

		gotData := eGot.ToArray()

		assert.Len(t, gotData, 4)
		assert.Equal(t, "{Owner=Higa, Pet=Scruffy}", gotData[0])
		assert.Equal(t, "{Owner=Higa, Pet=Sam}", gotData[1])
		assert.Equal(t, "{Owner=Ashkenazi, Pet=Sugar}", gotData[2])
		assert.Equal(t, "{Owner=Price, Pet=Scratches}", gotData[3])
	})
}

func Test_example_5(t *testing.T) {
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

	var compareClassNameFunc goe.OptionalEqualsFunc[string] = func(key1, key2 string) bool {
		return key1 == key2
	}

	var groups goe.IEnumerable[goe.Group[string, goe.IEnumerable[string]]]

	groups = goe_helper.GroupBy(ieSrc, func(src Student) string {
		return src.Class
	}, func(src Student) string {
		return src.Name
	}, compareClassNameFunc)
	assert.Equal(t, 2, groups.Count(nil))

	group1 := groups.ElementAt(0, false)
	group2 := groups.ElementAt(1, false)

	assert.Equal(t, "A", group1.Key)
	assert.Equal(t, 3, group1.Elements.Count(nil))
	assert.Equal(t, "John", group1.Elements.ElementAt(0, false))
	assert.Equal(t, "Camila", group1.Elements.ElementAt(1, false))
	assert.Equal(t, "Stephen", group1.Elements.ElementAt(2, false))
	assert.Equal(t, "B", group2.Key)
	assert.Equal(t, 2, group2.Elements.Count(nil))
	assert.Equal(t, "Hope", group2.Elements.ElementAt(0, false))
	assert.Equal(t, "Paul", group2.Elements.ElementAt(1, false))
}

func Test_example_6(t *testing.T) {
	type Person struct {
		Name string
	}
	type Pet struct {
		Name  string
		Owner Person
	}
	type PetInfo struct {
		OwnerName string
		Pet       string
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

	var compareOwnerFunc goe.OptionalEqualsFunc[Person] = func(person1, person2 Person) bool {
		return person1.Name == person2.Name
	}

	var ieGot goe.IEnumerable[PetInfo]

	ieGot = goe_helper.Join(iePeople, iePets, func(person Person) Person {
		return person
	}, func(pet Pet) Person {
		return pet.Owner
	}, func(person Person, pet Pet) PetInfo {
		return PetInfo{
			OwnerName: person.Name,
			Pet:       pet.Name,
		}
	}, compareOwnerFunc)
	got := ieGot.ToArray()
	assert.Len(t, got, 4)

	getString := func(pet PetInfo) string {
		return fmt.Sprintf("%s - %s", pet.OwnerName, pet.Pet)
	}

	assert.Equal(t, "Hedlund, Magnus - Daisy", getString(got[0]))
	assert.Equal(t, "Adams, Terry - Barley", getString(got[1]))
	assert.Equal(t, "Adams, Terry - Boots", getString(got[2]))
	assert.Equal(t, "Weiss, Charlotte - Whiskers", getString(got[3]))
}

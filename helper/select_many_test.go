package helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestSelectMany(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]([]int8{2, 3}, []int8{6, 7}, []int8{4, 5})

		eGot := SelectMany(eSrc, func(a []int8) []int8 {
			return []int8{a[0] * 2, a[1] * 2}
		})
		gotData := eGot.ToArray()
		assert.Len(t, gotData, 6)
		assert.Equal(t, int8(4), gotData[0])
		assert.Equal(t, int8(6), gotData[1])
		assert.Equal(t, int8(12), gotData[2])
		assert.Equal(t, int8(14), gotData[3])
		assert.Equal(t, int8(8), gotData[4])
		assert.Equal(t, int8(10), gotData[5])
	})

	t.Run("empty", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]()

		eGot := SelectMany(eSrc, func(i []int8) []int64 {
			return []int64{int64(i[0])}
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 0)
	})

	t.Run("skip empty result from selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[[]int8]([]int8{1, 2}, []int8{0, 2}, []int8{1, 0}, []int8{0, 0})

		eGot := SelectMany(eSrc, func(i []int8) []int8 {
			result := make([]int8, 0)
			for _, iv := range i {
				if iv != 0 {
					result = append(result, iv)
				}
			}
			return result
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 4)
	})

	t.Run("nil selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[int8]()

		defer deferExpectPanicContains(t, "result selector function is nil", true)

		_ = SelectMany[int8, int](eSrc, nil)
	})

	t.Run("automatically inject type and comparer", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[[]int]([]int{3, 1})

		ieGot := SelectMany(ieSrc, func(i []int) []time.Duration {
			return []time.Duration{time.Duration(i[0]) * time.Minute, time.Duration(i[1]) * time.Minute}
		})

		gotArray := ieGot.ToArray()

		assert.Equal(t, 3*time.Minute, gotArray[0])
		assert.Equal(t, 1*time.Minute, gotArray[1])

		gotArray = ieGot.Order().GetOrderedEnumerable().ToArray()

		assert.Equal(t, 1*time.Minute, gotArray[0])
		assert.Equal(t, 3*time.Minute, gotArray[1])
	})

	t.Run("panic nil value as result of selector", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[[]int]([]int{9, 3})

		defer deferExpectPanicContains(t, "result array can not be nil", true)

		_ = SelectMany(ieSrc, func(i []int) []int {
			return nil
		})
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := goe.NewIEnumerable[[]int]([]int{3, 1, 2, 6}, []int{})

		gotArray := SelectMany[[]int, *MyInt](ieSrc, func(i []int) []*MyInt {
			result := make([]*MyInt, len(i))
			for idx, iv := range i {
				if iv == 2 {
					result[idx] = nil
				} else {
					result[idx] = &MyInt{
						Value: iv,
					}
				}
			}
			return result
		}).ToArray()

		assert.Len(t, gotArray, 4)
		assert.Equal(t, 3, gotArray[0].Value)
		assert.Equal(t, 1, gotArray[1].Value)
		assert.Nil(t, gotArray[2])
		assert.Equal(t, 6, gotArray[3].Value)
	})
}

func TestSelectManyTransform(t *testing.T) {
	type PetOwner struct {
		Name string
		Pets []string
	}

	t.Run("sample", func(t *testing.T) {
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

		eGot := SelectManyTransform(eSrc, func(petOwner PetOwner) []string {
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

	t.Run("empty", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[PetOwner]()

		eGot := SelectManyTransform(eSrc, func(petOwner PetOwner) []string {
			return petOwner.Pets
		}, func(petOwner PetOwner, petName string) goe.ValueTuple2[PetOwner, string] {
			return goe.ValueTuple2[PetOwner, string]{
				First:  petOwner,
				Second: petName,
			}
		})

		gotData := eGot.ToArray()

		assert.Len(t, gotData, 0)
	})

	t.Run("skip empty result from selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[PetOwner](
			PetOwner{
				Name: "Higa",
				Pets: []string{},
			},
			PetOwner{
				Name: "Ashkenazi",
				Pets: []string{"Walker", "Sugar"},
			},
		)

		eGot := SelectManyTransform(eSrc, func(petOwner PetOwner) []string {
			return petOwner.Pets
		}, func(petOwner PetOwner, petName string) goe.ValueTuple2[PetOwner, string] {
			return goe.ValueTuple2[PetOwner, string]{
				First:  petOwner,
				Second: petName,
			}
		})

		gotData := eGot.ToArray()
		assert.Len(t, gotData, 2)
		assert.Equal(t, "Ashkenazi", gotData[0].First.Name)
		assert.Equal(t, "Walker", gotData[0].Second)
		assert.Equal(t, "Ashkenazi", gotData[1].First.Name)
		assert.Equal(t, "Sugar", gotData[1].Second)
	})

	t.Run("nil collection selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[PetOwner]()

		defer deferExpectPanicContains(t, "collection selector function is nil", true)

		_ = SelectManyTransform[PetOwner, string, goe.ValueTuple2[PetOwner, string]](eSrc, nil, func(petOwner PetOwner, petName string) goe.ValueTuple2[PetOwner, string] {
			return goe.ValueTuple2[PetOwner, string]{
				First:  petOwner,
				Second: petName,
			}
		})
	})

	t.Run("nil result selector", func(t *testing.T) {
		eSrc := goe.NewIEnumerable[PetOwner]()

		defer deferExpectPanicContains(t, "result selector function is nil", true)

		_ = SelectManyTransform[PetOwner, string, goe.ValueTuple2[PetOwner, string]](eSrc, func(petOwner PetOwner) []string {
			return petOwner.Pets
		}, nil)
	})

	t.Run("panic nil value as result of selector", func(t *testing.T) {
		ieSrc := goe.NewIEnumerable[int](1, 2, 3, 4)

		defer deferExpectPanicContains(t, "result array can not be nil", true)

		_ = SelectManyTransform[int, *string, *string](ieSrc, func(i int) []*string {
			return nil
		}, func(i int, my *string) *string {
			return my
		}).ToArray()
	})

	t.Run("not panic if not able to detect comparer", func(t *testing.T) {
		type MyInt struct {
			Value int
		}

		ieSrc := goe.NewIEnumerable[int](1, 2, 3, 4)

		gotArray := SelectManyTransform[int, *MyInt, *MyInt](ieSrc, func(i int) []*MyInt {
			return []*MyInt{
				&MyInt{
					Value: i,
				},
				&MyInt{
					Value: i * 2,
				},
			}
		}, func(i int, my *MyInt) *MyInt {
			my.Value += i
			return my
		}).ToArray()

		assert.Len(t, gotArray, 8)
		assert.Equal(t, 2, gotArray[0].Value)
		assert.Equal(t, 3, gotArray[1].Value)
		assert.Equal(t, 4, gotArray[2].Value)
		assert.Equal(t, 6, gotArray[3].Value)
	})
}

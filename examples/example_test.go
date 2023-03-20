package examples

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_example_1(t *testing.T) {

	// This is example contains Where, OrderByDescending, Reverse and FirstOrDefault

	got := goe.NewIEnumerable[string]("Hello", "World").
		Where(func(v string) bool {
			return len(v) < 3
		}).OrderByDescending().
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
		Skip(1).Take(11).Select(transform).CastInt32().Append('"').
		AggregateWithAnySeed("\"", aggregate)
	fmt.Println(got)

	assert.Equal(t, "\"Hello World\"", got)
}

//func Test_example_3(t *testing.T) {
//	eSrc := goe.NewIEnumerable[string]("v2430", "v1530", "v3530", "v4530", "v2420", "v2160", "v3990")
//
//	comparatorLevel1 := func(l, r string) int {
//		return comparers.StringComparer.Compare(string(l[1]), string(r[1]))
//	}
//
//	comparatorLevel2 := func(l, r string) int {
//		return comparers.StringComparer.Compare(string(l[2]), string(r[2]))
//	}
//
//	comparatorLevel3 := func(l, r string) int {
//		return comparers.StringComparer.Compare(string(l[3]), string(r[3]))
//	}
//
//	// TODO
//}

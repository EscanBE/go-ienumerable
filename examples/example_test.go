package examples

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_example_1(t *testing.T) {
	got := goe.NewIEnumerable[string]("Hello", "World").
		Where(func(v string) bool {
			return len(v) < 3
		}).OrderByDescending().Reverse().
		FirstOrDefaultUsing("Oops")
	fmt.Println(got)

	assert.Equal(t, "Oops", got)
}

func Test_example_2(t *testing.T) {
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

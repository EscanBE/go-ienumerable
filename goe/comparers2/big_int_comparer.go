package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"math/big"
	"reflect"
)

// ensure implementation
var _ IComparer[*big.Int] = bigIntComparer{}

type bigIntComparer struct {
}

func NewBigIntComparer() IComparer[*big.Int] {
	return bigIntComparer{}
}

func (n bigIntComparer) CompareTyped(x, y *big.Int) int {
	if x == nil && y == nil {
		return 0
	}

	if x == nil {
		return -1
	}

	if y == nil {
		return 1
	}

	return x.Cmp(y)
}

func (n bigIntComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if voy.Kind() == reflect.Struct {
			if _, oky := voy.Interface().(big.Int); oky {
				return -1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to big.Int", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if vox.Kind() == reflect.Struct {
			if _, okx := vox.Interface().(big.Int); okx {
				return 1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to big.Int", vox.Kind().String(), vox.Type().String()))
	}

	bix := extractBigIntFromReflectValue(*vox)
	biy := extractBigIntFromReflectValue(*voy)

	return n.CompareTyped(bix, biy)
}

func extractBigIntFromReflectValue(voi reflect.Value) *big.Int {
	if voi.Kind() == reflect.Struct {
		if t, ok := voi.Interface().(big.Int); ok {
			return &t
		}
	}

	panic(fmt.Sprintf("%s %s can not be cast to big.Int", voi.Kind().String(), voi.Type().String()))
}

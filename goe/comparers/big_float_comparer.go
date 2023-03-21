package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"math/big"
	"reflect"
)

// ensure implementation
var _ IComparer[*big.Float] = bigFloatComparer{}

type bigFloatComparer struct {
}

func NewBigFloatComparer() IComparer[*big.Float] {
	return bigFloatComparer{}
}

func (n bigFloatComparer) CompareTyped(x, y *big.Float) int {
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

func (n bigFloatComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if voy.Kind() == reflect.Struct {
			if _, oky := voy.Interface().(big.Float); oky {
				return -1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to big.Float", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if vox.Kind() == reflect.Struct {
			if _, okx := vox.Interface().(big.Float); okx {
				return 1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to big.Float", vox.Kind().String(), vox.Type().String()))
	}

	bix := extractBigFloatFromReflectValue(*vox)
	biy := extractBigFloatFromReflectValue(*voy)

	return n.CompareTyped(bix, biy)
}

func extractBigFloatFromReflectValue(voi reflect.Value) *big.Float {
	if voi.Kind() == reflect.Struct {
		if t, ok := voi.Interface().(big.Float); ok {
			return &t
		}
	}

	panic(fmt.Sprintf("%s %s can not be cast to big.Float", voi.Kind().String(), voi.Type().String()))
}

package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"reflect"
)

// ensure implementation
var _ IComparer[bool] = boolComparer{}

type boolComparer struct {
}

func NewBoolComparer() IComparer[bool] {
	return boolComparer{}
}

func (t boolComparer) CompareTyped(x, y bool) int {
	if !x && y {
		return -1
	}

	if x && !y {
		return 1
	}

	return 0
}

func (t boolComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if voy.Kind() == reflect.Bool {
			if _, oky := voy.Interface().(bool); oky {
				return -1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to bool", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if vox.Kind() == reflect.Bool {
			if _, okx := vox.Interface().(bool); okx {
				return 1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to bool", vox.Kind().String(), vox.Type().String()))
	}

	tx := extractBoolFromReflectValue(*vox)
	ty := extractBoolFromReflectValue(*voy)

	return t.CompareTyped(tx, ty)
}

func extractBoolFromReflectValue(voi reflect.Value) bool {
	if voi.Kind() == reflect.Bool {
		if t, ok := voi.Interface().(bool); ok {
			return t
		}
	}

	panic(fmt.Sprintf("%s %s can not be cast to bool", voi.Kind().String(), voi.Type().String()))
}

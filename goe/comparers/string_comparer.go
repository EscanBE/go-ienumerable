package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"reflect"
	"strings"
)

// ensure implementation
var _ IComparer[string] = stringComparer{}

type stringComparer struct {
}

func NewStringComparer() IComparer[string] {
	return stringComparer{}
}

func (t stringComparer) CompareTyped(x, y string) int {
	return strings.Compare(x, y)
}

func (t stringComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if voy.Kind() == reflect.String {
			if _, oky := voy.Interface().(string); oky {
				return -1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to string", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if vox.Kind() == reflect.String {
			if _, okx := vox.Interface().(string); okx {
				return 1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to string", vox.Kind().String(), vox.Type().String()))
	}

	tx := extractStringFromReflectValue(*vox)
	ty := extractStringFromReflectValue(*voy)

	return t.CompareTyped(tx, ty)
}

func extractStringFromReflectValue(voi reflect.Value) string {
	if voi.Kind() == reflect.String {
		if t, ok := voi.Interface().(string); ok {
			return t
		}
	}

	panic(fmt.Sprintf("%s %s can not be cast to string", voi.Kind().String(), voi.Type().String()))
}

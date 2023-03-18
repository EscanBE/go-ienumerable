package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"reflect"
	"time"
)

var _ IComparer[time.Time] = timeTimeComparer{}

type timeTimeComparer struct {
}

func NewTimeComparer() IComparer[time.Time] {
	return timeTimeComparer{}
}

func (t timeTimeComparer) CompareTyped(x, y time.Time) int {
	if x.Before(y) {
		return -1
	}

	if x.After(y) {
		return 1
	}

	return 0
}

func (t timeTimeComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if voy.Kind() == reflect.Struct {
			if _, oky := voy.Interface().(time.Time); oky {
				return -1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to time.Time", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if vox.Kind() == reflect.Struct {
			if _, okx := vox.Interface().(time.Time); okx {
				return 1
			}
		}
		panic(fmt.Sprintf("%s %s can not be cast to time.Time", vox.Kind().String(), vox.Type().String()))
	}

	tx := extractTimeTimeFromReflectValue(*vox)
	ty := extractTimeTimeFromReflectValue(*voy)

	return t.CompareTyped(tx, ty)
}

func extractTimeTimeFromReflectValue(voi reflect.Value) time.Time {
	if voi.Kind() == reflect.Struct {
		if t, ok := voi.Interface().(time.Time); ok {
			return t
		}
	}

	panic(fmt.Sprintf("%s %s can not be cast to time.Time", voi.Kind().String(), voi.Type().String()))
}

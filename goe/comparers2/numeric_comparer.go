package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"math"
	"math/big"
	"reflect"
)

var _ IComparer[any] = numericComparer{}

type numericComparer struct {
}

func NewNumericComparer() IComparer[any] {
	return numericComparer{}
}

func (n numericComparer) CompareTyped(x, y any) int {
	if x == nil || y == nil {
		panic("nil can not be cast to a number")
	}

	ix, bfx, cx := extractNumericValueFromReflectValue(reflect.ValueOf(x))
	iy, bfy, cy := extractNumericValueFromReflectValue(reflect.ValueOf(y))

	assertExactOneValueNotNil(ix, bfx, cx)
	assertExactOneValueNotNil(iy, bfy, cy)

	if ix != nil {
		if bfy != nil {
			return n.internalCompareTyped(new(big.Float).SetInt64(*ix), bfy, _USCTM_BIGFLOAT)
		}

		if cy != nil {
			bfx = new(big.Float).SetInt64(*ix)
			bfy = new(big.Float).SetFloat64(real(*cy))
			cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
			if cmp != 0 {
				return cmp
			}
			return bfx.SetFloat64(imag(*cy)).Sign()
		}

		return n.internalCompareTyped(*ix, *iy, _USCTM_INT64)
	}

	if bfx != nil {
		if iy != nil {
			return n.internalCompareTyped(bfx, new(big.Float).SetInt64(*iy), _USCTM_BIGFLOAT)
		}

		if cy != nil {
			bfy = new(big.Float).SetFloat64(real(*cy))
			cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
			if cmp != 0 {
				return cmp
			}
			return bfy.SetFloat64(imag(*cy)).Sign()
		}

		return n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
	}

	// tx == _USCTM_COMPLEX128
	if iy != nil {
		bfx = new(big.Float).SetFloat64(real(*cx))
		bfy = new(big.Float).SetInt64(*iy)
		cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
		if cmp != 0 {
			return cmp
		}
		return bfx.SetFloat64(imag(*cx)).Sign()
	}

	if bfy != nil {
		bfx = new(big.Float).SetFloat64(real(*cx))
		cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
		if cmp != 0 {
			return cmp
		}
		return bfx.SetFloat64(imag(*cx)).Sign()
	}

	return n.internalCompareTyped(*cx, *cy, _USCTM_COMPLEX128)
}

func (n numericComparer) CompareAny(x, y any) int {
	vox, nilX := reflection.RootValueExtractor(x)
	voy, nilY := reflection.RootValueExtractor(y)

	if nilX && nilY {
		return 0
	}

	if nilX {
		if reflection.IsNumericKind(*voy) {
			return -1
		}
		panic(fmt.Sprintf("%s %s can not be cast to a number", voy.Kind().String(), voy.Type().String()))
	}

	if nilY {
		if reflection.IsNumericKind(*vox) {
			return 1
		}
		panic(fmt.Sprintf("%s %s can not be cast to a number", vox.Kind().String(), vox.Type().String()))
	}

	ix, bfx, cx := extractNumericValueFromReflectValue(*vox)
	iy, bfy, cy := extractNumericValueFromReflectValue(*voy)

	assertExactOneValueNotNil(ix, bfx, cx)
	assertExactOneValueNotNil(iy, bfy, cy)

	if ix != nil {
		if bfy != nil {
			return n.internalCompareTyped(new(big.Float).SetInt64(*ix), bfy, _USCTM_BIGFLOAT)
		}

		if cy != nil {
			bfx = new(big.Float).SetInt64(*ix)
			bfy = new(big.Float).SetFloat64(real(*cy))
			cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
			if cmp != 0 {
				return cmp
			}
			return bfx.SetFloat64(imag(*cy)).Sign()
		}

		return n.internalCompareTyped(*ix, *iy, _USCTM_INT64)
	}

	if bfx != nil {
		if iy != nil {
			return n.internalCompareTyped(bfx, new(big.Float).SetInt64(*iy), _USCTM_BIGFLOAT)
		}

		if cy != nil {
			bfy = new(big.Float).SetFloat64(real(*cy))
			cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
			if cmp != 0 {
				return cmp
			}
			return bfy.SetFloat64(imag(*cy)).Sign()
		}

		return n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
	}

	// tx == _USCTM_COMPLEX128
	if iy != nil {
		bfx = new(big.Float).SetFloat64(real(*cx))
		bfy = new(big.Float).SetInt64(*iy)
		cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
		if cmp != 0 {
			return cmp
		}
		return bfx.SetFloat64(imag(*cx)).Sign()
	}

	if bfy != nil {
		bfx = new(big.Float).SetFloat64(real(*cx))
		cmp := n.internalCompareTyped(bfx, bfy, _USCTM_BIGFLOAT)
		if cmp != 0 {
			return cmp
		}
		return bfx.SetFloat64(imag(*cx)).Sign()
	}

	return n.internalCompareTyped(*cx, *cy, _USCTM_COMPLEX128)
}

type unsafeCompareTypedMode byte

//goland:noinspection GoSnakeCaseUsage,SpellCheckingInspection
const (
	_USCTM_INT64      unsafeCompareTypedMode = 0
	_USCTM_BIGFLOAT   unsafeCompareTypedMode = 1
	_USCTM_COMPLEX128 unsafeCompareTypedMode = 2
)

func (n numericComparer) internalCompareTyped(x, y any, mode unsafeCompareTypedMode) int {
	if mode == _USCTM_BIGFLOAT {
		xf := x.(*big.Float)
		yf := y.(*big.Float)

		return xf.Cmp(yf)
	}

	if mode == _USCTM_COMPLEX128 {
		xc := x.(complex128)
		yc := y.(complex128)

		xf := new(big.Float)
		yf := new(big.Float)

		c := n.internalCompareTyped(xf.SetFloat64(real(xc)), yf.SetFloat64(real(yc)), _USCTM_BIGFLOAT)

		if c != 0 {
			return c
		}

		return n.internalCompareTyped(xf.SetFloat64(imag(xc)), yf.SetFloat64(imag(yc)), _USCTM_BIGFLOAT)
	}

	xi := x.(int64)
	yi := y.(int64)

	if xi < yi {
		return -1
	}

	if xi > yi {
		return 1
	}

	return 0
}

func extractNumericValueFromReflectValue(voi reflect.Value) (iv *int64, bfv *big.Float, cv *complex128) {
	switch voi.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i64 := voi.Int()
		iv = &i64
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ui64 := voi.Uint()
		if ui64 > math.MaxInt64 {
			bfv = new(big.Float).SetUint64(voi.Uint())
		} else {
			i64 := int64(ui64)
			iv = &i64
		}
		return
	case reflect.Float32, reflect.Float64:
		bfv = new(big.Float).SetFloat64(voi.Float())
		return
	case reflect.Complex64, reflect.Complex128:
		cpl := voi.Complex()
		cv = &cpl
		return
	default:
		panic(fmt.Sprintf("%s %s can not be cast to a number", voi.Kind().String(), voi.Type().String()))
	}
}

func assertExactOneValueNotNil(i *int64, bf *big.Float, cx *complex128) {
	cntNonNil := 0
	if i != nil {
		cntNonNil++
	}
	if bf != nil {
		cntNonNil++
	}
	if cx != nil {
		cntNonNil++
	}
	if cntNonNil != 1 {
		panic(fmt.Sprintf("expect exactly one param not nil, found %d/3", cntNonNil))
	}
}

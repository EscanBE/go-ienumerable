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

func (n numericComparer) CompareTyped(x, y any) int {
	if x == nil || y == nil {
		panic("nil can not be cast to a number")
	}

	ix, fx, cx, tx := extractValueDepends(reflect.ValueOf(x))
	iy, fy, cy, ty := extractValueDepends(reflect.ValueOf(y))

	if tx == _USCTM_INT64 {
		if ty == _USCTM_FLOAT64 {
			return n.internalCompareTyped(new(big.Float).SetInt64(ix), fy, _USCTM_FLOAT64)
		}

		if ty == _USCTM_COMPLEX128 {
			fx = new(big.Float).SetInt64(ix)
			fy = new(big.Float).SetFloat64(real(cy))
			cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
			if cmp != 0 {
				return cmp
			}
			return fx.SetFloat64(imag(cy)).Sign()
		}

		return n.internalCompareTyped(ix, iy, _USCTM_INT64)
	}

	if tx == _USCTM_FLOAT64 {
		if ty == _USCTM_INT64 {
			return n.internalCompareTyped(fx, new(big.Float).SetInt64(iy), _USCTM_FLOAT64)
		}

		if ty == _USCTM_COMPLEX128 {
			fy = new(big.Float).SetFloat64(real(cy))
			cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
			if cmp != 0 {
				return cmp
			}
			return fy.SetFloat64(imag(cy)).Sign()
		}

		return n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
	}

	// tx == _USCTM_COMPLEX128
	if ty == _USCTM_INT64 {
		fx = new(big.Float).SetFloat64(real(cx))
		fy = new(big.Float).SetInt64(iy)
		cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
		if cmp != 0 {
			return cmp
		}
		return fx.SetFloat64(imag(cx)).Sign()
	}

	if ty == _USCTM_FLOAT64 {
		fx = new(big.Float).SetFloat64(real(cx))
		cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
		if cmp != 0 {
			return cmp
		}
		return fx.SetFloat64(imag(cx)).Sign()
	}

	return n.internalCompareTyped(cx, cy, _USCTM_COMPLEX128)
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

	ix, fx, cx, tx := extractValueDepends(*vox)
	iy, fy, cy, ty := extractValueDepends(*voy)

	if tx == _USCTM_INT64 {
		if ty == _USCTM_FLOAT64 {
			return n.internalCompareTyped(new(big.Float).SetInt64(ix), fy, _USCTM_FLOAT64)
		}

		if ty == _USCTM_COMPLEX128 {
			fx = new(big.Float).SetInt64(ix)
			fy = new(big.Float).SetFloat64(real(cy))
			cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
			if cmp != 0 {
				return cmp
			}
			return fx.SetFloat64(imag(cy)).Sign()
		}

		return n.internalCompareTyped(ix, iy, _USCTM_INT64)
	}

	if tx == _USCTM_FLOAT64 {
		if ty == _USCTM_INT64 {
			return n.internalCompareTyped(fx, new(big.Float).SetInt64(iy), _USCTM_FLOAT64)
		}

		if ty == _USCTM_COMPLEX128 {
			fy = new(big.Float).SetFloat64(real(cy))
			cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
			if cmp != 0 {
				return cmp
			}
			return fy.SetFloat64(imag(cy)).Sign()
		}

		return n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
	}

	// tx == _USCTM_COMPLEX128
	if ty == _USCTM_INT64 {
		fx = new(big.Float).SetFloat64(real(cx))
		fy = new(big.Float).SetInt64(iy)
		cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
		if cmp != 0 {
			return cmp
		}
		return fx.SetFloat64(imag(cx)).Sign()
	}

	if ty == _USCTM_FLOAT64 {
		fx = new(big.Float).SetFloat64(real(cx))
		cmp := n.internalCompareTyped(fx, fy, _USCTM_FLOAT64)
		if cmp != 0 {
			return cmp
		}
		return fx.SetFloat64(imag(cx)).Sign()
	}

	return n.internalCompareTyped(cx, cy, _USCTM_COMPLEX128)
}

type unsafeCompareTypedMode byte

//goland:noinspection GoSnakeCaseUsage,SpellCheckingInspection
const (
	_USCTM_INT64      unsafeCompareTypedMode = 0
	_USCTM_FLOAT64    unsafeCompareTypedMode = 1
	_USCTM_COMPLEX128 unsafeCompareTypedMode = 2
)

func (n numericComparer) internalCompareTyped(x, y any, mode unsafeCompareTypedMode) int {
	if mode == _USCTM_FLOAT64 {
		xf := x.(*big.Float)
		yf := y.(*big.Float)

		return xf.Cmp(yf)
	}

	if mode == _USCTM_COMPLEX128 {
		xc := x.(complex128)
		yc := y.(complex128)

		xf := new(big.Float)
		yf := new(big.Float)

		c := n.internalCompareTyped(xf.SetFloat64(real(xc)), yf.SetFloat64(real(yc)), _USCTM_FLOAT64)

		if c != 0 {
			return c
		}

		return n.internalCompareTyped(xf.SetFloat64(imag(xc)), yf.SetFloat64(imag(yc)), _USCTM_FLOAT64)
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

func extractValueDepends(voi reflect.Value) (iv int64, fv *big.Float, cv complex128, mode unsafeCompareTypedMode) {
	switch voi.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		mode = _USCTM_INT64
		iv = voi.Int()
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ui64 := voi.Uint()
		if ui64 > math.MaxInt64 {
			mode = _USCTM_FLOAT64
			fv = new(big.Float).SetUint64(voi.Uint())
		} else {
			mode = _USCTM_INT64
			iv = int64(ui64)
		}
		return
	case reflect.Float32, reflect.Float64:
		mode = _USCTM_FLOAT64
		fv = new(big.Float).SetFloat64(voi.Float())
		return
	case reflect.Complex64, reflect.Complex128:
		mode = _USCTM_COMPLEX128
		cv = voi.Complex()
		return
	default:
		panic(fmt.Sprintf("%s %s can not be cast to a number", voi.Kind().String(), voi.Type().String()))
	}
}

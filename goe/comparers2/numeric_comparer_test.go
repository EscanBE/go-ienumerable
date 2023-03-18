package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type randomFunc func() any

var randomFuncsGenerateNumericTest = []randomFunc{
	func() any {
		return int8(rand.Intn(math.MaxUint8) - math.MaxInt8)
	},
	func() any {
		return uint8(rand.Intn(math.MaxUint8))
	},
	func() any {
		return int16(rand.Intn(math.MaxUint16) - math.MaxInt16)
	},
	func() any {
		return uint16(rand.Intn(math.MaxUint16))
	},
	func() any {
		//goland:noinspection GoRedundantConversion
		return int32(rand.Int31n(10_000))
	},
	func() any {
		return uint32(rand.Int31n(10_000))
	},
	func() any {
		return int64(rand.Int31n(10_000))
	},
	func() any {
		return uint64(rand.Int31n(10_000))
	},
	func() any {
		//goland:noinspection GoRedundantConversion
		return int(rand.Int())
	},
	func() any {
		return uint(rand.Int31n(10_000))
	},
	func() any {
		return float32(rand.Int31n(10_000))
	},
	func() any {
		return float64(rand.Int31n(10_000))
	},
	func() any {
		return complex(float32(rand.Int31n(10_000)), float32(0))
	},
	func() any {
		return complex(float64(rand.Int31n(10_000)), float64(0))
	},
	func() any {
		return fmt.Sprintf("%d", rand.Int31n(10_000))
	},
	func() any {
		x := rand.Intn(10) % 2
		if x == 0 {
			return true
		} else {
			return false
		}
	},
	func() any {
		return [2]int{1, 2}
	},
	func() any {
		return nil
	},
}

func Test_numericComparer_CompareTyped(t *testing.T) {
	t.Run("monkey test", func(t *testing.T) {
		report := make(map[int]int)
		radFuncSize := len(randomFuncsGenerateNumericTest)
		comparer := numericComparer{}

		for i := 0; i < 1_000_000; i++ {
			executeMonkeyTest := func() {
				v1 := randomFuncsGenerateNumericTest[rand.Intn(radFuncSize)]()
				v2 := randomFuncsGenerateNumericTest[rand.Intn(radFuncSize)]()

				vo1 := reflect.ValueOf(v1)
				vo2 := reflect.ValueOf(v2)

				v1isNumeric := vo1.CanInt() || vo1.CanUint() || vo1.CanFloat() || vo1.CanComplex()
				v2isNumeric := vo2.CanInt() || vo2.CanUint() || vo2.CanFloat() || vo2.CanComplex()

				//goland:noinspection GoDeferInLoop
				defer deferExpectPanicContains(t, "can not be cast to a number", !v1isNumeric || !v2isNumeric)

				ct := comparer.CompareTyped(v1, v2)

				var fv1, fv2 float64
				if vo1.CanInt() {
					fv1 = float64(vo1.Int())
				} else if vo1.CanUint() {
					fv1 = float64(vo1.Uint())
				} else if vo1.CanFloat() {
					fv1 = vo1.Float()
				} else if vo1.CanComplex() {
					fv1 = real(vo1.Complex())
				}

				if vo2.CanInt() {
					fv2 = float64(vo2.Int())
				} else if vo2.CanUint() {
					fv2 = float64(vo2.Uint())
				} else if vo2.CanFloat() {
					fv2 = vo2.Float()
				} else if vo2.CanComplex() {
					fv2 = real(vo2.Complex())
				}

				var ec int
				if math.Abs(fv1-fv2) <= math.Abs(fv1*AbsEpsilonRatio) {
					ec = 0
				} else if fv1 < fv2 {
					ec = -1
				} else if fv1 > fv2 {
					ec = 1
				} else {
					panic("unknown reason")
				}

				if !assert.Equal(t, ec, ct) {
					t.Errorf("expected value of compare typed %T %f vs %T %f is %d but got %d", v1, fv1, v2, fv2, ec, ct)
					return
				}

				if pc, found := report[ct]; found {
					report[ct] = pc + 1
				} else {
					report[ct] = 1
				}
			}

			executeMonkeyTest()
		}

		fmt.Printf("[Typed] -1: %d\n", report[-1])
		fmt.Printf("[Typed]  0: %d\n", report[00])
		fmt.Printf("[Typed]  1: %d\n", report[01])
	})

	t.Run("compare uint64 ok", func(t *testing.T) {
		u1 := uint64(math.MaxInt64 + 1)
		u2 := uint64(math.MaxInt64 + 2)
		assert.Equal(t, -1, NumericComparer.CompareTyped(u1, u2))
	})
}

func Test_numericComparer_CompareAny(t *testing.T) {
	t.Run("not numeric", func(t *testing.T) {
		defer deferExpectPanicContains(t, "can not be cast to a number", true)
		NumericComparer.CompareAny(true, true)
	})
	t.Run("not numeric", func(t *testing.T) {
		defer deferExpectPanicContains(t, "can not be cast to a number", true)
		NumericComparer.CompareAny(nil, "")
	})
	t.Run("nil pointer", func(t *testing.T) {
		var i16 *int8
		var f32 *float32
		assert.Zero(t, NumericComparer.CompareAny(i16, f32))
	})
	t.Run("monkey test", func(t *testing.T) {
		report := make(map[int]int)
		radFuncSize := len(randomFuncsGenerateNumericTest)

		for i := 0; i < 1_000_000; i++ {
			//fmt.Println(i)
			executeMonkeyTest := func() int {
				v1 := randomFuncsGenerateNumericTest[rand.Intn(radFuncSize)]()
				v2 := randomFuncsGenerateNumericTest[rand.Intn(radFuncSize)]()

				vo1 := reflect.ValueOf(v1)
				vo2 := reflect.ValueOf(v2)

				v1isNumeric := vo1.CanInt() || vo1.CanUint() || vo1.CanFloat() || vo1.CanComplex() || !vo1.IsValid()
				v2isNumeric := vo2.CanInt() || vo2.CanUint() || vo2.CanFloat() || vo2.CanComplex() || !vo2.IsValid()

				//goland:noinspection GoDeferInLoop
				defer deferExpectPanicContains(t, "can not be cast to a number", !v1isNumeric || !v2isNumeric)

				ca := NumericComparer.CompareAny(v1, v2)

				var fv1, fv2 float64
				var nilV1, nilV2 bool
				if vo1.CanInt() {
					fv1 = float64(vo1.Int())
				} else if vo1.CanUint() {
					fv1 = float64(vo1.Uint())
				} else if vo1.CanFloat() {
					fv1 = vo1.Float()
				} else if vo1.CanComplex() {
					fv1 = real(vo1.Complex())
				} else {
					nilV1 = true
				}

				if vo2.CanInt() {
					fv2 = float64(vo2.Int())
				} else if vo2.CanUint() {
					fv2 = float64(vo2.Uint())
				} else if vo2.CanFloat() {
					fv2 = vo2.Float()
				} else if vo2.CanComplex() {
					fv2 = real(vo2.Complex())
				} else {
					nilV2 = true
				}

				var ec int
				if nilV1 && nilV2 {
					ec = 0
				} else if nilV1 {
					ec = -1
				} else if nilV2 {
					ec = 1
				} else if math.Abs(fv1-fv2) <= math.Abs(fv1*AbsEpsilonRatio) {
					ec = 0
				} else if fv1 < fv2 {
					ec = -1
				} else if fv1 > fv2 {
					ec = 1
				} else {
					panic("unknown reason")
				}

				if !assert.Equal(t, ec, ca) {
					t.Errorf("expected value of compare any %T %f vs %T %f is %d but got %d", v1, fv1, v2, fv2, ec, ca)
					return -1
				}

				if pc, found := report[ca]; found {
					report[ca] = pc + 1
				} else {
					report[ca] = 1
				}

				return 0
			}

			if executeMonkeyTest() != 0 {
				return
			}
		}

		fmt.Printf("[Any] -1: %d\n", report[-1])
		fmt.Printf("[Any]  0: %d\n", report[00])
		fmt.Printf("[Any]  1: %d\n", report[01])
	})

	t.Run("compare uint64 ok", func(t *testing.T) {
		u1 := uint64(math.MaxInt64 + 1)
		u2 := uint64(math.MaxInt64 + 2)
		assert.Equal(t, -1, NumericComparer.CompareAny(u1, u2))
	})
}

//goland:noinspection GoRedundantConversion,SpellCheckingInspection
func Test_numericComparer_CompareTyped_CompareAny(t *testing.T) {
	lessParamCandiates := []any{int8(1), int16(1), int32(1), int64(1), int(1), uint8(1), uint16(1), uint32(1), uint64(1), uint(1), float32(1.1), float64(1.1), complex(float32(1.1), float32(1.1)), complex(float64(1.1), float64(1.1))}
	greaterParamCandiates := []any{int8(3), int16(3), int32(3), int64(3), int(3), uint8(3), uint16(3), uint32(3), uint64(3), uint(3), float32(3.3), float64(3.3), complex(float32(3.3), float32(3.3)), complex(float64(1.1), float64(2)), complex(float64(3.3), float64(3.3))}

	for _, less := range lessParamCandiates {
		// TODO resolve problem compare float vs complex and complex vs complex
		for _, greater := range greaterParamCandiates {
			t.Run(fmt.Sprintf("Typed (lesser) %T %v vs %T %v (greater)", less, less, greater, greater), func(t *testing.T) {
				assert.Equalf(t, -1, NumericComparer.CompareTyped(less, greater), "%v must < %v", less, greater)
				assert.Equalf(t, 1, NumericComparer.CompareTyped(greater, less), "%v must > %v", greater, less)
				assert.Zerof(t, NumericComparer.CompareTyped(less, less), "%v must equals to itself", less)
				assert.Zerof(t, NumericComparer.CompareTyped(greater, greater), "%v must equals to itself", greater)
			})
			t.Run(fmt.Sprintf("Any (lesser) %T %v vs %T %v (greater)", less, less, greater, greater), func(t *testing.T) {
				assert.Equalf(t, -1, NumericComparer.CompareAny(less, greater), "%v must < %v", less, greater)
				assert.Equalf(t, 1, NumericComparer.CompareAny(greater, less), "%v must > %v", greater, less)
				assert.Zerof(t, NumericComparer.CompareAny(less, less), "%v must equals to itself", less)
				assert.Zerof(t, NumericComparer.CompareAny(greater, greater), "%v must equals to itself", greater)
			})
		}
	}

	t.Run("sample confirm", func(t *testing.T) {
		i8 := int8(1)
		f6 := float64(1.1)
		assert.Equal(t, -1, NumericComparer.CompareTyped(i8, f6))
		assert.Equal(t, 1, NumericComparer.CompareTyped(f6, i8))
		assert.Equal(t, -1, NumericComparer.CompareAny(&i8, f6))
		assert.Equal(t, 1, NumericComparer.CompareAny(f6, &i8))
		anyI8 := any(&i8)
		anyF6 := any(&f6)
		assert.Equal(t, -1, NumericComparer.CompareAny(anyI8, anyF6))
		assert.Equal(t, 1, NumericComparer.CompareAny(anyF6, anyI8))
		assert.Equal(t, -1, NumericComparer.CompareAny(anyI8, &anyF6))
		assert.Equal(t, 1, NumericComparer.CompareAny(&anyF6, anyI8))
		assert.Equal(t, -1, NumericComparer.CompareAny(&anyI8, anyF6))
		assert.Equal(t, 1, NumericComparer.CompareAny(anyF6, &anyI8))
		assert.Equal(t, -1, NumericComparer.CompareAny(&anyI8, &anyF6))
		assert.Equal(t, 1, NumericComparer.CompareAny(&anyF6, &anyI8))
	})

	t.Run("accept type which defined on top of numeric type", func(t *testing.T) {
		assert.Zero(t, NumericComparer.CompareAny(time.Minute, time.Minute))
		assert.Zero(t, NumericComparer.CompareAny(_USCTM_COMPLEX128, _USCTM_COMPLEX128))
	})
}

func Test_assertExactOneParamNotNil(t *testing.T) {
	t.Run("zero of all", func(t *testing.T) {
		var i *int64
		var bf *big.Float
		var c *complex128
		defer deferExpectPanicContains(t, "expect exactly one param not nil, found 0/3", true)
		assertExactOneValueNotNil(i, bf, c)
	})
	t.Run("one of all", func(t *testing.T) {
		var i int64 = 1
		var bf *big.Float
		var c *complex128
		assertExactOneValueNotNil(&i, bf, c)
	})
	t.Run("two of all", func(t *testing.T) {
		var i int64 = 1
		var bf = new(big.Float).SetInt64(1)
		var c *complex128
		defer deferExpectPanicContains(t, "expect exactly one param not nil, found 2/3", true)
		assertExactOneValueNotNil(&i, bf, c)
	})
	t.Run("all", func(t *testing.T) {
		var i int64 = 1
		var bf = new(big.Float).SetInt64(1)
		var c = complex(1, 2)
		defer deferExpectPanicContains(t, "expect exactly one param not nil, found 3/3", true)
		assertExactOneValueNotNil(&i, bf, &c)
	})
}

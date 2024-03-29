package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"math/big"
	"reflect"
	"time"
)

// TODO embed big.Int, big.Float into numeric comparer

//goland:noinspection GoRedundantConversion
var mappedDefaultComparers = map[reflect.Type]IComparer[any]{
	getDefaultComparerKeyFromSampleValue(int8(1)):                             NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(int16(1)):                            NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(int32(1)):                            NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(int64(1)):                            NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(int(1)):                              NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(uint8(1)):                            NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(uint16(1)):                           NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(uint32(1)):                           NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(uint64(1)):                           NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(uint(1)):                             NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(float32(1.0)):                        NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(float64(1.0)):                        NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(complex(float32(1.0), float32(1.0))): NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(complex(float64(1.0), float64(1.0))): NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(new(big.Int).SetInt64(1)):            ConvertFromComparerIntoDefaultComparer[*big.Int](NewBigIntComparer()),
	getDefaultComparerKeyFromSampleValue(new(big.Float).SetInt64(1)):          ConvertFromComparerIntoDefaultComparer[*big.Float](NewBigFloatComparer()),
	getDefaultComparerKeyFromSampleValue("string"):                            ConvertFromComparerIntoDefaultComparer[string](NewStringComparer()),
	getDefaultComparerKeyFromSampleValue(true):                                ConvertFromComparerIntoDefaultComparer[bool](NewBoolComparer()),
	getDefaultComparerKeyFromSampleValue(time.Minute):                         NewNumericComparer(),
	getDefaultComparerKeyFromSampleValue(time.Now()):                          ConvertFromComparerIntoDefaultComparer[time.Time](NewTimeComparer()),
}

func getDefaultComparerKeyFromSampleValue(sampleValue any) reflect.Type {
	key, err := tryGetDefaultComparerKeyFromSampleValue(sampleValue)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func tryGetDefaultComparerKeyFromSampleValue(sampleValue any) (key reflect.Type, err error) {
	value, isNil := reflection.RootValueExtractor(sampleValue)

	if isNil || !value.IsValid() {
		to := reflect.TypeOf(sampleValue)
		if to != nil {
			if to.Kind() == reflect.Pointer {
				key = to.Elem()
				return
			}
		}
		err = fmt.Errorf("sample value can not be nil or invalid")
		return
	}

	key = value.Type()
	return
}

func TryGetDefaultComparer[T any]() (comparer IComparer[any], ok bool) {
	key, _ := tryGetDefaultComparerKeyFromSampleValue(*new(T))
	comparer, ok = mappedDefaultComparers[key]
	return
}

func GetDefaultComparer[T any]() IComparer[any] {
	comparer, found := TryGetDefaultComparer[T]()
	if !found {
		panic(fmt.Errorf("not found any default comparer for %T", *new(T)))
	}
	return comparer
}

func TryGetDefaultComparerFromValue(sampleValue any) (comparer IComparer[any], ok bool) {
	key, err := tryGetDefaultComparerKeyFromSampleValue(sampleValue)
	if err != nil {
		return
	}

	comparer, ok = mappedDefaultComparers[key]
	return
}

func RegisterDefaultComparer[T any](comparer IComparer[T]) {
	key, _ := tryGetDefaultComparerKeyFromSampleValue(*new(T))
	mappedDefaultComparers[key] = ConvertFromComparerIntoDefaultComparer[T](comparer)
}

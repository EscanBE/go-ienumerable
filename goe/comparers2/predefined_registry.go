package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"reflect"
	"time"
)

// TODO implement comparer for big.Int, big.Float

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
	getDefaultComparerKeyFromSampleValue("string"):                            ConvertFromComparerIntoDefaultComparer[string](NewStringComparer()),
	getDefaultComparerKeyFromSampleValue(true):                                ConvertFromComparerIntoDefaultComparer[bool](NewBoolComparer()),
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

func TryGetDefaultComparerFromValue(sampleValue any) (comparer IComparer[any], ok bool) {
	key, err := tryGetDefaultComparerKeyFromSampleValue(sampleValue)
	if err != nil {
		return
	}

	comparer, ok = mappedDefaultComparers[key]
	return
}

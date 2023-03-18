package comparers

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe/reflection"
	"reflect"
	"time"
)

type DefaultComparerKey struct {
	Kind reflect.Kind
	Type reflect.Type
}

// TODO implement comparer for big.Int, big.Float

//goland:noinspection GoRedundantConversion
var mappedDefaultComparers = map[DefaultComparerKey]IComparer[any]{
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

func getDefaultComparerKeyFromSampleValue(sampleValue any) DefaultComparerKey {
	key, err := tryGetDefaultComparerKeyFromSampleValue(sampleValue)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func tryGetDefaultComparerKeyFromSampleValue(sampleValue any) (key DefaultComparerKey, err error) {
	value, isNil := reflection.RootValueExtractor(sampleValue)

	if isNil || !value.IsValid() {
		err = fmt.Errorf("sample value can not be nil or invalid")
		return
	}

	key = DefaultComparerKey{
		Kind: value.Kind(),
		Type: value.Type(),
	}
	return
}

func TryGetDefaultComparer[T any]() (comparer IComparer[any], ok bool) {
	defer func() {
		err := recover()
		if err != nil {
			ok = false
			// silent error
		}
	}()

	key := getDefaultComparerKeyFromSampleValue(*new(T))
	comparer, ok = mappedDefaultComparers[key]
	return
}

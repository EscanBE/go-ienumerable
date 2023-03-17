package reflection

import (
	"fmt"
	"reflect"
)

func RootValueExtractor(v any) (value *reflect.Value, isNil bool) {
	if v == nil {
		value = nil
		isNil = true
		return
	}

	defer func() {
		if !isNil {
			if value == nil {
				isNil = true
			} else if canValueNil(value) {
				isNil = value.IsNil()
			}
		}
	}()

	vo := reflect.ValueOf(v)

	recursiveLevel := 0
	for {
		recursiveLevel++
		panicMaxRecursiveLoopReached(recursiveLevel)

		switch vo.Kind() {
		case reflect.Invalid:
			value = nil
			return
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
			reflect.Float32, reflect.Float64,
			reflect.Complex64, reflect.Complex128,
			reflect.Array,
			reflect.Chan,
			reflect.Func:
			value = &vo
			return
		case reflect.Interface:
			vo = vo.Elem()
			continue
		case reflect.Map:
			value = &vo
			isNil = vo.IsNil()
			return
		case reflect.Pointer:
			if vo.IsNil() {
				value = &vo
				isNil = true
				return
			}

			vo = vo.Elem()
			continue
		case reflect.Slice, reflect.String, reflect.Struct:
			value = &vo
			return
		//case reflect.UnsafePointer: fall down
		default:
			panic(fmt.Sprintf("not yet supported kind %s", vo.Kind().String()))
		}
	}
}

func canValueNil(value interface{}) bool {
	var v reflect.Value
	if vnp, ok1 := value.(*reflect.Value); ok1 {
		v = *vnp
	} else {
		v = value.(reflect.Value)
	}

	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return true
	default:
		return false
	}
}

const maxRecursiveLevelForRootValueExtractor = 100

func panicMaxRecursiveLoopReached(level int) {
	if level >= maxRecursiveLevelForRootValueExtractor {
		panic("max recursive reached")
	}
}

package go_ienumerable

import (
	"fmt"
	"strings"
)

func (src *enumerable[T]) WithEqualsComparer(equalsComparer func(v1 T, v2 T) bool) IEnumerable[T] {
	src.equalityComparer = equalsComparer
	return src
}

func (src *enumerable[T]) WithLessComparer(lessComparer func(left T, right T) bool) IEnumerable[T] {
	src.lessComparer = lessComparer
	return src
}

func (src *enumerable[T]) WithDefaultComparers() IEnumerable[T] {
	_type := fmt.Sprintf("%T", *new(T))
	switch _type {
	case "int8":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(int8) == any(v2).(int8)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(int8) < any(v2).(int8)
		}
		break
	case "uint8":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uint8) == any(v2).(uint8)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uint8) < any(v2).(uint8)
		}
		break
	case "int16":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(int16) == any(v2).(int16)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(int16) < any(v2).(int16)
		}
		break
	case "uint16":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uint16) == any(v2).(uint16)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uint16) < any(v2).(uint16)
		}
		break
	case "int32":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(int32) == any(v2).(int32)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(int32) < any(v2).(int32)
		}
		break
	case "uint32":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uint32) == any(v2).(uint32)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uint32) < any(v2).(uint32)
		}
		break
	case "int64":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(int64) == any(v2).(int64)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(int64) < any(v2).(int64)
		}
		break
	case "uint64":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uint64) == any(v2).(uint64)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uint64) < any(v2).(uint64)
		}
		break
	case "int":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(int) == any(v2).(int)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(int) < any(v2).(int)
		}
		break
	case "uint":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uint) == any(v2).(uint)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uint) < any(v2).(uint)
		}
		break
	case "uintptr":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(uintptr) == any(v2).(uintptr)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(uintptr) < any(v2).(uintptr)
		}
		break
	case "float32":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(float32) == any(v2).(float32)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(float32) < any(v2).(float32)
		}
		break
	case "float64":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(float64) == any(v2).(float64)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(float64) < any(v2).(float64)
		}
		break
	case "complex64":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(complex64) == any(v2).(complex64)
		}
		break
	case "complex128":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(complex128) == any(v2).(complex128)
		}
		break
	case "string":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(string) == any(v2).(string)
		}
		src.lessComparer = func(v1, v2 T) bool {
			return any(v1).(string) < any(v2).(string)
		}
		break
	case "bool":
		src.equalityComparer = func(v1, v2 T) bool {
			return any(v1).(bool) == any(v2).(bool)
		}
		src.lessComparer = func(v1, v2 T) bool {
			b1 := any(v1).(bool)
			b2 := any(v2).(bool)
			return !b1 && b2
		}
		break
	default:
		panic(fmt.Errorf("not yet supported inject comparers for type %s", _type))
	}
	return src
}

type requireWithExtraFunc byte

const (
	requireEqualityComparer requireWithExtraFunc = 0b01
	requireLessComparer     requireWithExtraFunc = 0b10
)

func panicRequire(require requireWithExtraFunc) {
	requiresName := getRequireName(require)
	panic(fmt.Errorf("the following comparer must be set: [%s]", strings.Join(requiresName, ",")))
}

func getRequireName(require requireWithExtraFunc) []string {
	result := make([]string, 0)

	if require&requireEqualityComparer == requireEqualityComparer {
		result = append(result, "Equals Comparer (can be set using WithEqualsComparer)")
	}

	if require&requireLessComparer == requireLessComparer {
		result = append(result, "Less Comparer (can be set using WithLessComparer)")
	}

	return result
}

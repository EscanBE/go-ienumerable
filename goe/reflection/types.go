package reflection

import "fmt"

type UnboxResultState int8

//goland:noinspection GoSnakeCaseUsage
const (
	UNBOX_FAILED   UnboxResultState = -1
	UNBOX_NIL      UnboxResultState = 0
	UNBOX_OVERFLOW UnboxResultState = 1
	UNBOX_SUCCESS  UnboxResultState = 2
)

// TODO remove if not used
func (s UnboxResultState) AssertIsSuccess() {
	if s != UNBOX_SUCCESS {
		panic(fmt.Sprintf("unbox result status is %v, require success", s))
	}
}

type UnboxFloat64DataType int8

//goland:noinspection GoSnakeCaseUsage
const (
	UF64_TYPE_FAILED  UnboxFloat64DataType = -1
	UF64_TYPE_NIL     UnboxFloat64DataType = 0
	UF64_TYPE_FLOAT64 UnboxFloat64DataType = 1
	UF64_TYPE_INT64   UnboxFloat64DataType = 2
)

// TODO remove if not used
func (t UnboxFloat64DataType) AssertHasNonNilResult() {
	if t != UF64_TYPE_FLOAT64 && t != UF64_TYPE_INT64 {
		panic(fmt.Sprintf("unbox data type result status is %v, require non-nil int64 or float", t))
	}
}

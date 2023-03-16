package reflection

type UnboxResultState int8

//goland:noinspection GoSnakeCaseUsage
const (
	UNBOX_FAILED   UnboxResultState = -1
	UNBOX_NIL      UnboxResultState = 0
	UNBOX_OVERFLOW UnboxResultState = 1
	UNBOX_SUCCESS  UnboxResultState = 2
)

type UnboxFloat64DataType int8

//goland:noinspection GoSnakeCaseUsage
const (
	UF64_TYPE_FAILED  UnboxFloat64DataType = -1
	UF64_TYPE_NIL     UnboxFloat64DataType = 0
	UF64_TYPE_FLOAT64 UnboxFloat64DataType = 1
	UF64_TYPE_INT64   UnboxFloat64DataType = 2
)

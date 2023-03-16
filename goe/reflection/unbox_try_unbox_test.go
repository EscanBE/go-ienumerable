package reflection

import (
	"math"
	"testing"
)

func TestTryUnboxAnyAsInt64(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name       string
		input      any
		wantResult int64
		wantState  UnboxResultState
	}{
		{
			name:      "nil",
			input:     nil,
			wantState: UNBOX_NIL,
		},
		{
			name:       "unsigned ok",
			input:      uint64(math.MaxInt64),
			wantResult: int64(math.MaxInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min",
			input:      int64(math.MinInt64),
			wantResult: int64(math.MinInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow unsigned",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "float64",
			input:     float64(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "float32",
			input:     float32(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "string",
			input:     "0",
			wantState: UNBOX_FAILED,
		},
		{
			name:      "bool",
			input:     true,
			wantState: UNBOX_FAILED,
		},
		{
			name:       "positive",
			input:      int64(8),
			wantResult: int64(8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "negative",
			input:      int64(-8),
			wantResult: int64(-8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name: "nil pointer int64",
			input: func() *int64 {
				var nilInt64 *int64
				return nilInt64
			}(),
			wantState: UNBOX_NIL,
		},
		{
			name:      "array",
			input:     []int64{},
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil array",
			input: func() []int64 {
				var arr []int64
				return arr
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "slice",
			input:     []int64{1, 2, 3, 4}[1:],
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil slice",
			input: func() []int64 {
				s := []int64{1, 2, 3, 4}[1:]
				s = nil
				return s
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:       "max int8",
			input:      int8(math.MaxInt8),
			wantResult: int64(math.MaxInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int8",
			input:      int8(math.MinInt8),
			wantResult: int64(math.MinInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint8",
			input:      uint8(math.MaxUint8),
			wantResult: int64(math.MaxUint8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int16",
			input:      int16(math.MaxInt16),
			wantResult: int64(math.MaxInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int16",
			input:      int16(math.MinInt16),
			wantResult: int64(math.MinInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint16",
			input:      uint16(math.MaxUint16),
			wantResult: int64(math.MaxUint16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int32",
			input:      int32(math.MaxInt32),
			wantResult: int64(math.MaxInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int32",
			input:      int32(math.MinInt32),
			wantResult: int64(math.MinInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint32",
			input:      uint32(math.MaxUint32),
			wantResult: int64(math.MaxUint32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int64",
			input:      int64(math.MaxInt64),
			wantResult: int64(math.MaxInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int64",
			input:      int64(math.MinInt64),
			wantResult: int64(math.MinInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max uint64",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "max int",
			input:      int(math.MaxInt),
			wantResult: int64(math.MaxInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int",
			input:      int(math.MinInt),
			wantResult: int64(math.MinInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max uint",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotResult int64
			var gotState UnboxResultState
			var testName string

			assertResult := func() {
				if gotResult != tt.wantResult {
					t.Errorf("[%s] TryUnboxAnyAsInt64() gotResult = %v, want %v", testName, gotResult, tt.wantResult)
				}
				if gotState != tt.wantState {
					t.Errorf("[%s] TryUnboxAnyAsInt64() gotState = %v, want %v", testName, gotState, tt.wantState)
				}
			}

			testName = "L0"
			gotResult, gotState = TryUnboxAnyAsInt64(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotResult, gotState = TryUnboxAnyAsInt64(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt64(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt64(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt64(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt64(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt64(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

func TestTryUnboxAnyAsInt32(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name       string
		input      any
		wantResult int32
		wantState  UnboxResultState
	}{
		{
			name:      "nil",
			input:     nil,
			wantState: UNBOX_NIL,
		},
		{
			name:       "unsigned ok",
			input:      uint32(math.MaxInt32),
			wantResult: int32(math.MaxInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min",
			input:      int32(math.MinInt32),
			wantResult: int32(math.MinInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow unsigned",
			input:     uint32(math.MaxUint32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "float64",
			input:     float64(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "float32",
			input:     float32(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "string",
			input:     "0",
			wantState: UNBOX_FAILED,
		},
		{
			name:      "bool",
			input:     true,
			wantState: UNBOX_FAILED,
		},
		{
			name:       "positive",
			input:      int32(8),
			wantResult: int32(8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "negative",
			input:      int32(-8),
			wantResult: int32(-8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow",
			input:     uint32(math.MaxUint32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name: "nil pointer int32",
			input: func() *int32 {
				var nilInt32 *int32
				return nilInt32
			}(),
			wantState: UNBOX_NIL,
		},
		{
			name:      "array",
			input:     []int64{},
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil array",
			input: func() []int {
				var arr []int
				return arr
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "slice",
			input:     []int{1, 2, 3, 4}[1:],
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil slice",
			input: func() []int {
				s := []int{1, 2, 3, 4}[1:]
				s = nil
				return s
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:       "max int8",
			input:      int8(math.MaxInt8),
			wantResult: int32(math.MaxInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int8",
			input:      int8(math.MinInt8),
			wantResult: int32(math.MinInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint8",
			input:      uint8(math.MaxUint8),
			wantResult: int32(math.MaxUint8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int16",
			input:      int16(math.MaxInt16),
			wantResult: int32(math.MaxInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int16",
			input:      int16(math.MinInt16),
			wantResult: int32(math.MinInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint16",
			input:      uint16(math.MaxUint16),
			wantResult: int32(math.MaxUint16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int32",
			input:      int32(math.MaxInt32),
			wantResult: int32(math.MaxInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int32",
			input:      int32(math.MinInt32),
			wantResult: int32(math.MinInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max uint32",
			input:     uint32(math.MaxUint32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max int64",
			input:     int64(math.MaxInt64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int64",
			input:     int64(math.MinInt64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint64",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max int",
			input:     int(math.MaxInt),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int",
			input:     int(math.MinInt),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotResult int32
			var gotState UnboxResultState
			var testName string

			assertResult := func() {
				if gotResult != tt.wantResult {
					t.Errorf("[%s] TryUnboxAnyAsInt32() gotResult = %v, want %v", testName, gotResult, tt.wantResult)
				}
				if gotState != tt.wantState {
					t.Errorf("[%s] TryUnboxAnyAsInt32() gotState = %v, want %v", testName, gotState, tt.wantState)
				}
			}

			testName = "L0"
			gotResult, gotState = TryUnboxAnyAsInt32(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotResult, gotState = TryUnboxAnyAsInt32(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt32(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt32(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt32(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt32(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt32(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

func TestTryUnboxAnyAsInt(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name       string
		input      any
		wantResult int
		wantState  UnboxResultState
	}{
		{
			name:      "nil",
			input:     nil,
			wantState: UNBOX_NIL,
		},
		{
			name:       "unsigned ok",
			input:      uint(math.MaxInt),
			wantResult: int(math.MaxInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min",
			input:      int(math.MinInt),
			wantResult: int(math.MinInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow unsigned",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "float64",
			input:     float64(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "float32",
			input:     float32(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "string",
			input:     "0",
			wantState: UNBOX_FAILED,
		},
		{
			name:      "bool",
			input:     true,
			wantState: UNBOX_FAILED,
		},
		{
			name:       "positive",
			input:      int(8),
			wantResult: int(8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "negative",
			input:      int(-8),
			wantResult: int(-8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "overflow",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name: "nil pointer int",
			input: func() *int {
				var nilInt *int
				return nilInt
			}(),
			wantState: UNBOX_NIL,
		},
		{
			name:      "array",
			input:     []int{},
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil array",
			input: func() []int {
				var arr []int
				return arr
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "slice",
			input:     []int{1, 2, 3, 4}[1:],
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil slice",
			input: func() []int {
				s := []int{1, 2, 3, 4}[1:]
				s = nil
				return s
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:       "max int8",
			input:      int8(math.MaxInt8),
			wantResult: int(math.MaxInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int8",
			input:      int8(math.MinInt8),
			wantResult: int(math.MinInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint8",
			input:      uint8(math.MaxUint8),
			wantResult: int(math.MaxUint8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int16",
			input:      int16(math.MaxInt16),
			wantResult: int(math.MaxInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int16",
			input:      int16(math.MinInt16),
			wantResult: int(math.MinInt16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint16",
			input:      uint16(math.MaxUint16),
			wantResult: int(math.MaxUint16),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int32",
			input:      int32(math.MaxInt32),
			wantResult: int(math.MaxInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int32",
			input:      int32(math.MinInt32),
			wantResult: int(math.MinInt32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max uint32",
			input:      uint32(math.MaxUint32),
			wantResult: int(math.MaxUint32),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "max int64",
			input:      int64(math.MaxInt64),
			wantResult: int(math.MaxInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int64",
			input:      int64(math.MinInt64),
			wantResult: int(math.MinInt64),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max uint64",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "max int",
			input:      int(math.MaxInt),
			wantResult: int(math.MaxInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min int",
			input:      int(math.MinInt),
			wantResult: int(math.MinInt),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max uint",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotResult int
			var gotState UnboxResultState
			var testName string

			assertResult := func() {
				if gotResult != tt.wantResult {
					t.Errorf("[%s] TryUnboxAnyAsInt() gotResult = %v, want %v", testName, gotResult, tt.wantResult)
				}
				if gotState != tt.wantState {
					t.Errorf("[%s] TryUnboxAnyAsInt() gotState = %v, want %v", testName, gotState, tt.wantState)
				}
			}

			testName = "L0"
			gotResult, gotState = TryUnboxAnyAsInt(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotResult, gotState = TryUnboxAnyAsInt(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotResult, gotState = TryUnboxAnyAsInt(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsInt(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

func TestTryUnboxAnyAsByte(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name       string
		input      any
		wantResult byte
		wantState  UnboxResultState
	}{
		{
			name:      "nil",
			input:     nil,
			wantState: UNBOX_NIL,
		},
		{
			name:       "unsigned ok",
			input:      byte(math.MaxUint8),
			wantResult: byte(math.MaxUint8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:       "min",
			input:      byte(0),
			wantResult: byte(0),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "float64",
			input:     float64(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "float32",
			input:     float32(1.0),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "string",
			input:     "0",
			wantState: UNBOX_FAILED,
		},
		{
			name:      "bool",
			input:     true,
			wantState: UNBOX_FAILED,
		},
		{
			name:       "positive",
			input:      byte(8),
			wantResult: byte(8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name: "nil pointer byte",
			input: func() *byte {
				var nilByte *byte
				return nilByte
			}(),
			wantState: UNBOX_NIL,
		},
		{
			name:      "array",
			input:     []byte{},
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil array",
			input: func() []byte {
				var arr []byte
				return arr
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:      "slice",
			input:     []byte{1, 2, 3, 4}[1:],
			wantState: UNBOX_FAILED,
		},
		{
			name: "nil slice",
			input: func() []byte {
				s := []byte{1, 2, 3, 4}[1:]
				s = nil
				return s
			}(),
			wantState: UNBOX_FAILED,
		},
		{
			name:       "max int8",
			input:      byte(math.MaxInt8),
			wantResult: byte(math.MaxInt8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "min int8",
			input:     int8(math.MinInt8),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "max uint8",
			input:      uint8(math.MaxUint8),
			wantResult: byte(math.MaxUint8),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max int16",
			input:     int16(math.MaxInt16),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int16",
			input:     int16(math.MinInt16),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint16",
			input:     uint16(math.MaxUint16),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "min uint16",
			input:      uint16(0),
			wantResult: byte(0),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max int32",
			input:     int32(math.MaxInt32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int32",
			input:     int32(math.MinInt32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint32",
			input:     uint32(math.MaxUint32),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "min uint32",
			input:      uint32(0),
			wantResult: byte(0),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max int64",
			input:     int64(math.MaxInt64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int64",
			input:     int64(math.MinInt64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint64",
			input:     uint64(math.MaxUint64),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "min uint64",
			input:      uint64(0),
			wantResult: byte(0),
			wantState:  UNBOX_SUCCESS,
		},
		{
			name:      "max int",
			input:     int(math.MaxInt),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "min int",
			input:     int(math.MinInt),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:      "max uint",
			input:     uint(math.MaxUint),
			wantState: UNBOX_OVERFLOW,
		},
		{
			name:       "min uint",
			input:      uint(0),
			wantResult: byte(0),
			wantState:  UNBOX_SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotResult byte
			var gotState UnboxResultState
			var testName string

			assertResult := func() {
				if gotResult != tt.wantResult {
					t.Errorf("[%s] TryUnboxAnyAsByte() gotResult = %v, want %v", testName, gotResult, tt.wantResult)
				}
				if gotState != tt.wantState {
					t.Errorf("[%s] TryUnboxAnyAsByte() gotState = %v, want %v", testName, gotState, tt.wantState)
				}
			}

			testName = "L0"
			gotResult, gotState = TryUnboxAnyAsByte(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotResult, gotState = TryUnboxAnyAsByte(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotResult, gotState = TryUnboxAnyAsByte(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotResult, gotState = TryUnboxAnyAsByte(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsByte(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotResult, gotState = TryUnboxAnyAsByte(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotResult, gotState = TryUnboxAnyAsByte(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

func TestTryUnboxAnyAsInt64OrFloat64(t *testing.T) {
	//goland:noinspection GoRedundantConversion
	tests := []struct {
		name               string
		input              any
		wantResultInt64    int64
		wantResultFloat64  float64
		wantResultDataType UnboxFloat64DataType
		wantState          UnboxResultState
	}{
		{
			name:               "nil",
			input:              nil,
			wantResultDataType: UF64_TYPE_NIL,
			wantState:          UNBOX_NIL,
		},
		{
			name:               "unsigned ok",
			input:              uint64(math.MaxInt64),
			wantResultInt64:    int64(math.MaxInt64),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min",
			input:              int64(math.MinInt64),
			wantResultInt64:    int64(math.MinInt64),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "overflow int64, should returns float",
			input:              uint64(math.MaxUint64),
			wantResultFloat64:  float64(math.MaxUint64),
			wantResultDataType: UF64_TYPE_FLOAT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "float64",
			input:              float64(1.0),
			wantResultFloat64:  float64(1.0),
			wantResultDataType: UF64_TYPE_FLOAT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "float32",
			input:              float32(1.0),
			wantResultFloat64:  float64(1.0),
			wantResultDataType: UF64_TYPE_FLOAT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "string",
			input:              "0",
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name:               "bool",
			input:              true,
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name:               "positive",
			input:              int64(8),
			wantResultInt64:    int64(8),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "negative",
			input:              int64(-8),
			wantResultInt64:    int64(-8),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name: "nil pointer int64",
			input: func() *int64 {
				var nilInt64 *int64
				return nilInt64
			}(),
			wantResultDataType: UF64_TYPE_NIL,
			wantState:          UNBOX_NIL,
		},
		{
			name: "nil pointer float64",
			input: func() *float64 {
				var nilFloat64 *float64
				return nilFloat64
			}(),
			wantResultDataType: UF64_TYPE_NIL,
			wantState:          UNBOX_NIL,
		},
		{
			name:               "array",
			input:              []float64{},
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name: "nil array",
			input: func() []float64 {
				var arr []float64
				return arr
			}(),
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name:               "slice",
			input:              []float64{1, 2, 3, 4}[1:],
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name: "nil slice",
			input: func() []float64 {
				s := []float64{1, 2, 3, 4}[1:]
				s = nil
				return s
			}(),
			wantResultDataType: UF64_TYPE_FAILED,
			wantState:          UNBOX_FAILED,
		},
		{
			name:               "max int8",
			input:              int8(math.MaxInt8),
			wantResultInt64:    int64(math.MaxInt8),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min int8",
			input:              int8(math.MinInt8),
			wantResultInt64:    int64(math.MinInt8),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max uint8",
			input:              uint8(math.MaxUint8),
			wantResultInt64:    int64(math.MaxUint8),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max int16",
			input:              int16(math.MaxInt16),
			wantResultInt64:    int64(math.MaxInt16),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min int16",
			input:              int16(math.MinInt16),
			wantResultInt64:    int64(math.MinInt16),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max uint16",
			input:              uint16(math.MaxUint16),
			wantResultInt64:    int64(math.MaxUint16),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max int32",
			input:              int32(math.MaxInt32),
			wantResultInt64:    int64(math.MaxInt32),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min int32",
			input:              int32(math.MinInt32),
			wantResultInt64:    int64(math.MinInt32),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max uint32",
			input:              uint32(math.MaxUint32),
			wantResultInt64:    int64(math.MaxUint32),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max int64",
			input:              int64(math.MaxInt64),
			wantResultInt64:    int64(math.MaxInt64),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min int64",
			input:              int64(math.MinInt64),
			wantResultInt64:    int64(math.MinInt64),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max uint64",
			input:              uint64(math.MaxUint64),
			wantResultFloat64:  float64(math.MaxUint64),
			wantResultDataType: UF64_TYPE_FLOAT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max int",
			input:              int(math.MaxInt),
			wantResultInt64:    int64(math.MaxInt),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "min int",
			input:              int(math.MinInt),
			wantResultInt64:    int64(math.MinInt),
			wantResultDataType: UF64_TYPE_INT64,
			wantState:          UNBOX_SUCCESS,
		},
		{
			name:               "max uint",
			input:              uint(math.MaxUint),
			wantResultFloat64:  float64(math.MaxUint),
			wantResultDataType: UF64_TYPE_FLOAT64,
			wantState:          UNBOX_SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotResultInt64 int64
			var gotResultFloat64 float64
			var gotDataType UnboxFloat64DataType
			var gotState UnboxResultState
			var testName string

			assertResult := func() {
				if gotResultInt64 != tt.wantResultInt64 {
					t.Errorf("[%s] TryUnboxAnyAsInt64OrFloat64() gotResultInt64 = %v, want %v", testName, gotResultInt64, tt.wantResultInt64)
				}
				if gotResultFloat64 != tt.wantResultFloat64 {
					t.Errorf("[%s] TryUnboxAnyAsInt64OrFloat64() gotResultFloat64 = %v, want %v", testName, gotResultFloat64, tt.wantResultFloat64)
				}
				if gotDataType != tt.wantResultDataType {
					t.Errorf("[%s] TryUnboxAnyAsInt64OrFloat64() gotDataType = %v, want %v", testName, gotDataType, tt.wantResultDataType)
				}
				if gotState != tt.wantState {
					t.Errorf("[%s] TryUnboxAnyAsInt64OrFloat64() gotState = %v, want %v", testName, gotState, tt.wantState)
				}
			}

			testName = "L0"
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotResultInt64, gotResultFloat64, gotDataType, gotState = TryUnboxAnyAsInt64OrFloat64(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

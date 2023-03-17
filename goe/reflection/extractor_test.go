package reflection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"unsafe"
)

func TestRootValueExtractor(t *testing.T) {
	type nilResultType int8
	var notNil nilResultType = -1
	var nilResult nilResultType = 0
	var nilValue nilResultType = 1

	type MyStruct struct {
		x string
	}

	tests := []struct {
		name         string
		input        any
		wantType     reflect.Type
		wantKind     reflect.Kind
		wantNil      nilResultType
		wantZero     bool
		wantOk       bool
		notSupported bool
	}{
		{
			name:    "nil",
			input:   nil,
			wantNil: nilResult,
			wantOk:  true,
		},
		{
			name: "nil interface",
			input: func() interface{} {
				var x interface{}
				return x
			}(),
			wantNil:  nilResult,
			wantZero: true,
			wantOk:   true,
		},
		{
			name: "non-nil interface",
			input: func() interface{} {
				var x interface{} = int64(200)
				return x
			}(),
			wantType: reflect.TypeOf(int64(0)),
			wantKind: reflect.Int64,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "int",
			input:    6,
			wantType: reflect.TypeOf(rand.Int()),
			wantKind: reflect.Int,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name: "*int with value",
			input: func() *int {
				i := 6
				return &i
			}(),
			wantType: reflect.TypeOf(rand.Int()),
			wantKind: reflect.Int,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name: "*int with zero value",
			input: func() *int {
				i := 0
				return &i
			}(),
			wantType: reflect.TypeOf(rand.Int()),
			wantKind: reflect.Int,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "string",
			input:    "6",
			wantType: reflect.TypeOf("6"),
			wantKind: reflect.String,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "empty",
			input:    "",
			wantType: reflect.TypeOf(""),
			wantKind: reflect.String,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name: "nil *string",
			input: func() *string {
				var ns *string
				return ns
			}(),
			wantType: reflect.TypeOf(new(string)),
			wantKind: reflect.Pointer,
			wantNil:  nilValue,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "empty slice",
			input:    []string{},
			wantType: reflect.TypeOf(*new([]string)),
			wantKind: reflect.Slice,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name: "nil slice",
			input: func() []string {
				var s []string
				return s
			}(),
			wantType: reflect.TypeOf(*new([]string)),
			wantKind: reflect.Slice,
			wantNil:  nilValue,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "non-empty slice",
			input:    []string{"1", "2"},
			wantType: reflect.TypeOf(*new([]string)),
			wantKind: reflect.Slice,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "non-empty array",
			input:    [2]string{},
			wantType: reflect.TypeOf(*new([2]string)),
			wantKind: reflect.Array,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "non-empty array",
			input:    [2]string{"1", "2"},
			wantType: reflect.TypeOf(*new([2]string)),
			wantKind: reflect.Array,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "type",
			input:    notNil,
			wantType: reflect.TypeOf(*new(nilResultType)),
			wantKind: reflect.Int8,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name: "struct",
			input: MyStruct{
				x: "a",
			},
			wantType: reflect.TypeOf(*new(MyStruct)),
			wantKind: reflect.Struct,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "empty struct",
			input:    MyStruct{},
			wantType: reflect.TypeOf(*new(MyStruct)),
			wantKind: reflect.Struct,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "function",
			input:    func(x, y int) int { return x + y },
			wantType: reflect.TypeOf(*new(func(x, y int) int)),
			wantKind: reflect.Func,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "channel",
			input:    make(chan int, 0),
			wantType: reflect.TypeOf(*new(chan int)),
			wantKind: reflect.Chan,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "map",
			input:    make(map[string]bool),
			wantType: reflect.TypeOf(*new(map[string]bool)),
			wantKind: reflect.Map,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "complex",
			input:    complex(1, 3),
			wantType: reflect.TypeOf(*new(complex128)),
			wantKind: reflect.Complex128,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "complex",
			input:    complex(0, 0),
			wantType: reflect.TypeOf(*new(complex128)),
			wantKind: reflect.Complex128,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name:     "bool",
			input:    true,
			wantType: reflect.TypeOf(false),
			wantKind: reflect.Bool,
			wantNil:  notNil,
			wantZero: false,
			wantOk:   true,
		},
		{
			name:     "bool",
			input:    false,
			wantType: reflect.TypeOf(false),
			wantKind: reflect.Bool,
			wantNil:  notNil,
			wantZero: true,
			wantOk:   true,
		},
		{
			name: "unsafe pointer",
			input: func() unsafe.Pointer {
				var f = 3.3
				return unsafe.Pointer(&f)
			}(),
			notSupported: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer deferExpectPanicContains(t, "not yet supported kind unsafe.Pointer", tt.notSupported)

			var gotValue *reflect.Value
			var isNil, gotOk bool
			var testName string

			assertResult := func() {
				fmt.Println(testName, "of", tt.name)
				assert.Equalf(t, tt.wantOk, gotOk, "[ok] expect %t, got %t", tt.wantOk, gotOk)

				if tt.wantNil == nilResult {
					assert.Nil(t, gotValue)
					assert.True(t, isNil)
				} else if tt.wantNil == nilValue {
					assert.NotNil(t, gotValue)
					assert.True(t, isNil)
					assert.True(t, gotValue.IsNil())
				} else {
					assert.NotNil(t, gotValue)
					assert.False(t, isNil)
					if canValueNil(gotValue) {
						assert.False(t, gotValue.IsNil())
					}
				}

				if gotValue != nil {
					assert.Equalf(t, tt.wantZero, gotValue.IsZero(), "[zero] expect %t, got %t", tt.wantZero, gotValue.IsZero())
					assert.Equalf(t, tt.wantType, gotValue.Type(), "[type] expect %v, got %v", tt.wantType, gotValue.Type())
					assert.Equalf(t, tt.wantKind, gotValue.Kind(), "[kind] expect %v, got %v", tt.wantKind, gotValue.Kind())

					if gotOk {
						if !gotValue.IsValid() {
							t.Errorf("got ok but value is not valid")
						}
					}
				}
			}

			testName = "L0"
			gotValue, isNil, gotOk = RootValueExtractor(tt.input)
			assertResult()

			testName = "L1 (pointer of input)"
			var pointerOfInput *any
			pointerOfInput = &tt.input
			gotValue, isNil, gotOk = RootValueExtractor(pointerOfInput)
			assertResult()

			testName = "L2-1 (any of pointer of input)"
			var anyOfPointerOfInput any
			anyOfPointerOfInput = any(pointerOfInput)
			gotValue, isNil, gotOk = RootValueExtractor(anyOfPointerOfInput)
			assertResult()

			testName = "L2-2 (pointer of pointer of input)"
			var pointerOfPointerOfInput **any
			pointerOfPointerOfInput = &pointerOfInput
			gotValue, isNil, gotOk = RootValueExtractor(pointerOfPointerOfInput)
			assertResult()

			testName = "L3-1-1 (pointer of any of pointer of input)"
			var pointerOfAnyOfPointerOfInput any
			pointerOfAnyOfPointerOfInput = &anyOfPointerOfInput
			gotValue, isNil, gotOk = RootValueExtractor(pointerOfAnyOfPointerOfInput)
			assertResult()

			testName = "L3-2-1 (any of pointer of pointer of input)"
			var anyOfPointerOfPointerOfInput any
			anyOfPointerOfPointerOfInput = any(pointerOfPointerOfInput)
			gotValue, isNil, gotOk = RootValueExtractor(anyOfPointerOfPointerOfInput)
			assertResult()

			testName = "L3-2-2 (pointer of pointer of pointer of input)"
			var pointerOfPointerOfPointerOfInput ***any
			pointerOfPointerOfPointerOfInput = &pointerOfPointerOfInput
			gotValue, isNil, gotOk = RootValueExtractor(pointerOfPointerOfPointerOfInput)
			assertResult()
		})
	}
}

func Test_canValueNil(t *testing.T) {
	t.Run("accept struct & pointer Value, not other", func(t *testing.T) {
		vo := reflect.ValueOf([]int{})
		assert.True(t, canValueNil(vo))
		assert.True(t, canValueNil(&vo))

		defer deferExpectPanicContains(t, "interface conversion", true)
		_ = canValueNil(1)
	})
}

func Test_panicMaxRecursiveLoopReached(t *testing.T) {
	for i := 0; i <= maxRecursiveLevelForRootValueExtractor+10; i++ {
		t.Run("max recursive extractor", func(t *testing.T) {
			defer deferExpectPanicContains(t, "max recursive reached", i >= maxRecursiveLevelForRootValueExtractor)
			panicMaxRecursiveLoopReached(i)
		})
	}
}

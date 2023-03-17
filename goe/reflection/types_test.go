package reflection

import (
	"fmt"
	"testing"
)

func TestUnboxResultState_AssertIsSuccess(t *testing.T) {
	tests := []struct {
		state     UnboxResultState
		wantPanic bool
	}{
		{
			state:     UNBOX_FAILED,
			wantPanic: true,
		},
		{
			state:     UNBOX_NIL,
			wantPanic: true,
		},
		{
			state:     UNBOX_OVERFLOW,
			wantPanic: true,
		},
		{
			state:     UNBOX_SUCCESS,
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.state), func(t *testing.T) {
			defer deferExpectPanicContains(t, "require success", tt.wantPanic)
			tt.state.AssertIsSuccess()
		})
	}
}

func TestUnboxFloat64DataType_AssertHasNonNilResult(t *testing.T) {
	tests := []struct {
		dataType  UnboxFloat64DataType
		wantPanic bool
	}{
		{
			dataType:  UF64_TYPE_FAILED,
			wantPanic: true,
		},
		{
			dataType:  UF64_TYPE_NIL,
			wantPanic: true,
		},
		{
			dataType:  UF64_TYPE_INT64,
			wantPanic: false,
		},
		{
			dataType:  UF64_TYPE_FLOAT64,
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.dataType), func(t *testing.T) {
			defer deferExpectPanicContains(t, "require non-nil int64 or float", tt.wantPanic)
			tt.dataType.AssertHasNonNilResult()
		})
	}
}

package goe_helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func deferExpectPanicContains(t *testing.T, msgPart string, wantPanic bool) {
	if len(msgPart) < 1 {
		t.Errorf("empty msg part was passed")
	}

	err := recover()

	if wantPanic {
		if err == nil {
			t.Errorf("expect error")
			return
		}

		assert.Contains(t, fmt.Sprintf("%v", err), msgPart)
	} else {
		if err != nil {
			t.Errorf("not expect error, but got: %v", err)
			return
		}
	}
}

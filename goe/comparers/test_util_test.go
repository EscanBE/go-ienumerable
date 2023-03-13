package comparers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func deferExpectPanicContains(t *testing.T, msgPart string) {
	err := recover()
	if err == nil {
		t.Errorf("expect error")
		return
	}

	assert.Contains(t, fmt.Sprintf("%v", err), msgPart)
}

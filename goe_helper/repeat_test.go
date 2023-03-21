package goe_helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("non-empty", func(t *testing.T) {
		e := Repeat(9, 4)

		data := e.ToArray()
		assert.Len(t, data, 4)
		assert.Equal(t, 9, data[0])
		assert.Equal(t, 9, data[1])
		assert.Equal(t, 9, data[2])
		assert.Equal(t, 9, data[3])
	})

	t.Run("empty", func(t *testing.T) {
		e := Repeat(9, 0)
		assert.Empty(t, e.ToArray())
	})

	t.Run("panic invalid size", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "count is less than 0")
		}()
		_ = Repeat(9, -1)
	})
}

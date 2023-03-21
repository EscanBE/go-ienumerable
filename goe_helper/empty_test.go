package goe_helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Run("create empty", func(t *testing.T) {
		e := Empty[int]()

		assert.Empty(t, e.ToArray())

		defaultValue := e.FirstOrDefault(nil, nil)
		assert.Zero(t, defaultValue)
		assert.Equal(t, "int", fmt.Sprintf("%T", defaultValue))
	})
	t.Run("create empty", func(t *testing.T) {
		e := Empty[string]()

		assert.Empty(t, e.ToArray())

		defaultValue := e.FirstOrDefault(nil, nil)
		assert.Equal(t, "", defaultValue)
		assert.Equal(t, "string", fmt.Sprintf("%T", defaultValue))
	})
}

package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnumerableHelper_Empty(t *testing.T) {
	Enumerable := Enumerable[int]()
	e := Enumerable.Empty()
	assert.Empty(t, e.ToArray())

	defaultValue := e.FirstOrDefault(nil, nil)
	assert.Zero(t, defaultValue)
	assert.Equal(t, "int", fmt.Sprintf("%T", defaultValue))
}

func TestEnumerableHelper_Repeat(t *testing.T) {
	Enumerable := Enumerable[int]()
	e := Enumerable.Repeat(9, 4)

	data := e.ToArray()
	assert.Len(t, data, 4)
	assert.Equal(t, 9, data[0])
	assert.Equal(t, 9, data[1])
	assert.Equal(t, 9, data[2])
	assert.Equal(t, 9, data[3])

	e = Enumerable.Repeat(9, 0)
	assert.Empty(t, e.ToArray())

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("expect error")
			return
		}
		assert.Contains(t, fmt.Sprintf("%v", err), "count is less than 0")
	}()
	_ = Enumerable.Repeat(9, -1)
}

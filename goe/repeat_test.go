package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Repeat(t *testing.T) {
	eSrc := NewIEnumerable[int](1, 2, 3, 4, 5, 6, 7, 8).WithDefaultComparers()
	bSrc := backupForAssetUnchanged(eSrc)

	eGot := eSrc.Repeat(9, 4)

	data := eGot.ToArray()
	assert.Len(t, data, 4)
	assert.Equal(t, 9, data[0])
	assert.Equal(t, 9, data[1])
	assert.Equal(t, 9, data[2])
	assert.Equal(t, 9, data[3])

	bSrc.assertUnchanged(t, eSrc)
	bSrc.assertUnchangedIgnoreData(t, eGot)

	eGot = eSrc.Repeat(9, 0)
	assert.Empty(t, eGot.ToArray())

	bSrc.assertUnchanged(t, eSrc)
	bSrc.assertUnchangedIgnoreData(t, eGot)

	defer func() {
		bSrc.assertUnchanged(t, eSrc)
	}()

	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("expect error")
			return
		}
		assert.Contains(t, fmt.Sprintf("%v", err), "count is less than 0")
	}()
	_ = eSrc.Repeat(9, -1)
}

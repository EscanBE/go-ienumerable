package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Empty(t *testing.T) {
	srcS := NewIEnumerable[string]("hello", "world")
	bSrc := backupForAssetUnchanged(srcS)

	gotS := srcS.Empty()
	eGot := e[string](gotS)
	assert.Zero(t, gotS.Count(nil))
	assert.Equal(t, "string", eGot.dataType)

	bSrc.assertUnchanged(t, srcS)
	bSrc.assertUnchangedIgnoreData(t, gotS)

	gotI := NewIEnumerable[int](99, 999).Empty()
	eGotI := e[int](gotI)
	assert.Zero(t, gotI.Count(nil))
	assert.Equal(t, "int", eGotI.dataType)

	gotA := NewIEnumerable[any](99, 999, "string", 0.0).Empty()
	eGotA := e[any](gotA)
	assert.Zero(t, gotA.Count(nil))
	assert.Equal(t, "", eGotA.dataType)
}

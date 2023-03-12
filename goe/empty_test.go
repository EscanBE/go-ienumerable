package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Empty(t *testing.T) {
	srcS := NewIEnumerable[string]("hello", "world").WithDefaultComparers()
	bSrc := backupForAssetUnchanged(srcS)
	gotS := srcS.Empty()
	assert.Zero(t, gotS.len())
	assert.Equal(t, "string", gotS.exposeDataType())
	bSrc.assertUnchanged(t, srcS)
	bSrc.assertUnchangedIgnoreData(t, gotS)

	gotI := NewIEnumerable[int](99, 999).Empty()
	assert.Zero(t, gotI.len())
	assert.Equal(t, "int", gotI.exposeDataType())

	gotA := NewIEnumerable[any](99, 999, "string", 0.0).Empty()
	assert.Zero(t, gotA.len())
	assert.Equal(t, "", gotA.exposeDataType())
}

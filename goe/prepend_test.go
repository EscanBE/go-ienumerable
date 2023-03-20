package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Prepend(t *testing.T) {
	tests := []struct {
		name    string
		src     IEnumerable[int]
		element int
		want    IEnumerable[int]
	}{
		{
			name:    "empty",
			src:     createEmptyIntEnumerable(),
			element: 1,
			want:    NewIEnumerable[int](1),
		},
		{
			name:    "prepend",
			src:     createIntEnumerable(91, 99),
			element: 90,
			want:    createIntEnumerable(90, 99),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.src)

			eGot := tt.src.Prepend(tt.element)
			gotData := eGot.ToArray()
			assert.Len(t, gotData, tt.src.Count(nil)+1)
			assert.Equal(t, tt.element, gotData[0])

			bSrc.assertUnchanged(t, tt.src)
			bSrc.assertUnchangedIgnoreData(t, eGot)
		})
	}

	t.Run("details", func(t *testing.T) {
		eSrc := createIntEnumerable(61, 64)
		bSrc := backupForAssetUnchanged(eSrc)

		eGot := eSrc.Prepend(60)
		gotData := eGot.ToArray()
		assert.Len(t, gotData, eSrc.Count(nil)+1)
		assert.Equal(t, 60, gotData[0])
		assert.Equal(t, 61, gotData[1])
		assert.Equal(t, 62, gotData[2])
		assert.Equal(t, 63, gotData[3])
		assert.Equal(t, 64, gotData[4])

		bSrc.assertUnchanged(t, eSrc)
		bSrc.assertUnchangedIgnoreData(t, eGot)
	})
}

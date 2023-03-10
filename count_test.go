package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Count_CountBy(t *testing.T) {
	tests := []struct {
		name        string
		src         IEnumerable[any]
		wantCount   int
		countBy     func(any any) bool
		wantCountBy int
	}{
		{
			name:      "nil",
			src:       createNilEnumerable(),
			wantCount: 0,
			countBy: func(any any) bool {
				return true
			},
			wantCountBy: 0,
		},
		{
			name:      "empty",
			src:       createEmptyEnumerable(),
			wantCount: 0,
			countBy: func(any any) bool {
				return true
			},
			wantCountBy: 0,
		},
		{
			name:      "with element",
			src:       NewIEnumerable[any](1, 2, 3, 4, 5, 6, 7, 8, 9),
			wantCount: 9,
			countBy: func(any any) bool {
				return any.(int) >= 6
			},
			wantCountBy: 4,
		},
		{
			name:      "with element",
			src:       NewIEnumerable[any](1, 2, 3, 4, 5, 6, 7, 8),
			wantCount: 8,
			countBy: func(any any) bool {
				return any.(int) >= 10
			},
			wantCountBy: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantCount, tt.src.Count(), "Count()")
			assert.Equalf(t, tt.wantCountBy, tt.src.CountBy(tt.countBy), "CountBy()")
		})
	}

	i1 := createRandomIntEnumerable(5)
	bi1 := backupForAssetUnchanged(i1)
	assert.Equalf(t, 5, i1.Count(), "Count()")
	assert.Equalf(t, 5, i1.CountBy(func(i int) bool {
		return i >= 0
	}), "CountBy()")
	bi1.assertUnchanged(t, i1)
}

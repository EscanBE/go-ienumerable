package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_Count(t *testing.T) {
	tests := []struct {
		name        string
		src         IEnumerable[any]
		wantCount   int
		countBy     OptionalPredicate[any]
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
			assert.Equalf(t, tt.wantCount, tt.src.Count(nil), "Count(nil)")
			assert.Equalf(t, tt.wantCountBy, tt.src.Count(tt.countBy), "Count(predicate)")
		})
	}
}

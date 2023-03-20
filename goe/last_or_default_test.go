package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_LastOrDefault(t *testing.T) {
	var defaultValue = 99

	tests := []struct {
		name         string
		src          IEnumerable[int]
		predicate    OptionalPredicate[int]
		defaultValue *int
		wantResult   int
	}{
		{
			name:       "last",
			src:        createIntEnumerable(5, 7),
			predicate:  nil,
			wantResult: 7,
		},
		{
			name:         "not any without predicate",
			src:          createEmptyIntEnumerable(),
			predicate:    nil,
			defaultValue: nil,
			wantResult:   0,
		},
		{
			name:         "not any without predicate",
			src:          createEmptyIntEnumerable(),
			predicate:    nil,
			defaultValue: &defaultValue,
			wantResult:   defaultValue,
		},
		{
			name: "last",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			wantResult: 7,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: nil,
			wantResult:   0,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantResult: 0,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: &defaultValue,
			wantResult:   defaultValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult := tt.src.LastOrDefault(tt.predicate, tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}

	t.Run("default string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("hello", "world")
		bSrc := backupForAssetUnchanged(eSrc)

		assert.Equal(t, "world", eSrc.LastOrDefault(nil, nil))

		var predicate OptionalPredicate[string] = func(s string) bool {
			return len(s) < 3
		}
		assert.Equal(t, "", eSrc.LastOrDefault(predicate, nil))

		var defaultValue string = "default"
		assert.Equal(t, "default", eSrc.LastOrDefault(predicate, &defaultValue))

		bSrc.assertUnchanged(t, eSrc)
	})
}

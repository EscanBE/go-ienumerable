package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_LastOrDefaultSafeBy_LastOrDefaultBy(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		predicate    func(int) bool
		defaultValue int
		wantResult   int
		wantErr      bool
	}{
		{
			name:      "nil predicate",
			src:       createRandomIntEnumerable(3),
			predicate: nil,
			wantErr:   true,
		},
		{
			name: "last",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			defaultValue: 999,
			wantResult:   7,
			wantErr:      false,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: 1,
			wantResult:   1,
			wantErr:      false,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: 9,
			wantResult:   9,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult, err := tt.src.LastOrDefaultSafeBy(tt.predicate, tt.defaultValue)
			assert.Equalf(t, tt.wantErr, err != nil, "error state is different, expect err %t, got %t", tt.wantErr, err != nil)
			if err == nil && !tt.wantErr {
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}

			backSrc.assertUnchanged(t, tt.src)

			if tt.wantErr {
				defer deferWantPanicDepends(t, true)
				_ = tt.src.LastOrDefaultBy(tt.predicate, tt.defaultValue)
			} else if err == nil {
				gotResult = tt.src.LastOrDefaultBy(tt.predicate, tt.defaultValue)
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}
		})
	}
}

func Test_enumerable_LastOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		defaultValue int
		wantResult   int
		wantErr      bool
	}{
		{
			name:         "last",
			src:          createIntEnumerable(5, 7),
			defaultValue: 99,
			wantResult:   7,
			wantErr:      false,
		},
		{
			name:         "not any",
			src:          createEmptyIntEnumerable(),
			defaultValue: 99,
			wantResult:   99,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult := tt.src.LastOrDefault(tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)

			backSrc.assertUnchanged(t, tt.src)
		})
	}
}

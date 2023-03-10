package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_LastSafeBy_LastBy(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		predicate  func(int) bool
		wantResult int
		wantErr    bool
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
			wantResult: 7,
			wantErr:    false,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantErr: true,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult, err := tt.src.LastSafeBy(tt.predicate)
			assert.Equalf(t, tt.wantErr, err != nil, "error state is different, expect err %t, got %t", tt.wantErr, err != nil)
			if err == nil && !tt.wantErr {
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}

			backSrc.assertUnchanged(t, tt.src)

			if tt.wantErr {
				defer deferWantPanicDepends(t, true)
				_ = tt.src.LastBy(tt.predicate)
			} else if err == nil {
				gotResult = tt.src.LastBy(tt.predicate)
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}
		})
	}
}

func Test_enumerable_LastSafe_Last(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		wantResult int
		wantErr    bool
	}{
		{
			name:       "last",
			src:        createIntEnumerable(5, 7),
			wantResult: 7,
			wantErr:    false,
		},
		{
			name:    "not any",
			src:     createEmptyIntEnumerable(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult, err := tt.src.LastSafe()
			assert.Equalf(t, tt.wantErr, err != nil, "error state is different, expect err %t, got %t", tt.wantErr, err != nil)
			if err == nil && !tt.wantErr {
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}

			backSrc.assertUnchanged(t, tt.src)

			if tt.wantErr {
				defer deferWantPanicDepends(t, true)
				_ = tt.src.Last()
			} else if err == nil {
				gotResult = tt.src.Last()
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}
		})
	}
}

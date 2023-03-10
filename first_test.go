package go_ienumerable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_FirstSafeBy_FirstBy(t *testing.T) {
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
			name: "first",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			wantResult: 6,
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

			gotResult, err := tt.src.FirstSafeBy(tt.predicate)
			assert.Equalf(t, tt.wantErr, err != nil, "error state is different, expect err %t, got %t", tt.wantErr, err != nil)
			if err == nil && !tt.wantErr {
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}

			backSrc.assertUnchanged(t, tt.src)

			if tt.wantErr {
				defer deferWantPanicDepends(t, true)
				_ = tt.src.FirstBy(tt.predicate)
			} else if err == nil {
				gotResult = tt.src.FirstBy(tt.predicate)
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}
		})
	}
}

func Test_enumerable_FirstSafe_First(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		wantResult int
		wantErr    bool
	}{
		{
			name:       "first",
			src:        createIntEnumerable(5, 7),
			wantResult: 5,
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

			gotResult, err := tt.src.FirstSafe()
			assert.Equalf(t, tt.wantErr, err != nil, "error state is different, expect err %t, got %t", tt.wantErr, err != nil)
			if err == nil && !tt.wantErr {
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}

			backSrc.assertUnchanged(t, tt.src)

			if tt.wantErr {
				defer deferWantPanicDepends(t, true)
				_ = tt.src.First()
			} else if err == nil {
				gotResult = tt.src.First()
				assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
			}
		})
	}
}

package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_LastBy(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		predicate  func(int) bool
		wantResult int
		wantPanic  bool
	}{
		{
			name:      "nil predicate",
			src:       createRandomIntEnumerable(3),
			predicate: nil,
			wantPanic: true,
		},
		{
			name: "last",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			wantResult: 7,
			wantPanic:  false,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantPanic: true,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			defer deferWantPanicDepends(t, tt.wantPanic)

			gotResult := tt.src.LastBy(tt.predicate)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}
}

func Test_enumerable_Last(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		wantResult int
		wantPanic  bool
	}{
		{
			name:       "last",
			src:        createIntEnumerable(5, 7),
			wantResult: 7,
			wantPanic:  false,
		},
		{
			name:      "not any",
			src:       createEmptyIntEnumerable(),
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			defer deferWantPanicDepends(t, tt.wantPanic)

			gotResult := tt.src.Last()
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}
}

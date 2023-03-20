package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_FirstOrDefault(t *testing.T) {
	tests := []struct {
		name       string
		src        IEnumerable[int]
		predicate  OptionalPredicate[int]
		wantResult int
	}{
		{
			name:       "first",
			src:        createIntEnumerable(5, 7),
			predicate:  nil,
			wantResult: 5,
		},
		{
			name:       "not any",
			src:        createEmptyIntEnumerable(),
			predicate:  nil,
			wantResult: 0,
		},
		{
			name: "first",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			wantResult: 6,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
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
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult := tt.src.FirstOrDefault(tt.predicate)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}

	t.Run("default string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("hello", "world")
		bSrc := backupForAssetUnchanged(eSrc)

		assert.Equal(t, "hello", eSrc.FirstOrDefault(nil))

		var predicate OptionalPredicate[string] = func(s string) bool {
			return len(s) < 3
		}
		assert.Equal(t, "", eSrc.FirstOrDefault(predicate))

		bSrc.assertUnchanged(t, eSrc)
	})
}

func Test_enumerable_FirstOrDefaultUsing(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		defaultValue int
		wantResult   int
		wantErr      bool
	}{
		{
			name:         "first",
			src:          createIntEnumerable(5, 7),
			defaultValue: 99,
			wantResult:   5,
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

			gotResult := tt.src.FirstOrDefaultUsing(tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)

			backSrc.assertUnchanged(t, tt.src)
		})
	}
}

func Test_enumerable_FirstOrDefaultByUsing(t *testing.T) {
	tests := []struct {
		name         string
		src          IEnumerable[int]
		predicate    func(int) bool
		defaultValue int
		wantResult   int
		wantPanic    bool
	}{
		{
			name:      "nil predicate",
			src:       createRandomIntEnumerable(3),
			predicate: nil,
			wantPanic: true,
		},
		{
			name: "first",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			defaultValue: 999,
			wantResult:   6,
			wantPanic:    false,
		},
		{
			name: "not any match",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: 1,
			wantResult:   1,
			wantPanic:    false,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			defaultValue: 9,
			wantResult:   9,
			wantPanic:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			defer deferWantPanicDepends(t, tt.wantPanic)

			gotResult := tt.src.FirstOrDefaultByUsing(tt.predicate, tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}
}

package goe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_enumerable_LastOrDefault(t *testing.T) {
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
			name:       "not any",
			src:        createEmptyIntEnumerable(),
			wantResult: 0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			gotResult := tt.src.LastOrDefault()
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)

			backSrc.assertUnchanged(t, tt.src)
		})
	}

	t.Run("default string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]()
		bSrc := backupForAssetUnchanged(eSrc)

		assert.Equal(t, "", eSrc.LastOrDefault())

		bSrc.assertUnchanged(t, eSrc)
	})
}

func Test_enumerable_LastOrDefaultBy(t *testing.T) {
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
			wantResult: 0,
			wantPanic:  false,
		},
		{
			name: "sequence contains no element",
			src:  createEmptyIntEnumerable(),
			predicate: func(i int) bool {
				return i >= 8
			},
			wantResult: 0,
			wantPanic:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backSrc := backupForAssetUnchanged(tt.src)

			defer func() {
				backSrc.assertUnchanged(t, tt.src)
			}()

			defer deferWantPanicDepends(t, tt.wantPanic)

			gotResult := tt.src.LastOrDefaultBy(tt.predicate)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}

	t.Run("default string", func(t *testing.T) {
		eSrc := NewIEnumerable[string]("hello", "world")
		bSrc := backupForAssetUnchanged(eSrc)

		assert.Equal(t, "", eSrc.LastOrDefaultBy(func(s string) bool {
			return len(s) < 3
		}))

		bSrc.assertUnchanged(t, eSrc)
	})
}

func Test_enumerable_LastOrDefaultUsing(t *testing.T) {
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

			gotResult := tt.src.LastOrDefaultUsing(tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)

			backSrc.assertUnchanged(t, tt.src)
		})
	}
}

func Test_enumerable_LastOrDefaultByUsing(t *testing.T) {
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
			name: "last",
			src:  createIntEnumerable(5, 7),
			predicate: func(i int) bool {
				return i >= 6
			},
			defaultValue: 999,
			wantResult:   7,
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

			gotResult := tt.src.LastOrDefaultByUsing(tt.predicate, tt.defaultValue)
			assert.Equalf(t, tt.wantResult, gotResult, "expected result %d, got %d", tt.wantResult, gotResult)
		})
	}
}

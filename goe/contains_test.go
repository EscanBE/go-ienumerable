package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_enumerable_Contains(t *testing.T) {
	t.Run("returns correctly", func(t *testing.T) {
		radCap := rand.Intn(10) + 10
		eSrc := createIntEnumerable(0, radCap)

		nContains := rand.Intn(radCap)
		nNotContains := rand.Intn(radCap) + radCap + 1

		assert.True(t, eSrc.Contains(nContains))
		assert.False(t, eSrc.Contains(nNotContains))

		eSrcP := eSrc.Select(func(v int) any {
			return &v
		})

		assert.True(t, eSrcP.Contains(&nContains))
		assert.False(t, eSrcP.Contains(&nNotContains))
	})

	t.Run("empty returns false", func(t *testing.T) {
		assert.False(t, createEmptyIntEnumerable().Contains(1))
	})

	t.Run("retry resolve if comparer not set", func(t *testing.T) {
		eSrc := createEmptyIntEnumerable()
		e[int](eSrc).defaultComparer = nil
		assert.False(t, eSrc.Contains(1))
	})

	t.Run("panic if type not registered for default comparer", func(t *testing.T) {
		type MyInt64 struct{}

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]", true)

		NewIEnumerable[MyInt64]().Contains(MyInt64{})
	})
}

func Test_enumerable_ContainsBy(t *testing.T) {
	fEquals := func(t1, t2 int) bool {
		return t1 == t2
	}
	var tests = []struct {
		name    string
		source  IEnumerable[int]
		check   int
		fEquals func(t1, t2 int) bool
		want    bool
	}{
		{
			name:    "empty source",
			source:  createEmptyIntEnumerable(),
			fEquals: fEquals,
			want:    false,
		},
		{
			name:    "single",
			source:  NewIEnumerable[int](2),
			check:   1,
			fEquals: fEquals,
			want:    false,
		},
		{
			name:    "single",
			source:  NewIEnumerable[int](2),
			check:   2,
			fEquals: fEquals,
			want:    true,
		},
		{
			name:    "no equality comparer still ok since int has default comparer",
			source:  NewIEnumerable[int](2),
			check:   2,
			fEquals: nil,
			want:    true,
		},
		{
			name:    "many",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:   3,
			fEquals: nil,
			want:    true,
		},
		{
			name:    "many",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:   99,
			fEquals: fEquals,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.ContainsBy(tt.check, tt.fEquals)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4)
		eSrc := e[int](ieSrc)
		eSrc.defaultComparer = nil

		assert.True(t, ieSrc.ContainsBy(3, nil))
	})

	t.Run("panic if no default resolver", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.ContainsBy(MyInt64{}, nil)
	})
}

func Test_enumerable_ContainsByComparer(t *testing.T) {
	var tests = []struct {
		name     string
		source   IEnumerable[int]
		check    int
		comparer comparers.IComparer[int]
		want     bool
	}{
		{
			name:     "empty source",
			source:   createEmptyIntEnumerable(),
			comparer: comparers.IntComparer,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    1,
			comparer: comparers.IntComparer,
			want:     false,
		},
		{
			name:     "single",
			source:   NewIEnumerable[int](2),
			check:    2,
			comparer: comparers.IntComparer,
			want:     true,
		},
		{
			name:     "no equality comparer still ok since int has default comparer",
			source:   NewIEnumerable[int](2),
			check:    2,
			comparer: nil,
			want:     true,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    3,
			comparer: nil,
			want:     true,
		},
		{
			name:     "many",
			source:   NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:    99,
			comparer: comparers.IntComparer,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSrc := backupForAssetUnchanged(tt.source)

			got := tt.source.ContainsByComparer(tt.check, tt.comparer)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)
		})
	}

	t.Run("auto-resolve comparer if default comparer not set", func(t *testing.T) {
		ieSrc := NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4)
		eSrc := e[int](ieSrc)
		eSrc.defaultComparer = nil

		assert.True(t, ieSrc.ContainsByComparer(3, nil))
	})

	t.Run("panic if no default resolver", func(t *testing.T) {
		type MyInt64 struct{}
		ieSrc := NewIEnumerable[MyInt64]()

		defer deferExpectPanicContains(t, "no default comparer registered", true)

		ieSrc.ContainsByComparer(MyInt64{}, nil)
	})
}

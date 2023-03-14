package goe

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_enumerable_Contains(t *testing.T) {
	t.Run("returns correctly", func(t *testing.T) {
		cap := rand.Intn(10) + 10
		eSrc := createIntEnumerable(0, cap)

		nContains := rand.Intn(cap)
		nNotContains := rand.Intn(cap) + cap + 1

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

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyInt64]")

		NewIEnumerable[MyInt64]().Contains(MyInt64{})
	})
}

func Test_enumerable_Contains_ContainsBy(t *testing.T) {
	fEquals := func(t1, t2 int) bool {
		return t1 == t2
	}
	var tests = []struct {
		name      string
		source    IEnumerable[int]
		check     int
		fEquals   func(t1, t2 int) bool
		want      bool
		wantPanic bool
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
			name:      "panic due to no equality comparer",
			source:    NewIEnumerable[int](2),
			check:     2,
			fEquals:   nil,
			wantPanic: true,
		},
		{
			name:    "many",
			source:  NewIEnumerable[int](1, 2, 2, 3, 3, 6, 6, 6, 5, 4, 4),
			check:   3,
			fEquals: fEquals,
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

			defer deferWantPanicDepends(t, tt.wantPanic)

			got := tt.source.ContainsBy(tt.check, tt.fEquals)

			assert.Equal(t, tt.want, got)

			bSrc.assertUnchanged(t, tt.source)

		})
	}

	t.Run("use default comparer", func(t *testing.T) {
		assert.True(t, NewIEnumerable[int](1, 5, 2, 345, 65, 4574, 1).Contains(5))
	})
}

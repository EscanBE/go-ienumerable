package goe

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

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

	t.Run("equality comparer not set", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)

		NewIEnumerable[int](1, 5, 2, 345, 65, 4574).Contains(5)
	})

	t.Run("equality comparer set", func(t *testing.T) {
		assert.True(t, NewIEnumerable[int](1, 5, 2, 345, 65, 4574, 1).WithDefaultComparers().Contains(5))
	})
}

func Test_enumerable_Contains2(t *testing.T) {
	cap := rand.Intn(10) + 10
	eSrc := createIntEnumerable(0, cap)

	c1 := rand.Intn(cap)
	c2 := rand.Intn(cap) + cap

	assert.True(t, eSrc.Contains2(c1))
	assert.False(t, eSrc.Contains2(c2))

	eSrcP := eSrc.Select(func(v int) any {
		return &v
	})

	assert.True(t, eSrcP.Contains2(&c1))
	assert.False(t, eSrcP.Contains2(&c2))
}

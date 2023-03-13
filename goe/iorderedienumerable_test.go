package goe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func Test_IOrderedIEnumerable(t *testing.T) {
	type s3 struct {
		value int
	}
	type s2 struct {
		nested s3
		value  int
	}
	type s1 struct {
		nested s2
		value  int
	}

	v2430 := s1{
		value: 2000,
		nested: s2{
			value: 400,
			nested: s3{
				value: 30,
			},
		},
	}

	v1530 := s1{
		value: 1000,
		nested: s2{
			value: 500,
			nested: s3{
				value: 30,
			},
		},
	}

	v3530 := s1{
		value: 3000,
		nested: s2{
			value: 500,
			nested: s3{
				value: 30,
			},
		},
	}

	v4530 := s1{
		value: 4000,
		nested: s2{
			value: 500,
			nested: s3{
				value: 30,
			},
		},
	}

	v2420 := s1{
		value: 2000,
		nested: s2{
			value: 400,
			nested: s3{
				value: 20,
			},
		},
	}

	v2160 := s1{
		value: 2000,
		nested: s2{
			value: 100,
			nested: s3{
				value: 60,
			},
		},
	}

	v3990 := s1{
		value: 3000,
		nested: s2{
			value: 900,
			nested: s3{
				value: 90,
			},
		},
	}

	eSrc := NewIEnumerable[s1](v2430, v1530, v3530, v4530, v2420, v2160, v3990)
	bSrc := backupForAssetUnchanged(eSrc)

	comparatorLevel1 := func(v1, v2 s1) int {
		return v1.value - v2.value
	}

	comparatorLevel2 := func(v1, v2 s1) int {
		return v1.nested.value - v2.nested.value
	}

	comparatorLevel3 := func(v1, v2 s1) int {
		return v1.nested.nested.value - v2.nested.nested.value
	}

	tests := []struct {
		name    string
		ordered IOrderedEnumerable[s1]
		want    IEnumerable[s1]
	}{
		{
			name:    "asc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenBy(comparatorLevel2).ThenBy(comparatorLevel3),
			want:    NewIEnumerable[s1](v1530, v2160, v2420, v2430, v3530, v3990, v4530),
		},
		{
			name:    "asc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenBy(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[s1](v1530, v2160, v2430, v2420, v3530, v3990, v4530),
		},
		{
			name:    "asc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenByDescending(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[s1](v1530, v2430, v2420, v2160, v3990, v3530, v4530),
		},
		{
			name:    "desc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenByDescending(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[s1](v4530, v3990, v3530, v2430, v2420, v2160, v1530),
		},
		{
			name:    "desc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenBy(comparatorLevel2).ThenBy(comparatorLevel3),
			want:    NewIEnumerable[s1](v4530, v3530, v3990, v2160, v2420, v2430, v1530),
		},
		{
			name:    "desc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenBy(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[s1](v4530, v3530, v3990, v2160, v2430, v2420, v1530),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eGot := tt.ordered.GetEnumerable()

			assert.Truef(t, reflect.DeepEqual(tt.want.exposeData(), eGot.exposeData()), "got %v, want %v", eGot.exposeData(), tt.want.exposeData())

			bSrc.assertUnchanged(t, eSrc)
			bSrc.assertUnchangedIgnoreData(t, eGot)
		})
	}
}

func Test_IOrderedIEnumerable2(t *testing.T) {
	eSrc := NewIEnumerable[string]("v2430", "v1530", "v3530", "v4530", "v2420", "v2160", "v3990")
	bSrc := backupForAssetUnchanged(eSrc)

	comparatorLevel1 := func(l, r string) int {
		return strings.Compare(string(l[1]), string(r[1]))
	}

	comparatorLevel2 := func(l, r string) int {
		return strings.Compare(string(l[2]), string(r[2]))
	}

	comparatorLevel3 := func(l, r string) int {
		return strings.Compare(string(l[3]), string(r[3]))
	}

	tests := []struct {
		name    string
		ordered IOrderedEnumerable[string]
		want    IEnumerable[string]
	}{
		{
			name:    "asc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenBy(comparatorLevel2).ThenBy(comparatorLevel3),
			want:    NewIEnumerable[string]("v1530", "v2160", "v2420", "v2430", "v3530", "v3990", "v4530"),
		},
		{
			name:    "asc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenBy(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[string]("v1530", "v2160", "v2430", "v2420", "v3530", "v3990", "v4530"),
		},
		{
			name:    "asc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_ASC).ThenByDescending(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[string]("v1530", "v2430", "v2420", "v2160", "v3990", "v3530", "v4530"),
		},
		{
			name:    "desc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenByDescending(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[string]("v4530", "v3990", "v3530", "v2430", "v2420", "v2160", "v1530"),
		},
		{
			name:    "desc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenBy(comparatorLevel2).ThenBy(comparatorLevel3),
			want:    NewIEnumerable[string]("v4530", "v3530", "v3990", "v2160", "v2420", "v2430", "v1530"),
		},
		{
			name:    "desc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, comparatorLevel1, CLC_DESC).ThenBy(comparatorLevel2).ThenByDescending(comparatorLevel3),
			want:    NewIEnumerable[string]("v4530", "v3530", "v3990", "v2160", "v2430", "v2420", "v1530"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eGot := tt.ordered.GetEnumerable()

			assert.Truef(t, reflect.DeepEqual(tt.want.exposeData(), eGot.exposeData()), "got %v, want %v", eGot.exposeData(), tt.want.exposeData())

			bSrc.assertUnchanged(t, eSrc)
			bSrc.assertUnchangedIgnoreData(t, eGot)
		})
	}
}

func Test_IOrderedIEnumerable3(t *testing.T) {
	t.Run("previous IOrderedComparer not changes when chaining more", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(3)
		bSrc := backupForAssetUnchanged(eSrc)

		oe0 := newIOrderedEnumerable(eSrc, func(v1, v2 int) int {
			return 0
		}, CLC_ASC)

		_oe0 := oe0.(*orderedEnumerable[int])
		_oe1 := _oe0.ThenByDescending(func(v1, v2 int) int {
			return 0
		}).(*orderedEnumerable[int])
		_oe2 := _oe1.ThenBy(func(v1, v2 int) int {
			return 0
		}).(*orderedEnumerable[int])

		assert.Len(t, _oe0.chainableComparers, 1)
		assert.Len(t, _oe1.chainableComparers, 2)
		assert.Len(t, _oe2.chainableComparers, 3)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("previous IOrderedComparer not changes when chaining more and comparers copied to new", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(3)
		bSrc := backupForAssetUnchanged(eSrc)

		oe0 := newIOrderedEnumerable(eSrc, func(v1, v2 int) int {
			return 0
		}, CLC_ASC)

		_oe0 := oe0.(*orderedEnumerable[int])
		_oe1 := _oe0.ThenByDescending(func(v1, v2 int) int {
			return 0
		}).(*orderedEnumerable[int])
		_oe2 := _oe1.ThenBy(func(v1, v2 int) int {
			return 0
		}).(*orderedEnumerable[int])

		assert.Len(t, _oe2.chainableComparers, 3)

		assert.Equal(t, CLC_ASC, _oe0.chainableComparers[0].orderType)
		assert.Equal(t, CLC_ASC, _oe1.chainableComparers[0].orderType)
		assert.Equal(t, CLC_DESC, _oe1.chainableComparers[1].orderType)
		assert.Equal(t, CLC_ASC, _oe2.chainableComparers[0].orderType)
		assert.Equal(t, CLC_DESC, _oe2.chainableComparers[1].orderType)
		assert.Equal(t, CLC_ASC, _oe2.chainableComparers[2].orderType)

		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("when all same", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)
		_ = newIOrderedEnumerable(eSrc, func(v1, v2 int) int {
			return 0
		}, CLC_ASC).GetEnumerable()
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("when src nil", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		oe := newIOrderedEnumerable(eSrc, func(v1, v2 int) int {
			return 0
		}, CLC_ASC).(*orderedEnumerable[int])

		oe = nil

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "source is nil")
		}()

		_ = oe.GetEnumerable()
	})

	t.Run("when comparer nil", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)

		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		oe := newIOrderedEnumerable(eSrc, func(v1, v2 int) int {
			return 0
		}, CLC_ASC).(*orderedEnumerable[int])

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expect error")
				return
			}
			assert.Contains(t, fmt.Sprintf("%v", err), "comparer is nil")
		}()

		_ = oe.ThenBy(nil)
	})
}

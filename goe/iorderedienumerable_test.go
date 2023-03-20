package goe

import (
	"github.com/EscanBE/go-ienumerable/goe/comparers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func test_getSelfSelector[T any]() KeySelector[T] {
	return func(value T) any {
		return value
	}
}

func Test_IOrderedIEnumerable1(t *testing.T) {
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

	keySelector1 := func(v s1) any {
		return v.value
	}

	keySelector2 := func(v s1) any {
		return v.nested.value
	}

	keySelector3 := func(v s1) any {
		return v.nested.nested.value
	}

	tests := []struct {
		name    string
		ordered IOrderedEnumerable[s1]
		want    IEnumerable[s1]
	}{
		{
			name:    "asc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenBy(keySelector2, nil).ThenBy(keySelector3, nil),
			want:    NewIEnumerable[s1](v1530, v2160, v2420, v2430, v3530, v3990, v4530),
		},
		{
			name:    "asc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenBy(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[s1](v1530, v2160, v2430, v2420, v3530, v3990, v4530),
		},
		{
			name:    "asc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenByDescending(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[s1](v1530, v2430, v2420, v2160, v3990, v3530, v4530),
		},
		{
			name:    "desc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenByDescending(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[s1](v4530, v3990, v3530, v2430, v2420, v2160, v1530),
		},
		{
			name:    "desc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenBy(keySelector2, nil).ThenBy(keySelector3, nil),
			want:    NewIEnumerable[s1](v4530, v3530, v3990, v2160, v2420, v2430, v1530),
		},
		{
			name:    "desc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenBy(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[s1](v4530, v3530, v3990, v2160, v2430, v2420, v1530),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eGot := tt.ordered.GetOrderedEnumerable()

			assert.Truef(t, reflect.DeepEqual(tt.want.ToArray(), eGot.ToArray()), "got %v, want %v", eGot.ToArray(), tt.want.ToArray())

			bSrc.assertUnchanged(t, eSrc)
			bSrc.assertUnchangedIgnoreData(t, eGot)
		})
	}
}

func Test_IOrderedIEnumerable2(t *testing.T) {
	eSrc := NewIEnumerable[string]("v2430", "v1530", "v3530", "v4530", "v2420", "v2160", "v3990")
	bSrc := backupForAssetUnchanged(eSrc)

	keySelector1 := func(v string) any {
		return v[1]
	}

	keySelector2 := func(v string) any {
		return v[2]
	}

	keySelector3 := func(v string) any {
		return v[3]
	}

	tests := []struct {
		name    string
		ordered IOrderedEnumerable[string]
		want    IEnumerable[string]
	}{
		{
			name:    "asc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenBy(keySelector2, nil).ThenBy(keySelector3, nil),
			want:    NewIEnumerable[string]("v1530", "v2160", "v2420", "v2430", "v3530", "v3990", "v4530"),
		},
		{
			name:    "asc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenBy(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[string]("v1530", "v2160", "v2430", "v2420", "v3530", "v3990", "v4530"),
		},
		{
			name:    "asc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_ASC).ThenByDescending(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[string]("v1530", "v2430", "v2420", "v2160", "v3990", "v3530", "v4530"),
		},
		{
			name:    "desc-desc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenByDescending(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[string]("v4530", "v3990", "v3530", "v2430", "v2420", "v2160", "v1530"),
		},
		{
			name:    "desc-asc-asc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenBy(keySelector2, nil).ThenBy(keySelector3, nil),
			want:    NewIEnumerable[string]("v4530", "v3530", "v3990", "v2160", "v2420", "v2430", "v1530"),
		},
		{
			name:    "desc-asc-desc",
			ordered: newIOrderedEnumerable(eSrc, keySelector1, nil, CLC_DESC).ThenBy(keySelector2, nil).ThenByDescending(keySelector3, nil),
			want:    NewIEnumerable[string]("v4530", "v3530", "v3990", "v2160", "v2430", "v2420", "v1530"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eGot := tt.ordered.GetOrderedEnumerable()

			assert.Truef(t, reflect.DeepEqual(tt.want.ToArray(), eGot.ToArray()), "got %v, want %v", eGot.ToArray(), tt.want.ToArray())

			bSrc.assertUnchanged(t, eSrc)
			bSrc.assertUnchangedIgnoreData(t, eGot)
		})
	}
}

func Test_IOrderedIEnumerable3(t *testing.T) {
	t.Run("previous IOrderedComparer not changes when chaining more", func(t *testing.T) {
		eSrc := createRandomIntEnumerable(3)
		bSrc := backupForAssetUnchanged(eSrc)

		oe0 := newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}, CLC_ASC)

		_oe0 := oe0.(*orderedEnumerable[int])
		_oe1 := _oe0.ThenByDescending(test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}).(*orderedEnumerable[int])
		_oe2 := _oe1.ThenBy(test_getSelfSelector[int](), func(_, _ any) int {
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

		oe0 := newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}, CLC_ASC)

		_oe0 := oe0.(*orderedEnumerable[int])
		_oe1 := _oe0.ThenByDescending(test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}).(*orderedEnumerable[int])
		_oe2 := _oe1.ThenBy(test_getSelfSelector[int](), func(_, _ any) int {
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
		_ = newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}, CLC_ASC).GetOrderedEnumerable()
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("when single element", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3)
		bSrc := backupForAssetUnchanged(eSrc)
		got := eSrc.Order().GetOrderedEnumerable()
		assert.Len(t, got.ToArray(), 1)
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("compare using compare func", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)
		f := func(v1, v2 any) int {
			return 0
		}
		_ = newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), f, CLC_ASC).GetOrderedEnumerable()
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("compare using IComparer", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)
		_ = newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(x, y any) int {
			return comparers.NumericComparer.CompareAny(x, y)
		}, CLC_ASC).GetOrderedEnumerable()
		bSrc.assertUnchanged(t, eSrc)
	})
}

func Test_IOrderedIEnumerable4(t *testing.T) {
	t.Run("use compare func if exists", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)
		output := e[int](newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(x, y any) int {
			return comparers.NumericComparer.CompareAny(x, y) * -1
		}, CLC_ASC).GetOrderedEnumerable())
		assert.Len(t, output.data, 4)
		assert.Equal(t, 3, output.data[0])
		assert.Equal(t, 2, output.data[1])
		bSrc.assertUnchanged(t, eSrc)
	})

	t.Run("save & re-use use cached compare func, rather than resolve everytime", func(t *testing.T) {
		eSrc := NewIEnumerable[int](3, 1, 1, 2)
		bSrc := backupForAssetUnchanged(eSrc)
		output := e[int](newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), nil, CLC_ASC).GetOrderedEnumerable())
		assert.Len(t, output.data, 4)
		assert.Equal(t, 1, output.data[0])
		assert.Equal(t, 1, output.data[1])
		bSrc.assertUnchanged(t, eSrc)
	})
}

func Test_IOrderedIEnumerable5_panic(t *testing.T) {
	eSrc := NewIEnumerable[int](3, 1, 1, 2)
	bSrc := backupForAssetUnchanged(eSrc)

	oe := newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
		return 0
	}, CLC_ASC).(*orderedEnumerable[int])

	t.Run("panic when src nil", func(t *testing.T) {
		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		oe2 := newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}, CLC_ASC).(*orderedEnumerable[int])

		oe2 = nil

		defer deferExpectPanicContains(t, getErrorSourceIsNil().Error(), true)

		_ = oe2.GetOrderedEnumerable()
	})

	t.Run("panic when no default comparer", func(t *testing.T) {
		type MyStruct struct{}
		eSrc2 := NewIEnumerable[MyStruct](MyStruct{}, MyStruct{})
		bSrc2 := backupForAssetUnchanged(eSrc2)

		defer func() {
			bSrc2.assertUnchanged(t, eSrc2)
		}()

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyStruct]", true)

		oe2 := eSrc2.Order()

		_ = oe2.GetOrderedEnumerable()
	})

	t.Run("panic when no default comparer", func(t *testing.T) {
		type MyStruct struct{}
		eSrc2 := NewIEnumerable[MyStruct](MyStruct{}, MyStruct{})
		bSrc2 := backupForAssetUnchanged(eSrc2)

		defer func() {
			bSrc2.assertUnchanged(t, eSrc2)
		}()

		defer deferExpectPanicContains(t, "no default comparer registered for [goe.MyStruct]", true)

		oe2 := eSrc2.OrderByDescending()

		_ = oe2.GetOrderedEnumerable()
	})

	t.Run("panic when src nil", func(t *testing.T) {
		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		oe2 := newIOrderedEnumerable(eSrc, test_getSelfSelector[int](), func(_, _ any) int {
			return 0
		}, CLC_ASC).(*orderedEnumerable[int])

		oe2 = nil

		defer deferExpectPanicContains(t, getErrorSourceIsNil().Error(), true)

		_ = oe2.GetOrderedEnumerable()
	})

	t.Run("panic when key selector nil", func(t *testing.T) {
		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorKeySelectorNotNil().Error(), true)

		_ = oe.ThenBy(nil, nil)
	})

	t.Run("panic when key selector nil", func(t *testing.T) {
		defer func() {
			bSrc.assertUnchanged(t, eSrc)
		}()

		defer deferExpectPanicContains(t, getErrorKeySelectorNotNil().Error(), true)

		_ = oe.ThenByDescending(nil, nil)
	})
}

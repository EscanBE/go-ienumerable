package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

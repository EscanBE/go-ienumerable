package goe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_enumerable_Except_ExceptBy(t *testing.T) {
	equalityComparer := func(i1, i2 int) bool {
		return i1 == i2
	}

	tests := []struct {
		name              string
		source            IEnumerable[int]
		second            IEnumerable[int]
		want              IEnumerable[int]
		wantPanicExcept   bool
		equalityComparer  func(int, int) bool
		wantPanicExceptBy bool
	}{
		{
			name:             "except not any",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
		},
		{
			name:              "except not any",
			source:            NewIEnumerable[int](1, 2, 3),
			second:            NewIEnumerable[int](4, 5, 6, 7),
			wantPanicExcept:   true,
			wantPanicExceptBy: true,
		},
		{
			name:             "except one",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3, 4)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
		},
		{
			name:             "except some",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3, 5, 6)),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
		},
		{
			name:             "except when source empty",
			source:           injectIntComparers(NewIEnumerable[int]()),
			second:           NewIEnumerable[int](4, 5, 6, 7),
			want:             NewIEnumerable[int](),
			equalityComparer: equalityComparer,
		},
		{
			name:             "except when second empty",
			source:           injectIntComparers(NewIEnumerable[int](1, 2, 3)),
			second:           NewIEnumerable[int](),
			want:             NewIEnumerable[int](1, 2, 3),
			equalityComparer: equalityComparer,
		},
		{
			name:              "panic with nil src",
			source:            nil,
			second:            NewIEnumerable[int](4, 5, 6, 7),
			wantPanicExcept:   true,
			equalityComparer:  equalityComparer,
			wantPanicExceptBy: true,
		},
		{
			name:              "panic with nil second",
			source:            NewIEnumerable[int](1, 2, 3),
			second:            nil,
			wantPanicExcept:   true,
			equalityComparer:  equalityComparer,
			wantPanicExceptBy: true,
		},
		{
			name:              "panic with both nil",
			source:            nil,
			second:            nil,
			wantPanicExcept:   true,
			equalityComparer:  equalityComparer,
			wantPanicExceptBy: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Except", func(t *testing.T) {
			//copyOfSource := tt.source.copy()
			//copyOfSecond := tt.second.copy()
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.wantPanicExcept)

			// Except
			resultOfExcept2 := tt.source.Except(tt.second)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfExcept2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
		t.Run(tt.name+"_ExceptBy", func(t *testing.T) {
			//copyOfSource := tt.source.copy()
			//copyOfSecond := tt.second.copy()
			bSource := backupForAssetUnchanged(tt.source)
			bSecond := backupForAssetUnchanged(tt.second)

			defer deferWantPanicDepends(t, tt.wantPanicExceptBy)

			// Except
			resultOfExcept2 := tt.source.ExceptBy(tt.second, tt.equalityComparer)

			assert.True(t, reflect.DeepEqual(tt.want.ToArray(), resultOfExcept2.ToArray()))

			bSource.assertUnchanged(t, tt.source)
			bSecond.assertUnchanged(t, tt.second)
		})
	}
}

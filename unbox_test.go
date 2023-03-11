package go_ienumerable

//func Test_enumerable_Unbox(t *testing.T) {
//	t.Run("int8", func(t *testing.T) {
//		eSrc := NewIEnumerable[any](int8(5), int8(3)).WithDefaultComparers()
//
//		eGot := eSrc.UnboxInt8().(*enumerable[int8])
//		assert.Len(t, eGot.data, 2)
//		assert.Equal(t, 5, eGot.data[0])
//		assert.Equal(t, 3, eGot.data[1])
//
//		assert.Nil(t, eGot.equalityComparer)
//		assert.Nil(t, eGot.lessComparer)
//	})
//}

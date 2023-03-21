package goe

import (
	"testing"
)

func Test_enumerable_Zip_ImplementedAsHelper(t *testing.T) {
	t.Run("unable to call", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)
		createEmptyIntEnumerable().Zip_ImplementedInHelper()
	})
}

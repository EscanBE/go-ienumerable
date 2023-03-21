package goe

import (
	"testing"
)

func Test_enumerable_ImplementedAsHelper(t *testing.T) {
	t.Run("unable to call", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)
		createEmptyIntEnumerable().Aggregate_ImplementedInHelper()
	})
	t.Run("unable to call", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)
		createEmptyIntEnumerable().Chunk_ImplementedInHelper()
	})
	t.Run("unable to call", func(t *testing.T) {
		defer deferWantPanicDepends(t, true)
		createEmptyIntEnumerable().Zip_ImplementedInHelper()
	})
}

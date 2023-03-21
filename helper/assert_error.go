package helper

import (
	"fmt"
	"github.com/EscanBE/go-ienumerable/goe"
)

func assertCollectionNotNil[T any](collection goe.IEnumerable[T], collectionName string) {
	if collection == nil {
		panic(fmt.Sprintf("%s collection is nil", collectionName))
	}
}

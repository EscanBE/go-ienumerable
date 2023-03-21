package goe_helper

import (
	"github.com/EscanBE/go-ienumerable/goe"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestToDictionary(t *testing.T) {
	type Package struct {
		Company        string
		Weight         float64
		TrackingNumber int64
	}

	iePackages := goe.NewIEnumerable(
		Package{
			Company:        "Coho Vineyard",
			Weight:         25.2,
			TrackingNumber: 89453312,
		},
		Package{
			Company:        "Lucerne Publishing",
			Weight:         18.7,
			TrackingNumber: 89112755,
		},
		Package{
			Company:        "Wingtip Toys",
			Weight:         6.0,
			TrackingNumber: 299456122,
		},
		Package{
			Company:        "Adventure Works",
			Weight:         33.8,
			TrackingNumber: 4665518773,
		},
	)

	t.Run("C# example", func(t *testing.T) {
		dictionary := ToDictionary(iePackages, func(src Package) int64 {
			return src.TrackingNumber
		}, func(src Package) Package {
			return src
		})

		assert.Len(t, dictionary, 4)
		v1, found1 := dictionary[89453312]
		v2, found2 := dictionary[89112755]
		v3, found3 := dictionary[299456122]
		v4, found4 := dictionary[4665518773]
		assert.True(t, found1)
		assert.True(t, found2)
		assert.True(t, found3)
		assert.True(t, found4)
		assert.Equal(t, "Coho Vineyard", v1.Company)
		assert.Equal(t, "Lucerne Publishing", v2.Company)
		assert.Equal(t, "Wingtip Toys", v3.Company)
		assert.Equal(t, "Adventure Works", v4.Company)
	})

	t.Run("nil key selector", func(t *testing.T) {
		defer deferExpectPanicContains(t, "key selector is nil", true)

		_ = ToDictionary[Package, int64, Package](iePackages, nil, func(src Package) Package {
			return src
		})
	})

	t.Run("nil element selector", func(t *testing.T) {
		defer deferExpectPanicContains(t, "element selector is nil", true)

		_ = ToDictionary[Package, int64, Package](iePackages, func(src Package) int64 {
			return src.TrackingNumber
		}, nil)
	})

	t.Run("duplicated key", func(t *testing.T) {
		radKey := rand.Int63()

		defer deferExpectPanicContains(t, "duplicated key", true)

		_ = ToDictionary[Package, int64, Package](goe.NewIEnumerable(
			Package{
				TrackingNumber: radKey,
			},
			Package{
				TrackingNumber: radKey,
			},
		), func(src Package) int64 {
			return src.TrackingNumber
		}, func(src Package) Package {
			return src
		})
	})
}

# Golang's IEnumerable[T]
### with inspiration from C#'s famous IEnumerable&lt;T&gt;

# go-ienumerable [![GoDoc](https://godoc.org/github.com/EscanBE/go-ienumerable?status.svg)](https://godoc.org/github.com/EscanBE/go-ienumerable) [![Go Report Card](https://goreportcard.com/badge/github.com/EscanBE/go-ienumerable)](https://goreportcard.com/report/github.com/EscanBE/go-ienumerable)

> Code coverage: 100% files, 96.9% statements covered by 11k+ lines of test over 4k+ lines of code

> Require: Go 1.18+

Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerable-1) IEnumerable[T] [here in the enumerable interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerable_interface.go) definition

In addition: Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerator-1) IEnumerator[T] [here in the enumerator interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerator_interface.go) definition

```go
got := goe.NewIEnumerable[string]("Hello", "World").
    Where(func(v string) bool {
        return len(v) < 3
    }).
    OrderByDescending().GetOrderedEnumerable().
    Reverse().
    FirstOrDefault(nil, goe.Ptr("Oops"))

fmt.Println(got)
// "Oops"
```
```go
array := []byte{0, 70, 99, 106, 106, 109, 30, 85, 109, 112, 106, 98, 99, 66, 88, 69}
got := goe.NewIEnumerable[byte](array...).
    Skip(1).
    Take(11).
    Select(transform).
    CastInt32().
    Append('"').
    AggregateWithAnySeed("\"", aggregate)

fmt.Println(got)
// "Hello World"
```

## Notices about value comparison:
Some `IEnumerable[T]` methods like `Order`, `Distinct`, `Except`, `Union`, `Intersect`,... need value comparator to compare between each element in the sequence and the `IComparer[T]` is needed.

This is definition of `IComparer[T]` in `comparers` package:
```go
// IComparer use methods Compare* to compare value of 2 input values.
//
// If left is less than right, returns -1.
//
// If left is equals to right, returns 0.
//
// If left is greater than right, returns 1.
type IComparer[T any] interface {
    // CompareTyped compares value from params.
    //
    // If x is less than y, returns -1.
    //
    // If x is equals to y, returns 0.s
    //
    // If x is greater than y, returns 1.
    CompareTyped(x, y T) int

    // CompareAny accept any params.
    //
    // If both x and y are nil, return 0.
    //
    // If x is nil and y is not nil, return -1.
    //
    // If x is not nil and y is nil, return 1.
    //
    // The rest, implement in your own way, since type any means you can pass everything here,
    // and you should handle them carefully
    CompareAny(x, y any) int
}
```
See implementation sample in `example`.

`go-ienumerable` will attempts to resolve a default comparer using predefined comparers for some type. You can register a comparer for `YourType` by implement your own `IComparer[YourType]`.
See sample of implement and default comparer registration for custom types and other types in `example`.

Predefined `IComparer[T]`: `string`, `bool`, `numeric` (`int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `float32/64`, `complex64/128`), `*big.Int`, `*big.Float`, `time.Time` with corresponding initialized comparer instance, eg: `compares.StringComparer` is a string comparer, `compares.NumericComparer` is comparer of numeric and so on.
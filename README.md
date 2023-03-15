# Golang's IEnumerable[T]
### with inspiration from C#'s famous IEnumerable&lt;T&gt;

# go-ienumerable [![GoDoc](https://godoc.org/github.com/EscanBE/go-ienumerable?status.svg)](https://godoc.org/github.com/EscanBE/go-ienumerable) [![Go Report Card](https://goreportcard.com/badge/github.com/EscanBE/go-ienumerable)](https://goreportcard.com/report/github.com/EscanBE/go-ienumerable)

> Code coverage: 100% files, 100% statements

Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerable-1) IEnumerable[T] [here in the enumerable interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerable_interface.go) definition

In addition: Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerator-1) IEnumerator[T] [here in the enumerator interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerator_interface.go) definition

```go
got := goe.NewIEnumerable[string]("Hello", "World").
    Where(func(v string) bool {
        return len(v) < 3
    }).OrderByDescending().Reverse().
    FirstOrDefaultUsing("\"Oops\"")
fmt.Println(got)
// "Oops"
```
```go
array := []byte{0, 70, 99, 106, 106, 109, 30, 85, 109, 112, 106, 98, 99, 66, 88, 69}
got := goe.NewIEnumerable[byte](array...).
        Skip(1).Take(11).Select(func(v byte) any {
        return v + 2
    }).CastInt32().Append('"').
    AggregateWithAnySeed("\"", func(str any, v int32) any {
        return fmt.Sprintf("%s%c", str, v)
    })
fmt.Println(got)
// "Hello World"
```

## Notices about value comparison:
Some `IEnumerable[T]` methods like `Order`, `Distinct`, `Except`, `Union`, `Intersect`,... need value comparator to compare between each element in the sequence and the `IComparer[T]` is needed.

This is definition of `IComparer[T]` in `comparers` package:
```go
type IComparer[T any] interface {
    // Compare compares value from params.
    //
    // If x is less than y, returns -1.
    //
    // If x is equals to y, returns 0.
    //
    // If x is greater than y, returns 1.
    Compare(x, y T) int

    // ComparePointerMode compares two params but in pointer presentation (like *int vs *int).
    //
    // If both x and y are nil, return 0.
    //
    // If x is nil and y is not nil, return -1.
    //
    // If x is not nil and y is nil, return 1.
    //
    // If both x and y are not nil, do like Compare does.
    //
    // Implementation must support both type of input param *T or *any (*interface{}) of T.
    ComparePointerMode(x, y *T) int
}
```
See implementation sample in `example`.

`go-ienumerable` will attempts to resolve a default comparer using predefined comparers for some type. You can register a comparer for `YourType` by implement your own `IComparer[YourType]`.
See sample of implement and default comparer registration for custom types and other types in `example`.

Predefined `IComparer[T]`: `string`, `bool`, `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64`, `float32/64`, `*big.Int`, `complex64/128`, `time.Time`, `time.Duration` with corresponding initialized comparer instance, eg: `compares.StringComparer` is a string comparer, `compares.Uint32Comparer` is uint32 comparer and so on.
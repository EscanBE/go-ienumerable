# Golang's IEnumerable[T]
### with inspiration from C#'s famous IEnumerable&lt;T&gt;

# go-ienumerable [![GoDoc](https://godoc.org/github.com/EscanBE/go-ienumerable?status.svg)](https://godoc.org/github.com/EscanBE/go-ienumerable) [![Go Report Card](https://goreportcard.com/badge/github.com/EscanBE/go-ienumerable)](https://goreportcard.com/report/github.com/EscanBE/go-ienumerable)

> Code coverage: 100% files, 100% statements

Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerable-1) IEnumerable[T] [here in the enumerable interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerable_interface.go) definition

In addition: Check the methods [ported from C#](https://learn.microsoft.com/en-us/dotnet/api/system.collections.generic.ienumerator-1) IEnumerator[T] [here in the enumerator interface](https://github.com/EscanBE/go-ienumerable/blob/main/goe/ienumerator_interface.go) definition

```go
got := goe.NewIEnumerable[string]("Hello", "World").WithDefaultComparers().
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
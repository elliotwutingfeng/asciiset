# asciiset

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/elliotwutingfeng/asciiset)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/asciiset?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/asciiset)
[![Codecov Coverage](https://img.shields.io/codecov/c/github/elliotwutingfeng/asciiset?color=bright-green&logo=codecov&style=for-the-badge&token=5ukdyK4pOG)](https://codecov.io/gh/elliotwutingfeng/asciiset)

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)

## Summary

**asciiset** is an [ASCII](https://simple.wikipedia.org/wiki/ASCII) character bitset.

Bitsets are fast and memory-efficient data structures for storing and retrieving information using bitwise operations.

**asciiset** is an extension of the **asciiSet** data structure from the Go Standard library [source code](https://cs.opensource.google/go/go/+/master:src/bytes/bytes.go).

Possible applications include checking strings for prohibited ASCII characters, and counting unique ASCII characters in a string.

Spot any bugs? Report them [here](https://github.com/elliotwutingfeng/asciiset/issues).

![ASCII Table](ASCII-Table.svg)

## Installation

```bash
go get github.com/elliotwutingfeng/asciiset
```

## Testing

```bash
make tests

# Alternatively, run tests without race detection
# Useful for systems that do not support the -race flag like windows/386
# See https://tip.golang.org/src/cmd/dist/test.go
make tests_without_race
```

## Benchmarks

Benchmarks comparing performance between **asciiset** and **map[byte]struct{}** sets are provided.

On average, compared to **map[byte]struct{}** sets, **asciiset** has 11 times the element addition speed, 29 times the element lookup speed, 1.5 times the element removal speed, and equivalent set length lookup speed.

```bash
make bench
```

### Results

```bash
go test -bench . -benchmem -cpu 1
goos: linux
goarch: amd64
pkg: github.com/elliotwutingfeng/asciiset
cpu: AMD Ryzen 7 5800X 8-Core Processor
BenchmarkASCIISet/ASCIISet_Add()                 1388647               869.9 ns/op             0 B/op          0 allocs/op
BenchmarkASCIISet/ASCIISet_Contains()            2038735               586.0 ns/op             0 B/op          0 allocs/op
BenchmarkASCIISet/ASCIISet_Remove()               742290              1624 ns/op               0 B/op          0 allocs/op
BenchmarkASCIISet/ASCIISet_Size()                3744616               319.9 ns/op             0 B/op          0 allocs/op
BenchmarkMapSet/map_Add                           121694              9998 ns/op               0 B/op          0 allocs/op
BenchmarkMapSet/map_Contains                       69081             17179 ns/op               0 B/op          0 allocs/op
BenchmarkMapSet/map_Remove                        468333              2561 ns/op               0 B/op          0 allocs/op
BenchmarkMapSet/map_Size                         3440536               337.4 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/elliotwutingfeng/asciiset    12.037s
```

# Sets
[![codecov](https://codecov.io/gh/clarkmcc/go-set/branch/master/graph/badge.svg?token=rwZFoBpjiu)](https://codecov.io/gh/clarkmcc/go-set)

A simple, zero-dependency generic set implementation for Go extracted from the Kubernetes apimachinery project and converted to use Go 1.18 generics. This package contains the following set methods:
* `Insert`
* `Delete`
* `Has`
* `HasAll`
* `HasAny`
* `Difference`
* `Union`
* `Intersection`
* `IsSuperset`
* `Equal`
* `List`
* `PopAny`
* `Len`

## Installation

    go get github.com/clarkmcc/go-set

## Example
```go
s := sets.New[int](1, 2, 3)
s.Insert(3, 4)
assert.Len(s.Len(), 4)
```
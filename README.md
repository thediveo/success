# `success`

[![Go Reference](https://pkg.go.dev/badge/github.com/thediveo/success.svg)](https://pkg.go.dev/github.com/thediveo/success)
![GitHub](https://img.shields.io/github/license/thediveo/success)
![build and test](https://github.com/TheDiveO/success/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/success)](https://goreportcard.com/report/github.com/thediveo/success)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

This tiny Go module improves your testing feng shui tremendiously by shimming in
Gomega error detection whenever trying to be `Successful()`. All your software
will immediately look so much better by importing success!

#### Before

```go
sekret, err := Foo(42)
Expect(err).NotTo(HaveOccured())
```

#### After

```go
// You might want to dot-import for convenience.
import . "github.com/thediveo/success"

sekret := Successful(Foo(42))

// analogous...
sekret, moresekret := Successful2R(Bar(12345))
sekret, moresekret, nosekretanymore := Successful3R(Baz())
```

## Go Version Support

`success` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Copyright and License

`success` is Copyright 2023 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
